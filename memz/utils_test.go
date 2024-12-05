package memz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/memz"
)

type UtilsSuite struct {
	// intentionally empty
}

func TestUtilsSuite(t *testing.T) {
	fixturez.RunSuite(t, &UtilsSuite{})
}

func (s *UtilsSuite) TestMin(g *WithT) {
	g.Expect(memz.Min(1)).To(BeNumerically("==", 1))
	g.Expect(memz.Min(1, 2)).To(BeNumerically("==", 1))
	g.Expect(memz.Min(1, 2, 3)).To(BeNumerically("==", 1))
	g.Expect(memz.Min(2, 2)).To(BeNumerically("==", 2))
	g.Expect(memz.Min(2, 2, 3)).To(BeNumerically("==", 2))
	g.Expect(memz.Min(3, 2)).To(BeNumerically("==", 2))
}

func (s *UtilsSuite) TestMax(g *WithT) {
	g.Expect(memz.Max(1)).To(BeNumerically("==", 1))
	g.Expect(memz.Max(1, 2)).To(BeNumerically("==", 2))
	g.Expect(memz.Max(1, 2, 3)).To(BeNumerically("==", 3))
	g.Expect(memz.Max(2, 2)).To(BeNumerically("==", 2))
	g.Expect(memz.Max(3, 2)).To(BeNumerically("==", 3))
}

func (*UtilsSuite) TestTernary(g *WithT) {
	g.Expect(memz.Ternary(true, 1, 2)).To(BeNumerically("==", 1))
	g.Expect(memz.Ternary(false, 1, 2)).To(BeNumerically("==", 2))
}

func (s *UtilsSuite) TestIsAnyNil(g *WithT) {
	g.Expect(memz.IsAnyNil(nil)).To(BeTrue())
	g.Expect(memz.IsAnyNil(memz.PtrZeroToNil(""))).To(BeTrue())
	g.Expect(memz.IsAnyNil((*WithT)(nil))).To(BeTrue())
	g.Expect(memz.IsAnyNil(g)).To(BeFalse())
	g.Expect(memz.IsAnyNil("")).To(BeFalse())
	g.Expect(memz.IsAnyNil(1)).To(BeFalse())
}

func (s *UtilsSuite) TestPredicateIsZeroValue(g *WithT) {
	g.Expect(memz.FilterSlice([]int{0, 1, 2}, memz.PredicateIsZeroValue[int])).To(Equal([]int{0}))
	g.Expect(memz.FilterSlice([]string{"", "1", "2"}, memz.PredicateIsZeroValue[string])).To(Equal([]string{""}))
}

func (s *UtilsSuite) TestTransformSprintf(g *WithT) {
	g.Expect(memz.TransformSprintfV(1)).To(Equal("1"))
	g.Expect(memz.TransformSprintfKV(0, 1)).To(Equal("1"))
}
