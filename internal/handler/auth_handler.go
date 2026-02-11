package handler

import (
	"delivery/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authentication HTTP handlers

func Login(c *gin.Context) {
	var req service.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := service.Login(req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Publish login event asynchronously in goroutine
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("❌ Error publishing event: %v", r)
			}
		}()
		// kafka.P(resp.UserID)
	}()

	c.JSON(http.StatusOK, resp)
}
