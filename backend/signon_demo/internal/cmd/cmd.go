package cmd

import (
	"context"
	"signon_demo/internal/router"
	"signon_demo/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(MiddlewareCORS)                  // 允许跨域

			router.Register(s)                     // 路由注册
			service.JwtStorage().PullTokenFromDB() // 从数据库拉取jwt令牌到Redis
			
			s.SetServerRoot("resource/public")     // 静态文件服务支持
			s.SetIndexFolder(true)
			
			s.Run()
			return nil
		},
	}
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}