package outz_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/outz"
)

type StylesSuite struct {
	// intentionally empty
}

func TestStylesSuite(t *testing.T) {
	fixturez.RunSuite(t, &StylesSuite{})
}

func (s *StylesSuite) TestStyles(g *WithT) {
	g.Expect(outz.DefaultStyles.Default()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.Highlight()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.SecondaryHighlight()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.Secondary()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.Info()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.Success()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.Warning()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.Error()).ToNot(BeNil())
	g.Expect(outz.DefaultStyles.LogLevel(logrus.TraceLevel)).To(Equal(outz.DefaultStyles.Secondary()))
	g.Expect(outz.DefaultStyles.LogLevel(logrus.DebugLevel)).To(Equal(outz.DefaultStyles.Secondary()))
	g.Expect(outz.DefaultStyles.LogLevel(logrus.InfoLevel)).To(Equal(outz.DefaultStyles.Info()))
	g.Expect(outz.DefaultStyles.LogLevel(logrus.WarnLevel)).To(Equal(outz.DefaultStyles.Warning()))
	g.Expect(outz.DefaultStyles.LogLevel(logrus.ErrorLevel)).To(Equal(outz.DefaultStyles.Error()))
	g.Expect(outz.DefaultStyles.LogLevel(logrus.FatalLevel)).To(Equal(outz.DefaultStyles.Error()))
	g.Expect(outz.DefaultStyles.LogLevel(logrus.PanicLevel)).To(Equal(outz.DefaultStyles.Error()))
	g.Expect(outz.DefaultStyles.LogLevel(10)).To(Equal(outz.DefaultStyles.Default()))
}

func (*StylesSuite) TestRestoreDefaultStyles(g *WithT) {
	defaultStyles := outz.DefaultStyles
	outz.DefaultStyles = nil
	outz.RestoreDefaultStyles()
	g.Expect(outz.DefaultStyles).To(Equal(defaultStyles))
}
