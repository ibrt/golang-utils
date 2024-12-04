package stringz

import (
	"strings"
)

// EnsureSuffix ensures that the string has the given (non-repeated) suffix.
func EnsureSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s
	}

	return s + suffix
}

// EnsurePrefix ensures that the string has the given (non-repeated) prefix.
func EnsurePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s
	}

	return prefix + s
}
