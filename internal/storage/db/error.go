// Package db contains custom errors for database
package db

import (
	"fmt"
	"time"
)

type Error struct {
	Err   error
	Retry int
}

func NewError(err error, retry int) error {
	if err == nil {
		return nil
	}

	return &Error{
		Err:   err,
		Retry: retry,
	}
}

func (e *Error) Error() string {
	if e.Retry == 0 {
		return fmt.Sprintf("%v %s", e.Err, time.Now().Format("2006/01/02 15:04:05"))
	}

	return fmt.Sprintf("%v attempt: %d %s", e.Err, e.Retry, time.Now().Format("2006/01/02 15:04:05"))
}

func (e *Error) Unwrap() error {
	return e.Err
}
