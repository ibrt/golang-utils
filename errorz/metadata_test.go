package errorz_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/errorz/terrorz"
)

func TestMetadata(t *testing.T) {
	g := NewWithT(t)

	type mk1 int
	type mk2 int

	const k1 mk1 = 0
	const k2 mk2 = 0

	{
		err := fmt.Errorf("test error")
		errorz.MaybeSetMetadata(err, k1, "v1")
		_, ok := errorz.MaybeGetMetadata[string](err, k1)
		g.Expect(ok).To(BeFalse())
		g.Expect(func() { errorz.MustGetMetadata[string](err, k1) }).To(Panic())
	}

	{
		err := errorz.Errorf("test error")
		errorz.MaybeSetMetadata(err, k1, "v1")
		errorz.MaybeSetMetadata(err, k2, "v2")

		v, ok := errorz.MaybeGetMetadata[string](err, k1)
		g.Expect(ok).To(BeTrue())
		g.Expect(v).To(Equal("v1"))

		g.Expect(func() {
			g.Expect(errorz.MustGetMetadata[string](err, k1)).To(Equal("v1"))
		}).ToNot(Panic())

		v, ok = errorz.MaybeGetMetadata[string](err, k2)
		g.Expect(ok).To(BeTrue())
		g.Expect(v).To(Equal("v2"))
		g.Expect(func() { errorz.MustGetMetadata[int](err, k2) }).To(Panic())

		_, ok = errorz.MaybeGetMetadata[string](err, 0)
		g.Expect(ok).To(BeFalse())
		g.Expect(func() { errorz.MustGetMetadata[string](err, 0) }).To(Panic())
	}
}

func TestSummary(t *testing.T) {
	g := NewWithT(t)

	e1a := terrorz.NewTestDetailedError("e1a-err", "e1a", 101, map[string]any{"e1a": true})
	e1b := fmt.Errorf("e1b: %w", e1a)
	e2a := terrorz.NewTestDetailedError("e2a-err", "e2a", 102, map[string]any{"e2a": true})
	e2b := terrorz.NewTestDetailedErrorUnwrapSingle("e2b-err", "", 0, nil, e2a)
	e3 := fmt.Errorf("e3")
	e4 := errors.Join(e2b, e3)

	g.Expect(errorz.GetSummary(nil, true)).To(BeNil())
	g.Expect(errorz.GetSummary(nil, false)).To(BeNil())

	g.Expect(errorz.GetSummary(errorz.Wrap(e1b, e4), true)).
		To(Equal(&errorz.Summary{
			Name:       "e1a",
			Message:    "e2b-err\ne3: e1b: e1a-err",
			HTTPStatus: 102,
			Details: map[string]any{
				"e1a": true,
				"e2a": true,
			},
			Components: []*errorz.Summary{
				{
					Name: "[wrap]",
					Components: []*errorz.Summary{
						{
							Message: "e1b: e1a-err",
							Components: []*errorz.Summary{
								{
									Name:       "e1a",
									Message:    "e1a-err",
									HTTPStatus: 101,
									Details: map[string]any{
										"e1a": true,
									},
								},
							},
						},
						{
							Name: "[join]",
							Components: []*errorz.Summary{
								{

									Name:    "*terrorz.testDetailedErrorUnwrapSingle",
									Message: "e2b-err",
									Components: []*errorz.Summary{
										{
											Name:       "e2a",
											Message:    "e2a-err",
											HTTPStatus: 102,
											Details: map[string]any{
												"e2a": true,
											},
										},
									},
								},
								{
									Message: "e3",
								},
							},
						},
					},
				},
			},
		}))

	g.Expect(errorz.GetSummary(errors.Join(e1b, e4), true)).
		To(Equal(&errorz.Summary{
			Name:       "e1a",
			Message:    "e1b: e1a-err\ne2b-err\ne3",
			HTTPStatus: 102,
			Details: map[string]any{
				"e1a": true,
				"e2a": true,
			},
			Components: []*errorz.Summary{
				{
					Name: "[join]",
					Components: []*errorz.Summary{
						{
							Message: "e1b: e1a-err",
							Components: []*errorz.Summary{
								{
									Name:       "e1a",
									Message:    "e1a-err",
									HTTPStatus: 101,
									Details: map[string]any{
										"e1a": true,
									},
								},
							},
						},
						{
							Name: "[join]",
							Components: []*errorz.Summary{
								{

									Name:    "*terrorz.testDetailedErrorUnwrapSingle",
									Message: "e2b-err",
									Components: []*errorz.Summary{
										{
											Name:       "e2a",
											Message:    "e2a-err",
											HTTPStatus: 102,
											Details: map[string]any{
												"e2a": true,
											},
										},
									},
								},
								{
									Message: "e3",
								},
							},
						},
					},
				},
			},
		}))

	g.Expect(errorz.GetSummary(e1a, true)).
		To(Equal(&errorz.Summary{
			Name:       "e1a",
			Message:    "e1a-err",
			HTTPStatus: 101,
			Details: map[string]any{
				"e1a": true,
			},
			Components: []*errorz.Summary{
				{
					Name:       "e1a",
					Message:    "e1a-err",
					HTTPStatus: 101,
					Details: map[string]any{
						"e1a": true,
					},
				},
			},
		}))

	g.Expect(errorz.GetSummary(e1a, false)).
		To(Equal(&errorz.Summary{
			Name:       "e1a",
			Message:    "e1a-err",
			HTTPStatus: 101,
			Details: map[string]any{
				"e1a": true,
			},
		}))

	g.Expect(errorz.GetSummary(e1b, true)).
		To(Equal(&errorz.Summary{
			Name:       "e1a",
			Message:    "e1b: e1a-err",
			HTTPStatus: 101,
			Details: map[string]any{
				"e1a": true,
			},
			Components: []*errorz.Summary{
				{
					Message: "e1b: e1a-err",
					Components: []*errorz.Summary{
						{
							Name:       "e1a",
							Message:    "e1a-err",
							HTTPStatus: 101,
							Details: map[string]any{
								"e1a": true,
							},
						},
					},
				},
			},
		}))

	g.Expect(errorz.GetSummary(e1b, false)).
		To(Equal(&errorz.Summary{
			Name:       "e1a",
			Message:    "e1b: e1a-err",
			HTTPStatus: 101,
			Details: map[string]any{
				"e1a": true,
			},
		}))
}
