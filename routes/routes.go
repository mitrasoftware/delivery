package routes

import (
	"delivery/internal/handler"
	"delivery/internal/handler/orders"
	"delivery/internal/kafka"
	"delivery/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

// API routes configuration

func RegisterRoute(r *gin.Engine) {
	// Initialize Kafka producer in goroutine
	go func() {
		kafka.InitProducer()
		kafka.StartConsumer()

	}()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")

	// Public routes
	api.POST("/", handler.Login)

	// Protected routes with JWT middleware
	protected := api.Group("/protected")
	protected.Use(middleware.Auth())
	protected.GET("/profile", func(c *gin.Context) {
		// This runs asynchronously with logging
		go func() {
			log.Println("📊 Fetching user profile in background")
		}()
		c.JSON(200, gin.H{"message": "profile data"})
	})

	protected.GET("/orders", orders.GetOrders)
	protected.POST("/pick_order", orders.PickOrders)

}
