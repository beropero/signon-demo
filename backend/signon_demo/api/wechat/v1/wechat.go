package v1

import "github.com/gogf/gf/v2/frame/g"

type GetQrCodeReq struct {
	g.Meta `method:"GET" path:"/getqrcode"`
}

type GetQrCodeRes struct {
	Ticket string `json:"ticket"`
	Url    string `json:"url"`
}

type GetAuthQrCodeReq struct {
	g.Meta `method:"GET" path:"/getauthqrcode"`
}

type GetAuthQrCodeRes struct {
	Ticket string `json:"ticket"`
	Url    string `json:"url"`
}

type CheckLoginReq struct {
	g.Meta `method:"POST" path:"/checklogin"`
	Ticket string `json:"ticket"`
}

type CheckLoginRes struct {
	Token    string `json:"token"`
}

type AccessUserCodeReq struct {
	g.Meta `method:"GET" path:"/accessusercode"`
	Code string `json:"code"`
	State string `json:"state"`
}

type AccessUserCodeRes struct {}