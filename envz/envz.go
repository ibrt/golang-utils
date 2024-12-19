// Package envz provides various utilities for working with environment variables.
package envz

import (
	"fmt"
	"os"
	"strings"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
)

// UnmarshalEnviron converts an OS-like environ to env map.
func UnmarshalEnviron(environ []string, filterPrefix string) map[string]string {
	env := make(map[string]string, len(environ))

	for _, e := range environ {
		p := strings.SplitN(e, "=", 2)

		if len(p) == 2 && p[0] != "" && strings.HasPrefix(p[0], filterPrefix) {
			env[p[0]] = p[1]
		}
	}

	return env
}

// MarshalEnviron converts an env map to OS-like environ.
func MarshalEnviron(env map[string]string) []string {
	environ := make([]string, 0, len(env))

	for k, v := range env {
		environ = append(environ, fmt.Sprintf("%v=%v", k, v))
	}

	return environ
}

// MustSetenv is like [os.Setenv], but panics on error.
func MustSetenv(key, value string) {
	errorz.MaybeMustWrap(os.Setenv(key, value))
}

// MustUnsetenv is like [os.Unsetenv], but panics on error.
func MustUnsetenv(key string) {
	errorz.MaybeMustWrap(os.Unsetenv(key))
}

// WithEnv runs the closure after setting the given env, cleans up after.
func WithEnv(env map[string]string, f func()) {
	origEnv := make(map[string]*string)

	defer func() {
		for k, v := range origEnv {
			if v != nil {
				MustSetenv(k, *v)
			} else {
				MustUnsetenv(k)
			}
		}
	}()

	for k, v := range env {
		if v, ok := os.LookupEnv(k); ok {
			origEnv[k] = memz.Ptr(v)
		} else {
			origEnv[k] = nil
		}

		MustSetenv(k, v)
	}

	f()
}
