package main

import (
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/client/queue"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/cache"
	"github.com/eneskzlcn/softdare/internal/core/html/template/renderer"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/mail"
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
	//
	internalServiceLogger := logger.NewZapLoggerAdapter(env, 1)
	//when there is an external service on use, we should caller skip 2 to see the exact place of the log.
	externalServiceLogger := logger.NewZapLoggerAdapter(env, 2)
	defer internalServiceLogger.Sync()
	defer externalServiceLogger.Sync()

	configs, err := config.LoadConfig[config.Config](".dev/", env, "yaml")
	if err != nil {
		return err
	}
	secretConfigs, err := config.LoadConfig[config.Secrets](".secrets/", "secrets", "yaml")
	if err != nil {
		return err
	}

	db, err := postgres.New(configs.Db)
	if err != nil {
		return err
	}

	renderer := renderer.New(externalServiceLogger)
	session := session.NewCollegeSessionAdapter(externalServiceLogger, configs.Session)
	rabbitmqClient := rabbitmq.New(configs.RabbitMQ, externalServiceLogger)
	mailService := mail.NewGomailServiceAdapter(secretConfigs.MailService, externalServiceLogger)
	cache := cache.NewGCacheAdapter(5)

	repository := repository.New(internalServiceLogger, db)
	service := service.New(repository, internalServiceLogger, session, rabbitmqClient, cache)
	webHandler := web.NewHandler(internalServiceLogger, session, service, renderer)
	client := queue.New(rabbitmqClient, internalServiceLogger, service, mailService)

	go client.ConsumeCommentCreated()
	go client.ConsumePostCreated()
	go client.ConsumeUserFollowCreated()
	go client.ConsumeUserFollowDeleted()
	go client.ConsumePostLikeCreated()
	go client.ConsumeCommentLikeCreated()
	go client.ConsumeUserCreated()

	if err != nil {
		return err
	}
	server := server.New(configs.Server, webHandler, internalServiceLogger)
	defer server.Close()

	internalServiceLogger.Infof("server started to listening and serve at address %s", configs.Server.Address)
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve %w", err)
	}
	return nil
}
