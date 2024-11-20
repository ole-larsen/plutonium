package prometheus

import (
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// NewResponseWriter creates a new instance of prometheus.ResponseWriter.
// A pointer to a new prometheus.ResponseWriter, initialized with the provided http.ResponseWriter
//
//	and a default HTTP status code of 200 (http.StatusOK).
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK}
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}
