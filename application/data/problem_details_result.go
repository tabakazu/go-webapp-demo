package data

import (
	"github.com/go-playground/validator/v10"
)

const (
	// Error Type
	ErrorTypeValidationError = "validation-error"

	// Error Title
	ErrorTitleValidationError = "Your request parameters didn't validate."
)

// BadRequestErrorResult for 400 Error
type BadRequestErrorResult struct {
	Type          string
	Title         string
	InvalidParams InvalidParams
}

type InvalidParams []InvalidParam

type InvalidParam struct {
	Name   string
	Reason string
}

func NewValidationError(err error) *BadRequestErrorResult {
	var params []InvalidParam

	for _, e := range err.(validator.ValidationErrors) {
		param := InvalidParam{e.StructField(), e.ActualTag()}
		params = append(params, param)
	}

	return &BadRequestErrorResult{
		Type:          ErrorTypeValidationError,
		Title:         ErrorTitleValidationError,
		InvalidParams: params,
	}
}
