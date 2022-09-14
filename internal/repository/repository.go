package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/entity"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository struct {
	db     DB
	logger logger.Logger
}

func New(logger logger.Logger, db DB) *Repository {
	if logger == nil {
		fmt.Println("given logger is nil")
		return nil
	}
	if db == nil {
		logger.Error(entity.NilDatabase)
		return nil
	}
	return &Repository{db: db, logger: logger}
}
