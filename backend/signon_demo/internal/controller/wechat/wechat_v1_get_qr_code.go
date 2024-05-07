package wechat

import (
	"context"

	"signon_demo/api/wechat/v1"
	"signon_demo/utility/wechat"
)

func (c *ControllerV1) GetQrCode(ctx context.Context, req *v1.GetQrCodeReq) (res *v1.GetQrCodeRes, err error) {
	res = &v1.GetQrCodeRes{}
	ticket, err := wechat.Wx.GetQRTicketReq("QR_SCENE", 01)
	if err != nil {
		return 
	}
	url := wechat.Wx.GetQrImageUrl(ticket)
	res.Url = url
	res.Ticket = ticket
	return 
}
