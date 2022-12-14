package main

import (
	"context"
	_ "embed"
)

//go:embed schema.sql
var schemaSql string

func MigrateTables(ctx context.Context, db DB) error {
	_, err := db.ExecContext(ctx, schemaSql)
	return err
}
