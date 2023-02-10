package pack

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/pkg/errno"
)

type BaseResp struct {
	StatusCode int32  `json:"status_code,omitempty"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func BuildBaseResp(err error) *BaseResp {
	return baseResp(errno.ConvertErr(err))
}

func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
