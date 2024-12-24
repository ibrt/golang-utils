package idz_test

import (
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/idz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestMustNewRandomUUID(g *WithT) {
	id := idz.MustNewRandomUUID()
	g.Expect(id).ToNot(BeEmpty())
	g.Expect(uuid.Parse(id)).Error().To(Succeed())
}

func (*Suite) TestIsValidUUID(g *WithT) {
	g.Expect(idz.IsValidUUID(idz.MustNewRandomUUID())).To(BeTrue())
	g.Expect(idz.IsValidUUID(idz.MustNewRandomUUID() + "x")).To(BeFalse())
	g.Expect(idz.IsValidUUID("")).To(BeFalse())
}
