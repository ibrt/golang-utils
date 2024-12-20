package errorz

// MaybeSetMetadata sets the given metadata (k, v) on the error if it has been wrapped, does nothing otherwise.
func MaybeSetMetadata(err error, k, v any) {
	if e, ok := err.(*wrappedError); ok { //nolint:errorlint
		e.setMetadata(k, v)
	}
}

// MustGetMetadata gets the given metadata key from the error, panics if not found or wrong type.
func MustGetMetadata[T any](err error, k any) T {
	return err.(*wrappedError).metadata[k].(T) //nolint:errorlint
}

// MaybeGetMetadata tries to get the given metadata key from the error.
func MaybeGetMetadata[T any](err error, k any) (T, bool) {
	if e, ok := err.(*wrappedError); ok { //nolint:errorlint
		if m, ok := e.getMetadata(k); ok {
			if v, ok := m.(T); ok {
				return v, true
			}
		}
	}

	var v T
	return v, false
}
