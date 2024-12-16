package handlers

import "net/http"

type HandlerErr interface {
	error
	Code() int
}

type HandlerError struct {
	HandlerErr
	code    int
	message string
}

func NewError(code int, message string) HandlerError {
	return HandlerError{code: code, message: message}
}

func NewErrorWithoutMessage(code int) HandlerError {
	return HandlerError{code: code}
}

func (h HandlerError) Error() string {
	if h.message != "" {
		return h.message
	} else {
		return http.StatusText(h.code)
	}
}

func (h HandlerError) Code() int {
	return h.code
}
