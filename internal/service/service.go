//go:generate mockgen -destination=./mock_service.go -source=./service.go -package=service
package service

import (
	"context"
	"github.com/wordict/w-api/internal/core/logger"
	"github.com/wordict/w-api/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, signupRequestDto entity.SignupRequestDTO) error
	IsUserExist(ctx context.Context, signInRequestDto entity.SignInRequestDTO) (entity.User, error)
}

type Service struct {
	logger     logger.Logger
	repository Repository
}

func New(repository Repository, logger logger.Logger) *Service {
	return &Service{repository: repository, logger: logger}
}
