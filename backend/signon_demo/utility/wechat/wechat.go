package wechat

import (
	"signon_demo/utility/config"

	"github.com/beropero/mysdk/wxgo"
)

var (
	cfg = &wxgo.WechatCfg{
		Token:     config.Wechat().Token,
		Appid:     config.Wechat().AppId,
		Appsecret: config.Wechat().AppSecret,
	}
	Wx = wxgo.NewWechat(cfg)
)
