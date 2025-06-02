package routes

import (
	"time"

	"api-rest-with-go/internal/core/ports"
	"api-rest-with-go/internal/infrastructure/server/http/handlers"
	"api-rest-with-go/internal/infrastructure/server/http/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Setup configura todas las rutas de la aplicación
func Setup(router *gin.Engine, services *ports.Service, db *gorm.DB) {
	// Middleware global
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.TimeoutMiddleware(10 * time.Second))
	router.Use(middleware.RateLimiter(100))

	// Handlers
	itemHandler := handlers.NewItemHandler(services.Items)
	healthHandler := handlers.NewHealthHandler(db)

	// Health check endpoints
	router.GET("/health", healthHandler.CheckHealth)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// API routes
	v1 := router.Group("/api/v2")
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
}
