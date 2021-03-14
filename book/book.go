package book

import (
	"context"
	"fmt"

	"http-server/proto"
)

type Book interface {
	Get(c context.Context) (*proto.Book, error)
	Create(c context.Context) (*proto.Book, error)
	Update(c context.Context) (*proto.Book, error)
	Delete(c context.Context) (*proto.Book, error)
}

type book struct {
	book *proto.Book
	grpc proto.BookServiceClient
}

func NewBook(b *proto.Book, g proto.BookServiceClient) Book {
	return &book{book: b, grpc: g}
}

func (b *book) Get(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Get(c, b.book)
	if err != nil {
		return nil, fmt.Errorf("get: unable to retrieve book: %s", err.Error())
	}

	return res, nil
}

func (b *book) Create(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Create(c, b.book)
	if err != nil {
		return nil, fmt.Errorf("create: unable to create book: %s", err.Error())
	}

	return res, nil
}

func (b *book) Update(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Update(c, b.book)
	if err != nil {
		return res, fmt.Errorf("update: unable to update book: %s", err.Error())
	}

	return res, nil
}

func (b *book) Delete(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Delete(c, b.book)
	if err != nil {
		return nil, fmt.Errorf("delete: unable to delete book: %s", err.Error())
	}

	return res, nil
}
