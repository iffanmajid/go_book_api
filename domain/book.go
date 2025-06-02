package domain

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookRepository interface {
	Create(book *Book) error
	Fetch() ([]Book, error)
	GetByID(id uint) (Book, error)
	Update(book *Book) error
	Delete(id uint) error
}