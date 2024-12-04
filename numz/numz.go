package numz

import (
	"reflect"
	"strconv"

	"github.com/ibrt/golang-utils/errorz"
)

// Signed describes signed integers.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned describes unsigned integers.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer describes signed and unsigned integers.
type Integer interface {
	Signed | Unsigned
}

// Float describes floating point numbers.
type Float interface {
	~float32 | ~float64
}

// Number describes a number.
type Number interface {
	Integer | Float
}

// Parse a number of the given generic type from string using default settings.
// Note that this method uses reflection.
func Parse[T Number](s string) (T, error) {
	var z T
	rt := reflect.TypeOf(z)

	switch rt.Kind() {
	case reflect.Float32, reflect.Float64:
		t, err := strconv.ParseFloat(s, rt.Bits())
		if err != nil {
			return z, errorz.Wrap(err)
		}
		return T(t), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t, err := strconv.ParseInt(s, 0, rt.Bits())
		if err != nil {
			return z, errorz.Wrap(err)
		}
		return T(t), nil
	default:
		t, err := strconv.ParseUint(s, 0, rt.Bits())
		if err != nil {
			return z, errorz.Wrap(err)
		}
		return T(t), nil
	}
}

// MustParse is like parse, but panics on error.
func MustParse[T Number](s string) T {
	t, err := Parse[T](s)
	errorz.MaybeMustWrap(err)
	return t
}
