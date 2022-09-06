package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"softdare/config"
	"softdare/postgres"
	"softdare/web"
	"softdare/web/login"
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

	config, err := config.LoadConfig(".dev/", "local", "yaml")
	if err != nil {
		return err
	}
	logger.Infof("Config Loaded %v", config)
	db, err := postgres.New(config.Db)
	if err != nil {
		return err
	}
	if err = postgres.MigrateTables(context.Background(), db); err != nil {
		return err
	}

	loginRepository := login.NewRepository(db)
	loginService := login.NewService(loginRepository)

	handler := web.NewHandler(logger, loginService, []byte(config.App.SessionKey))

	if err != nil {
		return err
	}
	srv := http.Server{
		Addr:    config.App.Address,
		Handler: handler,
	}

	defer srv.Close()

	logger.Infof("server started to listening and serve at address %s", config.App.Address)
	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve %w", err)
	}
	return nil
}
