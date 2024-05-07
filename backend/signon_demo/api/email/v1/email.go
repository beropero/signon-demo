package v1

import "github.com/gogf/gf/v2/frame/g"

// 是否已注册
type IsExistReq struct {
	g.Meta `method:"POST" path:"/isexist"`
	Email  string `json:"email" v:"required"`
}

type IsExistRes struct {
	IsExist bool `json:"exists"`
}

// 发送验证码
type SendCodeReq struct {
	g.Meta `method:"POST" path:"/sendcode"`
	Email  string `json:"email"`
}

type SendCodeRes struct {
	States bool `json:"states"`
}

// 注册
type SignupReq struct {
	g.Meta           `method:"POST" path:"/signup"`
	Email            string `json:"email" v:"required"`
	Password         string `json:"password" v:"required|ci|same:Password2"`
	Password2        string `json:"password2" v:"required"`
	VerificationCode string `json:"verification_code" v:"required"`
}

type SignupRes struct {
	Token string `json:"token"`
}

// 登陆
type LoginReq struct {
	g.Meta           `method:"POST" path:"/login"`
	Email            string `json:"email" v:"required"`
	Password         string `json:"password"`
	VerificationCode string `json:"verification_code"`
}

type LoginRes struct {
	Token string `json:"token"`
}
