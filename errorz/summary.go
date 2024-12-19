package errorz

import (
	"reflect"
)

// Summary presents an error in human-readable, debug form.
type Summary struct {
	Name       string         `json:"name,omitempty"`
	Message    string         `json:"message,omitempty"`
	HTTPStatus int            `json:"httpStatus,omitempty"`
	Details    map[string]any `json:"details,omitempty"`
	Components []*Summary     `json:"components,omitempty"`
}

// GetSummary returns a summary of the error.
func GetSummary(err error, includeComponents bool) *Summary {
	if err == nil {
		return nil
	}

	s := &Summary{
		Message: err.Error(),
		Details: make(map[string]any),
		Components: []*Summary{
			getSummaryInternal(err),
		},
	}

	walkSummary(s.Components[0], func(iS *Summary) {
		if iS.Name != "" && s.Name == "" {
			s.Name = iS.Name
		}

		if iS.HTTPStatus != 0 {
			s.HTTPStatus = iS.HTTPStatus
		}

		for k, v := range iS.Details {
			s.Details[k] = v
		}
	})

	if !includeComponents {
		s.Components = nil
	}

	return s
}

func getSummaryInternal(err error) *Summary {
	s := &Summary{
		Name:       maybeGetName(err),
		Message:    err.Error(),
		HTTPStatus: maybeGetHTTPStatus(err),
		Details:    maybeGetDetails(err),
	}

	switch {
	case isWrapError(err):
		s.Name = "[wrap]"
		s.Message = ""
	case isJoinError(err):
		s.Name = "[join]"
		s.Message = ""
	}

	switch e := err.(type) {
	case UnwrapMulti:
		for _, uErr := range e.Unwrap() {
			if uErr != nil {
				s.Components = append(s.Components, getSummaryInternal(uErr))
			}
		}
	case UnwrapSingle:
		if uErr := e.Unwrap(); uErr != nil {
			s.Components = append(s.Components, getSummaryInternal(uErr))

		}
	}

	return s
}

func walkSummary(s *Summary, f func(iS *Summary)) {
	for _, cS := range s.Components {
		walkSummary(cS, f)
	}

	f(s)
}

func maybeGetName(err error) string {
	if e, ok := err.(ErrorName); ok {
		if n := e.GetErrorName(); n != "" {
			return n
		}
	}

	if !isWrapError(err) && !isGenericError(err) {
		return reflect.TypeOf(err).String()
	}

	return ""
}

func maybeGetHTTPStatus(err error) int {
	if e, ok := err.(ErrorHTTPStatus); ok {
		return e.GetErrorHTTPStatus()
	}

	return 0
}

func maybeGetDetails(err error) map[string]any {
	if e, ok := err.(ErrorDetails); ok {
		if d := e.GetErrorDetails(); len(d) > 0 {
			return d
		}
	}

	return nil
}
