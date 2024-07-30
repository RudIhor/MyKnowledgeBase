package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidateError struct {
	FieldName      string
	ValidateType   string
	ExpectedResult string
	GotResult      interface{}
	Description    string
	fieldError     validator.FieldError
	Details        interface{}
}

type ValidateErrors struct {
	validator.ValidationErrors
}

func (errors ValidateErrors) TransformError() *ValidateError {
	validateError := ValidateError{}
	if len(errors.ValidationErrors) > 0 {
		e := errors.ValidationErrors[0]
		validateError.FieldName = e.Field()
		validateError.ValidateType = e.ActualTag()
		validateError.GotResult = e.Value()
		validateError.ExpectedResult = e.Param()
		validateError.fieldError = e

		e1 := ValidationError{
			Namespace:       e.Namespace(),
			Field:           e.Field(),
			StructNamespace: e.StructNamespace(),
			StructField:     e.StructField(),
			Tag:             e.Tag(),
			ActualTag:       e.ActualTag(),
			Kind:            fmt.Sprintf("%v", e.Kind()),
			Type:            fmt.Sprintf("%v", e.Type()),
			Value:           fmt.Sprintf("%v", e.Value()),
			Param:           e.Param(),
			Message:         e.Error(),
		}

		validateError.Details = e1
	}

	return &validateError
}
