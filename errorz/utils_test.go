package errorz_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/ioz/tioz"
)

func TestIgnoreClose_Success(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	c := tioz.NewMockTestReadCloser(ctrl)

	c.EXPECT().
		Close().
		Return(nil).
		Times(1)

	g.Expect(func() { errorz.IgnoreClose(c) }).ToNot(Panic())
}

func TestIgnoreClose_Error(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	c := tioz.NewMockTestReadCloser(ctrl)

	c.EXPECT().
		Close().
		Return(errorz.Errorf("test error")).
		Times(1)

	g.Expect(func() { errorz.IgnoreClose(c) }).ToNot(Panic())
}

func TestIgnoreClose_Panic(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	c := tioz.NewMockTestReadCloser(ctrl)

	c.EXPECT().
		Close().
		DoAndReturn(func() error {
			errorz.MustErrorf("test error")
			return nil
		}).
		Times(1)

	g.Expect(func() { errorz.IgnoreClose(c) }).ToNot(Panic())
}

func TestMustClose_Success(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	c := tioz.NewMockTestReadCloser(ctrl)

	c.EXPECT().
		Close().
		Return(nil).
		Times(1)

	g.Expect(func() { errorz.MustClose(c) }).ToNot(Panic())
}

func TestMustClose_Error(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	c := tioz.NewMockTestReadCloser(ctrl)

	c.EXPECT().
		Close().
		Return(errorz.Errorf("test error")).
		Times(1)

	g.Expect(func() { errorz.MustClose(c) }).To(PanicWith(MatchError("test error")))
}
