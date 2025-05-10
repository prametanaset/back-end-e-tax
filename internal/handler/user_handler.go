// internal/handler/user_handler.go
package handler

import (
	"strconv"

	"back-end-e-tax/internal/dto"
	"back-end-e-tax/internal/service"
	"back-end-e-tax/pkg/model"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

// ลงทะเบียนเส้นทาง
func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	// public
	app.Post("/register", h.Register)

	// protected (ต้องมี X-Tenant-ID หรือ JWT แล้วแต่ระบบ)
	users := app.Group("/users")
	users.Get("/", h.GetAll)
	users.Get("/:id", h.GetByID)
	users.Post("/", h.Create)      // สร้าง user ภายใน tenant (admin)
	users.Put("/:id", h.Update)    // แก้ไข user
	users.Delete("/:id", h.Delete) // ลบ user
}

// GET /users
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	tenantID := getTenantID(c)
	users, err := h.svc.GetAll(tenantID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to fetch users"})
	}
	return c.JSON(users)
}

// GET /users/:id
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	tenantID := getTenantID(c)
	user, err := h.svc.GetByID(uint(id), tenantID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

// POST /users        (admin สร้าง user ภายใน tenant)
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var in dto.CreateUserDTO
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid input"})
	}
	if in.Username == "" || in.Password == "" || in.Email == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "username, password, email are required"})
	}
	tenantID := getTenantID(c)
	user := &model.User{
		TenantID: tenantID,
		Username: in.Username,
		Password: in.Password, // จะถูก hash ใน service
		FullName: in.FullName,
		Email:    in.Email,
	}
	if err := h.svc.Register(user); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// PUT /users/:id
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var in dto.UpdateUserDTO
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid payload"})
	}
	tenantID := getTenantID(c)
	user := &model.User{
		ID:       uint(id),
		TenantID: tenantID,
		Username: in.Username,
		FullName: in.FullName,
		Email:    in.Email,
		// Password: in.Password, // ถ้าเปลี่ยน password ให้ส่งมา, ถ้าไม่ส่ง ให้เว้นเป็น ""
	}
	if in.Password != "" {
		user.Password = in.Password
	}
	if err := h.svc.Update(user); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

// DELETE /users/:id
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	tenantID := getTenantID(c)
	if err := h.svc.Delete(uint(id), tenantID); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// POST /register     (public สมัคร user ครั้งแรก)
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var in dto.RegisterRequest
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid input"})
	}
	if in.Username == "" || in.Password == "" || in.Email == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "username, password and email are required"})
	}
	// ฝั่ง register จะต้องมีการสร้าง tenant ขึ้นมาก่อน (ดู service.Register เลย)
	tenantID := h.svc.Register(&model.User{
		Username: in.Username,
		Password: in.Password,
		FullName: in.FullName,
		Email:    in.Email,
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "user registered successfully",
		"tenant_id": tenantID,
	})
}
