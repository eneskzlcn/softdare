package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/comment"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/home"
	"github.com/eneskzlcn/softdare/internal/login"
	"github.com/eneskzlcn/softdare/internal/post"
	"github.com/eneskzlcn/softdare/internal/server"
	osUtil "github.com/eneskzlcn/softdare/internal/util/os"
	"github.com/eneskzlcn/softdare/postgres"
	"github.com/eneskzlcn/softdare/rabbitmq"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	env := osUtil.GetEnv("DEPLOYMENT_ENVIRONMENT", "local")
	logger := logger.NewZapLoggerAdapter(env)

	defer logger.Sync()

	configs, err := config.LoadConfig(".dev/", env, "yaml")
	if err != nil {
		return err
	}

	db, err := postgres.New(configs.Db)
	if err != nil {
		return err
	}
	if err = postgres.MigrateTables(context.Background(), db); err != nil {
		return err
	}
	session := session.NewCollegeSessionAdapter(logger, configs.Session)
	rabbitmqClient := rabbitmq.New(configs.RabbitMQ, logger)

	loginRepository := login.NewRepository(logger, db)
	loginService := login.NewService(logger, loginRepository)
	loginHandler := login.NewHandler(logger, loginService, session)

	commentRepository := comment.NewRepository(db, logger)
	commentService := comment.NewService(logger, commentRepository, rabbitmqClient)
	commentHandler := comment.NewHandler(logger, commentService, session)

	postRepository := post.NewRepository(db, logger)
	postService := post.NewService(postRepository, logger)
	postHandler := post.NewHandler(logger, postService, session, commentService)

	homeService := home.NewService(postService, logger)
	homeHandler := home.NewHandler(logger, session, homeService)

	handler, err := server.NewHandler(logger, []server.RouteHandler{
		loginHandler,
		homeHandler,
		postHandler,
		commentHandler,
	}, session)

	go post.IncreasePostCommentCountConsumer(rabbitmqClient, postService, logger)
	if err != nil {
		return err
	}
	server := server.New(configs.Server, handler, logger)
	defer server.Close()

	logger.Infof("server started to listening and serve at address %s", configs.Server.Address)
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve %w", err)
	}
	return nil
}
