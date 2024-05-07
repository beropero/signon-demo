package wechat

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	v1 "signon_demo/api/wechat/v1"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
	"signon_demo/utility/config"
)

func (c *ControllerV1) CheckLogin(ctx context.Context, req *v1.CheckLoginReq) (res *v1.CheckLoginRes, err error) {
	res = &v1.CheckLoginRes{}
	search, err := g.Redis().Do(ctx, "GET", fmt.Sprintf("ticket.%s", req.Ticket))
	if err != nil {
		return
	}
	openid := search.String()
	if openid == "" {
		return
	}
	acc := entity.ThirdPartyAccount{}
	if err = dao.ThirdPartyAccount.Ctx(ctx).Where("name= ? and JSON_CONTAINS(account_info,JSON_OBJECT(\"openid\",?))", "Wechat", openid).Scan(&acc); err != nil {
		return
	}
	// 保存jwt令牌
	token, err := service.JwtStorage().GetJwtAndSave(ctx, acc.UserId)
	if err != nil {
		return
	}
	g.Redis().Do(ctx, "SETEX", fmt.Sprintf("jwt.token.%d", acc.UserId), config.Jwt().Duration*60*60, token)
	res.Token = token
	return
}
