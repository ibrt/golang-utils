package hashz_test

import (
	"crypto/sha256"
	"hash/fnv"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/hashz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestMustHash(g *WithT) {
	g.Expect(hashz.MustHash(sha256.New(), []byte("test"))).To(Equal("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
	g.Expect(hashz.MustHashSHA256([]byte("test"))).To(Equal("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
	g.Expect(hashz.MustHash(fnv.New128(), []byte("test"))).To(Equal("66ab2a8b6f757277b806e89c56faf339"))
	g.Expect(hashz.MustHashFNV1128([]byte("test"))).To(Equal("66ab2a8b6f757277b806e89c56faf339"))
}
