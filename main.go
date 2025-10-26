package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create Gin router
	router := gin.New()
	
	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "goldcore-api-poc",
		})
	})

	// Ping endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "pong",
			"service":   "goldcore-api-poc", 
			"timestamp": time.Now().Unix(),
		})
	})

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Goldcore API POC is running!",
			"service": "goldcore-api-poc",
			"endpoints": []string{"/health", "/ping"},
		})
	})

	// Start server
	address := fmt.Sprintf(":%s", port)
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(address, router))
}