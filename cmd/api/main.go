package main

import (
	"log"
	"os"

	"api-rest-with-go/internal/handlers"
	"api-rest-with-go/internal/models"
	"api-rest-with-go/internal/repository"
	"api-rest-with-go/internal/service"
	"api-rest-with-go/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize dependencies
	itemRepo := repository.NewItemRepository()
	itemService := service.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	// Setup Gin router
	router := gin.Default()

	// API routes
	v1 := router.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("/", itemHandler.CreateItem)
			items.GET("/", itemHandler.GetAllItems)
			items.GET("/:id", itemHandler.GetItem)
			items.PUT("/:id", itemHandler.UpdateItem)
			items.DELETE("/:id", itemHandler.DeleteItem)
		}
	}

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 