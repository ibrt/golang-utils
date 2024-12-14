package outz_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/outz"
)

type OutputSuite struct {
	// intentionally empty
}

func TestOutputSuite(t *testing.T) {
	fixturez.RunSuite(t, &OutputSuite{})
}

func (*OutputSuite) TestOutputCapture(g *WithT) {
	defer outz.ResetOutputCapture()
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(
		outz.OutputSetupStandard,
		outz.GetOutputSetupFatihColor(true),
		outz.OutputSetupRodaineTable)

	g.Expect(fmt.Fprint(os.Stdout, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(os.Stderr, "<err>")).Error().To(Succeed())

	outBuf, errBuf := outz.MustEndOutputCapture()

	g.Expect(outBuf).To(Equal("<out>"))
	g.Expect(errBuf).To(Equal("<err>"))

	g.Expect(func() {
		outz.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestResetOutputCapture(_ *WithT) {
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(
		outz.OutputSetupStandard,
		outz.GetOutputSetupFatihColor(false),
		outz.OutputSetupRodaineTable)

	fmt.Println("ignored")
}
