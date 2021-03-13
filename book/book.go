package book

import (
	"http-server/models"
)

type Book interface {
	Get() (models.Book, error)
	Create() (models.Book, error)
	Update() (models.Book, error)
	Delete() error
}

type book struct {
	book models.Book
}

func NewBook(b models.Book) Book {
	return &book{book: b}
}

func (b *book) Get() (models.Book, error) {
	return models.Book{}, nil
}

func (b *book) Create() (models.Book, error) {
	return models.Book{}, nil
}

func (b *book) Update() (models.Book, error) {
	return models.Book{}, nil
}

func (b *book) Delete() error {
	return nil
}
