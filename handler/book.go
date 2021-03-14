package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"http-server/factory"
	"http-server/proto"
	"http-server/response"
)

// Create function to handle POST requests
func Create(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload proto.Book
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("Create: invalid payload: %s", err.Error())
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if payload.Name == "" && payload.Author == "" {
			l.Errorf("Create: 'name'/'author' is requried")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		book := f.NewBook(&payload)
		res, err := book.Create(r.Context())
		if err != nil {
			l.WithError(err).Errorf("Create: unable to create book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

// Update function to process PATCH requests
func Update(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload proto.Book
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("Update: invalid payload: %s", err.Error())
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if payload.Id == "" {
			l.Errorf("Update: 'id' cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		book := f.NewBook(&payload)
		res, err := book.Update(r.Context())
		if err != nil {
			l.WithError(err).Errorf("Update: unable to update book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

// Delete function to handle a delete request
func Delete(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, exists := vars["id"]
		if !exists {
			l.Errorf("Delete: unable to read 'id' from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		b := proto.Book{Id: id}
		book := f.NewBook(&b)
		res, err := book.Delete(r.Context())
		if err != nil {
			l.WithError(err).Errorf("Delete: unable to delete book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

// Get function to handle get requests
func Get(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, exists := vars["id"]
		if !exists {
			l.Errorf("Get: unable to read 'id' from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		b := proto.Book{Id: id}
		book := f.NewBook(&b)
		res, err := book.Get(r.Context())
		if err != nil {
			l.WithError(err).Errorf("Get: unable to retrieve book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}
