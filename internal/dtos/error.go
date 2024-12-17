package dtos

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorBody struct {
	Status string `json:"status"`
	Error  string `json:"errors"`
}

const (
	StatusError = "Error"
)

func GeneralError(err error) ErrorBody {
	return ErrorBody{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) ErrorBody {

	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return ErrorBody{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}

}
