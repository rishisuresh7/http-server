package factory

import (
	"github.com/stretchr/testify/mock"

	"http-server/book"
	"http-server/proto"
)

type Factory struct {
	mock.Mock
}

func (f *Factory) NewBook(b *proto.Book) book.Book {
	ret := f.Called(b)

	var res book.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(book.Book)
	}

	return res
}

func (f *Factory)  NewGRPCClient() proto.BookServiceClient {
	ret := f.Called()

	var res proto.BookServiceClient
	if ret.Get(0) != nil {
		res = ret.Get(0).(proto.BookServiceClient)
	}

	return res
}
