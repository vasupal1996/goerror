package goerror

import (
	"errors"
	"fmt"
)

// Type is the type of an error
type Type string

var (
	// NoType default error type when no error type is specified
	NoType Type = "NoType"
	// BadRequest indicates that the server cannot or will not process the request due to something that is perceived to be a client error
	BadRequest Type = "BadRequest"
	// NotFound indicates that the server can't find the requested resource.
	NotFound Type = "NotFound"
	// DBError indicates that the server can't find the requested resource due to some error occurred while querying the database.
	DBError Type = "DBError"
	//Unauthorized indicates that the request lacks valid authentication credentials for the target resource.
	Unauthorized Type = "Unauthorized"
	// PermissionDenied indicates that client does not have access rights to the resource
	PermissionDenied Type = "PermissionDenied"
	// SomethingWentWrong indicates that server has encountered a situation it doesn't know how to handle
	SomethingWentWrong Type = "SomethingWentWrong"
)

// new creates a new custom error object
func (errType Type) new(msg string) error {
	return &goError{errorType: errType, originalError: errors.New(msg)}
}

// wrap wraps context with an error object
func (errType Type) wrap(err error, msg string) error {
	return &goError{errorType: errType, originalError: fmt.Errorf(fmt.Sprintf("%s: %s", msg, "%w"), err)}
}
