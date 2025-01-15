package httpclient

import (
	"fmt"
	"time"
)

// HTTPError - custom http client error.
type HTTPError struct {
	err   error
	retry int
}

func NewHTTPError(err error, retry int) error {
	if err == nil {
		return nil
	}

	return &HTTPError{
		err:   err,
		retry: retry,
	}
}

func (e *HTTPError) Error() string {
	if e.retry == 0 {
		return fmt.Sprintf("%v %s", e.err, time.Now().Format("2006/01/02 15:04:05"))
	}

	return fmt.Sprintf("%v attempt: %d %s", e.err, e.retry, time.Now().Format("2006/01/02 15:04:05"))
}

func (e *HTTPError) Unwrap() error {
	return e.err
}

type Error struct {
	err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("[client]: %v", e.err)
}

func NewError(err error) error {
	if err == nil {
		return nil
	}

	return &Error{
		err: err,
	}
}

func (e *Error) Unwrap() error {
	return e.err
}
