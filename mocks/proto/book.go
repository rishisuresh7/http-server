package proto

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"http-server/proto"
)

type BookServiceClient struct {
	mock.Mock
}

func (b *BookServiceClient) Create(ctx context.Context, in *proto.Book, _ ...grpc.CallOption) (*proto.Book, error) {
	ret := b.Called(ctx, in)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}

func (b *BookServiceClient) Update(ctx context.Context, in *proto.Book, _ ...grpc.CallOption) (*proto.Book, error) {
	ret := b.Called(ctx, in)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}

func (b *BookServiceClient) Delete(ctx context.Context, in *proto.Book, _ ...grpc.CallOption) (*proto.Book, error) {
	ret := b.Called(ctx, in)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}

func (b *BookServiceClient) Get(ctx context.Context, in *proto.Book, _ ...grpc.CallOption) (*proto.Book, error) {
	ret := b.Called(ctx, in)

	var res *proto.Book
	if ret.Get(0) != nil {
		res = ret.Get(0).(*proto.Book)
	}

	return res, ret.Error(1)
}
