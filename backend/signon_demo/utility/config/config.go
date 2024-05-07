package config

import (
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
)

// 配置结构体
type Cfg struct {
	Email  EmailCfg  `json:"email"`
	Aliyun AliyunCfg `json:"aliyun"`
	Wechat WechatCfg `json:"wechat"`
	Jwt    JwtCfg    `json:"jwt"`
	Code   CodeCfg   `json:"code"`
}
type EmailCfg struct {
	Username string `json:"username"`
	Password string `json:"Password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}
type AliyunCfg struct {
	RegionId        string `json:"regionid"`
	AccessKeyId     string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
	SignName        string `json:"signname"`
	TemplateCode    string `json:"templatecode"`
	Scheme          string `json:"scheme"`
}

type WechatCfg struct {
	Token     string `json:"username"`
	AppId     string `json:"appid"`
	AppSecret string `json:"appsecret"`
}

type JwtCfg struct {
	Key      string `json:"key"`
	Duration int    `json:"duration"`
}

type CodeCfg struct {
	MinLimit  int `json:"min_limit"`
	HourLimit int `json:"hour_limit"`
	DayLimit  int `json:"day_limit"`
}

var (
	cfg = NewCfg()
)

func NewCfg() (cfg *Cfg) {
	ctx := gctx.New()
	g.Cfg().MustGet(ctx, "utility").Scan(&cfg)
	return
}

func Email() EmailCfg {
	return cfg.Email
}

func Jwt() JwtCfg {
	return cfg.Jwt
}

func Aliyun() AliyunCfg {
	return cfg.Aliyun
}

func Wechat() WechatCfg {
	return cfg.Wechat
}

func Code() CodeCfg {
	return cfg.Code
}