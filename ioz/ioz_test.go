package ioz_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/ioz"
)

type readCloser struct {
	r        io.Reader
	closeErr error
	isClosed bool
}

func (c *readCloser) Read(p []byte) (n int, err error) {
	if c.isClosed {
		return 0, fmt.Errorf("already closed")
	}

	return c.r.Read(p)
}

func (c *readCloser) Close() error {
	c.isClosed = true
	return c.closeErr
}

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestCountingReader(g *WithT) {
	s := strings.Repeat("x", 1024)
	c := ioz.NewCountingReader(strings.NewReader(s))
	g.Expect(c.Count()).To(BeNumerically("==", 0))
	n, err := io.Copy(io.Discard, c)
	g.Expect(err).To(Succeed())
	g.Expect(n).To(BeNumerically("==", 1024))
	g.Expect(c.Count()).To(BeNumerically("==", 1024))
}

func (*Suite) TestMustReadAll(g *WithT) {
	g.Expect(func() {
		g.Expect(ioz.MustReadAll(strings.NewReader("r"))).To(Equal([]byte("r")))
	}).ToNot(Panic())

	g.Expect(func() { ioz.MustReadAll(iotest.ErrReader(fmt.Errorf("test error"))) }).
		To(PanicWith(MatchError("test error")))
}

func (*Suite) TestMustReadAllString(g *WithT) {
	g.Expect(func() {
		g.Expect(ioz.MustReadAllString(strings.NewReader("r"))).To(Equal("r"))
	}).ToNot(Panic())

	g.Expect(func() { ioz.MustReadAllString(iotest.ErrReader(fmt.Errorf("test error"))) }).
		To(PanicWith(MatchError("test error")))
}

func (*Suite) TestMustReadAllAndClose(g *WithT) {
	g.Expect(
		func() {
			rc := &readCloser{r: strings.NewReader("r")}
			g.Expect(ioz.MustReadAllAndClose(rc)).To(Equal([]byte("r")))
			g.Expect(rc.isClosed).To(BeTrue())
		}).
		ToNot(Panic())

	{
		rc := &readCloser{r: iotest.ErrReader(fmt.Errorf("test error"))}
		g.Expect(func() { ioz.MustReadAllAndClose(rc) }).To(PanicWith(MatchError("test error")))
		g.Expect(rc.isClosed).To(BeTrue())
	}

	{
		rc := &readCloser{r: strings.NewReader("r"), closeErr: fmt.Errorf("close error")}
		g.Expect(func() { ioz.MustReadAllAndClose(rc) }).To(PanicWith(MatchError("close error")))
		g.Expect(rc.isClosed).To(BeTrue())
	}
}

func (*Suite) TestMustReadAllAndCloseString(g *WithT) {
	g.Expect(
		func() {
			rc := &readCloser{r: strings.NewReader("r")}
			g.Expect(ioz.MustReadAllAndCloseString(rc)).To(Equal("r"))
			g.Expect(rc.isClosed).To(BeTrue())
		}).
		ToNot(Panic())

	{
		rc := &readCloser{r: iotest.ErrReader(fmt.Errorf("test error"))}
		g.Expect(func() { ioz.MustReadAllAndCloseString(rc) }).To(PanicWith(MatchError("test error")))
		g.Expect(rc.isClosed).To(BeTrue())
	}

	{
		rc := &readCloser{r: strings.NewReader("r"), closeErr: fmt.Errorf("close error")}
		g.Expect(func() { ioz.MustReadAllAndCloseString(rc) }).To(PanicWith(MatchError("close error")))
		g.Expect(rc.isClosed).To(BeTrue())
	}
}
