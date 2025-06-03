package postgres

import (
	"testing"
	"go_book_api/test"
	"github.com/stretchr/testify/assert"
)

func TestBookRepository_Create(t *testing.T) {
	db, err := test.NewTestDB()
	assert.NoError(t, err)

	repo := NewPostgresBookRepository(db)
	book := test.CreateTestBook()

	err = repo.Create(book)
	assert.NoError(t, err)
	assert.NotZero(t, book.ID)
}

func TestBookRepository_Fetch(t *testing.T) {
	db, err := test.NewTestDB()
	assert.NoError(t, err)

	repo := NewPostgresBookRepository(db)
	book := test.CreateTestBook()
	err = repo.Create(book)
	assert.NoError(t, err)

	books, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, book.Title, books[0].Title)
}

func TestBookRepository_GetByID(t *testing.T) {
	db, err := test.NewTestDB()
	assert.NoError(t, err)

	repo := NewPostgresBookRepository(db)
	book := test.CreateTestBook()
	err = repo.Create(book)
	assert.NoError(t, err)

	fetchedBook, err := repo.GetByID(book.ID)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, fetchedBook.Title)
}

func TestBookRepository_Update(t *testing.T) {
	db, err := test.NewTestDB()
	assert.NoError(t, err)

	repo := NewPostgresBookRepository(db)
	book := test.CreateTestBook()
	err = repo.Create(book)
	assert.NoError(t, err)

	book.Title = "Updated Title"
	err = repo.Update(book)
	assert.NoError(t, err)

	updatedBook, err := repo.GetByID(book.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updatedBook.Title)
}

func TestBookRepository_Delete(t *testing.T) {
	db, err := test.NewTestDB()
	assert.NoError(t, err)

	repo := NewPostgresBookRepository(db)
	book := test.CreateTestBook()
	err = repo.Create(book)
	assert.NoError(t, err)

	err = repo.Delete(book.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(book.ID)
	assert.Error(t, err) // should return error as book is deleted
}