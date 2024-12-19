package terrorz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz/terrorz"
)

func TestTestDetailedError(t *testing.T) {
	g := NewWithT(t)

	e := terrorz.NewTestDetailedError(
		"errorMessage",
		"name",
		500,
		map[string]any{"k": "v"})

	g.Expect(e.Error()).To(Equal("errorMessage"))
	g.Expect(e.GetErrorName()).To(Equal("name"))
	g.Expect(e.GetErrorHTTPStatus()).To(Equal(500))
	g.Expect(e.GetErrorDetails()).To(Equal(map[string]any{"k": "v"}))
}

func TestTestDetailedErrorUnwrapSingle(t *testing.T) {
	g := NewWithT(t)

	e := terrorz.NewTestDetailedErrorUnwrapSingle(
		"errorMessage",
		"name",
		500,
		map[string]any{"k": "v"},
		fmt.Errorf("innerErrorMessage"))

	g.Expect(e.Error()).To(Equal("errorMessage"))
	g.Expect(e.GetErrorName()).To(Equal("name"))
	g.Expect(e.GetErrorHTTPStatus()).To(Equal(500))
	g.Expect(e.GetErrorDetails()).To(Equal(map[string]any{"k": "v"}))
	g.Expect(e.Unwrap()).To(Equal(fmt.Errorf("innerErrorMessage")))
}

func TestTestDetailedErrorUnwrapMulti(t *testing.T) {
	g := NewWithT(t)

	e := terrorz.NewTestDetailedErrorUnwrapMulti(
		"errorMessage",
		"name",
		500,
		map[string]any{"k": "v"},
		[]error{fmt.Errorf("innerErrorMessage")})

	g.Expect(e.Error()).To(Equal("errorMessage"))
	g.Expect(e.GetErrorName()).To(Equal("name"))
	g.Expect(e.GetErrorHTTPStatus()).To(Equal(500))
	g.Expect(e.GetErrorDetails()).To(Equal(map[string]any{"k": "v"}))
	g.Expect(e.Unwrap()).To(Equal([]error{fmt.Errorf("innerErrorMessage")}))
}
