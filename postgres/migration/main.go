package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/postgres"
	"os"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func main() {
	if err := run(); err != nil {
		fmt.Println("DDL command not worked.")
		fmt.Println(err.Error())
	} else {
		fmt.Println("DDL command worked.")
	}
}

func run() error {
	conf, err := config.LoadConfig(".dev/", "local", "yaml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	postgresDB, err := postgres.New(conf.Db)
	if err != nil {
		fmt.Println("error when initializing database", err.Error())
		return err
	}
	var ddlType string

	fs := flag.NewFlagSet("softdare", flag.ExitOnError)
	fs.StringVar(&ddlType, "type", "migrate", `
		Enter the type of the operation you want to perform.
		migrate: will create all the tables
		drop: will drop all the tables`)
	if err = fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags error")
	}
	switch ddlType {
	case "migrate":
		if err = MigrateTables(context.Background(), postgresDB); err != nil {
			return fmt.Errorf("migration error")
		}
	case "drop":
		if err = DropTables(context.Background(), postgresDB); err != nil {
			return fmt.Errorf("drop error")
		}
	default:
		return errors.New("not valid flag type for action")
	}
	return nil
}
