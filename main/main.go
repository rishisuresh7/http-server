package main

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	formatter "github.com/x-cray/logrus-prefixed-formatter"

	"http-server/config"
	"http-server/factory"
	"http-server/middleware"
	"http-server/router"
	"http-server/wrapper"
)

var Version = "0.0.0"

const (
	TIME = "2006-01-02 15:04:05"
)

func main() {
	conf, err := config.NewConfig(wrapper.NewWrapper())
	if err != nil {
		log.Fatalf("Invalid configurations: %s" , err.Error())
	}

	logger := logrus.New()
	logger.Formatter = &formatter.TextFormatter{
		FullTimestamp:   true,
		ForceFormatting: true,
		TimestampFormat: TIME,
	}
	logger.Out = conf.LogFile
	logger.Infof("Running server version: %s", Version)
	n := negroni.New()
	f := factory.NewFactory(conf, logger)
	n.Use(middleware.NewLoggerMiddleware(logger))
	n.UseHandler(router.NewRouter(f, logger))
	n.Run(fmt.Sprintf(":%d", conf.Port))
}
