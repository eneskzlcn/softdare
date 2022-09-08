package mocks

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
)

func NewMockPostgresDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil
	}
	return db, mock
}
