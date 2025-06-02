package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go_book_api/delivery/http/handler"
	postgres "go_book_api/repository/postgres"
	"go_book_api/usecase/auth"
	"go_book_api/usecase/book"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository
	bookRepo := postgres.NewPostgresBookRepository(db)

	// Initialize usecases
	bookUseCase := book.NewBookUseCase(bookRepo)
	authUseCase := auth.NewAuthUseCase()

	// Initialize Gin router
	router := gin.Default()

	// Initialize handlers
	handler.NewAuthHandler(router, authUseCase)
	handler.NewBookHandler(router, bookUseCase, authUseCase)

	// Start server
	router.Run(":8080")
}
