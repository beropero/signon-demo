package user

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "signon_demo/api/user/v1"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
)

// 获取用户信息
func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	res = &v1.GetUserInfoRes{}
	userid := service.Context().Get(ctx).User.UserId
	user, err := service.User().UserInfo(ctx, entity.User{Id: userid})
	if err != nil {
		return
	}
	tpAcc, err := service.TpAccount().AllTpAccount(ctx, userid)
	if err != nil {
		return
	}
	gconv.Scan(user, &res.User)
	gconv.Scan(tpAcc, &res.TpAccount)
	return
}
