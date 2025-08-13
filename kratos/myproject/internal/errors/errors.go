package errors

import (
	"myproject/internal/proto/hello"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func New(reason hello.ErrorReason, format string, a ...any) error {
	// todo: add multiple language support
	return errors.Errorf(int(reason), reason.String(), format, a...)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target error) bool {
	return errors.As(err, target)
}

func Reason(err error) string {
	return errors.Reason(err)
}

func Wrap(err error, format string, a ...any) error {
	e := errors.FromError(err)
	e.Code = int32(hello.ErrorReason_INTERNAL_SERVER_ERROR)
	e.Message = fmt.Sprintf(format, a...)

	// todo: add multiple language support
	return e
}
