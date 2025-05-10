package handler

import (
	"back-end-e-tax/internal/service"
	"back-end-e-tax/pkg/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// 🔹 Utility
func getTenantID(c *fiber.Ctx) uint {
	id, _ := strconv.ParseUint(c.Get("X-Tenant-ID"), 10, 32)
	return uint(id)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	tenantID := getTenantID(c)

	products, total, err := h.service.GetPaginatedProducts(page, limit, search, tenantID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching products"})
	}

	return c.JSON(fiber.Map{
		"data":  products,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	tenantID := getTenantID(c)

	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := h.service.CreateProduct(&product, tenantID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating product",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}
	tenantID := getTenantID(c)

	var input model.Product
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}
	updated, err := h.service.UpdateProduct(id, &input, tenantID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Product updated", "product": updated})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}
	tenantID := getTenantID(c)

	if err := h.service.DeleteProduct(id, tenantID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
