package postgres

import (
	"go_book_api/domain"
	"gorm.io/gorm"
)

type postgresBookRepository struct {
	db *gorm.DB
}

func NewPostgresBookRepository(db *gorm.DB) domain.BookRepository {
	return &postgresBookRepository{db}
}

func (r *postgresBookRepository) Create(book *domain.Book) error {
	return r.db.Create(book).Error
}

func (r *postgresBookRepository) Fetch() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *postgresBookRepository) GetByID(id uint) (domain.Book, error) {
	var book domain.Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *postgresBookRepository) Update(book *domain.Book) error {
	return r.db.Save(book).Error
}

func (r *postgresBookRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Book{}, id).Error
}