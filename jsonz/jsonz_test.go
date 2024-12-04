package jsonz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/jsonz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestMustMarshal(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshal(map[string]int{"a": 1})).To(Equal([]byte(`{"a":1}`)))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshal(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalString(map[string]int{"a": 1})).To(Equal(`{"a":1}`))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalString(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalPretty(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalPretty(map[string]int{"a": 1})).To(Equal([]byte("{\n  \"a\": 1\n}")))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalPretty(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalPrettyString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalPrettyString(map[string]int{"a": 1})).To(Equal("{\n  \"a\": 1\n}"))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalPrettyString(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestUnmarshal(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(jsonz.Unmarshal[*testStruct]([]byte(`{"k": "v"}`))).
		To(Equal(&testStruct{K: "v"}))

	g.Expect(jsonz.Unmarshal[*testStruct]([]byte(`bad`))).Error().
		To(MatchError("invalid character 'b' looking for beginning of value"))
}

func (*Suite) TestMustUnmarshal(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(func() {
		g.Expect(jsonz.MustUnmarshal[*testStruct]([]byte(`{"k": "v"}`))).
			To(Equal(&testStruct{K: "v"}))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustUnmarshal[*testStruct]([]byte(`bad`))
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}

func (*Suite) TestUnmarshalString(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(jsonz.UnmarshalString[*testStruct](`{"k": "v"}`)).
		To(Equal(&testStruct{K: "v"}))

	g.Expect(jsonz.UnmarshalString[*testStruct](`bad`)).Error().
		To(MatchError("invalid character 'b' looking for beginning of value"))
}

func (*Suite) TestMustUnmarshalString(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(func() {
		g.Expect(jsonz.MustUnmarshalString[*testStruct](`{"k": "v"}`)).
			To(Equal(&testStruct{K: "v"}))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustUnmarshalString[*testStruct](`bad`)
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}
