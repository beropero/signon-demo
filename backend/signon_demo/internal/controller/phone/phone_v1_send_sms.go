package phone

import (
	"context"
	"fmt"

	v1 "signon_demo/api/phone/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/utility/aliyun"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error) {
	res = &v1.SendSmsRes{}
	if allow, _ := dao.Allow(ctx, req.Mobile); !allow {
		return res, gerror.New("验证码发送过于频繁")
	}
	code, err := aliyun.SendSMS(req.Mobile)
	if err != nil {
		return
	}
	_, err = g.Redis().Do(ctx, "SETEX", fmt.Sprintf("code.%s", req.Mobile), 60, code)
	dao.CodeRecord.Ctx(ctx).Data(entity.CodeRecord{
		Account: req.Mobile,
		Code:    code,
	}).OmitEmpty().Insert()
	if err != nil {
		return
	}
	res.States = true
	return
}
