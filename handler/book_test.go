package handler

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"http-server/mocks/book"
	"http-server/mocks/factory"
	"http-server/proto"
	"net/http/httptest"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		body         []byte
		expectedErr  string
		expectedCode int
		expectedBody string
		entries      int
		mocks        func(f *factory.Factory, b *book.Book)
	}{
		{
			body:         nil,
			expectedErr:  "Create: invalid payload: EOF",
			expectedBody: `{"error": "invalid request"}`,
			expectedCode: 400,
			entries:      1,
			mocks:        func(f *factory.Factory, b *book.Book) {},
		},
		{
			body:         []byte(`{}`),
			expectedErr:  "Create: 'name'/'author' is required",
			expectedBody: `{"error": "invalid request"}`,
			expectedCode: 400,
			entries:      1,
			mocks:        func(f *factory.Factory, b *book.Book) {},
		},
		{
			body:         []byte(`{"name": "foo"}`),
			expectedErr:  "Create: unable to create book",
			expectedBody: `{"error": "unexpected error happened"}`,
			expectedCode: 500,
			entries:      1,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Create", mock.Anything).Return(nil, fmt.Errorf("some-error"))
			},
		},
		{
			body:         []byte(`{"name": "foo"}`),
			expectedErr:  "",
			expectedBody: `{"success": {"name": "foo"}}`,
			expectedCode: 200,
			entries:      0,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Create", mock.Anything).Return(&proto.Book{Name: "foo"}, nil)
			},
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/foo", bytes.NewBuffer(tt.body))
		logger, hook := test.NewNullLogger()
		f := &factory.Factory{}
		b := &book.Book{}

		tt.mocks(f, b)
		handler := Create(f, logger)
		handler.ServeHTTP(w, r)

		assert.Equal(t, tt.expectedCode, w.Code)
		assert.JSONEq(t, tt.expectedBody, w.Body.String())
		assert.Equal(t, tt.entries, len(hook.Entries))
		if tt.entries > 0 {
			assert.Equal(t, tt.expectedErr, hook.LastEntry().Message)
		}

		b.AssertExpectations(t)
		f.AssertExpectations(t)
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		params       map[string]string
		expectedErr  string
		expectedCode int
		expectedBody string
		entries      int
		mocks        func(f *factory.Factory, b *book.Book)
	}{
		{
			params:       map[string]string{},
			expectedBody: `{"error": "invalid request"}`,
			expectedCode: 400,
			expectedErr:  "Get: unable to read 'id' from path params",
			entries:      1,
			mocks:        func(f *factory.Factory, b *book.Book) {},
		},
		{
			params:       map[string]string{"id": "123"},
			expectedBody: `{"error": "unexpected error happened"}`,
			expectedCode: 500,
			expectedErr:  "Get: unable to retrieve book",
			entries:      1,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Get", mock.Anything).Return(nil, fmt.Errorf("some-error"))
			},
		},
		{
			params:       map[string]string{"id": "123"},
			expectedBody: `{"success": {"name": "foo"}}`,
			expectedCode: 200,
			expectedErr:  "",
			entries:      0,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Get", mock.Anything).Return(&proto.Book{Name: "foo"}, nil)
			},
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/foo/id", nil)
		logger, hook := test.NewNullLogger()
		f := &factory.Factory{}
		b := &book.Book{}

		r = mux.SetURLVars(r, tt.params)

		tt.mocks(f, b)
		handler := Get(f, logger)
		handler.ServeHTTP(w, r)

		assert.Equal(t, tt.expectedCode, w.Code)
		assert.JSONEq(t, tt.expectedBody, w.Body.String())
		assert.Equal(t, tt.entries, len(hook.Entries))
		if tt.entries > 0 {
			assert.Equal(t, tt.expectedErr, hook.LastEntry().Message)
		}

		b.AssertExpectations(t)
		f.AssertExpectations(t)
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		params       map[string]string
		expectedErr  string
		expectedCode int
		expectedBody string
		entries      int
		mocks        func(f *factory.Factory, b *book.Book)
	}{
		{
			params:       map[string]string{},
			expectedBody: `{"error": "invalid request"}`,
			expectedCode: 400,
			expectedErr:  "Delete: unable to read 'id' from path params",
			entries:      1,
			mocks:        func(f *factory.Factory, b *book.Book) {},
		},
		{
			params:       map[string]string{"id": "123"},
			expectedBody: `{"error": "unexpected error happened"}`,
			expectedCode: 500,
			expectedErr:  "Delete: unable to delete book",
			entries:      1,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Delete", mock.Anything).Return(nil, fmt.Errorf("some-error"))
			},
		},
		{
			params:       map[string]string{"id": "123"},
			expectedBody: `{"success": {"name": "foo"}}`,
			expectedCode: 200,
			expectedErr:  "",
			entries:      0,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Delete", mock.Anything).Return(&proto.Book{Name: "foo"}, nil)
			},
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/foo/id", nil)
		logger, hook := test.NewNullLogger()
		f := &factory.Factory{}
		b := &book.Book{}

		r = mux.SetURLVars(r, tt.params)

		tt.mocks(f, b)
		handler := Delete(f, logger)
		handler.ServeHTTP(w, r)

		assert.Equal(t, tt.expectedCode, w.Code)
		assert.JSONEq(t, tt.expectedBody, w.Body.String())
		assert.Equal(t, tt.entries, len(hook.Entries))
		if tt.entries > 0 {
			assert.Equal(t, tt.expectedErr, hook.LastEntry().Message)
		}

		b.AssertExpectations(t)
		f.AssertExpectations(t)
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		body         []byte
		expectedErr  string
		expectedCode int
		expectedBody string
		entries      int
		mocks        func(f *factory.Factory, b *book.Book)
	}{
		{
			body:         nil,
			expectedErr:  "Update: invalid payload: EOF",
			expectedBody: `{"error": "invalid request"}`,
			expectedCode: 400,
			entries:      1,
			mocks:        func(f *factory.Factory, b *book.Book) {},
		},
		{
			body:         []byte(`{}`),
			expectedErr:  "Update: 'id' cannot be empty",
			expectedBody: `{"error": "invalid request"}`,
			expectedCode: 400,
			entries:      1,
			mocks:        func(f *factory.Factory, b *book.Book) {},
		},
		{
			body:         []byte(`{"id": "123", "name": "foo"}`),
			expectedErr:  "Update: unable to update book",
			expectedBody: `{"error": "unexpected error happened"}`,
			expectedCode: 500,
			entries:      1,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Update", mock.Anything).Return(nil, fmt.Errorf("some-error"))
			},
		},
		{
			body:         []byte(`{"id": "123", "name": "foo"}`),
			expectedErr:  "",
			expectedBody: `{"success": {"name": "foo"}}`,
			expectedCode: 200,
			entries:      0,
			mocks: func(f *factory.Factory, b *book.Book) {
				f.On("NewBook", mock.Anything).Return(b)
				b.On("Update", mock.Anything).Return(&proto.Book{Name: "foo"}, nil)
			},
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PATCH", "/foo", bytes.NewBuffer(tt.body))
		logger, hook := test.NewNullLogger()
		f := &factory.Factory{}
		b := &book.Book{}

		tt.mocks(f, b)
		handler := Update(f, logger)
		handler.ServeHTTP(w, r)

		assert.Equal(t, tt.expectedCode, w.Code)
		assert.JSONEq(t, tt.expectedBody, w.Body.String())
		assert.Equal(t, tt.entries, len(hook.Entries))
		if tt.entries > 0 {
			assert.Equal(t, tt.expectedErr, hook.LastEntry().Message)
		}

		b.AssertExpectations(t)
		f.AssertExpectations(t)
	}
}
