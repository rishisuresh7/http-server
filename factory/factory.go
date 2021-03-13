package factory

import (
	"http-server/book"
	"http-server/models"
)

type Factory interface {
	NewBook(b models.Book) book.Book
}

type factory struct {}

func NewFactory() Factory {
	return &factory{}
}

func (f *factory) NewBook(b models.Book) book.Book {
	return book.NewBook(b)
}
