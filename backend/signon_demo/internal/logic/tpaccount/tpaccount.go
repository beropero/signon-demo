package tpaccount

import (
	"context"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sTpAccount struct{}

func init() {
	service.RegisterTpAccount(&sTpAccount{})
}
// 用户绑定第三方账号

// 获取用户指定第三方平台绑定信息
func (s *sTpAccount) OneTpAccount(ctx context.Context, userId int, tpName string) (ThirdPartyAccount entity.ThirdPartyAccount, err error) {
	ThirdPartyAccount, err = dao.OneTpAccount(ctx, entity.ThirdPartyAccount{
		UserId: userId,
		Name:   tpName,
	})
	return
}

// 获取所有第三方绑定信息
func (s *sTpAccount) AllTpAccount(ctx context.Context, userId int) (ThirdPartyAccount []entity.ThirdPartyAccount, err error) {
	ThirdPartyAccount, err = dao.AllTpAccount(ctx, userId)
	return
}

// 查询第三方账号绑定情况
func (s *sTpAccount) QueryTpAccount(ctx context.Context, userId int) (ThirdPartyAccount []entity.ThirdPartyAccount, err error) {
	ThirdPartyAccount, err = dao.AllTpAccount(ctx, userId)
	return
}

// 微信账号注册
func (s *sTpAccount) CreateWxAccountIfNotExists(ctx context.Context, wxUser entity.WxUser) error {
	// 开启事务
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		count, err := dao.ThirdPartyAccount.Ctx(ctx).Where("name= ? and JSON_CONTAINS(account_info,JSON_OBJECT(\"openid\",?))", "Wechat", wxUser.OpenId).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			// 创建新用户并获取用户id
			userid, err := dao.AddUser(ctx, entity.User{
				Username: wxUser.OpenId,
			})
			if err != nil {
				return err
			}
			// 为用户添加微信账号信息
			_, err = dao.ThirdPartyAccount.Ctx(ctx).Data(entity.ThirdPartyAccount{
				UserId:      int(userid),
				Name:        "Wechat",
				AccountInfo: gconv.String(wxUser),
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
