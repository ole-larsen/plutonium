package repository

import (
	"fmt"
)

type Error struct {
	Err error
}

func NewError(err error) error {
	if err == nil {
		return nil
	}

	return &Error{
		Err: err,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[repository]: %s", e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}
