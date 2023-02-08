package handle

import (
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/constants"
	"github.com/aldlss/MiniTikTok-Relation/app/pkg/pack"
	"github.com/cloudwego/hertz/pkg/app"
)

func SendResponse(c *app.RequestContext, resp any) {
	c.JSON(constants.OK, resp)
}

func SendBaseResponse(c *app.RequestContext, err error) {
	c.JSON(constants.OK, pack.BuildBaseResp(err))
}
