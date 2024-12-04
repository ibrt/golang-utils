package vldz

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/ibrt/golang-utils/errorz"
)

var (
	validate = validator.New()
)

func init() {
	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		if name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]; name != "" && name != "-" {
			return name
		}
		return f.Name
	})

	MustRegisterValidator("kind-struct-or-struct-ptr", func(fl validator.FieldLevel) bool {
		return fl.Field().Kind() == reflect.Struct || (fl.Field().Kind() == reflect.Pointer && fl.Field().Elem().Kind() == reflect.Struct)
	})
}

// MustRegisterValidator registers a validator.
func MustRegisterValidator(tag string, validator validator.Func) {
	errorz.MaybeMustWrap(validate.RegisterValidation(tag, validator))
}

// RegexpValidatorFactory creates a validator that matches a string against the given regexp.
func RegexpValidatorFactory(regexp *regexp.Regexp) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return regexp.MatchString(fl.Field().String())
	}
}

var (
	_ error               = (*ValidationError)(nil)
	_ errorz.UnwrapSingle = (*ValidationError)(nil)
)

// ValidationError describes a validation error.
type ValidationError struct {
	fieldsSummary    map[string]any
	validationErrors validator.ValidationErrors
	otherError       error
}

// NewValidationError initializes a new validation error.
func NewValidationError(err error) *ValidationError {
	if validationErrs, ok := errorz.As[validator.ValidationErrors](err); ok {
		fieldsSummary := make(map[string]any)

		for _, err := range validationErrs {
			n := err.Namespace()
			if i := strings.Index(err.Namespace(), "."); i >= 0 {
				n = err.Namespace()[i+1:]
			}
			fieldsSummary[n] = err.Tag()
		}

		return &ValidationError{
			fieldsSummary:    fieldsSummary,
			validationErrors: validationErrs,
			otherError:       nil,
		}
	}

	return &ValidationError{
		otherError: err,
	}
}

// MaybeGetFieldsSummary returns the fields summary if available.
func (e *ValidationError) MaybeGetFieldsSummary() map[string]any {
	return e.fieldsSummary
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	switch {
	case e.validationErrors != nil:
		w := strings.Builder{}
		w.WriteString("validation error(s):")

		for _, vErr := range e.validationErrors {
			w.WriteString("\n- ")
			w.WriteString(vErr.Error())
		}

		return w.String()
	case e.otherError != nil:
		return "validation error: " + e.otherError.Error()
	default:
		return "validation error: unknown"
	}
}

// Unwrap implements the errorz.UnwrapSingle interface.
func (e *ValidationError) Unwrap() error {
	if e == nil {
		return nil
	}

	if e.validationErrors != nil {
		return e.validationErrors
	}

	return e.otherError
}

// ValidateStruct validates a struct.
func ValidateStruct(v any) error {
	if err := validate.Struct(v); err != nil {
		return NewValidationError(err)
	}

	return nil
}

// MustValidateStruct is like ValidateStruct but panics on error.
func MustValidateStruct(v any) {
	errorz.MaybeMustWrap(ValidateStruct(v))
}
