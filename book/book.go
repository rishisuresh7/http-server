package book

import (
	"context"
	"fmt"

	"http-server/models"
	"http-server/proto"
)

type Book interface {
	Get(c context.Context) (models.Book, error)
	Create(c context.Context) (models.Book, error)
	Update(c context.Context) (models.Book, error)
	Delete(c context.Context) (models.Book, error)
}

type book struct {
	book models.Book
	grpc proto.BookServiceClient
}

func NewBook(b models.Book, g proto.BookServiceClient) Book {
	return &book{book: b, grpc: g}
}

func (b *book) Get(c context.Context) (models.Book, error) {
	res, err := b.grpc.Get(c, &proto.Book{Id: b.book.Id})
	if err != nil {
		return models.Book{}, fmt.Errorf("get: unable to retrieve book: %s", err.Error())
	}

	return models.Book{Id: res.Id, Name: res.Name, Author: res.Author}, nil
}

func (b *book) Create(c context.Context) (models.Book, error) {
	res, err := b.grpc.Create(c, &proto.Book{Id: b.book.Id, Name: b.book.Id, Author: b.book.Author})
	if err != nil {
		return models.Book{}, fmt.Errorf("create: unable to create book: %s", err.Error())
	}

	return models.Book{Id: res.Id, Name: res.Name, Author: res.Author}, nil
}

func (b *book) Update(c context.Context) (models.Book, error) {
	res, err := b.grpc.Update(c, &proto.Book{Id: b.book.Id, Name: b.book.Id, Author: b.book.Author})
	if err != nil {
		return models.Book{}, fmt.Errorf("update: unable to update book: %s", err.Error())
	}

	return models.Book{Id: res.Id, Name: res.Name, Author: res.Author}, nil
}

func (b *book) Delete(c context.Context) (models.Book, error) {
	res, err := b.grpc.Delete(c, &proto.Book{Id: b.book.Id})
	if err != nil {
		return models.Book{}, fmt.Errorf("delete: unable to delete book: %s", err.Error())
	}

	return models.Book{Id: res.Id, Name: res.Name, Author: res.Author}, nil
}
