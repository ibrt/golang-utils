package terrorz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz/terrorz"
)

func TestTestDetailedError(t *testing.T) {
	g := NewWithT(t)

	e := &terrorz.SimpleMockTestDetailedError{
		ErrorMessage: "errorMessage",
		Name:         "name",
		HTTPStatus:   500,
		Details:      map[string]any{"k": "v"},
	}
	g.Expect(e.Error()).To(Equal("errorMessage"))
	g.Expect(e.GetErrorName()).To(Equal("name"))
	g.Expect(e.GetErrorHTTPStatus()).To(Equal(500))
	g.Expect(e.GetErrorDetails()).To(Equal(map[string]any{"k": "v"}))
}

func TestTestDetailedErrorUnwrapSingle(t *testing.T) {
	g := NewWithT(t)

	e := &terrorz.SimpleMockTestDetailedUnwrapSingleError{
		SimpleMockTestDetailedError: &terrorz.SimpleMockTestDetailedError{
			ErrorMessage: "errorMessage",
			Name:         "name",
			HTTPStatus:   500,
			Details:      map[string]any{"k": "v"},
		},
		UnwrapSingle: fmt.Errorf("innerErrorMessage"),
	}
	g.Expect(e.Error()).To(Equal("errorMessage"))
	g.Expect(e.GetErrorName()).To(Equal("name"))
	g.Expect(e.GetErrorHTTPStatus()).To(Equal(500))
	g.Expect(e.GetErrorDetails()).To(Equal(map[string]any{"k": "v"}))
	g.Expect(e.Unwrap()).To(Equal(fmt.Errorf("innerErrorMessage")))
}

func TestTestDetailedErrorUnwrapMulti(t *testing.T) {
	g := NewWithT(t)

	e := &terrorz.SimpleMockTestDetailedUnwrapMultiError{
		SimpleMockTestDetailedError: &terrorz.SimpleMockTestDetailedError{
			ErrorMessage: "errorMessage",
			Name:         "name",
			HTTPStatus:   500,
			Details:      map[string]any{"k": "v"},
		},
		UnwrapMulti: []error{
			fmt.Errorf("innerErrorMessage"),
		},
	}
	g.Expect(e.Error()).To(Equal("errorMessage"))
	g.Expect(e.GetErrorName()).To(Equal("name"))
	g.Expect(e.GetErrorHTTPStatus()).To(Equal(500))
	g.Expect(e.GetErrorDetails()).To(Equal(map[string]any{"k": "v"}))
	g.Expect(e.Unwrap()).To(Equal([]error{fmt.Errorf("innerErrorMessage")}))
}

func TestTestStringError(t *testing.T) {
	g := NewWithT(t)

	g.Expect(terrorz.TestStringError("errorMessage").Error()).To(Equal("errorMessage"))
}
