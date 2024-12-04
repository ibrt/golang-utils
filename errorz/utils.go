package errorz

import (
	"io"
)

// IgnoreClose calls Close on the given io.Closer, ignoring the returned error. Handy for the defer Close pattern.
func IgnoreClose(c io.Closer) {
	if c != nil {
		_ = c.Close()
	}
}

// MustClose calls Close on the given io.Closer, panicking in case of error. Handy for the defer Close pattern.
func MustClose(c io.Closer) {
	if c != nil {
		MaybeMustWrap(c.Close())
	}
}
