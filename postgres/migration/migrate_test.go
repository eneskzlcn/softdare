package main_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eneskzlcn/softdare/postgres"
	schema "github.com/eneskzlcn/softdare/postgres/schema"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestMigrateTables(t *testing.T) {
	db, mock := postgres.NewMockPostgres()
	fileBytes, err := ioutil.ReadFile("./schema.sql")
	assert.Nil(t, err)
	mock.ExpectExec(regexp.QuoteMeta(string(fileBytes))).WillReturnResult(sqlmock.NewResult(1, 1))
	err = schema.MigrateTables(context.Background(), db)
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}
