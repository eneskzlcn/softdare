package main

import (
	"context"
	_ "embed"
)

//go:embed drop.sql
var dropSql string

func DropTables(ctx context.Context, db DB) error {
	_, err := db.ExecContext(ctx, dropSql)
	return err
}
