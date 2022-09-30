package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wordict/w-api/internal/core/logger"
)

type Service interface {
}

type Handler struct {
	logger  logger.Logger
	service *Service
}

func New(service Service, logger logger.Logger) *Handler {
	return &Handler{service: &service, logger: logger}
}

func (h *Handler) Signup(ctx *fiber.Ctx) error {
	return ctx.SendString("signup")
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/signup", h.Signup)
}
