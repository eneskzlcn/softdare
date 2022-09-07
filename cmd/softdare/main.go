package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/home"
	"github.com/eneskzlcn/softdare/internal/login"
	"github.com/eneskzlcn/softdare/postgres"
	"github.com/eneskzlcn/softdare/server"
	"go.uber.org/zap"
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
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	configs, err := config.LoadConfig(".dev/", "local", "yaml")
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
	sessionProvider := server.NewSessionProvider(logger, configs.Server.SessionKey)

	loginRepository := login.NewRepository(db)
	loginService := login.NewService(loginRepository)
	loginHandler := login.NewHandler(logger, loginService, renderer, sessionProvider)

	homeHandler := home.NewHandler(logger, renderer, sessionProvider)
	handler, err := server.NewHandler(logger, []server.RouteHandler{
		loginHandler,
		homeHandler,
	}, sessionProvider)
	if err != nil {
		return err
	}
	server := server.New(configs.Server, handler)

	defer server.Close()

	logger.Infof("server started to listening and serve at address %s", configs.Server.Address)
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve %w", err)
	}
	return nil
}
