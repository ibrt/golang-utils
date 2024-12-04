package numz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/numz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestParse(g *WithT) {
	g.Expect(numz.MustParse[float32]("1.1")).To(Equal(float32(1.1)))
	g.Expect(numz.MustParse[float64]("1.1")).To(Equal(1.1))
	g.Expect(numz.MustParse[int]("1")).To(Equal(1))
	g.Expect(numz.MustParse[int8]("1")).To(Equal(int8(1)))
	g.Expect(numz.MustParse[int16]("1")).To(Equal(int16(1)))
	g.Expect(numz.MustParse[int32]("1")).To(Equal(int32(1)))
	g.Expect(numz.MustParse[int64]("1")).To(Equal(int64(1)))
	g.Expect(numz.MustParse[uint]("1")).To(Equal(uint(1)))
	g.Expect(numz.MustParse[uint8]("1")).To(Equal(uint8(1)))
	g.Expect(numz.MustParse[uint16]("1")).To(Equal(uint16(1)))
	g.Expect(numz.MustParse[uint32]("1")).To(Equal(uint32(1)))
	g.Expect(numz.MustParse[uint64]("1")).To(Equal(uint64(1)))
	g.Expect(numz.MustParse[uintptr]("1")).To(Equal(uintptr(1)))

	g.Expect(numz.MustParse[int]("0b111")).To(Equal(7))
	g.Expect(numz.MustParse[int]("0o111")).To(Equal(73))
	g.Expect(numz.MustParse[int]("0x111")).To(Equal(273))

	g.Expect(numz.MustParse[uint]("0b111")).To(Equal(uint(7)))
	g.Expect(numz.MustParse[uint]("0o111")).To(Equal(uint(73)))
	g.Expect(numz.MustParse[uint]("0x111")).To(Equal(uint(273)))

	g.Expect(func() {
		numz.MustParse[float32]("bad")
	}).To(PanicWith(MatchError(`strconv.ParseFloat: parsing "bad": invalid syntax`)))

	g.Expect(func() {
		numz.MustParse[int]("bad")
	}).To(PanicWith(MatchError(`strconv.ParseInt: parsing "bad": invalid syntax`)))

	g.Expect(func() {
		numz.MustParse[uint]("bad")
	}).To(PanicWith(MatchError(`strconv.ParseUint: parsing "bad": invalid syntax`)))
}
