package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	formatter "github.com/x-cray/logrus-prefixed-formatter"

	"http-server/factory"
	"http-server/middleware"
	"http-server/router"
)

var Version = "0.0.0"

const (
	PORT = 9001
	TIME = "2006-01-02 15:04:05"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &formatter.TextFormatter{
		FullTimestamp:   true,
		ForceFormatting: true,
		TimestampFormat: TIME,
	}

	logger.Infof("Running server version: %s", Version)
	n := negroni.New()
	n.Use(middleware.NewLoggerMiddleware(logger))
	n.UseHandler(router.NewRouter(factory.NewFactory(), logger))
	n.Run(fmt.Sprintf(":%d", PORT))
}
