package tests

import (
	"bytes"
	"encoding/json"
	"go_book_api/api"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("SECRET_TOKEN"))

func setupTestDB() {
	var err error
	api.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	api.DB.AutoMigrate(&api.Book{})
}

func addBook() api.Book {
	book := api.Book{Title: "Go Programming", Author: "John Doe", Year: 2023}
	api.DB.Create(&book)
	return book
}

func generateValidToken() string {
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func TestGenerateJWT(t *testing.T) {
	router := gin.Default()
	router.POST("/token", api.GenerateJWT)
	
	loginRequest := map[string]string{
		"username": "admin",
		"password": "password",
	}

	jsonValue, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(jsonValue))
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status!= http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["token"] == "" {
		t.Error("Expected token in response, got nil or empty")
	}
}

func TestCreateBook(t *testing.T) {
	setupTestDB()
	router := gin.Default()
	protected := router.Group("/", api.JWTAuthMiddleware())
	protected.POST("/book", api.CreateBook)

	token := generateValidToken()

	book := api.Book{Title: "Demo Book name", Author: "Demo Author name", Year: 2021}

	jsonValue, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil {
		t.Error("Expected book data, got nil")
	}
}

func TestGetBooks(t *testing.T) {
	setupTestDB()
	addBook()
	router := gin.Default()
	router.GET("/books", api.GetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if len(response.Data.([]interface{})) == 0 {
		t.Errorf("Expected non-empty books list")
	}
}

func TestGetBook(t *testing.T) {
	setupTestDB()
	book := addBook()
	router := gin.Default()
	router.GET("/books/:id", api.GetBook)

	req, _ := http.NewRequest("GET", "/books/"+strconv.Itoa(int(book.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["id"] != float64(book.ID) {
		t.Errorf("Expected book ID %d, got nil or wrong ID", book.ID)
	}
}

func TestUpdateBook(t *testing.T) {
	setupTestDB()
	book := addBook()
	router := gin.Default()
	router.PUT("/books/:id", api.UpdateBook)

	updatedBook := api.Book{Title: "Advance Go Programming", Author: "Demo Author name", Year: 2021}
	jsonValue, _ := json.Marshal(updatedBook)

	req, _ := http.NewRequest("PUT", "/books/"+strconv.Itoa(int(book.ID)), bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["title"] != "Advance Go Programming" {
		t.Errorf("Expected updated book title 'Advance Go Programming', got %v", response.Data)

	}
}

func TestDeleteBook(t *testing.T) {
	setupTestDB()
	book := addBook()
	router := gin.Default()
	router.DELETE("/books/:id", api.DeleteBook)

	req, _ := http.NewRequest("DELETE", "/books/"+strconv.Itoa(int(book.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Message != "Book deleted successfully" {
		t.Errorf("Expected message 'Book deleted successfully', got %v", response.Message)
	}

	// verify that the book is deleted
	var deletedBook api.Book
	result := api.DB.First(&deletedBook, book.ID)
	if result.Error == nil {
		t.Errorf("Expected book to be deleted, but it still exists")
	}
}
