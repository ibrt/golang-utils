package errorz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

func TestMetadata(t *testing.T) {
	g := NewWithT(t)

	type mk1 int
	type mk2 int

	const k1 mk1 = 0
	const k2 mk2 = 0

	{
		err := fmt.Errorf("test error")
		errorz.MaybeSetMetadata(err, k1, "v1")
		_, ok := errorz.MaybeGetMetadata[string](err, k1)
		g.Expect(ok).To(BeFalse())
		g.Expect(func() { errorz.MustGetMetadata[string](err, k1) }).To(Panic())
	}

	{
		err := errorz.Errorf("test error")
		errorz.MaybeSetMetadata(err, k1, "v1")
		errorz.MaybeSetMetadata(err, k2, "v2")

		v, ok := errorz.MaybeGetMetadata[string](err, k1)
		g.Expect(ok).To(BeTrue())
		g.Expect(v).To(Equal("v1"))

		g.Expect(func() {
			g.Expect(errorz.MustGetMetadata[string](err, k1)).To(Equal("v1"))
		}).ToNot(Panic())

		v, ok = errorz.MaybeGetMetadata[string](err, k2)
		g.Expect(ok).To(BeTrue())
		g.Expect(v).To(Equal("v2"))
		g.Expect(func() { errorz.MustGetMetadata[int](err, k2) }).To(Panic())

		_, ok = errorz.MaybeGetMetadata[string](err, 0)
		g.Expect(ok).To(BeFalse())
		g.Expect(func() { errorz.MustGetMetadata[string](err, 0) }).To(Panic())
	}
}
