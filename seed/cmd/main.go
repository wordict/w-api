package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wordict/w-api/internal/config"
	"github.com/wordict/w-api/seed"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func run() error {
	conf, err := config.LoadConfig(".dev/", "local", "yaml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Port, conf.Db.Username, conf.Db.Password, conf.Db.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("error when initializing database", err.Error())
		return err
	}

	var ddlType string

	fs := flag.NewFlagSet("wordict", flag.ExitOnError)
	fs.StringVar(&ddlType, "type", "migrate", `
		Enter the type of the operation you want to perform.
		migrate: will create all the tables
		drop: will drop all the tables`)
	if err = fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags error")
	}
	switch ddlType {
	case "migrate":
		if err = seed.MigrateTables(context.Background(), db); err != nil {
			return fmt.Errorf("migration error")
		}
		break
	case "drop":
		if err = seed.DropTables(context.Background(), db); err != nil {
			return fmt.Errorf("drop error")
		}
		break
	default:
		return errors.New("not valid flag type for action")
	}
	return nil
}
