package errorz_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
)

func TestCatch0(t *testing.T) {
	g := NewWithT(t)

	g.Expect(
		errorz.Catch0(func() error {
			return nil
		})).To(Succeed())

	g.Expect(
		errorz.Catch0(func() error {
			return errorz.Errorf("test error")
		})).To(
		MatchError("test error"))

	g.Expect(
		errorz.Catch0(func() error {
			panic(errorz.Errorf("test error"))
		})).To(
		MatchError("test error"))
}

func TestCatch0Ctx(t *testing.T) {
	g := NewWithT(t)

	g.Expect(
		errorz.Catch0Ctx(context.Background(), func(ctx context.Context) error {
			return nil
		})).To(Succeed())

	g.Expect(
		errorz.Catch0Ctx(context.Background(), func(ctx context.Context) error {
			return errorz.Errorf("test error")
		})).To(
		MatchError("test error"))

	g.Expect(
		errorz.Catch0Ctx(context.Background(), func(ctx context.Context) error {
			panic(errorz.Errorf("test error"))
		})).To(
		MatchError("test error"))
}

func TestCatch1(t *testing.T) {
	g := NewWithT(t)

	out, err := errorz.Catch1(func() (string, error) {
		return "test", nil
	})
	g.Expect(err).To(Succeed())
	g.Expect(out).To(Equal("test"))

	out, err = errorz.Catch1(func() (string, error) {
		return "", errorz.Errorf("test error")
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out).To(Equal(""))

	out, err = errorz.Catch1(func() (string, error) {
		panic(errorz.Errorf("test error"))
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out).To(Equal(""))
}

func TestCatch1Ctx(t *testing.T) {
	g := NewWithT(t)

	out, err := errorz.Catch1Ctx(context.Background(), func(ctx context.Context) (string, error) {
		return "test", nil
	})
	g.Expect(err).To(Succeed())
	g.Expect(out).To(Equal("test"))

	out, err = errorz.Catch1Ctx(context.Background(), func(ctx context.Context) (string, error) {
		return "", errorz.Errorf("test error")
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out).To(Equal(""))

	out, err = errorz.Catch1Ctx(context.Background(), func(ctx context.Context) (string, error) {
		panic(errorz.Errorf("test error"))
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out).To(Equal(""))
}

func TestCatch2(t *testing.T) {
	g := NewWithT(t)

	out1, out2, err := errorz.Catch2(func() (string, *int, error) {
		return "test", memz.Ptr(1), nil
	})
	g.Expect(err).To(Succeed())
	g.Expect(out1).To(Equal("test"))
	g.Expect(out2).To(Equal(memz.Ptr(1)))

	out1, out2, err = errorz.Catch2(func() (string, *int, error) {
		return "", nil, errorz.Errorf("test error")
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())

	out1, out2, err = errorz.Catch2(func() (string, *int, error) {
		panic(errorz.Errorf("test error"))
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())
}

func TestCatch2Ctx(t *testing.T) {
	g := NewWithT(t)

	out1, out2, err := errorz.Catch2Ctx(context.Background(), func(ctx context.Context) (string, *int, error) {
		return "test", memz.Ptr(1), nil
	})
	g.Expect(err).To(Succeed())
	g.Expect(out1).To(Equal("test"))
	g.Expect(out2).To(Equal(memz.Ptr(1)))

	out1, out2, err = errorz.Catch2Ctx(context.Background(), func(ctx context.Context) (string, *int, error) {
		return "", nil, errorz.Errorf("test error")
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())

	out1, out2, err = errorz.Catch2Ctx(context.Background(), func(ctx context.Context) (string, *int, error) {
		panic(errorz.Errorf("test error"))
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())
}

func TestCatch3(t *testing.T) {
	g := NewWithT(t)

	out1, out2, out3, err := errorz.Catch3(func() (string, *int, struct{}, error) {
		return "test", memz.Ptr(1), struct{}{}, nil
	})
	g.Expect(err).To(Succeed())
	g.Expect(out1).To(Equal("test"))
	g.Expect(out2).To(Equal(memz.Ptr(1)))
	g.Expect(out3).To(Equal(struct{}{}))

	out1, out2, out3, err = errorz.Catch3(func() (string, *int, struct{}, error) {
		return "", nil, struct{}{}, errorz.Errorf("test error")
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())
	g.Expect(out3).To(Equal(struct{}{}))

	out1, out2, out3, err = errorz.Catch3(func() (string, *int, struct{}, error) {
		panic(errorz.Errorf("test error"))
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())
	g.Expect(out3).To(Equal(struct{}{}))
}

func TestCatch3Ctx(t *testing.T) {
	g := NewWithT(t)

	out1, out2, out3, err := errorz.Catch3Ctx(context.Background(), func(ctx context.Context) (string, *int, struct{}, error) {
		return "test", memz.Ptr(1), struct{}{}, nil
	})
	g.Expect(err).To(Succeed())
	g.Expect(out1).To(Equal("test"))
	g.Expect(out2).To(Equal(memz.Ptr(1)))
	g.Expect(out3).To(Equal(struct{}{}))

	out1, out2, out3, err = errorz.Catch3Ctx(context.Background(), func(ctx context.Context) (string, *int, struct{}, error) {
		return "", nil, struct{}{}, errorz.Errorf("test error")
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())
	g.Expect(out3).To(Equal(struct{}{}))

	out1, out2, out3, err = errorz.Catch3Ctx(context.Background(), func(ctx context.Context) (string, *int, struct{}, error) {
		panic(errorz.Errorf("test error"))
	})
	g.Expect(err).To(MatchError("test error"))
	g.Expect(out1).To(Equal(""))
	g.Expect(out2).To(BeNil())
	g.Expect(out3).To(Equal(struct{}{}))
}
