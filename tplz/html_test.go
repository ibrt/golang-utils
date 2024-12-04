package tplz_test

import (
	htpl "html/template"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/tplz"
)

type HTMLSuite struct {
	// intentionally empty
}

func TestHTMLSuite(t *testing.T) {
	fixturez.RunSuite(t, &HTMLSuite{})
}

func (*HTMLSuite) TestExecuteHTML(g *WithT) {
	okTpl, err := htpl.New("").Parse("<html><body>{{ . }}</body></html>")
	g.Expect(err).To(Succeed())

	errTpl, err := htpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	g.Expect(tplz.ExecuteHTML(okTpl, "<a>&</a>")).
		To(Equal([]byte("<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n")))

	g.Expect(tplz.ExecuteHTML(errTpl, nil)).Error().
		To(MatchError(`html/template::1:12: no such template "x"`))
}

func (*HTMLSuite) TestMustExecuteHTML(g *WithT) {
	okTpl, err := htpl.New("").Parse("<html><body>{{ . }}</body></html>")
	g.Expect(err).To(Succeed())

	errTpl, err := htpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	g.Expect(func() {
		g.Expect(tplz.MustExecuteHTML(okTpl, "<a>&</a>")).
			To(Equal([]byte("<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustExecuteHTML(errTpl, nil)
	}).To(PanicWith(MatchError(`html/template::1:12: no such template "x"`)))
}

func (*HTMLSuite) TestParseAndExecuteHTML(g *WithT) {
	g.Expect(tplz.ParseAndExecuteHTML("<html><body>{{ . }}</body></html>", "<a>&</a>")).
		To(Equal([]byte("<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n")))

	g.Expect(tplz.ParseAndExecuteHTML("{{ bad }}", "<a>&</a>")).Error().
		To(MatchError(`template: :1: function "bad" not defined`))

	g.Expect(tplz.ParseAndExecuteHTML(`{{ template "x" }}`, nil)).Error().
		To(MatchError(`html/template::1:12: no such template "x"`))
}

func (*HTMLSuite) TestMustParseAndExecuteHTML(g *WithT) {
	g.Expect(func() {
		g.Expect(tplz.MustParseAndExecuteHTML("<html><body>{{ . }}</body></html>", "<a>&</a>")).
			To(Equal([]byte("<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustParseAndExecuteHTML("{{ bad }}", "<a>&</a>")
	}).To(PanicWith(MatchError(`template: :1: function "bad" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteHTML(`{{ template "x" }}`, nil)
	}).To(PanicWith(MatchError(`html/template::1:12: no such template "x"`)))
}
