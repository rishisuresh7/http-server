package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"http-server/handler"
)

func NewRouter(_ interface{}, l *logrus.Logger) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health()).Methods("GET")

	return router
}
