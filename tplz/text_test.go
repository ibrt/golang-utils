package tplz_test

import (
	"testing"
	ttpl "text/template"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/tplz"
)

type TextSuite struct {
	// intentionally empty
}

func TestTextSuite(t *testing.T) {
	fixturez.RunSuite(t, &TextSuite{})
}

func (*TextSuite) TestExecuteText(g *WithT) {
	okTpl, err := ttpl.New("").Parse("{{ . }}")
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	g.Expect(tplz.ExecuteText(okTpl, "<a>&</a>")).To(Equal([]byte("<a>&</a>")))
	g.Expect(tplz.ExecuteText(errTpl, nil)).Error().To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))
}

func (*TextSuite) TestMustExecuteText(g *WithT) {
	okTpl, err := ttpl.New("").Parse("{{ . }}")
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	g.Expect(func() {
		g.Expect(tplz.MustExecuteText(okTpl, "<a>&</a>")).To(Equal([]byte("<a>&</a>")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustExecuteText(errTpl, nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))
}

func (*TextSuite) TestParseAndExecuteText(g *WithT) {
	g.Expect(tplz.ParseAndExecuteText("{{ . }}", "<a>&</a>")).To(Equal([]byte("<a>&</a>")))

	g.Expect(tplz.ParseAndExecuteText("{{ bad }}", "<a>&</a>")).Error().To(
		MatchError(`template: :1: function "bad" not defined`))

	g.Expect(tplz.ParseAndExecuteText(`{{ template "x" }}`, nil)).Error().To(
		MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))
}

func (*TextSuite) TestMustParseAndExecuteText(g *WithT) {
	g.Expect(func() {
		g.Expect(tplz.MustParseAndExecuteText("{{ . }}", "<a>&</a>")).To(Equal([]byte("<a>&</a>")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustParseAndExecuteText("{{ bad }}", "<a>&</a>")
	}).To(PanicWith(MatchError(`template: :1: function "bad" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteText(`{{ template "x" }}`, nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))
}
