package errorz

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var (
	spewConfig = &spew.ConfigState{
		Indent:                  "    ",
		DisableMethods:          true,
		DisablePointerMethods:   true,
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}
)

// SDump converts the error to an extremely detailed string representation for debug purposes.
func SDump(err error) string {
	if err == nil {
		return "[nil]"
	}

	if e, ok := err.(*wrappedError); ok {
		type dump struct {
			Message string
			Debug   []error
			Frames  []string
		}

		return strings.TrimSuffix(
			spewConfig.Sdump(dump{
				Message: e.Error(),
				Debug:   e.errs,
				Frames:  e.frames.ToSummaries(),
			}),
			"\n")
	}

	type dump struct {
		Message string
		Debug   []error
	}

	return strings.TrimSuffix(
		spewConfig.Sdump(dump{
			Message: err.Error(),
			Debug:   []error{err},
		}),
		"\n")
}
