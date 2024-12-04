package errorz_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

type structError struct {
	k string
}

func (e *structError) Error() string {
	return e.k
}

type stringError string

func (e stringError) Error() string {
	return string(e)
}

func TestWrap(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}
	e3 := stringError("o2")

	err := errorz.Wrap(e1, e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))
	g.Expect(err.(errorz.UnwrapMulti).Unwrap()).To(HaveExactElements(e2))

	err = errorz.Wrap(err, nil, e3)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o2: o1: e"))
	g.Expect(err.(errorz.UnwrapMulti).Unwrap()).To(HaveExactElements(e2, e3))

	g.Expect(func() { _ = errorz.Wrap(nil) }).To(PanicWith(MatchError("err is nil")))
}

func TestMaybeWrap(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}

	err := errorz.MaybeWrap(e1, e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))

	err = errorz.MaybeWrap(nil)
	g.Expect(err).To(BeNil())
}

func TestMustWrap(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}

	g.Expect(func() { errorz.MustWrap(e1, e2) }).To(PanicWith(MatchError("o1: e")))
	g.Expect(func() { errorz.MustWrap(nil) }).To(PanicWith(MatchError("err is nil")))
}

func TestMaybeMustWrap(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}

	g.Expect(func() { errorz.MaybeMustWrap(e1, e2) }).To(PanicWith(MatchError("o1: e")))
	g.Expect(func() { errorz.MaybeMustWrap(nil) }).ToNot(Panic())
}

func TestWrapRecover(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}

	err := errorz.WrapRecover(e1, e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))

	err = errorz.WrapRecover("e", e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))

	g.Expect(func() { _ = errorz.WrapRecover(nil) }).To(PanicWith(MatchError("r is nil")))
}

func TestMaybeWrapRecover(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}

	err := errorz.MaybeWrapRecover(e1, e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))

	err = errorz.MaybeWrapRecover("e", e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))

	g.Expect(func() { _ = errorz.MaybeWrapRecover(nil) }).ToNot(Panic())
}

func TestErrorsIsInterop(t *testing.T) {
	g := NewWithT(t)

	e1 := &structError{k: "e"}
	e2 := fmt.Errorf("o1")
	e3 := stringError("o2")
	ew := fmt.Errorf("x: %w", e1)

	err := errorz.Wrap(e1, e2, e3)
	g.Expect(errors.Is(err, e1)).To(BeTrue())
	g.Expect(errors.Is(err, e2)).To(BeTrue())
	g.Expect(errors.Is(err, e3)).To(BeTrue())
	g.Expect(errors.Is(err, &structError{k: "e"})).To(BeFalse()) // errors.Is requires "==" true or an "Is() bool" method returning true
	g.Expect(errors.Is(err, fmt.Errorf("o1"))).To(BeFalse())
	g.Expect(errors.Is(err, stringError("o2"))).To(BeTrue())
	g.Expect(errors.Is(err, stringError(""))).To(BeFalse())

	err = errorz.Wrap(ew, e2, e3)
	g.Expect(errors.Is(err, e1)).To(BeTrue())
	g.Expect(errors.Is(err, e2)).To(BeTrue())
	g.Expect(errors.Is(err, e3)).To(BeTrue())
	g.Expect(errors.Is(err, &structError{k: "e"})).To(BeFalse()) // errors.Is requires "==" true or an "Is() bool" method returning true
	g.Expect(errors.Is(err, fmt.Errorf("o1"))).To(BeFalse())
	g.Expect(errors.Is(err, stringError("o2"))).To(BeTrue())
	g.Expect(errors.Is(err, stringError(""))).To(BeFalse())

	err = errorz.Wrap(e2, ew, e3)
	g.Expect(errors.Is(err, e1)).To(BeTrue())
	g.Expect(errors.Is(err, e2)).To(BeTrue())
	g.Expect(errors.Is(err, e3)).To(BeTrue())
	g.Expect(errors.Is(err, &structError{k: "v"})).To(BeFalse()) // errors.Is requires "==" true or an "Is() bool" method returning true
	g.Expect(errors.Is(err, stringError(""))).To(BeFalse())
	g.Expect(errors.Is(err, fmt.Errorf("o1"))).To(BeFalse())
}

func TestErrorsAsInterop(t *testing.T) {
	g := NewWithT(t)

	e1 := &structError{k: "e"}
	e2 := fmt.Errorf("o1")
	e3 := stringError("o2")
	ew := fmt.Errorf("x: %w", e1)

	err := errorz.Wrap(e1, e2, e3)
	{
		var e *structError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(e1))
	}
	{
		var e stringError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(e3))
	}

	err = errorz.Wrap(ew, e2, e3)
	{
		var e *structError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(e1))
	}
	{
		var e stringError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(e3))
	}

	err = errorz.Wrap(e2, ew, e3)
	{
		var e *structError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(e1))
	}
	{
		var e stringError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(e3))
	}
}

func TestAs(t *testing.T) {
	g := NewWithT(t)

	e1 := &structError{k: "e"}
	e2 := fmt.Errorf("o1")
	e3 := stringError("o2")
	ew := fmt.Errorf("x: %w", e1)

	err := errorz.Wrap(e1, e2, e3)
	{
		e, ok := errorz.As[*structError](err)
		g.Expect(ok).To(BeTrue())
		g.Expect(e).To(Equal(e1))
	}
	{
		e, ok := errorz.As[stringError](err)
		g.Expect(ok).To(BeTrue())
		g.Expect(e).To(Equal(e3))
	}

	err = errorz.Wrap(ew, e2, e3)
	{
		e, ok := errorz.As[*structError](err)
		g.Expect(ok).To(BeTrue())
		g.Expect(e).To(Equal(e1))
	}
	{
		e, ok := errorz.As[stringError](err)
		g.Expect(ok).To(BeTrue())
		g.Expect(e).To(Equal(e3))
	}

	err = errorz.Wrap(e2, ew, e3)
	{
		e, ok := errorz.As[*structError](err)
		g.Expect(ok).To(BeTrue())
		g.Expect(e).To(Equal(e1))
	}
	{
		e, ok := errorz.As[stringError](err)
		g.Expect(ok).To(BeTrue())
		g.Expect(e).To(Equal(e3))
	}
}
