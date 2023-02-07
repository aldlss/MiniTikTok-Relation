package errno

import "fmt"

const (
	SuccessCode     = 0
	ParamErrCode    = 1001
	ServiceErrCode  = 1002
	DatabaseErrCode = 1003
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

var (
	Success     = NewErrNo(SuccessCode, "Success")
	ParamErr    = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	ServiceErr  = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	DatabaseErr = NewErrNo(DatabaseErrCode, "Error when operate database")
)
