package outz

import (
	"bytes"
	"cmp"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ibrt/golang-utils/jsonz"
	"github.com/ibrt/golang-utils/memz"
)

// NewLogger initializes new [*logrus.Logger] and registers it for possible output caputre.
func NewLogger() *logrus.Logger {
	m.Lock()
	defer m.Unlock()

	logger := logrus.New()
	loggers = append(loggers, logger)

	if isCapturing {
		logger.SetOutput(errW)
	} else {
		logger.SetOutput(sysStderr)
	}

	return logger
}

var (
	_ logrus.Formatter = (*HumanLogFormatter)(nil)
)

// HumanLogFormatter is a human-readable [logrus.Formatter] suitable for colored, interactive console output.
type HumanLogFormatter struct {
	initTime time.Time
	styles   Styles
}

// NewHumanLogFormatter initializes a new [*HumanLogFormatter].
func NewHumanLogFormatter() *HumanLogFormatter {
	return &HumanLogFormatter{
		initTime: time.Now(),
		styles:   DefaultStyles,
	}
}

// SetInitTime sets the init [time.Time].
func (f *HumanLogFormatter) SetInitTime(initTime time.Time) *HumanLogFormatter {
	f.initTime = initTime
	return f
}

// SetStyles sets the [Styles].
func (f *HumanLogFormatter) SetStyles(styles Styles) *HumanLogFormatter {
	f.styles = styles
	return f
}

// Format implements the [logrus.Formatter] interface.
func (f *HumanLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buf := &bytes.Buffer{}

	_, _ = f.styles.LogLevel(entry.Level).Fprintf(buf, "┌─ %v[%04v]",
		strings.ToUpper(entry.Level.String())[:4],
		int64(entry.Time.Sub(f.initTime).Seconds()))

	_, _ = fmt.Fprintf(buf, " %s\n",
		strings.TrimSuffix(entry.Message, "\n"))

	lines := make([]string, 0)

	for _, rk := range memz.GetSortedMapKeys(entry.Data, cmp.Less) {
		if k, v, ok := f.prepareField(rk, entry.Data[rk]); ok {
			lines = append(lines, strings.Split(
				fmt.Sprintf("%v=%v", f.styles.LogLevel(entry.Level).Sprintf("%v", k), v),
				"\n")...)
		}
	}

	for i, line := range lines {
		if i == len(lines)-1 {
			_, _ = f.styles.LogLevel(entry.Level).Fprintf(buf, "└─   ")
		} else {
			_, _ = f.styles.LogLevel(entry.Level).Fprintf(buf, "│    ")
		}

		_, _ = fmt.Fprintln(buf, line)
	}

	return buf.Bytes(), nil
}

func (f *HumanLogFormatter) prepareField(k string, v any) (string, any, bool) {
	switch k {
	case "name":
		return "", nil, false
	case "duration_ms":
		if vv, ok := v.(float64); ok {
			k = "duration"
			v = time.Duration(vv) * time.Millisecond
		}
	}

	if vv := reflect.Indirect(reflect.ValueOf(v)); vv.IsValid() {
		switch vv.Type().Kind() {
		case reflect.Map, reflect.Slice:
			if !vv.IsNil() {
				return k, jsonz.MustMarshalPrettyString(v), true
			}
		case reflect.Struct:
			return k, jsonz.MustMarshalString(v), true
		default:
			// continue
		}
	}

	switch tv := v.(type) {
	case time.Duration:
		return k, tv.String(), true
	default:
		// continue
	}

	return k, v, true
}
