package errorz_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
)

func TestNewFrame(t *testing.T) {
	type testCase struct {
		frameFunction string
		file          string
		line          int
		expected      *errorz.Frame
	}

	for i, tc := range []testCase{
		{
			frameFunction: "",
			file:          "",
			line:          0,
			expected: &errorz.Frame{
				Summary: "<unknown>",
			},
		},
		{
			frameFunction: "a.b",
			file:          "",
			line:          0,
			expected: &errorz.Frame{
				Summary:       "a.b",
				Location:      "a.b",
				ShortLocation: "a.b",
				Package:       "a",
				ShortPackage:  "a",
				Function:      "b",
				FileAndLine:   "",
				File:          "",
				Line:          0,
			},
		},
		{
			frameFunction: "",
			file:          "f.go",
			line:          10,
			expected: &errorz.Frame{
				Summary:     "f.go:10",
				FileAndLine: "f.go:10",
				File:        "f.go",
				Line:        10,
			},
		},
		{
			frameFunction: "a.b",
			file:          "f.go",
			line:          10,
			expected: &errorz.Frame{
				Summary:       "a.b (f.go:10)",
				Location:      "a.b",
				ShortLocation: "a.b",
				Package:       "a",
				ShortPackage:  "a",
				Function:      "b",
				FileAndLine:   "f.go:10",
				File:          "f.go",
				Line:          10,
			},
		},
		{
			frameFunction: "a.b.c",
			file:          "f.go",
			line:          10,
			expected: &errorz.Frame{
				Summary:       "a.b.c (f.go:10)",
				Location:      "a.b.c",
				ShortLocation: "a.b.c",
				Package:       "a",
				ShortPackage:  "a",
				Function:      "b.c",
				FileAndLine:   "f.go:10",
				File:          "f.go",
				Line:          10,
			},
		},
		{
			frameFunction: "a.(*b).c",
			file:          "f.go",
			line:          10,
			expected: &errorz.Frame{
				Summary:       "a.(*b).c (f.go:10)",
				Location:      "a.(*b).c",
				ShortLocation: "a.(*b).c",
				Package:       "a",
				ShortPackage:  "a",
				Function:      "(*b).c",
				FileAndLine:   "f.go:10",
				File:          "f.go",
				Line:          10,
			},
		},
		{
			frameFunction: "a/b/c.d",
			file:          "f.go",
			line:          10,
			expected: &errorz.Frame{
				Summary:       "c.d (f.go:10)",
				Location:      "a/b/c.d",
				ShortLocation: "c.d",
				Package:       "a/b/c",
				ShortPackage:  "c",
				Function:      "d",
				FileAndLine:   "f.go:10",
				File:          "f.go",
				Line:          10,
			},
		},
		{
			frameFunction: "a.com/b.c/d.(*e).f",
			file:          "f.go",
			line:          10,
			expected: &errorz.Frame{
				Summary:       "d.(*e).f (f.go:10)",
				Location:      "a.com/b.c/d.(*e).f",
				ShortLocation: "d.(*e).f",
				Package:       "a.com/b.c/d",
				ShortPackage:  "d",
				Function:      "(*e).f",
				FileAndLine:   "f.go:10",
				File:          "f.go",
				Line:          10,
			},
		},
	} {
		t.Run(fmt.Sprintf("%03v", i+1), func(t *testing.T) {
			g := NewWithT(t)
			g.Expect(errorz.NewFrame(tc.frameFunction, tc.file, tc.line)).To(Equal(tc.expected))
		})
	}
}

func TestFrames(t *testing.T) {
	g := NewWithT(t)

	g.Expect((errorz.Frames)(nil).ToSummaries()).To(BeEmpty())
	g.Expect(errorz.Frames{}.ToSummaries()).To(BeEmpty())
	g.Expect(errorz.Frames{{Summary: "f1"}, {Summary: "f2"}}.ToSummaries()).To(HaveExactElements("f1", "f2"))
}

func TestGetFrames(t *testing.T) {
	g := NewWithT(t)

	frames := errorz.GetFrames(nil)
	g.Expect(len(frames)).To(BeNumerically(">", 0))
	g.Expect(frames[0].ShortLocation).To(Equal("errorz_test.TestGetFrames"))

	frames = errorz.GetFrames(fmt.Errorf("err"))
	g.Expect(len(frames)).To(BeNumerically(">", 0))
	g.Expect(frames[0].ShortLocation).To(Equal("errorz_test.TestGetFrames"))

	frames = errorz.GetFrames(errorz.Errorf("err"))
	g.Expect(len(frames)).To(BeNumerically(">", 0))
	g.Expect(frames[0].ShortLocation).To(Equal("errorz_test.TestGetFrames"))
}
