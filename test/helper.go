package test

import (
	"go_book_api/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewTestDB creates a new in-memory SQLite database for testing
func NewTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&domain.Book{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTestBook creates a test book record
func CreateTestBook() *domain.Book {
	return &domain.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2023,
	}
}