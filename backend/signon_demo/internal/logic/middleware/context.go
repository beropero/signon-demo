package middleware

import (
	ucontext "signon_demo/internal/model/context"
	"signon_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)
// 初始化用户上下文
func  (s *sMiddleware) Uctx(r *ghttp.Request) {
	uctx := &ucontext.Ucontext{
		User: &ucontext.User{},
		Data: make(g.Map),
	}
	service.Context().Set(r, uctx)
	r.Middleware.Next()
}
