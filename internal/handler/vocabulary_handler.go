package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wordict/w-api/internal/entity"
)

func (h *Handler) AddVocabulary(ctx *fiber.Ctx) error {
	userID := ctx.Get("userID", "")
	userNativeLanguage := ctx.Get("userNativeLang", "en")
	userTargetLanguage := ctx.Get("userTargetLang", "tr")

	var request entity.AddVocabularyRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.logger.Debug("Bad request for add vocabulary")
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	request.UserID = userID
	request.FromLanguage = userNativeLanguage
	request.ToLanguage = userTargetLanguage
	response, err := h.service.AddVocabulary(ctx.Context(), request)
	if err != nil {
		h.logger.Error("Error adding vocabulary on service")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
