package router

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	router := NewRouter(nil, nil)

	var paths []string
	var methods []string
	walk := func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		require.NoError(t, err, "Expected no error getting path")
		require.Nil(t, err, "Expected err to be nil")

		method, err := route.GetMethods()
		require.NoError(t, err, "Expected no error getting methods")
		require.Nil(t, err, "Expected err to be nil")

		paths = append(paths, path)
		methods = append(methods, method[0])
		return nil
	}

	router.Walk(walk)
	expectedPaths := []string{"/health", "/book", "/book", "/book/{id}", "/book/{id}"}
	expectedMethods := []string{GET, GET, POST, DELETE, PATCH}

	assert.ElementsMatch(t, expectedMethods, methods)
	assert.ElementsMatch(t, expectedPaths, paths)
}
