package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"http-server/mocks/wrapper"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		setPort     bool
		setToken    bool
		setGRPC     bool
		expectedErr error
		expectedRes *AppConfig
		mocks       func(w *wrapper.Wrapper)
	}{
		{
			setPort:     false,
			expectedErr: fmt.Errorf("NewConfig: invalid value for port: "),
			mocks: func(w *wrapper.Wrapper) {
				w.On("Atoi", "").Return(0, fmt.Errorf("some-error"))
			},
		},
		{
			setPort:     true,
			expectedErr: fmt.Errorf("NewConfig: unable to init config, missing TOKEN, GRPC_URI"),
			mocks: func(w *wrapper.Wrapper) {
				w.On("Atoi", "9000").Return(9000, nil)
			},
		},
		{
			setPort:     true,
			setToken:    true,
			setGRPC:     true,
			expectedErr: nil,
			expectedRes: &AppConfig{
				Port:    9000,
				Token:   "baz",
				GRPCUri: "foo:bar",
				LogFile: os.Stdout,
			},
			mocks: func(w *wrapper.Wrapper) {
				w.On("Atoi", "9000").Return(9000, nil)
			},
		},
	}

	for _, tt := range tests {
		w := &wrapper.Wrapper{}
		if tt.setPort {
			_ = os.Setenv("PORT", "9000")
		}

		if tt.setToken {
			_ = os.Setenv("TOKEN", "baz")
		}

		if tt.setGRPC {
			_ = os.Setenv("GRPC_URI", "foo:bar")
		}

		tt.mocks(w)
		defer os.Clearenv()
		conf, err := NewConfig(w)

		assert.Equal(t, tt.expectedErr, err)
		assert.Equal(t, tt.expectedRes, conf)
		w.AssertExpectations(t)
	}
}
