package gzipz

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/ibrt/golang-utils/errorz"
)

// MustCompress implements in-memory GZIP compression.
func MustCompress(buf []byte) []byte {
	cBuf := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(cBuf, gzip.BestCompression)
	errorz.MaybeMustWrap(err)
	defer func() { errorz.MaybeMustWrap(w.Close()) }()

	_, err = w.Write(buf)
	errorz.MaybeMustWrap(err)
	errorz.MaybeMustWrap(w.Close())

	return cBuf.Bytes()
}

// MustDecompress implements in-memory GZIP decompression.
func MustDecompress(buf []byte) []byte {
	r, err := gzip.NewReader(bytes.NewReader(buf))
	errorz.MaybeMustWrap(err)
	defer func() { errorz.MaybeMustWrap(r.Close()) }()

	dBuf, err := io.ReadAll(r)
	errorz.MaybeMustWrap(err)
	return dBuf
}
