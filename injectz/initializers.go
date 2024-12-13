package injectz

import (
	"context"

	"github.com/ibrt/golang-utils/errorz"
)

// Initializer initializes a module, returning a corresponding [Injector] and [Releaser].
type Initializer func(ctx context.Context) (Injector, Releaser)

// Bootstrap builds and manages a group of [Initializer].
type Bootstrap struct {
	initializers []Initializer
}

// NewBootstrap initializes a new [*Bootstrap].
func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		initializers: make([]Initializer, 0),
	}
}

// Add one or more [Initializer].
func (i *Bootstrap) Add(initializers ...Initializer) *Bootstrap {
	i.initializers = append(i.initializers, initializers...)
	return i
}

// MustInitialize runs all the [Initializer] in the group, returns a compound [Injector] and [Releaser].
func (i *Bootstrap) MustInitialize() (Injector, Releaser) {
	ctx := context.Background()
	injectors := make([]Injector, 0, len(i.initializers))
	releasers := make([]Releaser, 0, len(i.initializers))

	defer func() {
		if err := errorz.MaybeWrapRecover(recover()); err != nil {
			NewReleasers(releasers...)()
			errorz.MustWrap(err)
		}
	}()

	for _, initializer := range i.initializers {
		injector, releaser := initializer(ctx)
		injectors = append(injectors, injector)
		releasers = append(releasers, releaser)
		ctx = injector(ctx)
	}

	return NewInjectors(injectors...), NewReleasers(releasers...)
}
