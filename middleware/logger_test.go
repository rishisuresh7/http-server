package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

func TestNewLoggerMiddleware(t *testing.T) {
	w := httptest.NewRecorder()
	wr := negroni.NewResponseWriter(w)
	r := httptest.NewRequest("GET", "/foo", nil)

	logger, hook := test.NewNullLogger()
	handler := NewLoggerMiddleware(logger)
	handler.ServeHTTP(wr, r, func(w http.ResponseWriter, r *http.Request) {})

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Request", hook.Entries[0].Message)
	assert.Equal(t, "Response", hook.Entries[1].Message)
}
