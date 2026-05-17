package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"simpleorder/internal/config"
	"simpleorder/internal/delivery/http"
	"simpleorder/internal/delivery/http/handlers"
	"simpleorder/internal/delivery/http/middleware"
	"simpleorder/internal/infrastructure/database"
	"simpleorder/internal/repository"
	"simpleorder/internal/usecase"

	_ "simpleorder/docs" // Swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title SimpleOrder API
// @version 1.0
// @description This is a sample production-ready backend project for SimpleOrder.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@simpleorder.local

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Init DB
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, customerRepo)
	dashboardUsecase := usecase.NewDashboardUsecase(orderRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authUsecase, cfg)
	customerHandler := handlers.NewCustomerHandler(customerUsecase)
	productHandler := handlers.NewProductHandler(productUsecase)
	orderHandler := handlers.NewOrderHandler(orderUsecase)
	dashboardHandler := handlers.NewDashboardHandler(dashboardUsecase)

	// App Initialization
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	// Routes
	http.SetupRoutes(app, authHandler, customerHandler, productHandler, orderHandler, dashboardHandler, cfg.JWTSecret)

	// Graceful Shutdown
	go func() {
		if err := app.Listen(":" + cfg.AppPort); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}
