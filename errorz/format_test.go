package errorz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

func TestErrorf(t *testing.T) {
	g := NewWithT(t)

	err := errorz.Errorf("e: %v", "v")
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("e: v"))
	g.Expect(err.(errorz.UnwrapMulti).Unwrap()).To(BeEmpty())
}

func TestMustErrorf(t *testing.T) {
	g := NewWithT(t)

	g.Expect(func() { errorz.MustErrorf("e: %v", "v") }).To(PanicWith(MatchError("e: v")))
}

func TestAssertf(t *testing.T) {
	g := NewWithT(t)

	g.Expect(func() { errorz.Assertf(true, "e: %v", "v") }).ToNot(Panic())
	g.Expect(func() { errorz.Assertf(false, "e: %v", "v") }).To(PanicWith(MatchError("e: v")))
}