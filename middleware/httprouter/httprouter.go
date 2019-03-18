// Package httprouter is a helper package to get a httprouter compatible
// handler/middleware from the standatd net/http Middleware factory.
package httprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/slok/go-http-metrics/middleware"
)

// Handler returns a httprouter.Handler compatible middleware from a Middleware factory instance.
// The first handlerID argument is the same argument passed on Middleware.Handler method.
// The second argument is the handler that wants to be wrapped.
func Handler(handlerID string, next httprouter.Handle, m middleware.Middleware) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Dummy handler to wrap httprouter Handle type
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next(w, r, p)
		})

		m.Handler(handlerID, h).ServeHTTP(w, r)
	}
}
