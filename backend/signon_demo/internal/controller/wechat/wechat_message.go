package wechat

import (
	"fmt"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
	"signon_demo/utility/wechat"

	"github.com/beropero/mysdk/wxgo"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func Message(r *ghttp.Request) {
	vp := wxgo.VerifyParams{}
	r.ParseQuery(&vp)
	if flag, _ := wechat.Wx.VerifySignature(vp); !flag {
		return
	}
	r.Response.WriteString(vp.Echostr)
	uevent := &wxgo.UserEvent{}
	r.ParseForm(&uevent)
	// 用户事件处理
	if uevent.Ticket != "" {
		// 登陆注册场景码
		g.Dump(uevent)
		if uevent.EventKey == "1" || uevent.EventKey == "qrscene_1" {
			ctx := r.GetCtx()
			openid := uevent.FromUserName
			g.Redis().Do(ctx, "SETEX", fmt.Sprintf("ticket.%s", uevent.Ticket), 120, openid)
			wxUser := entity.WxUser{
				OpenId:  openid,
				Unionid: "",
			}
			err := service.TpAccount().CreateWxAccountIfNotExists(ctx, wxUser)
			if err != nil {
				return
			}
		}
		// 账号绑定场景码
		if uevent.EventKey == "2" || uevent.EventKey == "qrscene_2" {
			ctx := r.GetCtx()
			search, _ := g.Redis().Do(ctx, "GET", uevent.Ticket)
			userid := search.String()
			if userid == "" {
				return
			}
			count, _ := dao.ThirdPartyAccount.Ctx(ctx).Where("name= ? and JSON_CONTAINS(account_info,JSON_OBJECT(\"openid\",?))", "Wechat", uevent.FromUserName).Count()
			if count != 0 {
				return
			}
			// 添加微信账号信息
			tpacc := entity.ThirdPartyAccount{
				Name:        "Wechat",
				UserId:      gconv.Int(userid),
				AccountInfo: "",
			}
			tpacc.AccountInfo = gconv.String(g.Map{
				"openid": uevent.FromUserName,
			})
			g.Dump(tpacc.AccountInfo)
			dao.AddTpAccount(ctx, tpacc)
		}
	}
}
