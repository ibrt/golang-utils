package injectz

import (
	"io"

	"github.com/ibrt/golang-utils/errorz"
)

// Releaser releases a module (e.g. closes DB connections).
type Releaser func()

// NewNoopReleaser returns a [Releaser] that does nothing.
func NewNoopReleaser() Releaser {
	return func() {
		// intentionally empty
	}
}

// NewCloseReleaser returns a [Releaser] that calls [io.Closer.Close], ignoring any returned error.
func NewCloseReleaser(closer io.Closer) Releaser {
	return func() {
		_ = errorz.Catch0(closer.Close)
	}
}

// NewReleasers combines multiple [Releaser] into a compound one (which invokes them in reverse order).
func NewReleasers(releasers ...Releaser) Releaser {
	return func() {
		for i := len(releasers) - 1; i >= 0; i-- {
			_ = errorz.Catch0(func() error {
				releasers[i]()
				return nil
			})
		}
	}
}
