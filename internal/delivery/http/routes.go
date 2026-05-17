package http

import (
	"simpleorder/internal/delivery/http/handlers"
	"simpleorder/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App,
	authHandler *handlers.AuthHandler,
	customerHandler *handlers.CustomerHandler,
	productHandler *handlers.ProductHandler,
	orderHandler *handlers.OrderHandler,
	dashboardHandler *handlers.DashboardHandler,
	jwtSecret string) {

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Public routes
	auth := v1.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes
	v1.Use(middleware.Protected(jwtSecret))

	customers := v1.Group("/customers")
	customers.Post("/", customerHandler.Create)
	customers.Get("/", customerHandler.FindAll)

	products := v1.Group("/products")
	products.Post("/", middleware.RoleAdmin(), productHandler.Create)
	products.Get("/", productHandler.FindAll)

	orders := v1.Group("/orders")
	orders.Post("/", orderHandler.Create)
	orders.Get("/", orderHandler.FindAll)
	orders.Get("/:id/invoice", orderHandler.GenerateInvoice)

	dashboard := v1.Group("/dashboard")
	dashboard.Get("/stats", dashboardHandler.GetStats)
}
