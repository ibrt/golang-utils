//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -source ./fixtures.go -destination ./mocks.gen.go -package tinjectz

package tinjectz

import (
	"context"

	"github.com/ibrt/golang-utils/injectz"
)

// ContextKey represents a context key.
type ContextKey int

// Known context keys.
const (
	FirstContextKey  ContextKey = iota
	SecondContextKey ContextKey = iota
)

// Initializer allows to mock an Initializer func.
type Initializer interface {
	Initialize(ctx context.Context) (injectz.Injector, injectz.Releaser)
}

// Injector allows to mock an Injector func.
type Injector interface {
	Inject(context.Context) context.Context
}

// Releaser allows to mock a Releaser func.
type Releaser interface {
	Release()
}

// Closer allows to mock an io.Closer.
type Closer interface {
	Close() error
}
