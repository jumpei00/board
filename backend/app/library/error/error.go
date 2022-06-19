package error

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFound     = errors.New("data not found")
	ErrUnauthorized = errors.New("not authorized")
)

type BadRequest struct {
	Message string `json:"message"`
}

func (b *BadRequest) Error() string {
	return b.Message
}

func IsBadRequest(err error) bool {
	_, ok := err.(*BadRequest)
	return ok
}

func NewErrBadRequest(errContents ErrContents, format string, arg ...interface{}) error {
	r := &BadRequest{Message: errContents.message}
	return errors.Wrapf(r, format, arg...)
}

func NewErrNotFound(format string, arg ...interface{}) error {
	return errors.Wrapf(ErrNotFound, format, arg...)
}

func NewErrUnauthorized(format string, arg ...interface{}) error {
	return errors.Wrapf(ErrUnauthorized, format, arg...)
}
