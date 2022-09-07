package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	_ "github.com/lib/pq"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(config config.DB) (DB, error) {
	db, err := sql.Open("postgres", createDSN(config))
	if err != nil {
		return nil, errors.New("error opening the database")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.New("error pinging the database")
	}
	return db, nil
}

func createDSN(config config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DBName)
}
