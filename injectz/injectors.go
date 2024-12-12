package injectz

import (
	"context"
)

// Injector injects values into a Context.
type Injector func(ctx context.Context) context.Context

// NewNoopInjector returns an injector that does nothing.
func NewNoopInjector() Injector {
	return func(ctx context.Context) context.Context {
		return ctx
	}
}

// NewSingletonInjector always injects the given value using the given context key.
func NewSingletonInjector(contextKey, value interface{}) Injector {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, contextKey, value)
	}
}

// NewInjectors combines multiple injectors into one.
func NewInjectors(injectors ...Injector) Injector {
	return func(ctx context.Context) context.Context {
		for _, injector := range injectors {
			ctx = injector(ctx)
		}

		return ctx
	}
}
