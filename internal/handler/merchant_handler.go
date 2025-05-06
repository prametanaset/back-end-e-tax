// pkg/handler/merchant_handler.go
package handler

import (
	"back-end-e-tax/internal/dto"
	"back-end-e-tax/internal/service"

	"github.com/gofiber/fiber/v2"
)

type MerchantHandler struct {
	svc service.MerchantService
}

func NewMerchantHandler(s service.MerchantService) *MerchantHandler {
	return &MerchantHandler{svc: s}
}

func (h *MerchantHandler) Register(c *fiber.Ctx) error {
	var in dto.MerchantRegisterDTO
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	// basic required-field check
	if in.LegalName == "" ||
		in.Address.Street == "" ||
		in.Username == "" ||
		in.Password == "" ||
		in.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing required fields"})
	}

	tenantID, err := h.svc.RegisterMerchant(
		in.LegalName, in.BranchCode, in.TaxID, in.Address, in.LogoURL,
		in.Username, in.Password, in.FullName, in.Email,
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "registration successful",
		"tenant_id": tenantID,
	})
}
