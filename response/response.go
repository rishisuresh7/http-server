package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Success struct {
	Success interface{} `json:"success"`
}

type Error struct {
	Error interface{} `json:"error"`
}

func (s Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(s); err != nil {
		return fmt.Errorf("send: unable to encode interface: %s", err.Error())
	}

	return nil
}

func (e Error) ClientError(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(e); err != nil {
		return fmt.Errorf("clientError: unable to encode interface: %s", err.Error())
	}

	return nil
}

func (e Error) ServerError(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if err := json.NewEncoder(w).Encode(e); err != nil {
		return fmt.Errorf("serverError: unable to encode interface: %s", err.Error())
	}

	return nil
}
