// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package email

import (
	"context"

	"signon_demo/api/email/v1"
)

type IEmailV1 interface {
	IsExist(ctx context.Context, req *v1.IsExistReq) (res *v1.IsExistRes, err error)
	SendCode(ctx context.Context, req *v1.SendCodeReq) (res *v1.SendCodeRes, err error)
	Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
