/**
@author: yeebing
@date: 2022/9/24
**/

package errno

import (
	"errors"
	"fmt"
	pkgerr "github.com/pkg/errors"
)

type withCode struct {
	code int
	err  error
}

// WithCode create err with code and error.
func WithCode(code int, err error) error {
	return &withCode{
		code: code,
		err:  err,
	}
}

// WithCodef create err with code and string.
func WithCodef(code int, format string, args ...interface{}) error {
	return &withCode{
		code: code,
		//err:  fmt.Errorf(format, args...),
		err: errors.New(fmt.Sprint(format, args)),
	}
}

func (c *withCode) Code() int {
	return c.code
}

func (c *withCode) Error() string {
	return c.err.Error()
}

// StackError return the row error, can be used to print stack information.
func StackError(err error) error {
	if c, ok := err.(*withCode); ok {
		return c.err
	}
	return err
}

// Wrap a string message to error.
func Wrap(err error, msg string) error {
	if c, ok := err.(*withCode); ok {
		c.err = pkgerr.Wrap(c.err, msg)
		return c
	}

	return pkgerr.Wrap(err, msg)
}

// Is check if two error types are equal.
func Is(lhs, rhs error) bool {
	if c, ok := lhs.(*withCode); ok {
		lhs = c.err
	}
	if c, ok := rhs.(*withCode); ok {
		rhs = c.err
	}

	return errors.Is(lhs, rhs)
}

// Unwrap returns the result of calling the Unwrap method on err,
// if errs type contains an Unwrap method returning error. Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	if c, ok := err.(*withCode); ok {
		return errors.Unwrap(c.err)
	}

	return errors.Unwrap(err)
}
