//go:generate mockgen -destination=./mock_service.go -source=./service.go -package=service
package service

import (
	"context"
	"github.com/wordict/w-api/internal/core/logger"
	"github.com/wordict/w-api/internal/entity"
)

type TranslationClient interface {
	Translate(data, fromLanguage, toLanguage string) (string, error)
}

type Repository interface {
	CreateUser(ctx context.Context, signupRequestDto entity.SignupRequestDTO) error
	IsUserExist(ctx context.Context, signInRequestDto entity.SignInRequestDTO) (entity.User, error)
	AddVocabulary(ctx context.Context, dto entity.AddVocabularyRequestDTO) (entity.Vocabulary, error)
}

type Service struct {
	logger     logger.Logger
	repository Repository
	client     TranslationClient
}

func New(repository Repository, logger logger.Logger) *Service {
	return &Service{repository: repository, logger: logger}
}
