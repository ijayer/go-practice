/*
 * 说明：
 * 作者：zhe
 * 时间：2018-09-11 3:33 PM
 * 更新：
 */

package error

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

type Error struct {
	Code        int64
	Message     string
	Temporary   bool
	UserErrCode int64
}

func (e *Error) Error() string {
	return e.Message
}

func Errorf(code codes.Code, temporary bool, msg string, args ...interface{}) error {
	return &Error{
		Code:      int64(code),
		Message:   fmt.Sprintf(msg, args...),
		Temporary: temporary,
	}
}
