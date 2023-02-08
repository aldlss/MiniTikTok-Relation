package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode         = 0
	ParamErrCode        = 1001
	ServiceErrCode      = 1002
	DatabaseErrCode     = 1003
	NilValueErrCode     = 1004
	UnclassifiedErrCode = 1005
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("ErrCode:%d, ErrMsg:%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WriteMsg(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}
	return UnclassifiedErr.WriteMsg(err.Error())
}

var (
	Success         = NewErrNo(SuccessCode, "Success")
	ParamErr        = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	ServiceErr      = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	DatabaseErr     = NewErrNo(DatabaseErrCode, "Error when operate database")
	NilValueErr     = NewErrNo(NilValueErrCode, "Value is nil or doesn't exist")
	UnclassifiedErr = NewErrNo(UnclassifiedErrCode, "Maybe it's an error returned by the framework used")
)
