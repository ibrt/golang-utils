package tplz_test

import (
	"testing"
	ttpl "text/template"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/tplz"
)

type JSONSuite struct {
	// intentionally empty
}

func TestJSONSuite(t *testing.T) {
	fixturez.RunSuite(t, &JSONSuite{})
}

func (*JSONSuite) TestExecuteJSON(g *WithT) {
	okTpl, err := ttpl.New("").Parse(`{ "value": "{{ . }}" }`)
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	jsonErrTpl, err := ttpl.New("").Parse(`{ "value": {{ . }} }`)
	g.Expect(err).To(Succeed())

	g.Expect(tplz.ExecuteJSON(okTpl, "|", "  ", "Hello World")).
		To(Equal([]byte("{\n|  \"value\": \"Hello World\"\n|}")))

	g.Expect(tplz.ExecuteJSON(errTpl, "|", "  ", nil)).Error().
		To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))

	g.Expect(tplz.ExecuteJSON(jsonErrTpl, "|", "  ", "Hello World")).Error().
		To(MatchError("invalid character 'H' looking for beginning of value"))
}

func (*JSONSuite) TestMustExecuteJSON(g *WithT) {
	okTpl, err := ttpl.New("").Parse(`{ "value": "{{ . }}" }`)
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	jsonErrTpl, err := ttpl.New("").Parse(`{ "value": {{ . }} }`)
	g.Expect(err).To(Succeed())

	g.Expect(func() {
		g.Expect(tplz.MustExecuteJSON(okTpl, "|", "  ", "Hello World")).
			To(Equal([]byte("{\n|  \"value\": \"Hello World\"\n|}")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustExecuteJSON(errTpl, "|", "  ", nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))

	g.Expect(func() {
		tplz.MustExecuteJSON(jsonErrTpl, "|", "  ", "Hello World")
	}).To(PanicWith(MatchError("invalid character 'H' looking for beginning of value")))
}

func (*JSONSuite) TestParseAndExecuteJSON(g *WithT) {
	g.Expect(tplz.ParseAndExecuteJSON(`{ "value": "{{ . }}" }`, "|", "  ", "Hello World")).
		To(Equal([]byte("{\n|  \"value\": \"Hello World\"\n|}")))

	g.Expect(tplz.ParseAndExecuteJSON("{{ bad }}", "|", "  ", "Hello World")).Error().
		To(MatchError(`template: :1: function "bad" not defined`))

	g.Expect(tplz.ParseAndExecuteJSON(`{{ template "x" }}`, "|", "  ", nil)).Error().
		To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))

	g.Expect(tplz.ParseAndExecuteJSON(`{ "value": {{ . }} }`, "|", "  ", "Hello World")).Error().
		To(MatchError("invalid character 'H' looking for beginning of value"))
}

func (*JSONSuite) TestMustParseAndExecuteJSON(g *WithT) {
	g.Expect(func() {
		g.Expect(tplz.MustParseAndExecuteJSON(`{ "value": "{{ . }}" }`, "|", "  ", "Hello World")).
			To(Equal([]byte("{\n|  \"value\": \"Hello World\"\n|}")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustParseAndExecuteJSON("{{ bad }}", "|", "  ", "Hello World")
	}).To(PanicWith(MatchError(`template: :1: function "bad" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteJSON(`{{ template "x" }}`, "|", "  ", nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteJSON(`{ "value": {{ . }} }`, "|", "  ", "Hello World")
	}).To(PanicWith(MatchError("invalid character 'H' looking for beginning of value")))
}
