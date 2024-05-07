package email

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/email/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/utility/email"
)

func (c *ControllerV1) SendCode(ctx context.Context, req *v1.SendCodeReq) (res *v1.SendCodeRes, err error) {
	res = &v1.SendCodeRes{}
	if allow, _ := dao.Allow(ctx,req.Email); !allow {
		return res, gerror.New("验证码发送过于频繁")
	}
	code, err := email.EmailSendCode(ctx, req.Email)
	if err != nil {
		return
	}
	// 验证码记录
	_, err = g.Redis().Do(ctx, "SETEX", fmt.Sprintf("code.%s", req.Email), 120, code)
	dao.CodeRecord.Ctx(ctx).Data(entity.CodeRecord{
		Account: req.Email,
		Code: code,
	}).OmitEmpty().Insert()
	if err != nil {
		return
	}
	res.States = true
	return
}
