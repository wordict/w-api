package main

import (
	"fmt"
	"github.com/wordict/w-api/internal/config"
	"github.com/wordict/w-api/internal/core/logger"
	"github.com/wordict/w-api/internal/core/postgres"
	"github.com/wordict/w-api/internal/core/server"
	"github.com/wordict/w-api/internal/handler"
	"github.com/wordict/w-api/internal/repository"
	"github.com/wordict/w-api/internal/service"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	env, exists := os.LookupEnv("DEPLOYMENT_ENV")
	if !exists {
		env = "local"
	}
	conf, err := config.LoadConfig(".dev", env, "yaml")
	if err != nil {
		return err
	}

	logger, err := logger.NewZapLoggerForEnv(env, 0)
	if err != nil {
		return err
	}

	db, err := postgres.New(conf.Db)
	if err != nil {
		return err
	}
	repository := repository.New(db, logger)
	service := service.New(repository, logger)
	handler := handler.New(service, logger)
	server := server.New([]server.Handler{
		handler,
	}, conf.Server, logger)

	return server.Start()
}
