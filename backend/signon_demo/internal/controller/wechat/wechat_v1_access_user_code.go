package wechat

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "signon_demo/api/wechat/v1"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
	"signon_demo/utility/wechat"
)

func (c *ControllerV1) AccessUserCode(ctx context.Context, req *v1.AccessUserCodeReq) (res *v1.AccessUserCodeRes, err error) {
	res = &v1.AccessUserCodeRes{}
	uat, err := wechat.Wx.GetUserATReq(req.Code)
	if err != nil {
		return
	}
	uInfo, err := wechat.Wx.GetUserInfoReq(uat)
	if err != nil {
		return
	}
	_, err = g.Redis().Do(ctx, "SETEX", fmt.Sprintf("ticket.%s", req.State), 120, uInfo.OpenId)
	if err != nil {
		return
	}
	wxUser := entity.WxUser{}
	gconv.Scan(uInfo, &wxUser)
	err = service.TpAccount().CreateWxAccountIfNotExists(ctx, wxUser)
	if err != nil {
		return res, err
	}
	g.RequestFromCtx(ctx).Response.RedirectTo("http://39.101.78.10/html/success.html")
	return
}
