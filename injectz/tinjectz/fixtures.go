//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -write_package_comment=false -source ./fixtures.go -destination ./mocks.gen.go -package tinjectz

// Package tinjectz provides test fixtures for the "injectz" package.
package tinjectz

import (
	"context"

	"github.com/ibrt/golang-utils/injectz"
)

// TestContextKeyA is a test type for context keys.
type TestContextKeyA int

// Provided values for [TestContextKeyA].
const (
	TestContextKeyA0 TestContextKeyA = iota
	TestContextKeyA1 TestContextKeyA = iota
	TestContextKeyA2 TestContextKeyA = iota
	TestContextKeyA3 TestContextKeyA = iota
	TestContextKeyA4 TestContextKeyA = iota
	TestContextKeyA5 TestContextKeyA = iota
	TestContextKeyA6 TestContextKeyA = iota
	TestContextKeyA7 TestContextKeyA = iota
	TestContextKeyA8 TestContextKeyA = iota
	TestContextKeyA9 TestContextKeyA = iota
)

// TestContextKeyB is a test type for context keys.
type TestContextKeyB int

// Provided values for [TestContextKeyA].
const (
	TestContextKeyB0 TestContextKeyB = iota
	TestContextKeyB1 TestContextKeyB = iota
	TestContextKeyB2 TestContextKeyB = iota
	TestContextKeyB3 TestContextKeyB = iota
	TestContextKeyB4 TestContextKeyB = iota
	TestContextKeyB5 TestContextKeyB = iota
	TestContextKeyB6 TestContextKeyB = iota
	TestContextKeyB7 TestContextKeyB = iota
	TestContextKeyB8 TestContextKeyB = iota
	TestContextKeyB9 TestContextKeyB = iota
)

// TestInitializer is a mock blueprint.
type TestInitializer interface {
	// Initialize provides a [injectz.Initializer] for test purposes.
	Initialize(ctx context.Context) (injectz.Injector, injectz.Releaser)
}

// TestInjector is a mock blueprint.
type TestInjector interface {
	// Inject provides a [injectz.Injector] for test purposes.
	Inject(ctx context.Context) context.Context
}

// TestReleaser is a mock blueprint.
type TestReleaser interface {
	// Release provides a [injectz.Releaser] for test purposes.
	Release()
}
