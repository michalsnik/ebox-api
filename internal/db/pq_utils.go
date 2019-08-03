package db

import (
	"errors"

	"github.com/lib/pq"
)

const (
	// pqUniqueConstraintViolationCode is the code used by lib/pq to represent a postgres unique constraint violation error
	pqUniqueConstraintViolationCode = "23505"
)

var (
	// ErrUniqueConstraintViolation is the error returned when a unique constraint violation error is found
	ErrUniqueConstraintViolation = errors.New("unique constraint ")
)

// GetError takes a standard error and checks to see if it is a pq.Error.  If so, it attempts to create a clean
// error to be returned so it can be gracefully handled.  If a known error can not be found for the provided error,
// that error is returned.
func GetError(err error) error {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	var errToReturn error
	switch code := pqErr.Code; code {
	case pqUniqueConstraintViolationCode:
		errToReturn = ErrUniqueConstraintViolation
	default:
		errToReturn = err
	}

	return errToReturn
}
