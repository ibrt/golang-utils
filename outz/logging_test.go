package outz_test

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/outz"
)

type LoggingSuite struct {
	// intentionally empty
}

func TestLoggingSuite(t *testing.T) {
	fixturez.RunSuite(t, &LoggingSuite{})
}

func (*LoggingSuite) TestHumanLogFormatter(g *WithT) {
	defer outz.ResetOutputCapture()
	outz.MustBeginOutputCapture(outz.GetOutputSetupFatihColor(false), outz.OutputSetupSirupsenLogrus)

	t := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	logger := outz.NewLogger()
	logger.SetFormatter(outz.NewHumanLogFormatter().
		SetInitTime(t).
		SetStyles(outz.DefaultStyles))

	logrus.NewEntry(logger).
		WithTime(t.Add(time.Second)).
		WithFields(logrus.Fields{
			"name":        "ignore",
			"duration_ms": float64(1000),
			"struct": struct {
				K string `json:"k"`
			}{K: "v"},
			"map": map[string]string{
				"k": "v",
			},
			"slice": []string{
				"v",
			},
			"other": "v",
		}).Info("message")

	outStr, errStr := outz.MustEndOutputCapture()

	g.Expect(outStr).To(Equal(""))
	g.Expect(errStr).To(Equal("\x1b[36m┌─ INFO[0001]\x1b[0m message\n\x1b[36m│    \x1b[0m\x1b[36mduration\x1b[0m=1s\n\x1b[36m│    \x1b[0m\x1b[36mmap\x1b[0m={\n\x1b[36m│    \x1b[0m  \"k\": \"v\"\n\x1b[36m│    \x1b[0m}\n\x1b[36m│    \x1b[0m\x1b[36mother\x1b[0m=v\n\x1b[36m│    \x1b[0m\x1b[36mslice\x1b[0m=[\n\x1b[36m│    \x1b[0m  \"v\"\n\x1b[36m│    \x1b[0m]\n\x1b[36m└─   \x1b[0m\x1b[36mstruct\x1b[0m={\"k\":\"v\"}\n"))
}
