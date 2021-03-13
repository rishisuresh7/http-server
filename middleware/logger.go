package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type loggerMiddleware struct {
	logger *logrus.Logger
}

func NewLoggerMiddleware(l *logrus.Logger) Middleware {
	return &loggerMiddleware{logger: l}
}

func (l *loggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	method := r.Method
	url := r.URL.String()
	wn := w.(negroni.ResponseWriter)
	l.logger.WithFields(logrus.Fields{
		"Client":     r.RemoteAddr,
		"Method":     method,
		"URL":        url,
		"Referrer":   r.Referer(),
		"User-Agent": r.UserAgent(),
	}).Infof("Request")

	next(w, r)

	l.logger.WithFields(logrus.Fields{
		"Method":     method,
		"URL":        url,
		"StatusCode": wn.Status(),
		"Size":       wn.Size(),
		"Duration":   int64(time.Since(start) / time.Millisecond),
	}).Infof("Response")
}
