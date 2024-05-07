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
// 手机绑定
func (c *ControllerV1) BindMobile(ctx context.Context, req *v1.BindMobileReq) (res *v1.BindMobileRes, err error) {
	res = &v1.BindMobileRes{}
	search, err := g.Redis().Do(ctx, "GET", fmt.Sprintf("code.%s", req.Mobile))
	if err != nil {
		return
	}
	if search.String() != req.Code {
		return res, gerror.New("验证码错误")
	}
	// 邮箱是否已经注册
	count , err := dao.User.Ctx(ctx).Where("mobile=?",req.Mobile).Count()
	if err != nil {
		return
	}
	if count != 0 {
		return res, gerror.New("手机号已被注册")
	}
	// 从用户上下文获取用户id
	userid := service.Context().Get(ctx).User.UserId
	// 更新用户mobile
	if _, err = dao.User.Ctx(ctx).Where("id=?", userid).Update(g.Map{
		"mobile": req.Mobile,
	}); err != nil {
		return
	}
	res.Success = true
	return
}
