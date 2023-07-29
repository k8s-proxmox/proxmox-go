package rest

import "net/http"

const (
	NotFound      = "resource not found"
	NotAuthorized = "not authorized"
)

var (
	NotFoundErr      = &Error{code: http.StatusNotFound, returnMessage: NotFound}
	NotAuthorizedErr = &Error{code: http.StatusUnauthorized, returnMessage: NotAuthorized}
)

type Error struct {
	code          int
	returnMessage string
}

func (e *Error) Error() string {
	return e.returnMessage
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == NotFound
}

func IsNotAuthorized(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == NotAuthorized
}
