package routes

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/handler"
	"back-end-e-tax/internal/repository"
	"back-end-e-tax/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Product setup
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// User setup
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	tenantRepo := repository.NewTenantRepository(config.DB)
	addressRepo := repository.NewAddressRepository(config.DB)

	merchantSvc := service.NewMerchantService(addressRepo, tenantRepo, userRepo, config.DB)
	merchantHandler := handler.NewMerchantHandler(merchantSvc)

	v1 := app.Group("/v1")

	// Product routes
	v1.Get("/products", productHandler.GetProducts)
	v1.Post("/products", productHandler.CreateProduct)
	v1.Put("/products/:id", productHandler.UpdateProduct)
	v1.Delete("/products/:id", productHandler.DeleteProduct)

	// User routes
	v1.Get("/users", userHandler.GetAll)
	v1.Get("/users/:id", userHandler.GetByID)
	v1.Post("/users", userHandler.Create)
	v1.Put("/users/:id", userHandler.Update)
	v1.Delete("/users/:id", userHandler.Delete)
	v1.Post("/register", merchantHandler.Register)
}
