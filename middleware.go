package rollbar

import (
	"net/http"
)

// WrapHandler ...
func WrapHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Wrap(func() { h.ServeHTTP(w, r) })
	})
}

// WrapMiddleware ...
func WrapMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	Wrap(func() { next(w, r) })
}

// WrapHandlerFunc ...
func WrapHandlerFunc(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WrapMiddleware(w, r, next)
	})
}
