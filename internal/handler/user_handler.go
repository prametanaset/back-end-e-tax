// internal/handler/user_handler.go
package handler

import (
	"strconv"

	"back-end-e-tax/internal/service"
	"back-end-e-tax/pkg/model"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", h.GetAll)
	users.Get("/:id", h.GetByID)
	users.Post("/", h.Create)
	users.Put("/:id", h.Update)
	users.Delete("/:id", h.Delete)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch users"})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var input model.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// Optional: validate fields here

	// ใช้ service ในการสร้าง
	if err := h.service.CreateUser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(input)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid payload"})
	}
	user.ID = uint(id)
	err := h.service.UpdateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update user"})
	}
	return c.JSON(user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.service.DeleteUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete user"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
