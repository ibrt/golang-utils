package urlz_test

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/urlz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestMustParse(g *WithT) {
	g.Expect(func() {
		g.Expect(urlz.MustParse("https://test").String()).To(Equal("https://test"))
	}).ToNot(Panic())

	g.Expect(func() {
		urlz.MustParse("\b")
	}).To(PanicWith(MatchError(`parse "\b": net/url: invalid control character in URL`)))
}

func (*Suite) TestMustEdit(g *WithT) {
	g.Expect(func() {
		g.Expect(urlz.MustEdit("https://test", func(u *url.URL) {
			u.Path = "/test"
		})).To(Equal("https://test/test"))
	}).ToNot(Panic())

	g.Expect(func() {
		urlz.MustEdit("\b", func(u *url.URL) {})
	}).To(PanicWith(MatchError(`parse "\b": net/url: invalid control character in URL`)))
}

func (*Suite) TestGetValueDef(g *WithT) {
	g.Expect(urlz.GetValueDef(url.Values{"k": []string{"v1", "v2"}}, "k", "x")).To(Equal("v1"))
	g.Expect(urlz.GetValueDef(url.Values{}, "k", "x")).To(Equal("x"))
}

func (*Suite) TestEncodeValues(g *WithT) {
	g.Expect(urlz.EncodeValues(nil)).To(Equal(""))
	g.Expect(urlz.EncodeValues(map[string]string{})).To(Equal(""))
	g.Expect(urlz.EncodeValues(map[string]string{"k1": "v1", "k2": "v2"})).To(Equal("k1=v1&k2=v2"))
}
