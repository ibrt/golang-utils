package ioz_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/ioz"
	"github.com/ibrt/golang-utils/ioz/tioz"
)

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

func (*Suite) TestMustReadAllAndClose_Success(g *WithT, ctrl *gomock.Controller) {
	rc := tioz.NewMockTestReadCloser(ctrl)

	rc.EXPECT().
		Read(gomock.Any()).
		DoAndReturn(func(p []byte) (n int, err error) {
			p[0] = 'r'
			return 1, io.EOF
		}).
		Times(1)

	rc.EXPECT().
		Close().
		Return(nil).
		Times(1)

	g.Expect(ioz.MustReadAllAndClose(rc)).To(Equal([]byte("r")))
}

func (*Suite) TestMustReadAllAndClose_ReadError(g *WithT, ctrl *gomock.Controller) {
	rc := tioz.NewMockTestReadCloser(ctrl)

	rc.EXPECT().
		Read(gomock.Any()).
		Return(0, errorz.Errorf("test error")).
		Times(1)

	rc.EXPECT().
		Close().
		Return(nil).
		Times(1)

	g.Expect(func() { ioz.MustReadAllAndClose(rc) }).To(PanicWith(MatchError("test error")))
}

func (*Suite) TestMustReadAllAndClose_CloseError(g *WithT, ctrl *gomock.Controller) {
	rc := tioz.NewMockTestReadCloser(ctrl)

	rc.EXPECT().
		Read(gomock.Any()).
		DoAndReturn(func(p []byte) (n int, err error) {
			p[0] = 'r'
			return 1, io.EOF
		}).
		Times(1)

	rc.EXPECT().
		Close().
		Return(errorz.Errorf("test error")).
		Times(1)

	g.Expect(func() { ioz.MustReadAllAndClose(rc) }).To(PanicWith(MatchError("test error")))
}

func (*Suite) TestMustReadAllAndCloseString_Success(g *WithT, ctrl *gomock.Controller) {
	rc := tioz.NewMockTestReadCloser(ctrl)

	rc.EXPECT().
		Read(gomock.Any()).
		DoAndReturn(func(p []byte) (n int, err error) {
			p[0] = 'r'
			return 1, io.EOF
		}).
		Times(1)

	rc.EXPECT().
		Close().
		Return(nil).
		Times(1)

	g.Expect(ioz.MustReadAllAndCloseString(rc)).To(Equal("r"))
}

func (*Suite) TestMustReadAllAndCloseString_ReadError(g *WithT, ctrl *gomock.Controller) {
	rc := tioz.NewMockTestReadCloser(ctrl)

	rc.EXPECT().
		Read(gomock.Any()).
		Return(0, errorz.Errorf("test error")).
		Times(1)

	rc.EXPECT().
		Close().
		Return(nil).
		Times(1)

	g.Expect(func() { ioz.MustReadAllAndCloseString(rc) }).To(PanicWith(MatchError("test error")))
}

func (*Suite) TestMustReadAllAndCloseString_CloseError(g *WithT, ctrl *gomock.Controller) {
	rc := tioz.NewMockTestReadCloser(ctrl)

	rc.EXPECT().
		Read(gomock.Any()).
		DoAndReturn(func(p []byte) (n int, err error) {
			p[0] = 'r'
			return 1, io.EOF
		}).
		Times(1)

	rc.EXPECT().
		Close().
		Return(errorz.Errorf("test error")).
		Times(1)

	g.Expect(func() { ioz.MustReadAllAndCloseString(rc) }).To(PanicWith(MatchError("test error")))
}
