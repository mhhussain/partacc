package main

import (
	"net/http"
)

var rCount = 0

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*start := time.Now()*/
		rCount++
		inner.ServeHTTP(w, r)

		/*log.Printf(
			"%s\t%s\t\t%s\t\t\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)*/
	})
}
