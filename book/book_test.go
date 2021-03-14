package book

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockProto "http-server/mocks/proto"
	"http-server/proto"
)

func TestGetFails(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Get", mock.Anything, b).Return(nil, fmt.Errorf("some-error"))

	book := NewBook(b, p)
	res, err := book.Get(context.Background())

	assert.Nil(t, res, "expected result to be nil")
	assert.NotNil(t, err, "expected error not to be nil")
	assert.Equal(t, err.Error(), "get: unable to retrieve book: some-error")
	p.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Get", mock.Anything, b).Return(b, nil)

	book := NewBook(b, p)
	res, err := book.Get(context.Background())

	assert.NotNil(t, res, "expected result not to be nil")
	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, b, res)
	p.AssertExpectations(t)
}

func TestCreateFails(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Create", mock.Anything, b).Return(nil, fmt.Errorf("some-error"))

	book := NewBook(b, p)
	res, err := book.Create(context.Background())

	assert.Nil(t, res, "expected result to be nil")
	assert.NotNil(t, err, "expected error not to be nil")
	assert.Equal(t, err.Error(), "create: unable to create book: some-error")
	p.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Create", mock.Anything, b).Return(b, nil)

	book := NewBook(b, p)
	res, err := book.Create(context.Background())

	assert.NotNil(t, res, "expected result not to be nil")
	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, b, res)
	p.AssertExpectations(t)
}

func TestUpdateFails(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Update", mock.Anything, b).Return(nil, fmt.Errorf("some-error"))

	book := NewBook(b, p)
	res, err := book.Update(context.Background())

	assert.Nil(t, res, "expected result to be nil")
	assert.NotNil(t, err, "expected error not to be nil")
	assert.Equal(t, err.Error(), "update: unable to update book: some-error")
	p.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Update", mock.Anything, b).Return(b, nil)

	book := NewBook(b, p)
	res, err := book.Update(context.Background())

	assert.NotNil(t, res, "expected result not to be nil")
	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, b, res)
	p.AssertExpectations(t)
}

func TestDeleteFails(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Delete", mock.Anything, b).Return(nil, fmt.Errorf("some-error"))

	book := NewBook(b, p)
	res, err := book.Delete(context.Background())

	assert.Nil(t, res, "expected result to be nil")
	assert.NotNil(t, err, "expected error not to be nil")
	assert.Equal(t, err.Error(), "delete: unable to delete book: some-error")
	p.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	p := &mockProto.BookServiceClient{}
	b := &proto.Book{}
	p.On("Delete", mock.Anything, b).Return(b, nil)

	book := NewBook(b, p)
	res, err := book.Delete(context.Background())

	assert.NotNil(t, res, "expected result not to be nil")
	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, b, res)
	p.AssertExpectations(t)
}
