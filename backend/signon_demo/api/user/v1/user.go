package v1

import (
	"signon_demo/internal/model/vo"

	"github.com/gogf/gf/v2/frame/g"
)

type GetUserInfoReq struct {
	g.Meta `method:"Get" path:"/getuserinfo"`
}

type GetUserInfoRes struct {
	User      *vo.UserVO
	TpAccount []*vo.TpAccountVO
}

type BindEmailReq struct {
	g.Meta `method:"Post" path:"/bindemail"`
	Email  string `json:"email" v:"required"`
	Code   string `json:"code"  v:"required"`
}

type BindEmailRes struct {
	Success bool `json:"success"`
}

type BindMobileReq struct {
	g.Meta `method:"Post" path:"/bindmobile"`
	Mobile string `json:"mobile" v:"required"`
	Code   string `json:"code"   v:"required"`
}

type BindMobileRes struct {
	Success bool `json:"success"`
}

type BindWechatReq struct {
	g.Meta `method:"Get" path:"/bindwechat"`
}

type BindWechatRes struct {
	Ticket string `json:"ticket"`
	Url    string `json:"url"`
}

type LogoutReq struct {
	g.Meta `method:"Get" path:"/logout"`
}

type LogoutRes struct {}

