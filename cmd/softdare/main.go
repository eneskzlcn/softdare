package main

import (
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/client/queue"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/cache"
	"github.com/eneskzlcn/softdare/internal/core/html/template/renderer"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/repository"
	"github.com/eneskzlcn/softdare/internal/server"
	"github.com/eneskzlcn/softdare/internal/service"
	"github.com/eneskzlcn/softdare/internal/web"
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
	env, exists := os.LookupEnv("DEPLOYMENT_ENVIRONMENT")
	if !exists {
		env = "local"
	}
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

	renderer := renderer.New(logger)
	session := session.NewCollegeSessionAdapter(logger, configs.Session)
	rabbitmqClient := rabbitmq.New(configs.RabbitMQ, logger)
	cache := cache.NewGCacheAdapter(5)

	repository := repository.New(logger, db)
	service := service.New(repository, logger, session, rabbitmqClient, cache)
	webHandler := web.NewHandler(logger, session, service, renderer)
	client := queue.New(rabbitmqClient, logger, service)

	go client.ConsumeCommentCreated()
	go client.ConsumePostCreated()
	go client.ConsumeUserFollowCreated()
	go client.ConsumeUserFollowDeleted()
	go client.ConsumePostLikeCreated()
	go client.ConsumeCommentLikeCreated()

	if err != nil {
		return err
	}
	server := server.New(configs.Server, webHandler, logger)
	defer server.Close()

	logger.Infof("server started to listening and serve at address %s", configs.Server.Address)
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve %w", err)
	}
	return nil
}
