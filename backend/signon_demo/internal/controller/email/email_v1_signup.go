package email

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/email/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
)

func (c *ControllerV1) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {
	user := entity.User{
		Username: req.Email,
		Email:    req.Email,
		Password: req.Password,
	}
	localUser := service.User()
	// 从redis中查找是否存在code
	search, _ := g.Redis().Do(ctx, "GET", fmt.Sprintf("code.%s", req.Email))
	code := search.String()
	if req.VerificationCode != code {
		return res, gerror.New("验证码错误或未发送")
	}
	// 用户是否存在
	if exist := dao.IsExistUser(ctx, entity.User{
		Email: user.Email,
	}); exist{
		return res, gerror.New("邮箱已注册")
	}
	
	// 注册用户
	userId, err := localUser.UserSignUp(ctx, user)
	if err != nil {
		return
	}
	// 获取并保存jwt令牌
	token, err := service.JwtStorage().GetJwtAndSave(ctx, int(userId))
	// 返回jwt令牌
	res = &v1.SignupRes{
		Token: token,
	}
	return
}
