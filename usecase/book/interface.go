package book

import "go_book_api/domain"

type UseCase interface {
	Create(book *domain.Book) error
	Fetch() ([]domain.Book, error)
	GetByID(id uint) (domain.Book, error)
	Update(book *domain.Book) error
	Delete(id uint) error
}