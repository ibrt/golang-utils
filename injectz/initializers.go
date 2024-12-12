package injectz

import (
	"context"

	"github.com/ibrt/golang-utils/errorz"
)

// Initializer initializes a value, returning a corresponding Injector and Releaser.
type Initializer func(ctx context.Context) (Injector, Releaser)

// Bootstrap builds a list of initializers.
type Bootstrap struct {
	initializers []Initializer
}

// NewBootstrap initializes a new *Bootstrap.
func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		initializers: make([]Initializer, 0),
	}
}

// Add one or more initializers.
func (i *Bootstrap) Add(initializers ...Initializer) *Bootstrap {
	i.initializers = append(i.initializers, initializers...)
	return i
}

// AddGroup adds one or more initializers.
func (i *Bootstrap) AddGroup(initializers []Initializer) *Bootstrap {
	i.initializers = append(i.initializers, initializers...)
	return i
}

// MustInitialize the current initializers.
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
