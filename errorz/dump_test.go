package errorz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

func TestSDump(t *testing.T) {
	g := NewWithT(t)

	g.Expect(errorz.SDump(nil)).To(Equal("[nil]"))
	g.Expect(errorz.SDump(fmt.Errorf("e"))).ToNot(BeEmpty())
	g.Expect(errorz.SDump(errorz.Errorf("e"))).ToNot(BeEmpty())
}
