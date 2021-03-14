package book

import (
	"context"

	"github.com/stretchr/testify/mock"

	"http-server/proto"
)

type Book struct {
	mock.Mock
}

func (b *Book) Get(c context.Context) (*proto.Book, error) {
	ret := b.Called(c)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}

func (b *Book) Create(c context.Context) (*proto.Book, error) {
	ret := b.Called(c)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}
func (b *Book) Update(c context.Context) (*proto.Book, error) {
	ret := b.Called(c)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}

func (b *Book) Delete(c context.Context) (*proto.Book, error) {
	ret := b.Called(c)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}
