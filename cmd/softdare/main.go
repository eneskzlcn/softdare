package main

import (
	"errors"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"softdare/web"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	logger := zap.NewExample().Sugar()
	handler := web.NewHandler(logger)
	defer logger.Sync()

	var address string
	err := addAddressFlagVariable(&address)
	if err != nil {
		return err
	}
	srv := http.Server{
		Addr:    address,
		Handler: handler,
	}

	defer srv.Close()

	logger.Info("server started to listening and serve", zap.String("Address", address))
	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve %w", err)
	}
	return nil
}

func addAddressFlagVariable(address *string) error {
	fs := flag.NewFlagSet("softdare", flag.ExitOnError)
	fs.StringVar(address, "addr", ":4000", "HTTP Serve Address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags %v", err)
	}
	return nil
}
