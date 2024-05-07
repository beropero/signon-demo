package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/user/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/service"
)

// 邮箱绑定
func (c *ControllerV1) BindEmail(ctx context.Context, req *v1.BindEmailReq) (res *v1.BindEmailRes, err error) {
	res = &v1.BindEmailRes{}
	search, err := g.Redis().Do(ctx, "GET", fmt.Sprintf("code.%s", req.Email))
	if err != nil {
		return
	}
	if search.String() != req.Code {
		return res, gerror.New("验证码错误")
	}
	// 邮箱是否已经注册
	count , err := dao.User.Ctx(ctx).Where("email=?",req.Email).Count()
	if err != nil {
		return
	}
	if count != 0 {
		return res, gerror.New("邮箱已被注册")
	}
	// 从用户上下文获取用户id
	userid := service.Context().Get(ctx).User.UserId
	// 更新用户email
	if _, err = dao.User.Ctx(ctx).Where("id=?", userid).Update(g.Map{
		"email": req.Email,
	}); err != nil {
		return
	}
	res.Success = true
	return
}
