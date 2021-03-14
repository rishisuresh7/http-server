package response

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	tests := []struct {
		setInvalidMessage bool
		expectedErr       error
	}{
		{
			setInvalidMessage: false,
			expectedErr:       nil,
		},
		{
			setInvalidMessage: true,
			expectedErr:       fmt.Errorf("send: unable to encode interface: json: unsupported type: chan int"),
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		s := Success{Success: "test"}
		if tt.setInvalidMessage {
			s.Success = make(chan int)
		}

		err := s.Send(w)
		assert.Equal(t, tt.expectedErr, err)
		assert.Equal(t, "Application/json", w.Header().Get("Content-Type"))
	}
}

func TestClientError(t *testing.T) {
	tests := []struct {
		setInvalidMessage bool
		expectedErr       error
	}{
		{
			setInvalidMessage: false,
			expectedErr:       nil,
		},
		{
			setInvalidMessage: true,
			expectedErr:       fmt.Errorf("clientError: unable to encode interface: json: unsupported type: chan int"),
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		e := Error{Error: "test"}
		if tt.setInvalidMessage {
			e.Error = make(chan int)
		}

		err := e.ClientError(w)
		assert.Equal(t, tt.expectedErr, err)
		assert.Equal(t, "Application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	}
}

func TestServerError(t *testing.T) {
	tests := []struct {
		setInvalidMessage bool
		expectedErr       error
	}{
		{
			setInvalidMessage: false,
			expectedErr:       nil,
		},
		{
			setInvalidMessage: true,
			expectedErr:       fmt.Errorf("serverError: unable to encode interface: json: unsupported type: chan int"),
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		e := Error{Error: "test"}
		if tt.setInvalidMessage {
			e.Error = make(chan int)
		}

		err := e.ServerError(w)
		assert.Equal(t, tt.expectedErr, err)
		assert.Equal(t, "Application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	}
}
