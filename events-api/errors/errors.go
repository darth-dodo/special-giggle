package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (

	// InternalServerError
	ErrInternal = &Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
	}

	// UnprocessableEntity
	ErrUnprocessableEntity = &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
	}

	// BadRequest HTTP 400
	ErrBadRequest = &Error{
		Code:    http.StatusBadRequest,
		Message: "Error invalid argument",
	}
	// EventNotFound HTTP 404
	ErrEventNotFound = &Error{
		Code:    http.StatusNotFound,
		Message: "Event not found",
	}
	// ObjectIsRequired HTTP 400
	ErrObjectIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Request object should be provided",
	}
	// ValidEventIDIsRequired HTTP 400
	ErrValidEventIDIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "A valid event id is required",
	}
	// EventTimingIsRequired HTTP 400
	ErrEventTimingIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Event start time and end time should be provided",
	}
	// InvalidLimit HTTP 400
	ErrInvalidLimit = &Error{
		Code:    http.StatusBadRequest,
		Message: "Limit should be an integral value",
	}
	// InvalidTimeFormat HTTP 400
	ErrInvalidTimeFormat = &Error{
		Code:    http.StatusBadRequest,
		Message: "Time Should be passed in RFC3339 Format: " + time.RFC3339,
	}
)

// Error main object for error
type Error struct {
	Code    int
	Message str
}

// creating a method on the interface of error to return the error string method
func (err *Error) Error() string {
	return err.String()
}

func (err *Error) String() string {
	if err == nil {
		return ""
	}

	return fmt.Sprintf("error: code=%s message=%s", htt.StatusText(err.Code), err.Message)
}

// Convert error into JSON
func (err *Error) JSON() []byte {
	if err == nil {
		return []byte("{}")
	}

	res, _ := json.Marshall(err)
	return res
}

// Get Status Code
func (err *Error) StatusCode() int {
	if err == nil {
		return http.StatusOk
	}

	return err.Code
}
