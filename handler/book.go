package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"http-server/factory"
	"http-server/models"
	"http-server/response"
)

func Create(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Book
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

		book := f.NewBook(payload)
		res, err := book.Create()
		if err != nil {
			l.WithError(err).Errorf("Create: unable to create book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func Update(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Book
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

		book := f.NewBook(payload)
		res, err := book.Update()
		if err != nil {
			l.WithError(err).Errorf("Update: unable to update book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func Delete(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, exists := vars["id"]
		if !exists {
			l.Errorf("Delete: unable to read 'id' from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		b := models.Book{Id: id}
		book := f.NewBook(b)
		err := book.Delete()
		if err != nil {
			l.WithError(err).Errorf("Delete: unable to delete book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: "book deleted successfully"}.Send(w)
	}
}

func Get(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, exists := vars["id"]
		if !exists {
			l.Errorf("Get: unable to read 'id' from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		b := models.Book{Id: id}
		book := f.NewBook(b)
		res, err := book.Get()
		if err != nil {
			l.WithError(err).Errorf("Get: unable to retrieve book")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}