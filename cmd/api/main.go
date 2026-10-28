package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/SubhamSaha99/url-shortner/internal/database"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/controllers"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/repositories"
	urlroutes "github.com/SubhamSaha99/url-shortner/internal/modules/url/routes"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/services"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// 1. Create repository with the database.
	urlRepo := repositories.NewURLRepository(db)

	// 2. Inject repository into service.
	urlService := services.NewURLService(urlRepo)

	// 3. Inject service into controller.
	urlController := controllers.NewURLController(
		urlService,
		"http://localhost:8080",
	)

	// 4. Register controller methods as routes.
	router := gin.Default()
	urlroutes.RegisterURLRoutes(router, urlController)

	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}