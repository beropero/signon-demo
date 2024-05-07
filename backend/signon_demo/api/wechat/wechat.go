// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package wechat

import (
	"context"

	"signon_demo/api/wechat/v1"
)

type IWechatV1 interface {
	GetQrCode(ctx context.Context, req *v1.GetQrCodeReq) (res *v1.GetQrCodeRes, err error)
	GetAuthQrCode(ctx context.Context, req *v1.GetAuthQrCodeReq) (res *v1.GetAuthQrCodeRes, err error)
	CheckLogin(ctx context.Context, req *v1.CheckLoginReq) (res *v1.CheckLoginRes, err error)
	AccessUserCode(ctx context.Context, req *v1.AccessUserCodeReq) (res *v1.AccessUserCodeRes, err error)
}
