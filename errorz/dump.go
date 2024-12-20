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

// SDump converts the error to a string representation for debug purposes.
func SDump(err error) string {
	if err == nil {
		return "[nil]"
	}

	type dump struct {
		Summary *Summary
		Raw     any
	}

	return strings.TrimSuffix(
		spewConfig.Sdump(dump{
			Summary: GetSummary(err, true),
			Raw:     err,
		}),
		"\n")
}
