package outz

import (
	"io"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/sirupsen/logrus"

	"github.com/ibrt/golang-utils/errorz"
)

var (
	sysStderr = os.Stderr
)

var (
	m                      = &sync.Mutex{}
	isCapturing            = false
	outR, outW, errR, errW *os.File
	restoreFuncs           []func()
	loggers                = []*logrus.Logger{logrus.StandardLogger()}
)

// OutputSetupFunc describes a function that replaces some streams with mock ones for capturing.
type OutputSetupFunc func(outW, errW *os.File) OutputRestoreFunc

// OutputRestoreFunc describe a function that restores some real streams after capturing.
type OutputRestoreFunc func()

// OutputSetupStandard is a [OutputSetupFunc] that configures the stdout/stderr streams.
func OutputSetupStandard(outW, errW *os.File) OutputRestoreFunc {
	origOut := os.Stdout
	origErr := os.Stderr

	os.Stdout = outW
	os.Stderr = errW

	return func() {
		os.Stdout = origOut
		os.Stderr = origErr
	}
}

// GetOutputSetupFatihColor returns a [OutputSetupFunc] that configures the color streams
// (from "github.com/fatih/color").
func GetOutputSetupFatihColor(noColor bool) OutputSetupFunc {
	return func(outW, errW *os.File) OutputRestoreFunc {
		origNoColor := color.NoColor
		origOut := color.Output
		origErr := color.Error

		color.NoColor = noColor
		color.Output = outW
		color.Error = errW

		return func() {
			color.NoColor = origNoColor
			color.Output = origOut
			color.Error = origErr
		}
	}
}

// OutputSetupRodaineTable is a [OutputSetupFunc] that configures the table streams (from "github.com/rodaine/table").
func OutputSetupRodaineTable(outW, _ *os.File) OutputRestoreFunc {
	origOut := table.DefaultWriter
	table.DefaultWriter = outW

	return func() {
		table.DefaultWriter = origOut
	}
}

// OutputSetupSirupsenLogrus is a [OutputSetupFunc] that configures the logging streams
// (from "github.com/sirupsen/logrus").
func OutputSetupSirupsenLogrus(_, errW *os.File) OutputRestoreFunc {
	origErr := make(map[*logrus.Logger]io.Writer)

	for _, logger := range loggers {
		origErr[logger] = logger.Out
		logger.SetOutput(errW)
	}

	return func() {
		for _, logger := range loggers {
			if o, ok := origErr[logger]; ok {
				logger.SetOutput(o)
			} else {
				logger.SetOutput(sysStderr)
			}
		}
	}
}

// MustBeginOutputCapture sets up the mock streams and starts capturing the output.
// It panics if another output capture is already in progress.
// It is the caller's responsibility to ensure mutual exclusion.
func MustBeginOutputCapture(outputSetupFuncs ...OutputSetupFunc) {
	m.Lock()
	defer m.Unlock()

	errorz.Assertf(!isCapturing, "output capture already in progress")
	isCapturing = true
	var err error

	outR, outW, err = os.Pipe()
	errorz.MaybeMustWrap(err)

	errR, errW, err = os.Pipe()
	errorz.MaybeMustWrap(err)

	for _, outputSetupFunc := range outputSetupFuncs {
		restoreFuncs = append(restoreFuncs, outputSetupFunc(outW, errW))
	}
}

// MustEndOutputCapture restores the real streams and returns the captured data.
// It panics if no output capture is in progress.
func MustEndOutputCapture() (outStr, errStr string) {
	m.Lock()
	defer m.Unlock()

	errorz.Assertf(isCapturing, "output capture already in progress")
	return mustFlush()
}

// ResetOutputCapture ensures the output capture is cleared and reset (e.g. after a panic, error or test assertion
// that prevents [MustEndOutputCapture] from being called). Always defer [ResetOutputCapture] before using output
// captures.
func ResetOutputCapture() {
	m.Lock()
	defer m.Unlock()

	if isCapturing {
		func() {
			defer func() { _ = recover() }()
			mustFlush()
		}()
	}
}

func mustFlush() (outStr, errSr string) {
	defer func() {
		errorz.MaybeMustWrap(outR.Close())
		errorz.MaybeMustWrap(errR.Close())

		isCapturing = false
		outR = nil
		outW = nil
		errR = nil
		errW = nil
		restoreFuncs = nil
	}()

	for i := len(restoreFuncs) - 1; i >= 0; i-- {
		restoreFuncs[i]()
	}

	errorz.MaybeMustWrap(outW.Close())
	errorz.MaybeMustWrap(errW.Close())

	outBuf, err := io.ReadAll(outR)
	errorz.MaybeMustWrap(err)

	errBuf, err := io.ReadAll(errR)
	errorz.MaybeMustWrap(err)

	return string(outBuf), string(errBuf)
}
