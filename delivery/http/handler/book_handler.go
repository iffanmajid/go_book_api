package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_book_api/domain"
	"go_book_api/delivery/http/middleware"
	"go_book_api/usecase/book"
	"go_book_api/usecase/auth"
)

type BookHandler struct {
	bookUseCase book.UseCase
}

func NewBookHandler(router *gin.Engine, bookUseCase book.UseCase, authUseCase auth.UseCase) {
	handler := &BookHandler{bookUseCase: bookUseCase}

	// Protected routes
	protected := router.Group("/", middleware.JWTAuthMiddleware(authUseCase))
	{
		protected.POST("/book", handler.Create)
		protected.GET("/books", handler.Fetch)
		protected.GET("/books/:id", handler.GetByID)
		protected.PUT("/books/:id", handler.Update)
		protected.DELETE("/books/:id", handler.Delete)
	}
}

func (h *BookHandler) Create(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.bookUseCase.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) Fetch(c *gin.Context) {
	books, err := h.bookUseCase.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := h.bookUseCase.GetByID(uint(id))
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	var book domain.Book
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book.ID = uint(id)

	if err := h.bookUseCase.Update(&book); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.bookUseCase.Delete(uint(id)); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}


