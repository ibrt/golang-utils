// Package vldz provides various utilities for validating data.
package vldz

import (
	"fmt"
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
		return fl.Field().Kind() == reflect.Struct ||
			(fl.Field().Kind() == reflect.Pointer && fl.Field().Elem().Kind() == reflect.Struct)
	})
}

// MustRegisterValidator registers a validator.
func MustRegisterValidator(tag string, validator validator.Func) {
	errorz.MaybeMustWrap(validate.RegisterValidation(tag, validator))
}

// RegexpValidatorFactory creates a validator that matches a string against the given regexp.
func RegexpValidatorFactory(r *regexp.Regexp) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return r.MatchString(fl.Field().String())
	}
}

var (
	_ error               = (*ValidationError)(nil)
	_ errorz.ErrorName    = (*ValidationError)(nil)
	_ errorz.ErrorDetails = (*ValidationError)(nil)
	_ errorz.UnwrapSingle = (*ValidationError)(nil)
)

// ValidationError describes a validation error.
type ValidationError struct {
	invalidFieldKeys []string
	invalidFields    map[string]any
	validationErrors validator.ValidationErrors
	otherError       error
}

// NewValidationError initializes a new validation error.
func NewValidationError(err error) *ValidationError {
	if validationErrs, ok := errorz.As[validator.ValidationErrors](err); ok {
		invalidFieldKeys := make([]string, 0, len(validationErrs))
		invalidFields := make(map[string]any)

		for _, e := range validationErrs {
			invalidFieldKeys = append(invalidFieldKeys, e.Field())
			invalidFields[e.Field()] = e.Tag()
		}

		return &ValidationError{
			invalidFieldKeys: invalidFieldKeys,
			invalidFields:    invalidFields,
			validationErrors: validationErrs,
			otherError:       nil,
		}
	}

	return &ValidationError{
		invalidFieldKeys: nil,
		invalidFields:    nil,
		validationErrors: nil,
		otherError:       err,
	}
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	switch {
	case e.validationErrors != nil:
		return fmt.Sprintf("validation errors: invalid field(s): %v", strings.Join(e.invalidFieldKeys, ", "))
	case e.otherError != nil:
		return "validation error: " + e.otherError.Error()
	default:
		return "validation error: unknown"
	}
}

// GetErrorName implements the [errorz.ErrorName] interface.
func (*ValidationError) GetErrorName() string {
	return "validation-error"
}

// GetErrorDetails implements the [errorz.ErrorDetails] interface.
func (e *ValidationError) GetErrorDetails() map[string]any {
	if len(e.invalidFieldKeys) > 0 {
		return map[string]any{
			"fields": e.invalidFields,
		}
	}

	return nil
}

// Unwrap implements the [errorz.UnwrapSingle] interface.
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

// MustValidateStruct is like [ValidateStruct] but panics on error.
func MustValidateStruct(v any) {
	errorz.MaybeMustWrap(ValidateStruct(v))
}
