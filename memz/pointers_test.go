package memz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/memz"
)

type PointersSuite struct {
	// intentionally empty
}

func TestPointersSuite(t *testing.T) {
	fixturez.RunSuite(t, &PointersSuite{})
}

func (*PointersSuite) TestPtr(g *WithT) {
	v := ""
	g.Expect(memz.Ptr(v)).To(Equal(&v))

	v = "a"
	g.Expect(memz.Ptr(v)).To(Equal(&v))
}

func (*PointersSuite) TestPtrIfTrue(g *WithT) {
	v := "a"
	g.Expect(memz.PtrIfTrue(true, v)).To(Equal(&v))
	g.Expect(memz.PtrIfTrue(false, v)).To(BeNil())

	v = "a"
	g.Expect(memz.PtrIfTrue(true, v)).To(Equal(&v))
	g.Expect(memz.PtrIfTrue(false, v)).To(BeNil())
}

func (*PointersSuite) TestPtrZeroToNil(g *WithT) {
	v := ""
	g.Expect(memz.PtrZeroToNil(v)).To(BeNil())

	v = "a"
	g.Expect(memz.PtrZeroToNil(v)).To(Equal(&v))
}

func (*PointersSuite) TestPtrZeroToNilIfTrue(g *WithT) {
	v := ""
	g.Expect(memz.PtrZeroToNilIfTrue(true, v)).To(BeNil())
	g.Expect(memz.PtrZeroToNilIfTrue(false, v)).To(BeNil())

	v = "a"
	g.Expect(memz.PtrZeroToNilIfTrue(true, v)).To(Equal(&v))
	g.Expect(memz.PtrZeroToNilIfTrue(false, v)).To(BeNil())
}

func (*PointersSuite) TestValNilToZero(g *WithT) {
	g.Expect(memz.ValNilToZero(memz.Ptr(""))).To(BeEmpty())
	g.Expect(memz.ValNilToZero(memz.Ptr("a"))).To(Equal("a"))
	g.Expect(memz.ValNilToZero[string](nil)).To(BeEmpty())
}

func (*PointersSuite) TestValNilToDef(g *WithT) {
	g.Expect(memz.ValNilToDef(memz.Ptr(""), "d")).To(BeEmpty())
	g.Expect(memz.ValNilToDef(memz.Ptr("a"), "d")).To(Equal("a"))
	g.Expect(memz.ValNilToDef(nil, "d")).To(Equal("d"))
}

func (*PointersSuite) TestSlicePtr(g *WithT) {
	g.Expect(memz.SlicePtr[string](nil)).To(BeNil())
	g.Expect(memz.SlicePtr([]string{})).To(Equal([]*string{}))
	g.Expect(memz.SlicePtr([]string{"", "v"})).To(Equal([]*string{memz.Ptr(""), memz.Ptr("v")}))
}

func (*PointersSuite) TestSlicePtrZeroToNil(g *WithT) {
	g.Expect(memz.SlicePtrZeroToNil[string](nil)).To(BeNil())
	g.Expect(memz.SlicePtrZeroToNil([]string{})).To(Equal([]*string{}))
	g.Expect(memz.SlicePtrZeroToNil([]string{"", "v"})).To(Equal([]*string{nil, memz.Ptr("v")}))
}

func (*PointersSuite) TestSliceValNilToZero(g *WithT) {
	g.Expect(memz.SliceValNilToZero[string](nil)).To(BeNil())
	g.Expect(memz.SliceValNilToZero([]*string{})).To(Equal([]string{}))
	g.Expect(memz.SliceValNilToZero([]*string{nil, memz.Ptr(""), memz.Ptr("v")})).To(Equal([]string{"", "", "v"}))
}

func (*PointersSuite) TestSliceValNilToDef(g *WithT) {
	g.Expect(memz.SliceValNilToDef(nil, "d")).To(BeNil())
	g.Expect(memz.SliceValNilToDef([]*string{}, "d")).To(Equal([]string{}))
	g.Expect(memz.SliceValNilToDef([]*string{nil, memz.Ptr(""), memz.Ptr("v")}, "d")).To(Equal([]string{"d", "", "v"}))
}

func (*PointersSuite) TestMapPtr(g *WithT) {
	g.Expect(memz.MapPtr[int, string](nil)).To(BeNil())
	g.Expect(memz.MapPtr(map[int]string{})).To(Equal(map[int]*string{}))
	g.Expect(memz.MapPtr(map[int]string{1: "", 2: "v"})).To(Equal(map[int]*string{1: memz.Ptr(""), 2: memz.Ptr("v")}))
}

func (*PointersSuite) TestMapPtrZeroToNil(g *WithT) {
	g.Expect(memz.MapPtrZeroToNil[int, string](nil)).To(BeNil())
	g.Expect(memz.MapPtrZeroToNil(map[int]string{})).To(Equal(map[int]*string{}))
	g.Expect(memz.MapPtrZeroToNil(map[int]string{1: "", 2: "v"})).To(Equal(map[int]*string{1: nil, 2: memz.Ptr("v")}))
}

func (*PointersSuite) TestMapValNilToZero(g *WithT) {
	g.Expect(memz.MapValNilToZero[int, string](nil)).To(BeNil())
	g.Expect(memz.MapValNilToZero(map[int]*string{})).To(Equal(map[int]string{}))
	g.Expect(memz.MapValNilToZero(map[int]*string{1: nil, 2: memz.Ptr(""), 3: memz.Ptr("v")})).To(Equal(map[int]string{1: "", 2: "", 3: "v"}))
}

func (*PointersSuite) TestMapValNilToDef(g *WithT) {
	g.Expect(memz.MapValNilToDef[int](nil, "d")).To(BeNil())
	g.Expect(memz.MapValNilToDef(map[int]*string{}, "d")).To(Equal(map[int]string{}))
	g.Expect(memz.MapValNilToDef(map[int]*string{1: nil, 2: memz.Ptr(""), 3: memz.Ptr("v")}, "d")).To(Equal(map[int]string{1: "d", 2: "", 3: "v"}))
}
