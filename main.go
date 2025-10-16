package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configure standard logger to write to stdout
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Initialize Gin with no default logger
	r := gin.New()

	// Middleware to log HTTP requests
	r.Use(func(c *gin.Context) {
		c.Next() // process request

		log.Printf("HTTP request | method=%s path=%s status=%d client_ip=%s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			c.ClientIP(),
		)
	})

	// Basic endpoint
	r.GET("/", func(c *gin.Context) {
		log.Println("Hello world endpoint hit")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Start server
	port := "8081"
	log.Println("Starting server on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Println("Server failed:", err)
		os.Exit(1)
	}
}
