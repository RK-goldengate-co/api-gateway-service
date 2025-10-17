package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
	Version   string `json:"version"`
}

type RootResponse struct {
	Service   string              `json:"service"`
	Version   string              `json:"version"`
	Endpoints map[string]string   `json:"endpoints"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, RootResponse{
			Service: "API Gateway Service",
			Version: "1.0.0",
			Endpoints: map[string]string{
				"health": "/health",
				"proxy":  "/api/proxy?url=<target_url>",
			},
		})
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{
			Status:    "healthy",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Service:   "api-gateway",
			Version:   "1.0.0",
		})
	})

	// Reverse proxy endpoint
	router.GET("/api/proxy", func(c *gin.Context) {
		targetURL := c.Query("url")

		if targetURL == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Missing url parameter",
				Message: "Please provide a target URL via ?url=<target_url>",
			})
			return
		}

		// Make HTTP request to target URL
		resp, err := http.Get(targetURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Proxy request failed",
				Message: err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Failed to read response",
				Message: err.Error(),
			})
			return
		}

		// Forward the response
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	})

	// Start server
	port := ":8080"
	fmt.Println("🚀 API Gateway running on http://localhost" + port)
	fmt.Println("📊 Health check: http://localhost" + port + "/health")
	fmt.Println("🔄 Proxy endpoint: http://localhost" + port + "/api/proxy?url=<target_url>")

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
