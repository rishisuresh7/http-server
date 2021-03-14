package book

import (
	"context"
	"fmt"

	"http-server/proto"
)

// Book interface definition for all CRUD operations
type Book interface {
	Get(c context.Context) (*proto.Book, error)
	Create(c context.Context) (*proto.Book, error)
	Update(c context.Context) (*proto.Book, error)
	Delete(c context.Context) (*proto.Book, error)
}

// book implements Book
type book struct {
	book *proto.Book
	grpc proto.BookServiceClient
}

// NewBook constructor for Book interface
func NewBook(b *proto.Book, g proto.BookServiceClient) Book {
	return &book{book: b, grpc: g}
}

// Get is used to retrieve data from grpc server
func (b *book) Get(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Get(c, b.book)
	if err != nil {
		return nil, fmt.Errorf("get: unable to retrieve book: %s", err.Error())
	}

	return res, nil
}

// Create is used to post data onto grpc server
func (b *book) Create(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Create(c, b.book)
	if err != nil {
		return nil, fmt.Errorf("create: unable to create book: %s", err.Error())
	}

	return res, nil
}

// Update is used to patch data onto grpc server
func (b *book) Update(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Update(c, b.book)
	if err != nil {
		return res, fmt.Errorf("update: unable to update book: %s", err.Error())
	}

	return res, nil
}

// Delete is used to delete a resource from grpc server
func (b *book) Delete(c context.Context) (*proto.Book, error) {
	res, err := b.grpc.Delete(c, b.book)
	if err != nil {
		return nil, fmt.Errorf("delete: unable to delete book: %s", err.Error())
	}

	return res, nil
}
