package email

import (
	"context"

	v1 "signon_demo/api/email/v1"
	"signon_demo/internal/dao"
)

func (c *ControllerV1) IsExist(ctx context.Context, req *v1.IsExistReq) (res *v1.IsExistRes, err error) {
	res = &v1.IsExistRes{}
	model := dao.User.Ctx(ctx)
	count, err := model.Where("email=?", req.Email).Count()
	if count == 1 {
		res.IsExist = true
	}
	return
}
