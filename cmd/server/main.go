package main

// @title           OTP Service API
// @version         0.1.0
// @description     Gin-based skeleton for OTP auth service.
// @contact.name    Backend Team
// @contact.email   dev@example.com
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	_ "gotpservice/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
	"gotpservice/internal/handlers"
	"gotpservice/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Recovery(), middleware.RequestLogger())

	r.GET("/health", handlers.Health)
	r.GET("/version", handlers.Version)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	log.Printf("server listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
