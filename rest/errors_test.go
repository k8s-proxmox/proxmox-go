package rest

import (
	"testing"
)

func TestIsNotFound(t *testing.T) {
	err := NotFoundErr
	if !IsNotFound(err) {
		t.Error("failed to confirm it's \"not found error\"")
	}

	err = NewError(403, "", nil)
	if IsNotFound(err) {
		t.Errorf("failed to confirm err=%v is not \"not found error\"", err)
	}
}

func TestIsNotAuthorized(t *testing.T) {
	err := NewError(401, "", nil)
	if !IsNotAuthorized(err) {
		t.Error("failed to confirm it's \"not authorized error\"")
	}

	err = NotFoundErr
	if IsNotAuthorized(err) {
		t.Errorf("failed to confirm err=%v is not \"not authorized error\"", err)
	}
}
