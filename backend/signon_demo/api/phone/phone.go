// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package phone

import (
	"context"

	"signon_demo/api/phone/v1"
)

type IPhoneV1 interface {
	IsExist(ctx context.Context, req *v1.IsExistReq) (res *v1.IsExistRes, err error)
	SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error)
	Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
