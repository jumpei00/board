package error

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFound     = errors.New("data not found")
	ErrUnauthorized = errors.New("not authorized")
)

func NewErrNotFound(format string, arg ...interface{}) error {
	return errors.Wrapf(ErrNotFound, format, arg...)
}

func NewErrNoauthorized(format string, arg ...interface{}) error {
	return errors.Wrapf(ErrUnauthorized, format, arg...)
}
