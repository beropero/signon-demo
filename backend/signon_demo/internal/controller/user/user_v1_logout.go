package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/user/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/service"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res = &v1.LogoutRes{}
	var (
		token  = g.RequestFromCtx(ctx).Header.Get("Authorization")
		userid = service.Context().Get(ctx).User.UserId
	)
	// 删除redis中的键
	g.Redis().Do(ctx, "DEL", fmt.Sprintf("jwt.token.%d", userid))
	// 数据库logout字段设置为1
	dao.JwtStorage.Ctx(ctx).Data("logout", 1).Where("token=?", token).Update()
	return
}
