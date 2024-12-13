// Package urlz provides various utilities for working URLs.
package urlz

import (
	"net/url"

	"github.com/ibrt/golang-utils/errorz"
)

// MustParse is like [url.Parse] but panics on error.
func MustParse(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	errorz.MaybeMustWrap(err)
	return u
}

// MustEdit parses the given URL, calls f to allow mutations, and then converts the URL back to string.
func MustEdit(rawURL string, f func(*url.URL)) string {
	u := MustParse(rawURL)
	f(u)
	return u.String()
}

// GetValueDef tries to get a key from [url.Values], returns def if not found.
func GetValueDef(v url.Values, key, def string) string {
	if vv := v.Get(key); vv != "" {
		return vv
	}
	return def
}

// EncodeValues is a shorthand for building an encoded query string.
func EncodeValues(m map[string]string) string {
	v := url.Values{}
	for k, vv := range m {
		v.Set(k, vv)
	}
	return v.Encode()
}
