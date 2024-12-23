package errorz

import (
	"io"
)

// IgnoreClose calls [io.Closer.Close], ignoring the returned error. Handy for the "defer Close" pattern.
func IgnoreClose(c io.Closer) {
	defer func() {
		_ = recover()
	}()

	if c != nil {
		_ = c.Close()
	}
}

// MustClose calls [io.Closer.Close], panicking in case of error. Handy for the "defer Close" pattern.
func MustClose(c io.Closer) {
	if c != nil {
		MaybeMustWrap(c.Close())
	}
}
