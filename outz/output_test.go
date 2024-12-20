package outz_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fatih/color"
	. "github.com/onsi/gomega"
	"github.com/rodaine/table"
	"github.com/sirupsen/logrus"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/outz"
)

type OutputSuite struct {
	// intentionally empty
}

func TestOutputSuite(t *testing.T) {
	fixturez.RunSuite(t, &OutputSuite{})
}

func (*OutputSuite) TestOutputCapture_Standard(g *WithT) {
	defer outz.ResetOutputCapture()
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(outz.OutputSetupStandard)

	g.Expect(fmt.Fprint(os.Stdout, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(os.Stderr, "<err>")).Error().To(Succeed())

	outBuf, errBuf := outz.MustEndOutputCapture()

	g.Expect(outBuf).To(Equal("<out>"))
	g.Expect(errBuf).To(Equal("<err>"))

	g.Expect(func() {
		outz.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestOutputCapture_FatihColor_False(g *WithT) {
	defer outz.ResetOutputCapture()
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(outz.GetOutputSetupFatihColor(false))

	g.Expect(outz.DefaultStyles.Default().Printf("<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(color.Output, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(color.Error, "<err>")).Error().To(Succeed())

	outBuf, errBuf := outz.MustEndOutputCapture()

	g.Expect(outBuf).To(Equal("\x1b[0m<out>\x1b[0m<out>"))
	g.Expect(errBuf).To(Equal("<err>"))

	g.Expect(func() {
		outz.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestOutputCapture_FatihColor_True(g *WithT) {
	defer outz.ResetOutputCapture()
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(outz.GetOutputSetupFatihColor(true))

	g.Expect(outz.DefaultStyles.Default().Printf("<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(color.Output, "<out>")).Error().To(Succeed())
	g.Expect(fmt.Fprint(color.Error, "<err>")).Error().To(Succeed())

	outBuf, errBuf := outz.MustEndOutputCapture()

	g.Expect(outBuf).To(Equal("<out><out>"))
	g.Expect(errBuf).To(Equal("<err>"))

	g.Expect(func() {
		outz.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestOutputCapture_RodaineTable(g *WithT) {
	defer outz.ResetOutputCapture()
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(outz.OutputSetupRodaineTable)

	table.New("h1").AddRow("v1").Print()
	g.Expect(fmt.Fprint(table.DefaultWriter, "<out>")).Error().To(Succeed())

	outBuf, errBuf := outz.MustEndOutputCapture()

	g.Expect(outBuf).To(Equal("h1  \nv1  \n<out>"))
	g.Expect(errBuf).To(BeEmpty())

	g.Expect(func() {
		outz.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestOutputCapture_SirupsenLogrus(g *WithT) {
	defer outz.ResetOutputCapture()
	defer outz.ResetOutputCapture()

	t := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	l1 := outz.NewLogger()

	outz.MustBeginOutputCapture(outz.OutputSetupSirupsenLogrus)

	l2 := outz.NewLogger()

	logrus.WithTime(t).Print("<err>")
	logrus.StandardLogger().WithTime(t).Print("<err>")
	l1.WithTime(t).Print("<err>")
	l2.WithTime(t).Print("<err>")

	outBuf, errBuf := outz.MustEndOutputCapture()

	g.Expect(outBuf).To(BeEmpty())
	g.Expect(errBuf).To(Equal("time=\"2000-01-01T00:00:00Z\" level=info msg=\"<err>\"\ntime=\"2000-01-01T00:00:00Z\" level=info msg=\"<err>\"\ntime=\"2000-01-01T00:00:00Z\" level=info msg=\"<err>\"\ntime=\"2000-01-01T00:00:00Z\" level=info msg=\"<err>\"\n"))

	g.Expect(func() {
		outz.ResetOutputCapture()
	}).ToNot(Panic())
}

func (*OutputSuite) TestResetOutputCapture(_ *WithT) {
	defer outz.ResetOutputCapture()

	outz.MustBeginOutputCapture(
		outz.OutputSetupStandard,
		outz.GetOutputSetupFatihColor(false),
		outz.OutputSetupRodaineTable,
		outz.OutputSetupSirupsenLogrus)

	fmt.Println("ignored")
}
