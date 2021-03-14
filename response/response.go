package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Success struct to encode server success responses
type Success struct {
	Success interface{} `json:"success"`
}

// Error struct to encode server error responses
type Error struct {
	Error interface{} `json:"error"`
}

// Send method to send a response with status 200
func (s Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(s); err != nil {
		return fmt.Errorf("send: unable to encode interface: %s", err.Error())
	}

	return nil
}

// ClientError method to send a response with status 400
func (e Error) ClientError(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(e); err != nil {
		return fmt.Errorf("clientError: unable to encode interface: %s", err.Error())
	}

	return nil
}

// ServerError method to send a response with status 500
func (e Error) ServerError(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if err := json.NewEncoder(w).Encode(e); err != nil {
		return fmt.Errorf("serverError: unable to encode interface: %s", err.Error())
	}

	return nil
}
