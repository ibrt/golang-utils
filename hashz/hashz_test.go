package hashz_test

import (
	"crypto/md5"
	"crypto/sha1"
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
	g.Expect(hashz.MustHash(md5.New(), []byte("test"))).To(Equal("098f6bcd4621d373cade4e832627b4f6"))
	g.Expect(hashz.MustHashMD5([]byte("test"))).To(Equal("098f6bcd4621d373cade4e832627b4f6"))
	g.Expect(hashz.MustHash(sha256.New(), []byte("test"))).To(Equal("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
	g.Expect(hashz.MustHashSHA256([]byte("test"))).To(Equal("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
	g.Expect(hashz.MustHash(sha1.New(), []byte("test"))).To(Equal("a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"))
	g.Expect(hashz.MustHashSHA1([]byte("test"))).To(Equal("a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"))
	g.Expect(hashz.MustHash(fnv.New128(), []byte("test"))).To(Equal("66ab2a8b6f757277b806e89c56faf339"))
	g.Expect(hashz.MustHashFNV1128([]byte("test"))).To(Equal("66ab2a8b6f757277b806e89c56faf339"))
}
