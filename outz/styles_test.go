package outz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/outz"
)

type StylesSuite struct {
	// intentionally empty
}

func TestStylesSuite(t *testing.T) {
	fixturez.RunSuite(t, &StylesSuite{})
}

func (*StylesSuite) TestRestoreDefaultStyles(g *WithT) {
	defaultStyles := outz.DefaultStyles
	outz.DefaultStyles = nil
	outz.RestoreDefaultStyles()
	g.Expect(outz.DefaultStyles).To(Equal(defaultStyles))
}
