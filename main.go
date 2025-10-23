package main

import (
	"gin-gorm-swagger-api/database"
	"gin-gorm-swagger-api/routes"
	"gin-gorm-swagger-api/seed"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "gin-gorm-swagger-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin GORM PostgreSQL API with JWT
// @version 1.0
// @description API example with Gin, GORM, PostgreSQL, JWT Auth, and Swagger
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.ConnectDatabase()
	seed.SeedUsers()

	r := gin.Default()
	routes.RegisterAuthRoutes(r)
	routes.RegisterUserRoutes(r)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
