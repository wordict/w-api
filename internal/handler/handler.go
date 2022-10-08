package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/wordict/w-api/internal/core/logger"
	"github.com/wordict/w-api/internal/entity"
)

type UserService interface {
	SignUp(ctx context.Context, request entity.SignupRequest) error
	SignIn(ctx context.Context, request entity.SignInRequest) (entity.SignInResponse, error)
}
type VocabularyService interface {
	AddVocabulary(ctx context.Context, request entity.AddVocabularyRequest) (entity.AddVocabularyResponse, error)
}
type Service interface {
	UserService
	VocabularyService
}

type Handler struct {
	logger  logger.Logger
	service Service
}

func New(service Service, logger logger.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/signup", h.SignUp)
	app.Post("/signIn", h.SignIn)
	app.Post("/vocabulary", h.AuthenticationMiddleware(h.AddVocabulary))
}

func (h *Handler) AuthenticationMiddleware(next fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return next(ctx)
	}
}
