package errorz

import (
	"context"
)

// Catch0 catches panics in a "func() error" closure.
func Catch0(f func() error) (outErr error) {
	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			outErr = rErr
		}
	}()

	return MaybeWrap(f())
}

// Catch0Ctx catches panics in a "func(context.Context) error" closure.
func Catch0Ctx(ctx context.Context, f func(ctx context.Context) error) (outErr error) {
	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			outErr = rErr
		}
	}()

	return MaybeWrap(f(ctx))
}

// Catch1 catches panics in a "func() (T, error)" closure.
func Catch1[T any](f func() (T, error)) (outV T, outErr error) {
	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			var t T
			outV = t
			outErr = rErr
		}
	}()

	out, err := f()
	return out, MaybeWrap(err)
}

// Catch1Ctx catches panics in a "func(context.Context) (T, error)" closure.
func Catch1Ctx[T any](ctx context.Context, f func(ctx context.Context) (T, error)) (outV T, outErr error) {
	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			var t T
			outV = t
			outErr = rErr
		}
	}()

	out, err := f(ctx)
	return out, MaybeWrap(err)
}

// Catch2 catches panics in a "func() (T1, T2, error)" closure.
func Catch2[T1 any, T2 any](f func() (T1, T2, error)) (outV1 T1, outV2 T2, outErr error) {
	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			var t1 T1
			var t2 T2

			outV1 = t1
			outV2 = t2
			outErr = rErr
		}
	}()

	out1, out2, err := f()
	return out1, out2, MaybeWrap(err)
}

// Catch2Ctx catches panics in a "func(context.Context) (T1, T2, error)" closure.
func Catch2Ctx[T1 any, T2 any](
	ctx context.Context,
	f func(ctx context.Context) (T1, T2, error),
) (outV1 T1, outV2 T2, outErr error) {

	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			var t1 T1
			var t2 T2

			outV1 = t1
			outV2 = t2
			outErr = rErr
		}
	}()

	out1, out2, err := f(ctx)
	return out1, out2, MaybeWrap(err)
}

// Catch3 catches panics in a "func() (T1, T2, T3, error)" closure.
func Catch3[T1 any, T2 any, T3 any](f func() (T1, T2, T3, error)) (outV1 T1, outV2 T2, outV3 T3, outErr error) {
	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			var t1 T1
			var t2 T2
			var t3 T3

			outV1 = t1
			outV2 = t2
			outV3 = t3
			outErr = rErr
		}
	}()

	out1, out2, out3, err := f()
	return out1, out2, out3, MaybeWrap(err)
}

// Catch3Ctx catches panics in a "func(context.Context) (T1, T2, T3, error)" closure.
func Catch3Ctx[T1 any, T2 any, T3 any](
	ctx context.Context,
	f func(ctx context.Context) (T1, T2, T3, error),
) (outV1 T1, outV2 T2, outV3 T3, outErr error) {

	defer func() {
		if rErr := MaybeWrapRecover(recover()); rErr != nil {
			var t1 T1
			var t2 T2
			var t3 T3

			outV1 = t1
			outV2 = t2
			outV3 = t3
			outErr = rErr
		}
	}()

	out1, out2, out3, err := f(ctx)
	return out1, out2, out3, MaybeWrap(err)
}
