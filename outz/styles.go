package outz

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Styles describes a set of output styles.
type Styles interface {
	Default() *color.Color
	Highlight() *color.Color
	SecondaryHighlight() *color.Color
	Secondary() *color.Color
	Info() *color.Color
	Success() *color.Color
	Warning() *color.Color
	Error() *color.Color
	LogLevel(level logrus.Level) *color.Color
}

var (
	_ Styles = (*stylesImpl)(nil)
)

type stylesImpl struct {
	// intentionally empty
}

// Default implements the [Styles] interface.
func (s *stylesImpl) Default() *color.Color {
	return color.New(color.Reset)
}

// Highlight implements the [Styles] interface.
func (s *stylesImpl) Highlight() *color.Color {
	return color.New(color.Bold)
}

// SecondaryHighlight implements the [Styles] interface.
func (s *stylesImpl) SecondaryHighlight() *color.Color {
	return color.New(color.Bold, color.Faint)
}

// Secondary implements the [Styles] interface.
func (s *stylesImpl) Secondary() *color.Color {
	return color.New(color.Faint)
}

// Info implements the [Styles] interface.
func (s *stylesImpl) Info() *color.Color {
	return color.New(color.FgCyan)
}

// Success implements the [Styles] interface.
func (s *stylesImpl) Success() *color.Color {
	return color.New(color.FgGreen)
}

// Warning implements the [Styles] interface.
func (s *stylesImpl) Warning() *color.Color {
	return color.New(color.FgYellow)
}

// Error implements the [Styles] interface.
func (s *stylesImpl) Error() *color.Color {
	return color.New(color.FgHiRed)
}

// LogLevel implements the [Styles] interface.
func (s *stylesImpl) LogLevel(level logrus.Level) *color.Color {
	switch level {
	case logrus.TraceLevel, logrus.DebugLevel:
		return s.Secondary()
	case logrus.InfoLevel:
		return s.Info()
	case logrus.WarnLevel:
		return s.Warning()
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return s.Error()
	default:
		return s.Default()
	}
}

var (
	defaultStyles = &stylesImpl{}
)

// DefaultStyles is a default, shared instance of [Styles].
var (
	DefaultStyles = defaultStyles
)

// RestoreDefaultStyles restores the default value of [DefaultStyles].
func RestoreDefaultStyles() {
	DefaultStyles = defaultStyles
}
