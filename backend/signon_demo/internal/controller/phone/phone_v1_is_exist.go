package phone

import (
	"context"

	"signon_demo/api/phone/v1"
	"signon_demo/internal/dao"
)

func (c *ControllerV1) IsExist(ctx context.Context, req *v1.IsExistReq) (res *v1.IsExistRes, err error) {
	res = &v1.IsExistRes{}
	model := dao.User.Ctx(ctx)
	count, err := model.Where("mobile=?",req.Mobile).Count()
	if count == 1 {
		res.IsExist = true
	}
	return 
}
