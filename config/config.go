package config

import (
	"fmt"
	"os"
	"strings"

	"http-server/wrapper"
)

type AppConfig struct {
	Port    int
	Token   string
	GRPCUri string
	LogFile *os.File
}

func NewConfig(w wrapper.Wrapper) (*AppConfig, error) {
	var missing []string
	portString := os.Getenv("PORT")
	port, err := w.Atoi(portString)
	if err != nil {
		return nil, fmt.Errorf("NewConfig: invalid value for port: %s", portString)
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		missing = append(missing, "TOKEN")
	}

	grpcUri := os.Getenv("GRPC_URI")
	if grpcUri == "" {
		missing = append(missing, "GRPC_URI")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("NewConfig: unable to init config, missing %s", strings.Join(missing, ", "))
	}

	return &AppConfig{
		Port:    port,
		GRPCUri: grpcUri,
		Token:   token,
		LogFile: os.Stdout,
	}, nil
}
