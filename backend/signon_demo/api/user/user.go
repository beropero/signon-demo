// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"signon_demo/api/user/v1"
)

type IUserV1 interface {
	GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error)
	BindEmail(ctx context.Context, req *v1.BindEmailReq) (res *v1.BindEmailRes, err error)
	BindMobile(ctx context.Context, req *v1.BindMobileReq) (res *v1.BindMobileRes, err error)
	BindWechat(ctx context.Context, req *v1.BindWechatReq) (res *v1.BindWechatRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}
