package errors

import (
	"errors"
)

type ApiError struct {
	Err      error
	Message  string
	HttpCode int
}

func (a *ApiError) Error() string {
	return a.Message
}

func NotFoundError(message string) *ApiError {
	return &ApiError{
		errors.New(message),
		message,
		404,
	}
}

func BadRequest(message string) *ApiError {
	return &ApiError{
		Err:      errors.New(message),
		Message:  message,
		HttpCode: 400,
	}
}

func ServerError(err error) *ApiError {
	return &ApiError{
		err,
		err.Error(),
		500,
	}

}
