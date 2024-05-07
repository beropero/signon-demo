package user

import (
	"context"

	v1 "signon_demo/api/user/v1"
	"signon_demo/internal/service"
	"signon_demo/utility/wechat"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) BindWechat(ctx context.Context, req *v1.BindWechatReq) (res *v1.BindWechatRes, err error) {
	res = &v1.BindWechatRes{}
	userid := service.Context().Get(ctx).User.UserId
	ticket, err := wechat.Wx.GetQRTicketReq("QR_SCENE", 02)
	g.Redis().Do(ctx, "SETEX", ticket, 60, userid)
	if err != nil {
		return
	}
	url := wechat.Wx.GetQrImageUrl(ticket)
	res.Url = url
	res.Ticket = ticket
	return
}
