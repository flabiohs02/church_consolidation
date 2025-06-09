package main

import (
	"log"
	"os"

	"church_consolidation/config"
	"church_consolidation/domain"
	"church_consolidation/handler"
	"church_consolidation/repository"
	"church_consolidation/routers"
	"church_consolidation/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := config.InitDb()
	db.AutoMigrate(&domain.Consolidation{})

	// Dependency Injection
	consolidationRepo := repository.NewGormConsolidationRepository(db)
	consolidationService := usecase.NewConsolidationService(consolidationRepo)
	consolidationHandler := handler.NewConsolidationHandler(consolidationService)

	// Set up Fiber app
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // You can restrict this to specific origins, e.g., "https://yourfrontend.com"
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
	}))

	// Setup routes
	routers.SetupConsolidationRoutes(app, consolidationHandler)

	port := os.Getenv("CONSOLIDATION_PORT")
	if port == "" {
		port = "8586" // Default port if not set in .env
	}

	log.Printf("Consolidation server listening on :%s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Consolidation server failed to start: %v", err)
	}
}
