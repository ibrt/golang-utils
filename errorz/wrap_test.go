package errorz_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

var (
	_ error = (*structError)(nil)
	_ error = (*stringError)(nil)
	_ error = (*wrapSingle)(nil)
	_ error = (*wrapMulti)(nil)

	_ errorz.UnwrapSingle = (*wrapSingle)(nil)
	_ errorz.UnwrapMulti  = (*wrapMulti)(nil)
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

type wrapSingle struct {
	err error
}

func (e *wrapSingle) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return "<empty>"
}

func (e *wrapSingle) Unwrap() error {
	return e.err
}

type wrapMulti struct {
	errs []error
}

func (e *wrapMulti) Error() string {
	if len(e.errs) > 0 {
		ms := make([]string, 0, len(e.errs))

		for _, err := range e.errs {
			ms = append(ms, err.Error())
		}

		return strings.Join(ms, ": ")
	}

	return "<empty>"
}

func (e *wrapMulti) Unwrap() []error {
	return e.errs
}

func TestWrap(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	e2 := &structError{k: "o1"}
	e3 := stringError("o2")

	err := errorz.Wrap(e1, e2)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o1: e"))
	g.Expect(err.(errorz.UnwrapMulti).Unwrap()).To(HaveExactElements(e1, e2))

	err = errorz.Wrap(err, nil, e3)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("o2: o1: e"))
	g.Expect(err.(errorz.UnwrapMulti).Unwrap()).To(HaveExactElements(e1, e2, e3))

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

func TestUnwrap(t *testing.T) {
	g := NewWithT(t)

	g.Expect(errorz.Unwrap(nil)).To(BeNil())
	g.Expect(errorz.Unwrap(&wrapSingle{})).To(BeNil())
	g.Expect(errorz.Unwrap(&wrapMulti{})).To(BeNil())
	g.Expect(errorz.Unwrap(fmt.Errorf("e1"))).To(BeNil())

	g.Expect(errorz.Unwrap(errors.Join(fmt.Errorf("e1"), fmt.Errorf("e2")))).
		To(HaveExactElements(fmt.Errorf("e1"), fmt.Errorf("e2")))

	g.Expect(errorz.Unwrap(fmt.Errorf("e2: %w", fmt.Errorf("e1")))).
		To(HaveExactElements(fmt.Errorf("e1")))
}
