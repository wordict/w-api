package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/wordict/w-api/internal/core/logger"
	"github.com/wordict/w-api/internal/entity"
	"net/http"
)

type Service interface {
	SignUp(ctx context.Context, request entity.SignupRequest) error
	SignIn(ctx context.Context, request entity.SignInRequest) (entity.SignInResponse, error)
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
}

func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	var request entity.SignupRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	h.logger.Debugf("SignUp request arrived with User Email:", request.Email)
	if err := h.service.SignUp(ctx.Context(), request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *Handler) SignIn(ctx *fiber.Ctx) error {
	var request entity.SignInRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	h.logger.Debugf("SignUp request arrived with User Email:", request.Email)
	user, err := h.service.SignIn(ctx.Context(), request)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(user)
}
