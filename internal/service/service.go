//go:generate mockgen -destination=./mock_service.go -source=./service.go -package=service
package service

import "github.com/wordict/w-api/internal/core/logger"

type Repository interface {
}

type Service struct {
	logger     logger.Logger
	repository *Repository
}

func New(repository Repository, logger logger.Logger) *Service {
	return &Service{repository: &repository, logger: logger}
}
