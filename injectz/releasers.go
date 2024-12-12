package injectz

import (
	"io"
)

// Releaser releases an initialized resource.
type Releaser func()

// NewNoopReleaser returns a releaser that does nothing.
func NewNoopReleaser() Releaser {
	return func() {
		// intentionally empty
	}
}

// NewCloseReleaser returns a releaser that calls closer.Close, ignoring any returned error.
func NewCloseReleaser(closer io.Closer) Releaser {
	return func() {
		defer func() { recover() }()
		_ = closer.Close()
	}
}

// NewReleasers combines multiple releasers into one (invoking them in reverse order).
func NewReleasers(releasers ...Releaser) Releaser {
	return func() {
		for i := len(releasers) - 1; i >= 0; i-- {
			func() {
				defer func() { recover() }()
				releasers[i]()
			}()
		}
	}
}
