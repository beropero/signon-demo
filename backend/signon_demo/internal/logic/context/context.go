package context

import (
	"context"
	ucontext "signon_demo/internal/model/context"
	"signon_demo/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sContext struct{}

func init() {
	service.RegisterContext(&sContext{})
}

func (s *sContext) Set(r *ghttp.Request, ctx *ucontext.Ucontext) {
	r.SetCtxVar(ucontext.CtxKey, ctx)
}

func (s *sContext) Get(ctx context.Context) *ucontext.Ucontext {
	value := ctx.Value(ucontext.CtxKey)
	if value == nil {
		return nil
	}
	if uctx, ok := value.(*ucontext.Ucontext); ok {
		return uctx
	}
	return nil
}

func (s *sContext) SetUser(ctx context.Context, user *ucontext.User) {
	s.Get(ctx).User = user
}
