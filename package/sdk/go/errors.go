package kms

import "fmt"

type ErrorResponse struct {
	ErrorMsg string `json:"error"`
	Code     string `json:"code"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s (%s)", e.ErrorMsg, e.Code)
}
