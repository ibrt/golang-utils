package memz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/memz"
)

type SlicesSuite struct {
	// intentionally empty
}

func TestSlicesSuite(t *testing.T) {
	fixturez.RunSuite(t, &SlicesSuite{})
}

func (*SlicesSuite) TestSafeSliceIndexZero(g *WithT) {
	s := []string{"a", "b", "c"}

	g.Expect(memz.SafeSliceIndexZero[string](nil, 0)).To(BeEmpty())
	g.Expect(memz.SafeSliceIndexZero(s, 0)).To(Equal("a"))
	g.Expect(memz.SafeSliceIndexZero(s, 1)).To(Equal("b"))
	g.Expect(memz.SafeSliceIndexZero(s, 2)).To(Equal("c"))
	g.Expect(memz.SafeSliceIndexZero(s, -3)).To(Equal("a"))
	g.Expect(memz.SafeSliceIndexZero(s, -2)).To(Equal("b"))
	g.Expect(memz.SafeSliceIndexZero(s, -1)).To(Equal("c"))
	g.Expect(memz.SafeSliceIndexZero(s, 3)).To(BeEmpty())
	g.Expect(memz.SafeSliceIndexZero(s, 4)).To(BeEmpty())
}

func (*SlicesSuite) TestSafeSliceIndexDef(g *WithT) {
	s := []string{"a", "b", "c"}

	g.Expect(memz.SafeSliceIndexDef(nil, 0, "d")).To(Equal("d"))
	g.Expect(memz.SafeSliceIndexDef(s, 0, "d")).To(Equal("a"))
	g.Expect(memz.SafeSliceIndexDef(s, 1, "d")).To(Equal("b"))
	g.Expect(memz.SafeSliceIndexDef(s, 2, "d")).To(Equal("c"))
	g.Expect(memz.SafeSliceIndexDef(s, -3, "d")).To(Equal("a"))
	g.Expect(memz.SafeSliceIndexDef(s, -2, "d")).To(Equal("b"))
	g.Expect(memz.SafeSliceIndexDef(s, -1, "d")).To(Equal("c"))
	g.Expect(memz.SafeSliceIndexDef(s, 3, "d")).To(Equal("d"))
	g.Expect(memz.SafeSliceIndexDef(s, 4, "d")).To(Equal("d"))
}

func (*SlicesSuite) TestSafeSliceIndexPtr(g *WithT) {
	s := []string{"a", "b", "c"}

	g.Expect(memz.SafeSliceIndexPtr[string](nil, 0)).To(BeNil())
	g.Expect(memz.SafeSliceIndexPtr(s, 0)).To(Equal(memz.Ptr("a")))
	g.Expect(memz.SafeSliceIndexPtr(s, 1)).To(Equal(memz.Ptr("b")))
	g.Expect(memz.SafeSliceIndexPtr(s, 2)).To(Equal(memz.Ptr("c")))
	g.Expect(memz.SafeSliceIndexPtr(s, -3)).To(Equal(memz.Ptr("a")))
	g.Expect(memz.SafeSliceIndexPtr(s, -2)).To(Equal(memz.Ptr("b")))
	g.Expect(memz.SafeSliceIndexPtr(s, -1)).To(Equal(memz.Ptr("c")))
	g.Expect(memz.SafeSliceIndexPtr(s, 3)).To(BeNil())
	g.Expect(memz.SafeSliceIndexPtr(s, 4)).To(BeNil())
}

func (s *SlicesSuite) TestConcatSlices(g *WithT) {
	g.Expect(memz.ConcatSlices[string](nil, nil)).To(Equal([]string{}))
	g.Expect(memz.ConcatSlices[string](nil, []string{})).To(Equal([]string{}))
	g.Expect(memz.ConcatSlices[string]([]string{}, nil)).To(Equal([]string{}))
	g.Expect(memz.ConcatSlices[string]([]string{}, []string{})).To(Equal([]string{}))
	g.Expect(memz.ConcatSlices[string](nil, []string{"a"})).To(Equal([]string{"a"}))
	g.Expect(memz.ConcatSlices[string]([]string{}, []string{"a"})).To(Equal([]string{"a"}))
	g.Expect(memz.ConcatSlices[string]([]string{"a"}, nil)).To(Equal([]string{"a"}))
	g.Expect(memz.ConcatSlices[string]([]string{"a"}, []string{})).To(Equal([]string{"a"}))
	g.Expect(memz.ConcatSlices[string]([]string{"a"}, []string{"b"})).To(Equal([]string{"a", "b"}))
	g.Expect(memz.ConcatSlices[string]([]string{"b"}, []string{"a"})).To(Equal([]string{"b", "a"}))
}

func (*SlicesSuite) TestShallowCopySlice(g *WithT) {
	g.Expect(memz.ShallowCopySlice[string](nil)).To(BeNil())
	g.Expect(memz.ShallowCopySlice([]string{})).To(Equal([]string{}))
	g.Expect(memz.ShallowCopySlice([]string{"a", "b"})).To(Equal([]string{"a", "b"}))
}

func (s *SlicesSuite) TestFilterSlice(g *WithT) {
	g.Expect(memz.FilterSlice(nil, func(i int) bool {
		return true
	})).To(BeNil())

	g.Expect(memz.FilterSlice([]int{}, func(i int) bool {
		return true
	})).To(Equal([]int{}))

	g.Expect(memz.FilterSlice([]int{1, 2, 3, 4}, func(i int) bool {
		return i%2 == 0
	})).To(Equal([]int{2, 4}))
}

func (s *SlicesSuite) TestBatchSlice(g *WithT) {
	g.Expect(memz.BatchSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)).
		To(Equal([][]int{
			{0, 1},
			{2, 3},
			{4, 5},
			{6, 7},
			{8, 9},
		}))

	g.Expect(memz.BatchSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 2)).
		To(Equal([][]int{
			{0, 1},
			{2, 3},
			{4, 5},
			{6, 7},
			{8},
		}))

	g.Expect(memz.BatchSlice([]int{0, 1}, 2)).
		To(Equal([][]int{
			{0, 1},
		}))

	g.Expect(memz.BatchSlice([]int{0}, 2)).
		To(Equal([][]int{
			{0},
		}))

	g.Expect(memz.BatchSlice([]int{}, 2)).
		To(Equal([][]int{}))
}

func (*SlicesSuite) TestTransformSlice(g *WithT) {
	g.Expect(memz.TransformSlice(nil, memz.TransformSprintfV[int])).To(BeNil())
	g.Expect(memz.TransformSlice([]int{}, memz.TransformSprintfV)).To(Equal([]string{}))
	g.Expect(memz.TransformSlice([]int{1, 2, 3}, memz.TransformSprintfV)).To(Equal([]string{"1", "2", "3"}))
}
