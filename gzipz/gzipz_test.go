package gzipz_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/gzipz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestCompression(g *WithT) {
	var gBuf, uBuf []byte
	buf := []byte(strings.Repeat("x", 1000))

	g.Expect(func() {
		gBuf = gzipz.MustCompress(buf)
	}).ToNot(Panic())

	g.Expect(len(gBuf)).To(BeNumerically("<", len(buf)))

	g.Expect(func() {
		uBuf = gzipz.MustDecompress(gBuf)
	}).ToNot(Panic())

	g.Expect(uBuf).To(Equal(buf))
}
