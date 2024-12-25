package errorz

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"

	. "github.com/onsi/gomega"
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

func TestValueError(t *testing.T) {
	g := NewWithT(t)

	err := &valueError{
		value: "test value",
	}
	g.Expect(err.GetErrorName()).To(Equal("value-error"))
	g.Expect(err.GetErrorDetails()).To(Equal(map[string]any{"value": "test value"}))
	g.Expect(err.Unwrap()).To(Succeed())
	g.Expect(err.Error()).To(Equal("test value"))

	err = &valueError{
		value: fmt.Errorf("test error"),
	}
	g.Expect(err.GetErrorName()).To(Equal("value-error"))
	g.Expect(err.GetErrorDetails()).To(Equal(map[string]any{"value": fmt.Errorf("test error")}))
	g.Expect(err.Unwrap()).To(Equal(fmt.Errorf("test error")))
	g.Expect(err.Error()).To(Equal("test error"))

	g.Expect(func() { _ = (*valueError)(nil).Error() }).To(Panic())
	g.Expect((*valueError)(nil).Unwrap()).To(Succeed())
}

func TestWrappedError(t *testing.T) {
	g := NewWithT(t)

	e1 := fmt.Errorf("e")
	err := &wrappedError{
		m:        &sync.Mutex{},
		errs:     []error{e1},
		frames:   GetFrames(nil),
		metadata: map[any]any{},
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
	{
		type mk1 int
		type mk2 int

		const k1 mk1 = 0
		const k2 mk2 = 0

		v, ok := err.getMetadata(k1)
		g.Expect(ok).To(BeFalse())
		g.Expect(v).To(BeNil())

		err.setMetadata(k1, 1)
		err.setMetadata(k2, 2)

		v, ok = err.getMetadata(k1)
		g.Expect(ok).To(BeTrue())
		g.Expect(v).To(Equal(1))

		v, ok = err.getMetadata(k2)
		g.Expect(ok).To(BeTrue())
		g.Expect(v).To(Equal(2))
	}

	err = &wrappedError{
		m:        &sync.Mutex{},
		errs:     []error{stringError("e")},
		frames:   GetFrames(nil),
		metadata: map[any]any{},
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
		m:        &sync.Mutex{},
		errs:     []error{e2},
		frames:   GetFrames(nil),
		metadata: map[any]any{},
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
		m:        &sync.Mutex{},
		errs:     []error{e3, e4, e5},
		frames:   GetFrames(nil),
		metadata: map[any]any{},
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

func TestGenericErrors(t *testing.T) {
	g := NewWithT(t)

	g.Expect(genericErrorsErrorStringName).To(Equal("*errors.errorString"))
	g.Expect(genericErrorsJoinErrorName).To(Equal("*errors.joinError"))
	g.Expect(genericFmtWrapErrorName).To(Equal("*fmt.wrapError"))
	g.Expect(genericFmtWrapErrorsName).To(Equal("*fmt.wrapErrors"))
}

func TestIsGenericError(t *testing.T) {
	g := NewWithT(t)

	g.Expect(isGenericError(nil)).To(BeFalse())
	g.Expect(isGenericError(fmt.Errorf("e"))).To(BeTrue())
	g.Expect(isGenericError(errors.Join(fmt.Errorf("e")))).To(BeTrue())
	g.Expect(isGenericError(fmt.Errorf("%w", fmt.Errorf("e")))).To(BeTrue())
	g.Expect(isGenericError(fmt.Errorf("%w%w", fmt.Errorf("e"), fmt.Errorf("e")))).To(BeTrue())
	g.Expect(isGenericError(&os.PathError{})).To(BeFalse())
}

func TestIsJoinError(t *testing.T) {
	g := NewWithT(t)

	g.Expect(isJoinError(nil)).To(BeFalse())
	g.Expect(isJoinError(fmt.Errorf("e"))).To(BeFalse())
	g.Expect(isJoinError(errors.Join(fmt.Errorf("e")))).To(BeTrue())
}

func TestIsWrapError(t *testing.T) {
	g := NewWithT(t)

	g.Expect(isWrapError(nil)).To(BeFalse())
	g.Expect(isWrapError(fmt.Errorf("e"))).To(BeFalse())
	g.Expect(isWrapError(Errorf("e"))).To(BeTrue())
}
