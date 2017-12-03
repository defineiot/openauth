package exception

import (
	"fmt"
	"net/http"
)

// APIException is openauth error
type apiException struct {
	msg  string
	code int
}

func (e *apiException) Error() string {
	return e.msg
}

// Code exception's code
func (e *apiException) Code() int {
	return e.code
}

// InternalServerError exception
type InternalServerError struct {
	*apiException
}

// NotFound exception
type NotFound struct {
	*apiException
}

// BadRequest exception
type BadRequest struct {
	*apiException
}

// NewAPIException is openauth api error
func NewAPIException(text string, code int) error {
	return &apiException{msg: text, code: code}
}

// NewInternalServerError for 503
func NewInternalServerError(format string, args ...interface{}) error {
	excp := new(InternalServerError)
	excp.apiException = &apiException{msg: fmt.Sprintf(format, args...), code: http.StatusInternalServerError}
	return excp
}

// NewNotFound for 404
func NewNotFound(format string, args ...interface{}) error {
	excp := new(NotFound)
	excp.apiException = &apiException{msg: fmt.Sprintf(format, args...), code: http.StatusNotFound}
	return excp
}

// NewBadRequest for 400
func NewBadRequest(format string, args ...interface{}) error {
	excp := new(BadRequest)
	excp.apiException = &apiException{msg: fmt.Sprintf(format, args...), code: http.StatusBadRequest}
	return excp
}
