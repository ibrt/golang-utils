package errorz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

type closer struct {
	err      error
	isClosed bool
}

func (c *closer) Close() error {
	c.isClosed = true
	return c.err
}

func TestIgnoreClose(t *testing.T) {
	g := NewWithT(t)

	g.Expect(
		func() {
			c := &closer{err: nil}
			errorz.IgnoreClose(c)
			g.Expect(c.isClosed).To(BeTrue())
		}).
		ToNot(Panic())

	g.Expect(
		func() {
			c := &closer{err: fmt.Errorf("e")}
			errorz.IgnoreClose(c)
			g.Expect(c.isClosed).To(BeTrue())
		}).
		ToNot(Panic())
}

func TestMustClose(t *testing.T) {
	g := NewWithT(t)

	g.Expect(
		func() {
			c := &closer{err: nil}
			errorz.MustClose(c)
			g.Expect(c.isClosed).To(BeTrue())
		}).
		ToNot(Panic())

	g.Expect(
		func() {
			c := &closer{err: fmt.Errorf("e")}
			errorz.MustClose(c)
		}).
		To(PanicWith(MatchError("e")))
}
