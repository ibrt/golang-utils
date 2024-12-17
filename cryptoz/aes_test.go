package cryptoz_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/cryptoz"
	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/idz"
)

type AESSuite struct {
	// intentionally empty
}

func TestAESSuite(t *testing.T) {
	fixturez.RunSuite(t, &AESSuite{})
}

func (*AESSuite) TestAES(g *WithT) {
	key := cryptoz.MustGenerateRandomAESKey()

	for i := 0; i < 100; i++ {
		plaintext := idz.MustNewRandomUUID()
		ciphertext := cryptoz.MustEncryptAES(key, plaintext)
		g.Expect(strings.HasPrefix(ciphertext, "enc:")).To(BeTrue())
		g.Expect(cryptoz.MustDecryptAES(key, ciphertext)).To(Equal(plaintext))
	}
}
