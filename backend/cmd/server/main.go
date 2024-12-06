package main

import (
	"log"
	"os"
	"path/filepath"

	"guinx/internal/config"
	"guinx/internal/database"
	"guinx/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// CORS configuration (if needed)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		authConfig := config.NewAuthConfig()
		authHandler := handlers.NewAuthHandler(db, authConfig)

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// Add other API routes here
	}

	// Serve frontend static files
	r.Static("/assets", "./frontend/build/assets")

	// Serve index.html for all other routes
	r.NoRoute(func(c *gin.Context) {
		indexPath := filepath.Join(".", "frontend", "build", "index.html")
		if _, err := os.Stat(indexPath); os.IsNotExist(err) {
			c.String(404, "Frontend not built")
			return
		}
		c.File(indexPath)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
