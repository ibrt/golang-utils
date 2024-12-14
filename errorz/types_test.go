package errorz

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	. "github.com/onsi/gomega"
)

type unwrapSingle struct {
	k   string
	err error
}

func (u *unwrapSingle) String() string {
	return u.k
}

func (u *unwrapSingle) Unwrap() error {
	return u.err
}

type unwrapMulti struct {
	k    string
	errs []error
}

func (u *unwrapMulti) String() string {
	return u.k
}

func (u *unwrapMulti) Unwrap() []error {
	return u.errs
}

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

func TestValueError(t *testing.T) {
	g := NewWithT(t)

	err := &valueError{
		Value: "test value",
	}
	g.Expect(err.Error()).To(Equal("test value"))
	g.Expect(err.Unwrap()).To(BeNil())

	err = &valueError{
		Value: &unwrapSingle{
			k: "v",
		},
	}
	g.Expect(err.Error()).To(Equal("v"))
	g.Expect(err.Unwrap()).To(BeNil())

	err = &valueError{
		Value: &unwrapSingle{
			k:   "v",
			err: fmt.Errorf("w"),
		},
	}
	g.Expect(err.Error()).To(Equal("v"))
	g.Expect(err.Unwrap()).To(HaveExactElements(fmt.Errorf("w")))

	err = &valueError{
		Value: &unwrapMulti{
			k: "v",
		},
	}
	g.Expect(err.Error()).To(Equal("v"))
	g.Expect(err.Unwrap()).To(BeNil())

	err = &valueError{
		Value: &unwrapMulti{
			k: "v",
			errs: []error{
				fmt.Errorf("w1"),
				fmt.Errorf("w2"),
			},
		},
	}
	g.Expect(err.Error()).To(Equal("v"))
	g.Expect(err.Unwrap()).To(HaveExactElements(fmt.Errorf("w1"), fmt.Errorf("w2")))

	g.Expect(func() { _ = (*valueError)(nil).Error() }).To(Panic())
	g.Expect((*valueError)(nil).Unwrap()).To(BeNil())
}

func TestWrappedError(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	err := &wrappedError{
		m:      &sync.Mutex{},
		errs:   []error{e1},
		frames: GetFrames(nil),
	}
	g.Expect(err.Error()).To(Equal("e"))
	g.Expect(err.Unwrap()).To(HaveExactElements(e1))
	g.Expect(errors.Is(err, e1)).To(BeTrue())
	g.Expect(errors.Is(err, fmt.Errorf("e"))).To(BeFalse())
	g.Expect(errors.Is(err, nil)).To(BeFalse())
	{
		var e stringError
		g.Expect(errors.As(err, &e)).To(BeFalse())
	}
	{
		var e *structError
		g.Expect(errors.As(err, &e)).To(BeFalse())
	}

	err = &wrappedError{
		m:      &sync.Mutex{},
		errs:   []error{stringError("e")},
		frames: GetFrames(nil),
	}
	g.Expect(err.Error()).To(Equal("e"))
	g.Expect(err.Unwrap()).To(HaveExactElements(stringError("e")))
	g.Expect(errors.Is(err, stringError("e"))).To(BeTrue())
	g.Expect(errors.Is(err, stringError(""))).To(BeFalse())
	g.Expect(errors.Is(err, fmt.Errorf("e"))).To(BeFalse())
	{
		var e stringError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(stringError("e")))
	}
	{
		var e *structError
		g.Expect(errors.As(err, &e)).To(BeFalse())
	}

	e2 := &structError{k: "e"}
	err = &wrappedError{
		m:      &sync.Mutex{},
		errs:   []error{e2},
		frames: GetFrames(nil),
	}
	g.Expect(err.Error()).To(Equal("e"))
	g.Expect(err.Unwrap()).To(HaveExactElements(e2))
	g.Expect(errors.Is(err, e2)).To(BeTrue())
	g.Expect(errors.Is(err, &structError{k: "e"})).To(BeFalse()) // errors.Is requires "==" true or an "Is() bool" method returning true
	g.Expect(errors.Is(err, stringError(""))).To(BeFalse())
	g.Expect(errors.Is(err, fmt.Errorf("e"))).To(BeFalse())
	{
		var e *structError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(&structError{k: "e"}))
	}
	{
		var e stringError
		g.Expect(errors.As(err, &e)).To(BeFalse())
	}

	e3 := &structError{k: "e"}
	e4 := fmt.Errorf("o1")
	e5 := stringError("o2")
	err = &wrappedError{
		m:      &sync.Mutex{},
		errs:   []error{e3, e4, e5},
		frames: GetFrames(nil),
	}
	g.Expect(err.Error()).To(Equal("o2: o1: e"))
	g.Expect(err.Unwrap()).To(HaveExactElements(e3, e4, e5))
	g.Expect(errors.Is(err, e3)).To(BeTrue())
	g.Expect(errors.Is(err, e4)).To(BeTrue())
	g.Expect(errors.Is(err, e5)).To(BeTrue())
	g.Expect(errors.Is(err, &structError{k: "e"})).To(BeFalse()) // errors.Is requires "==" true or an "Is() bool" method returning true
	g.Expect(errors.Is(err, stringError("o2"))).To(BeTrue())
	g.Expect(errors.Is(err, stringError(""))).To(BeFalse())
	g.Expect(errors.Is(err, fmt.Errorf("o1"))).To(BeFalse())
	{
		var e *structError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(&structError{k: "e"}))
	}
	{
		var e stringError
		as := errors.As(err, &e)
		g.Expect(as).To(BeTrue())
		g.Expect(e).To(Equal(stringError("o2")))
	}

	g.Expect(func() { _ = (*wrappedError)(nil).Error() }).To(Panic())
	g.Expect((*wrappedError)(nil).Unwrap()).To(BeNil())
}
