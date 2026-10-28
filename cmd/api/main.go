package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/SubhamSaha99/url-shortner/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Starting URL Shortener...")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database handle: %v", err)
	}
	defer sqlDB.Close()

	log.Println("Database connected successfully!")

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	log.Println("Server starting on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
