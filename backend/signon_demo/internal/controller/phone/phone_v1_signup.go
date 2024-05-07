package phone

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/phone/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
)

func (c *ControllerV1) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {
	user := entity.User{
		Username: req.Mobile,
		Mobile:    req.Mobile,
		Password: req.Password,
	}
	g.Dump(req)
	localUser := service.User()
	// 从redis中查找是否存在code
	code, _ := g.Redis().Do(ctx, "GET", fmt.Sprintf("code.%s", req.Mobile))
	if req.VerificationCode != code.String() {
		return res, gerror.New("验证码错误或未发送")
	}
	// 用户是否存在
	if exist := dao.IsExistUser(ctx, entity.User{
		Mobile: user.Mobile,
	}); exist{
		return res, gerror.New("手机号已注册")
	}
	// 注册用户
	userId, err := localUser.UserSignUp(ctx, user)
	if err != nil {
		return
	}
	// 获取并保存jwt令牌
	token, err := service.JwtStorage().GetJwtAndSave(ctx, int(userId))
	// 保存jwt令牌
	res = &v1.SignupRes{
		Token: token,
	}
	return
}
