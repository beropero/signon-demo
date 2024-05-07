package email

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/email/v1"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if req.Password != "" {
		success, _ := service.User().UserLoginVerify(ctx, user)
		if !success {
			return res, gerror.New("邮箱或密码错误")
		}
	}
	if req.VerificationCode != "" {
		search, _ := g.Redis().Do(ctx, "GET", fmt.Sprintf("code.%s", req.Email))
		code := search.String()
		if code != req.VerificationCode {
			return res, gerror.New("验证码错误或未发送")
		}
	}
	search, _ := service.User().UserInfo(ctx,user)
	// 保存jwt令牌
	token, _ := service.JwtStorage().GetJwtAndSave(ctx, search.Id)
	res = &v1.LoginRes{
		Token: token,
	}
	return
}
