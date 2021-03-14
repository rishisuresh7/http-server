package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"http-server/factory"
	"http-server/handler"
)

const (
	GET = "GET"
	POST = "POST"
	PATCH = "PATCH"
	DELETE = "DELETE"
)

// NewRouter contains all the routes and route handlers
func NewRouter(f factory.Factory, l *logrus.Logger) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health()).Methods(GET)

	router.HandleFunc("/book/{id}", handler.Get(f, l)).Methods(GET)
	router.HandleFunc("/book", handler.Create(f, l)).Methods(POST)
	router.HandleFunc("/book", handler.Update(f, l)).Methods(PATCH)
	router.HandleFunc("/book/{id}", handler.Delete(f, l)).Methods(DELETE)

	return router
}
