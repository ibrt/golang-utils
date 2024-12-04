package tplz_test

import (
	"testing"
	ttpl "text/template"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/tplz"
)

type GoSuite struct {
	// intentionally empty
}

func TestGoSuite(t *testing.T) {
	fixturez.RunSuite(t, &GoSuite{})
}

func (*GoSuite) TestExecuteGo(g *WithT) {
	okTpl, err := ttpl.New("").Parse("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	goErrTpl, err := ttpl.New("").Parse("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	g.Expect(tplz.ExecuteGo(okTpl, "Hello World")).
		To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))

	g.Expect(tplz.ExecuteGo(errTpl, nil)).Error().
		To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))

	g.Expect(tplz.ExecuteGo(goErrTpl, "Hello World")).Error().
		To(MatchError("2:1: expected declaration, found funcmain"))
}

func (*GoSuite) TestMustExecuteGo(g *WithT) {
	okTpl, err := ttpl.New("").Parse("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	goErrTpl, err := ttpl.New("").Parse("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	g.Expect(func() {
		g.Expect(tplz.MustExecuteGo(okTpl, "Hello World")).
			To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustExecuteGo(errTpl, nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))

	g.Expect(func() {
		tplz.MustExecuteGo(goErrTpl, "Hello World")
	}).To(PanicWith(MatchError("2:1: expected declaration, found funcmain")))
}

func (*GoSuite) TestParseAndExecuteGo(g *WithT) {
	g.Expect(tplz.ParseAndExecuteGo("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }", "Hello World")).
		To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))

	g.Expect(tplz.ParseAndExecuteGo("{{ bad }}", "Hello World")).Error().
		To(MatchError(`template: :1: function "bad" not defined`))

	g.Expect(tplz.ParseAndExecuteGo(`{{ template "x" }}`, nil)).Error().
		To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))

	g.Expect(tplz.ParseAndExecuteGo("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }", "Hello World")).Error().
		To(MatchError("2:1: expected declaration, found funcmain"))
}

func (*GoSuite) TestMustParseAndExecuteGo(g *WithT) {
	g.Expect(func() {
		g.Expect(tplz.MustParseAndExecuteGo("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }", "Hello World")).
			To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustParseAndExecuteGo("{{ bad }}", "Hello World")
	}).To(PanicWith(MatchError(`template: :1: function "bad" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteGo(`{{ template "x" }}`, nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteGo("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }", "Hello World")
	}).To(PanicWith(MatchError("2:1: expected declaration, found funcmain")))
}
