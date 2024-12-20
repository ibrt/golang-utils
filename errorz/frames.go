package errorz

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Frame describes a frame.
type Frame struct {
	Summary       string `json:"summary,omitempty"`
	Location      string `json:"location,omitempty"`
	ShortLocation string `json:"shortLocation,omitempty"`
	Package       string `json:"fullPackage,omitempty"`
	ShortPackage  string `json:"package,omitempty"`
	Function      string `json:"function,omitempty"`
	FileAndLine   string `json:"fileAndLine,omitempty"`
	File          string `json:"file,omitempty"`
	Line          int    `json:"line,omitempty"`
}

// NewFrame initializes a new frame.
func NewFrame(frameFunction string, file string, line int) *Frame {
	fileAndLine := ""

	if file != "" {
		fileAndLine = fmt.Sprintf("%v:%v", file, line)
	}

	f := &Frame{
		Summary:       fileAndLine,
		Location:      "",
		ShortLocation: "",
		Package:       "",
		ShortPackage:  "",
		Function:      "",
		FileAndLine:   fileAndLine,
		File:          file,
		Line:          line,
	}

	if frameFunction != "" {
		dir := ""
		base := frameFunction

		if lastSlashIndex := strings.LastIndex(base, "/"); lastSlashIndex >= 0 {
			dir = frameFunction[:lastSlashIndex]
			base = frameFunction[lastSlashIndex+1:]
		}

		if dotIndex := strings.Index(base, "."); dotIndex >= 0 {
			base = base[:dotIndex]
		}

		if dir != "" {
			base = dir + "/" + base
		}

		shortPkg := path.Base(base)
		function := strings.TrimPrefix(strings.TrimPrefix(frameFunction, base), ".")
		shortLocation := fmt.Sprintf("%v.%v", shortPkg, function)

		if fileAndLine != "" {
			f.Summary = fmt.Sprintf("%v (%v)", shortLocation, fileAndLine)
		} else {
			f.Summary = shortLocation
		}

		f.Location = frameFunction
		f.ShortLocation = shortLocation
		f.Package = base
		f.ShortPackage = shortPkg
		f.Function = function
	}

	if f.Summary == "" {
		f.Summary = "<unknown>"
	}

	return f
}

// Frames describes a stack of frames.
type Frames []*Frame

// ToSummaries converts the frames to a slice of frame summaries.
func (f Frames) ToSummaries() []string {
	summaries := make([]string, 0, len(f))

	for _, frame := range f {
		summaries = append(summaries, frame.Summary)
	}

	return summaries
}

// GetFrames returns the frames from the error, or the current frames the error is not wrapped or is nil.
func GetFrames(err error) Frames {
	if e, ok := err.(*wrappedError); ok { //nolint:errorlint
		return e.frames
	}

	callers := make([]uintptr, 1024)
	callers = callers[:runtime.Callers(1, callers[:])]
	callersFrames := runtime.CallersFrames(callers)

	frames := make([]*Frame, 0, len(callers))

	for {
		callerFrame, more := callersFrames.Next()
		frame := NewFrame(callerFrame.Function, callerFrame.File, callerFrame.Line)

		if frame.ShortPackage == "errorz" {
			continue
		}

		frames = append(frames, frame)

		if !more {
			break
		}
	}

	return frames
}
