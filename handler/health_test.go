package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"net/http/httptest"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)

	handler := Health()
	handler.ServeHTTP(w, r)

	assert.Equal(t, "Iam alive", w.Body.String())
	assert.Equal(t, 200, w.Code)
}
