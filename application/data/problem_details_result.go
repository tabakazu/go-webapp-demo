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

type ErrorResult struct {
	Message string `json:"message"`
}

type ErrorDetailsResult struct {
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	InvalidParams InvalidParams `json:"invalid_params"`
}

type InvalidParams []InvalidParam

type InvalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func NewValidationErrorResult(err error) *ErrorDetailsResult {
	var params []InvalidParam

	for _, e := range err.(validator.ValidationErrors) {
		param := InvalidParam{e.StructField(), e.ActualTag()}
		params = append(params, param)
	}

	return &ErrorDetailsResult{
		Type:          ErrorTypeValidationError,
		Title:         ErrorTitleValidationError,
		InvalidParams: params,
	}
}
