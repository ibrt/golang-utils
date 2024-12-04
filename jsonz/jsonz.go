package jsonz

import (
	"encoding/json"
	"strings"

	"github.com/ibrt/golang-utils/errorz"
)

// MustMarshal is like json.Marshal but panics on error.
func MustMarshal(v any) []byte {
	buf, err := json.Marshal(v)
	errorz.MaybeMustWrap(err)
	return buf
}

// MustMarshalString is like MustMarshal but returns a string.
func MustMarshalString(v any) string {
	return string(MustMarshal(v))
}

// MustMarshalPretty is like json.MarshalIndent with prefix = "" and indent = "  " but panics on error.
func MustMarshalPretty(v any) []byte {
	buf, err := json.MarshalIndent(v, "", "  ")
	errorz.MaybeMustWrap(err)
	return buf
}

// MustMarshalPrettyString is like MustMarshalIndent but returns a string.
func MustMarshalPrettyString(v any) string {
	return string(MustMarshalPretty(v))
}

// Unmarshal is like json.Unmarshal but instantiates the target using a generic type.
func Unmarshal[T any](data []byte) (T, error) {
	var t T
	return t, errorz.MaybeWrap(json.Unmarshal(data, &t))
}

// MustUnmarshal is like Unmarshal but panics on error.
func MustUnmarshal[T any](data []byte) T {
	t, err := Unmarshal[T](data)
	errorz.MaybeMustWrap(err)
	return t
}

// UnmarshalString is like Unmarshal but accepts a string.
func UnmarshalString[T any](data string) (T, error) {
	var t T
	return t, errorz.MaybeWrap(json.NewDecoder(strings.NewReader(data)).Decode(&t))
}

// MustUnmarshalString is like UnmarshalString but panics on error.
func MustUnmarshalString[T any](data string) T {
	t, err := UnmarshalString[T](data)
	errorz.MaybeMustWrap(err)
	return t
}
