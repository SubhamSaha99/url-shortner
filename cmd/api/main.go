package main

import (
	"fmt"
	"log"

	"github.com/SubhamSaha99/url-shortner/internal/config"
	"github.com/SubhamSaha99/url-shortner/internal/database"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/controllers"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/repositories"
	urlroutes "github.com/SubhamSaha99/url-shortner/internal/modules/url/routes"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadEnv()

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
	urlService := services.NewURLService(urlRepo, cfg.Charset)

	// 3. Inject service into controller.
	urlController := controllers.NewURLController(
		urlService,
		fmt.Sprintf("http://localhost:%s", cfg.Port),
	)

	// 4. Register controller methods as routes.
	router := gin.Default()
	urlroutes.RegisterURLRoutes(router, urlController)

	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server running on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
