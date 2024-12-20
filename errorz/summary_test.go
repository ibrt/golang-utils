package errorz_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/errorz/terrorz"
)

func TestGetSummary(t *testing.T) {
	g := NewWithT(t)

	e1a := &terrorz.SimpleMockTestDetailedError{
		ErrorMessage: "e1a-err",
		Name:         "e1a",
		HTTPStatus:   101,
		Details:      map[string]any{"e1a": true},
	}

	e1b := fmt.Errorf("e1b: %w", e1a)

	e2a := &terrorz.SimpleMockTestDetailedError{
		ErrorMessage: "e2a-err",
		Name:         "e2a",
		HTTPStatus:   102,
		Details:      map[string]any{"e2a": true},
	}

	e2b := &terrorz.SimpleMockTestDetailedUnwrapSingleError{
		SimpleMockTestDetailedError: &terrorz.SimpleMockTestDetailedError{
			ErrorMessage: "e2b-err",
		},
		UnwrapSingle: e2a,
	}

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

									Name:    "*terrorz.SimpleMockTestDetailedUnwrapSingleError",
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

									Name:    "*terrorz.SimpleMockTestDetailedUnwrapSingleError",
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
