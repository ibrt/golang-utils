package injectz_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/injectz"
	"github.com/ibrt/golang-utils/injectz/tinjectz"
)

type InitializersSuite struct {
	// intentionally empty
}

func TestInitializersSuite(t *testing.T) {
	fixturez.RunSuite(t, &InitializersSuite{})
}

func (*InitializersSuite) TestMustInitialize_Success(g *WithT, ctrl *gomock.Controller) {
	firstInitializer := tinjectz.NewMockTestInitializer(ctrl)
	secondInitializer := tinjectz.NewMockTestInitializer(ctrl)
	firstInjector := tinjectz.NewMockTestInjector(ctrl)
	secondInjector := tinjectz.NewMockTestInjector(ctrl)
	firstReleaser := tinjectz.NewMockTestReleaser(ctrl)
	secondReleaser := tinjectz.NewMockTestReleaser(ctrl)
	isSecondReleased := false

	firstInitializer.EXPECT().Initialize(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) == nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(_ context.Context) (injectz.Injector, injectz.Releaser) {
			return firstInjector.Inject, firstReleaser.Release
		})

	secondInitializer.EXPECT().Initialize(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) != nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(_ context.Context) (injectz.Injector, injectz.Releaser) {
			return secondInjector.Inject, secondReleaser.Release
		})

	firstInjector.EXPECT().Inject(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) == nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(ctx context.Context) context.Context {
			return context.WithValue(ctx, tinjectz.TestContextKeyA0, "v1")
		}).
		Times(2)

	secondInjector.EXPECT().Inject(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) != nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(ctx context.Context) context.Context {
			return context.WithValue(ctx, tinjectz.TestContextKeyA1, "v2")
		}).
		Times(2)

	firstReleaser.EXPECT().Release().Do(func() { g.Expect(isSecondReleased).To(BeTrue()) })
	secondReleaser.EXPECT().Release().Do(func() { isSecondReleased = true })

	injector, releaser := injectz.NewBootstrap().
		Add(firstInitializer.Initialize, secondInitializer.Initialize).
		MustInitialize()

	ctx := injector(context.Background())
	g.Expect(ctx.Value(tinjectz.TestContextKeyA0)).To(Equal("v1"))
	g.Expect(ctx.Value(tinjectz.TestContextKeyA1)).To(Equal("v2"))
	releaser()
}

func (*InitializersSuite) TestMustInitialize_Error(g *WithT, ctrl *gomock.Controller) {
	firstInitializer := tinjectz.NewMockTestInitializer(ctrl)
	secondInitializer := tinjectz.NewMockTestInitializer(ctrl)
	firstInjector := tinjectz.NewMockTestInjector(ctrl)
	firstReleaser := tinjectz.NewMockTestReleaser(ctrl)

	firstInitializer.EXPECT().Initialize(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) == nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(_ context.Context) (injectz.Injector, injectz.Releaser) {
			return firstInjector.Inject, firstReleaser.Release
		})

	secondInitializer.EXPECT().Initialize(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) != nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(_ context.Context) (injectz.Injector, injectz.Releaser) {
			panic(fmt.Errorf("initializer error"))
		})

	firstInjector.EXPECT().Inject(
		gomock.Cond(func(ctx context.Context) bool {
			return ctx.Value(tinjectz.TestContextKeyA0) == nil && ctx.Value(tinjectz.TestContextKeyA1) == nil
		})).
		DoAndReturn(func(ctx context.Context) context.Context {
			return context.WithValue(ctx, tinjectz.TestContextKeyA0, "v1")
		})

	firstReleaser.EXPECT().Release()

	g.Expect(
		func() {
			injectz.NewBootstrap().
				Add(firstInitializer.Initialize, secondInitializer.Initialize).
				MustInitialize()
		}).
		To(PanicWith(MatchError("initializer error")))
}
