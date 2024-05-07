package middleware

import (
	"fmt"
	"net/http"
	ucontext "signon_demo/internal/model/context"
	"signon_demo/internal/service"
	"signon_demo/utility/config"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)
// JWT校验拦截器 
func (s *sMiddleware) JwtInterceptor(r *ghttp.Request) {
	var (
		jwtKey      = []byte(config.Jwt().Key)
		tokenString = r.Header.Get("Authorization")
	)
	// 解析JWT令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// token是否有效
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	// 获取用户id
	claims := token.Claims.(jwt.MapClaims)
	userid := gconv.Int(claims["userid"])
	// Redis中是否存在JWT
	ctx := r.GetCtx()
	search, _ := g.Redis().Do(ctx, "GET", fmt.Sprintf("jwt.token.%d", userid))
	if search.String() != tokenString {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	// 上下文用户信息设置
    service.Context().SetUser(r.GetCtx(),&ucontext.User{
		UserId: userid,
	})
	r.Middleware.Next()
}
