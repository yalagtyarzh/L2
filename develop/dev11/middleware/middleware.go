package middleware

import (
	"net/http"
	"time"

	"dev11/logging"
)

func EventLogger(next http.Handler, logger *logging.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			next.ServeHTTP(w, r)
			//log.Printf("Method: %s URI: %s Time: %s", r.Method, r.RequestURI, now)
			logger.Infof("Method: %s URI: %s Time: %s", r.Method, r.RequestURI, now)
		},
	)
}
