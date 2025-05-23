/*
Package api provides RESTful handlers and middleware for managing a library books.
It includes functionality for creating, reading, updating, and deleteing books, as well as generating JWT for authentication.
*/
package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection using environment variables.
// It loads the database configuration from a .env file and migrates the Book schema.
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// migrate the schema
	if err := DB.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}


// CreateBook handles the creation of a new book in the database.
// It expects a JSON payload with the book details and responds with the created book.
func CreateBook(c *gin.Context) {
	var book Book

	// bind the request body
	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	DB.Create(&book)
	ResponseJSON(c, http.StatusCreated, "Book created successfully", book)
}

// GetBooks retrieves all books from the database.
// It responds with a list of books.
func GetBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)
	ResponseJSON(c, http.StatusOK, "Books retrieved successfully", books)
}

func GetBook(c *gin.Context) {
	var book Book
	if err := DB.First(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Book retrieved successfully", book)
}

func UpdateBook(c *gin.Context) {
	var book Book
	if err := DB.First(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// bind the request body
	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	DB.Save(&book)
	ResponseJSON(c, http.StatusOK, "Book updated successfully", book)
}

func DeleteBook(c *gin.Context) {
	var book Book
	if err := DB.Delete(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Book deleted successfully", nil)
}

func GenerateJWT(c *gin.Context) {
var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if loginRequest.Username != "admin" || loginRequest.Password != "password" {
		ResponseJSON(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(jwtSecret)
	if err!= nil {
		ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}

	ResponseJSON(c, http.StatusOK, "Token generated successfully", gin.H{"token": tokenString})
}
