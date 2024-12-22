package injectz_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/injectz"
	"github.com/ibrt/golang-utils/injectz/tinjectz"
)

type InjectorsSuite struct {
	// intentionally empty
}

func TestInjectorsSuite(t *testing.T) {
	fixturez.RunSuite(t, &InjectorsSuite{})
}

func (*InjectorsSuite) TestNewNoopInjector(g *WithT) {
	ctx := injectz.NewNoopInjector()(context.Background())
	g.Expect(ctx).To(Equal(context.Background()))
}

func (*InjectorsSuite) TestNewSingletonInjector(g *WithT) {
	type contextKey int
	const myContextKey contextKey = iota

	ctx := injectz.NewSingletonInjector(myContextKey, "v1")(context.Background())
	g.Expect(ctx.Value(myContextKey)).To(Equal("v1"))
}

func (*InjectorsSuite) TestNewInjectors(g *WithT, ctrl *gomock.Controller) {
	firstInjector := tinjectz.NewMockTestInjector(ctrl)
	secondInjector := tinjectz.NewMockTestInjector(ctrl)

	firstInjector.EXPECT().Inject(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) == nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(ctx context.Context) context.Context {
			return context.WithValue(ctx, tinjectz.TestContextKeyA0, "v1")
		})

	secondInjector.EXPECT().Inject(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) != nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(ctx context.Context) context.Context {
			return context.WithValue(ctx, tinjectz.TestContextKeyA1, "v2")
		})

	ctx := injectz.NewInjectors(firstInjector.Inject, secondInjector.Inject)(context.Background())
	g.Expect(ctx.Value(tinjectz.TestContextKeyA0)).To(Equal("v1"))
	g.Expect(ctx.Value(tinjectz.TestContextKeyA1)).To(Equal("v2"))
}
