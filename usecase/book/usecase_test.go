package book

import (
	"testing"
	"go_book_api/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBookRepository is a mock type for domain.BookRepository
type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) Create(book *domain.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Fetch() ([]domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m *MockBookRepository) GetByID(id uint) (domain.Book, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Book), args.Error(1)
}

func (m *MockBookRepository) Update(book *domain.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestBookUseCase_Create(t *testing.T) {
	mockRepo := new(MockBookRepository)
	book := &domain.Book{Title: "Test Book", Author: "Test Author", Year: 2023}

	mockRepo.On("Create", book).Return(nil)

	useCase := NewBookUseCase(mockRepo)
	err := useCase.Create(book)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBookUseCase_Fetch(t *testing.T) {
	mockRepo := new(MockBookRepository)
	expectedBooks := []domain.Book{{ID: 1, Title: "Test Book"}}

	mockRepo.On("Fetch").Return(expectedBooks, nil)

	useCase := NewBookUseCase(mockRepo)
	books, err := useCase.Fetch()

	assert.NoError(t, err)
	assert.Equal(t, expectedBooks, books)
	mockRepo.AssertExpectations(t)
}