// Package ioz provides various utilities for working with input/output.
package ioz

import (
	"io"
	"sync/atomic"

	"github.com/ibrt/golang-utils/errorz"
)

var (
	_ io.Reader = (*CountingReader)(nil)
)

// CountingReader implements a [io.Reader] that counts bytes.
type CountingReader struct {
	r       io.Reader
	counter *atomic.Int64
}

// NewCountingReader initializes a new [*CountingReader].
func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader{
		r:       r,
		counter: &atomic.Int64{},
	}
}

// Read implements the [io.Reader] interface.
func (c *CountingReader) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	c.counter.Add(int64(n))
	return n, err
}

// Count returns the number of bytes read.
func (c *CountingReader) Count() int64 {
	return c.counter.Load()
}

// MustReadAll is like [io.ReadAll], but panics on error.
func MustReadAll(r io.Reader) []byte {
	buf, err := io.ReadAll(r)
	errorz.MaybeMustWrap(err)
	return buf
}

// MustReadAllString is like [MustReadAll], but returns a string.
func MustReadAllString(r io.Reader) string {
	return string(MustReadAll(r))
}

// MustReadAllAndClose is like MustReadAll but also always closes the ReadCloser.
func MustReadAllAndClose(r io.ReadCloser) []byte {
	defer errorz.MustClose(r)
	return MustReadAll(r)
}

// MustReadAllAndCloseString is like MustReadAllAndClose, but returns a string.
func MustReadAllAndCloseString(r io.ReadCloser) string {
	return string(MustReadAllAndClose(r))
}
