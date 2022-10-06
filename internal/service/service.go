//go:generate mockgen -destination=./mock_service.go -source=./service.go -package=service
package service

import (
	"context"
	"github.com/rs/xid"
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

func (s *Service) SignUp(ctx context.Context, request entity.SignupRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}
	userID := xid.New().String()
	return s.repository.CreateUser(ctx, entity.SignupRequestDTO{
		ID:       userID,
		Email:    request.Email,
		Password: request.Password,
	})
}

func (s *Service) SignIn(ctx context.Context, request entity.SignInRequest) (entity.SignInResponse, error) {
	user, err := s.repository.IsUserExist(ctx, entity.SignInRequestDTO{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return entity.SignInResponse{}, err
	}
	return entity.SignInResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
