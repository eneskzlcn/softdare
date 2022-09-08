package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/home"
	"github.com/eneskzlcn/softdare/internal/login"
	"github.com/eneskzlcn/softdare/internal/post"
	"github.com/eneskzlcn/softdare/internal/server"
	loggerUtil "github.com/eneskzlcn/softdare/internal/util/logger"
	osUtil "github.com/eneskzlcn/softdare/internal/util/os"
	"github.com/eneskzlcn/softdare/postgres"

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
	logger, err := loggerUtil.NewLoggerForEnv(env)
	if err != nil {
		return err
	}
	defer logger.Sync()

	configs, err := config.LoadConfig(".dev/", env, "yaml")
	if err != nil {
		return err
	}

	logger.Debugf("CONFIG LOADED %v", configs)
	db, err := postgres.New(configs.Db)
	if err != nil {
		return err
	}
	if err = postgres.MigrateTables(context.Background(), db); err != nil {
		return err
	}

	renderer := server.NewRenderer(logger)
	sessionProvider := server.NewSessionProvider(logger, configs.Session)

	loginRepository := login.NewRepository(logger, db)
	loginService := login.NewService(logger, loginRepository)
	loginHandler := login.NewHandler(logger, loginService, renderer, sessionProvider)

	homeHandler := home.NewHandler(logger, renderer, sessionProvider)

	postRepository := post.NewRepository(db, logger)
	postService := post.NewService(postRepository, logger)
	postHandler := post.NewHandler(logger, postService, renderer, sessionProvider)

	handler, err := server.NewHandler(logger, []server.RouteHandler{
		loginHandler,
		homeHandler,
		postHandler,
	}, sessionProvider)
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
