package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/wordict/w-api/internal/core/logger"
	"github.com/wordict/w-api/internal/entity"
)

type Service interface {
	Signup(ctx context.Context, request entity.SignupRequest) error
}

type Handler struct {
	logger  logger.Logger
	service Service
}

func New(service Service, logger logger.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) Signup(ctx *fiber.Ctx) error {
	var request entity.SignupRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	h.logger.Debugf("Signup request arrived with User Email:", request.Email)
	if err := h.service.Signup(ctx.Context(), request); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/signup", h.Signup)
}
