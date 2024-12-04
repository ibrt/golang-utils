package stringz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/stringz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestEnsurePrefix(g *WithT) {
	g.Expect(stringz.EnsurePrefix("ab", "a")).To(Equal("ab"))
	g.Expect(stringz.EnsurePrefix("b", "a")).To(Equal("ab"))
}

func (*Suite) TestEnsureSuffix(g *WithT) {
	g.Expect(stringz.EnsureSuffix("ab", "b")).To(Equal("ab"))
	g.Expect(stringz.EnsureSuffix("a", "b")).To(Equal("ab"))
}
