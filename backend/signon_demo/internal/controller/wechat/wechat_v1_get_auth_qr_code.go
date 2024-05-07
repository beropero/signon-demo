package wechat

import (
	"context"
	"fmt"
	"net/url"

	"github.com/beropero/mysdk/wxgo"

	v1 "signon_demo/api/wechat/v1"
	"signon_demo/utility/wechat"
)

func (c *ControllerV1) GetAuthQrCode(ctx context.Context, req *v1.GetAuthQrCodeReq) (res *v1.GetAuthQrCodeRes, err error) {
	res = &v1.GetAuthQrCodeRes{}
	ticket := wxgo.GenerateRandomTicket(20)
	savedir := "resource/public/resource/image"
	authurl := wechat.Wx.GetOauth2CodeUrl("http://39.101.78.10/signon/user/wechat/accessusercode","snsapi_userinfo",ticket)
	qrurl := fmt.Sprintf("%s?url=%s","http://39.101.78.10/html/auth.html",url.QueryEscape(authurl))
	err = wxgo.GenerateQrCode(qrurl, savedir, ticket)
	if err != nil {
		return
	}
	url := fmt.Sprintf("http://39.101.78.10/resource/image/%s.png", ticket)
	res = &v1.GetAuthQrCodeRes{
		Ticket: ticket,
		Url:    url,
	}
	return
}
