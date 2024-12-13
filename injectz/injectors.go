package injectz

import (
	"context"
)

// Injector injects modules into a [context.Context].
type Injector func(ctx context.Context) context.Context

// NewNoopInjector returns an [Injector] that does nothing.
func NewNoopInjector() Injector {
	return func(ctx context.Context) context.Context {
		return ctx
	}
}

// NewSingletonInjector returns a constant [Injector].
func NewSingletonInjector(contextKey, value interface{}) Injector {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, contextKey, value)
	}
}

// NewInjectors combines multiple [Injector] into a compound one.
func NewInjectors(injectors ...Injector) Injector {
	return func(ctx context.Context) context.Context {
		for _, injector := range injectors {
			ctx = injector(ctx)
		}

		return ctx
	}
}
