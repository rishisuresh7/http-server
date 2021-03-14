package factory

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"http-server/book"
	"http-server/config"
	"http-server/proto"
)

var once sync.Once

// Factory interface for object creation
type Factory interface {
	NewBook(b *proto.Book) book.Book
	NewGRPCClient() proto.BookServiceClient
}

// factory implements Factory
type factory struct {
	conf   *config.AppConfig
	logger *logrus.Logger
	client proto.BookServiceClient
}

// NewFactory constructor to initialize factory
func NewFactory(c *config.AppConfig, l *logrus.Logger) Factory {
	return &factory{conf: c, logger: l}
}

// NewBook object of book.Book interface
func (f *factory) NewBook(b *proto.Book) book.Book {
	return book.NewBook(b, f.NewGRPCClient())
}

// NewGRPCClient object of proto.BookServiceClient(grpc server)
func (f *factory) NewGRPCClient() proto.BookServiceClient {
	var err error
	once.Do(func() {
		conn, connErr := grpc.Dial(f.conf.GRPCUri, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithAuthority(f.conf.Token))
		if connErr != nil {
			err = fmt.Errorf("unable to establish connection")
			return
		}

		f.client = proto.NewBookServiceClient(conn)
	})

	if err != nil {
		 f.logger.Fatalf("NewGRPCClient: failed to connect to GRPC server: %s", err.Error())
	}

	return f.client
}
