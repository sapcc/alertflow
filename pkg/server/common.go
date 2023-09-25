package server

import (
	"net/http"
)

// LogReqInfo describes info about HTTP request
type Agent struct {
	method string
	uri    string
	code   int
}

// wrapper to catch response code
type responseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func (w *responseWriterWrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// server wrapper
func WrapHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resWrapper := responseWriterWrapper{w, 200}

		next.ServeHTTP(&resWrapper, r)

		agent := &Agent{
			method: r.Method,
			uri:    r.URL.String(),
			code:   resWrapper.status,
		}
		logger.Printf("served request: %+v \n", agent)
	}
}
