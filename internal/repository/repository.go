package repository

import (
	"context"
	"database/sql"
	"github.com/wordict/w-api/internal/core/logger"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository struct {
	logger logger.Logger
	db     DB
}

func New(db DB, logger logger.Logger) *Repository {
	return &Repository{logger: logger, db: db}
}
