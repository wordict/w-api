package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wordict/w-api/internal/entity"
	"net/http"
)

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
