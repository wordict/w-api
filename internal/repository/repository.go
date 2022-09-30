package repository

import "github.com/wordict/w-api/internal/core/logger"

type DB interface {
}

type Repository struct {
	logger logger.Logger
	db     DB
}

func New(db DB, logger logger.Logger) *Repository {
	return &Repository{logger: logger, db: db}
}
