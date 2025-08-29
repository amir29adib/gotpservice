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
	"fmt"
	"log"
	"os"

	"gotpservice/internal/handler"
	"gotpservice/internal/middleware"
	"gotpservice/internal/model"
	"gotpservice/internal/repository"
	"gotpservice/internal/service"

	docs "gotpservice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    db := setupDB()

    // Auto migrate tables
    if err := db.AutoMigrate(&model.User{}); err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    // Repositories
    otpRepo := repository.NewOTPRepository()
    userRepo := repository.NewUserRepository(db)

    // Services
    otpService := service.NewOTPService(otpRepo, userRepo)
    userService := service.NewUserService(userRepo)

    // Handlers
    otpHandler := handler.NewOTPHandler(otpService)
    userHandler := handler.NewUserHandler(userService)

    // Router
    r := gin.Default()

	// Logger Middleware
	r.Use(gin.Recovery(), middleware.RequestLogger())
    
    docs.SwaggerInfo.BasePath = "/"                     
	
    // Swagger Route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
    // Routes
    auth := r.Group("/auth")
    {
        auth.POST("/request-otp", otpHandler.RequestOTP)
        auth.POST("/verify-otp", otpHandler.VerifyOTP)
    }

    user := r.Group("/users")
    user.Use(middleware.JWTAuth())
    {
        user.GET("/", userHandler.ListUsers)
        user.GET("/:phone", userHandler.GetUserByPhone)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}

func setupDB() *gorm.DB {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        getEnv("DB_HOST", "db"),
        getEnv("DB_PORT", "5432"),
        getEnv("DB_USER", "otpuser"),
        getEnv("DB_PASSWORD", "otppass"),
        getEnv("DB_NAME", "otpdb"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    return db
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}
