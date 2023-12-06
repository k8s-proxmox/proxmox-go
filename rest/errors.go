package rest

import (
	"fmt"
	"net/http"
)

const (
	NotFound = "resource not found"
)

var (
	NotFoundErr = NewError(http.StatusNotFound, NotFound, []byte("from k8s-proxmox/proxmox-go"))
)

func NewError(code int, status string, body []byte) *Error {
	return &Error{code: code, status: status, body: string(body)}
}

type Error struct {
	code   int
	status string
	body   string
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf("%d - %s - %s", e.code, e.status, e.body)
}

func IsNotFound(err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.code == http.StatusNotFound
}

func IsNotAuthorized(err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.code == http.StatusUnauthorized
}
