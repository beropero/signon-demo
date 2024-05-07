// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	ucontext "signon_demo/internal/model/context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IContext interface {
		Set(r *ghttp.Request, ctx *ucontext.Ucontext)
		Get(ctx context.Context) *ucontext.Ucontext
		SetUser(ctx context.Context, user *ucontext.User)
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
