package injectz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/injectz"
	"github.com/ibrt/golang-utils/injectz/internal/tinjectz"
)

type ReleasersSuite struct {
	// intentionally empty
}

func TestReleasersSuite(t *testing.T) {
	fixturez.RunSuite(t, &ReleasersSuite{})
}

func (*ReleasersSuite) TestNewNoopReleaser(g *WithT) {
	g.Expect(func() {
		injectz.NewNoopReleaser()()
	}).ToNot(Panic())
}

func (*ReleasersSuite) TestNewCloseReleaser(ctrl *gomock.Controller) {
	closer := tinjectz.NewMockCloser(ctrl)
	closer.EXPECT().Close().Return(fmt.Errorf("close error"))
	injectz.NewCloseReleaser(closer)()
}

func (*ReleasersSuite) TestNewReleasers(g *WithT, ctrl *gomock.Controller) {
	firstReleaser := tinjectz.NewMockReleaser(ctrl)
	secondReleaser := tinjectz.NewMockReleaser(ctrl)
	isSecondReleased := false

	firstReleaser.EXPECT().Release().Do(func() { g.Expect(isSecondReleased).To(BeTrue()) })
	secondReleaser.EXPECT().Release().Do(func() { isSecondReleased = true })

	injectz.NewReleasers(firstReleaser.Release, secondReleaser.Release)()
}
