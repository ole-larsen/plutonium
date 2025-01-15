package compressor

import (
	"compress/gzip"
	"net/http"
)

// CompressWriter - gzip.Writer wrapper.
type CompressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

func NewCompressWriter(w http.ResponseWriter) *CompressWriter {
	return &CompressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

func (c *CompressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *CompressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

func (c *CompressWriter) WriteHeader(statusCode int) {
	const reqStatusCode = 300

	if statusCode < reqStatusCode {
		c.w.Header().Set("Content-Encoding", "gzip")
	}

	c.w.WriteHeader(statusCode)
}

func (c *CompressWriter) Close() error {
	return c.zw.Close()
}
