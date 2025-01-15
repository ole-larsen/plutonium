// Package compressor is a wrapper over gzip writer and reader
package compressor

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

// Compress - huge gzip compressor.
func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	_, err := w.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %w", err)
	}

	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("failed compress data: %w", err)
	}

	return b.Bytes(), nil
}

// Decompress - decompress incoming data.
func Decompress(reader io.Reader) ([]byte, error) {
	r, err := gzip.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed read decompresses data: %w", err)
	}

	defer func() {
		e := r.Close()
		if e != nil {
			return
		}
	}()

	var b bytes.Buffer

	_, err = b.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("failed decompress data: %w", err)
	}

	return b.Bytes(), nil
}
