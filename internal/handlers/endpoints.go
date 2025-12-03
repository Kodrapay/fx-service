package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/fx-service/internal/dto"
	"github.com/kodra-pay/fx-service/internal/services"
)

type FXHandler struct {
	svc *services.FXService
}

func NewFXHandler(svc *services.FXService) *FXHandler { return &FXHandler{svc: svc} }

func (h *FXHandler) Rates(c *fiber.Ctx) error {
	return c.JSON(h.svc.Rates(c.Context()))
}

func (h *FXHandler) Convert(c *fiber.Ctx) error {
	var req dto.FXConvertRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	return c.JSON(h.svc.Convert(c.Context(), req))
}
