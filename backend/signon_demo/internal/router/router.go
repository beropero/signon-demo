package router

import (
	"signon_demo/internal/controller/email"
	"signon_demo/internal/controller/phone"
	"signon_demo/internal/controller/user"
	"signon_demo/internal/controller/wechat"
	"signon_demo/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// 路由注册额
func Register(s *ghttp.Server) {
	// 全局用户上下文信息中间件
	s.Use(service.Middleware().Uctx)
	// 微信配置+事件推送接口
	s.BindHandler("/wechat/message", wechat.Message)
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 响应中间件
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		
		// 用户接口
		group.Group("/user",func(group *ghttp.RouterGroup) {
			// JWT校验拦截器
			group.Middleware(service.Middleware().JwtInterceptor)
            group.Bind(
				user.NewV1(),
			)
		})
		// 登录注册接口
		group.Group("/signon", func(group *ghttp.RouterGroup) {
			group.Group("/user", func(group *ghttp.RouterGroup) {			
				// email 相关接口
				group.Group("/email", func(group *ghttp.RouterGroup) {
					group.Bind(
						email.NewV1(),
					)
				})
				// moblie 相关接口
				group.Group("/phone", func(group *ghttp.RouterGroup) {
					group.Bind(
						phone.NewV1(),
					)
				})
				// wechat 相关接口
				group.Group("/wechat", func(group *ghttp.RouterGroup) {
					group.Bind(
						wechat.NewV1(),
					)
				})
			})
		})
	})
}
