package book

import "go_book_api/domain"

type bookUseCase struct {
	bookRepo domain.BookRepository
}

func NewBookUseCase(repo domain.BookRepository) UseCase {
	return &bookUseCase{bookRepo: repo}
}

func (u *bookUseCase) Create(book *domain.Book) error {
	return u.bookRepo.Create(book)
}

func (u *bookUseCase) Fetch() ([]domain.Book, error) {
	return u.bookRepo.Fetch()
}

func (u *bookUseCase) GetByID(id uint) (domain.Book, error) {
	return u.bookRepo.GetByID(id)
}

func (u *bookUseCase) Update(book *domain.Book) error {
	return u.bookRepo.Update(book)
}

func (u *bookUseCase) Delete(id uint) error {
	return u.bookRepo.Delete(id)
}