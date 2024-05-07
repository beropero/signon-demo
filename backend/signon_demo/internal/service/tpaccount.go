// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"signon_demo/internal/model/entity"
)

type (
	ITpAccount interface {
		// 获取用户指定第三方平台绑定信息
		OneTpAccount(ctx context.Context, userId int, tpName string) (ThirdPartyAccount entity.ThirdPartyAccount, err error)
		// 获取所有第三方绑定信息
		AllTpAccount(ctx context.Context, userId int) (ThirdPartyAccount []entity.ThirdPartyAccount, err error)
		// 查询第三方账号绑定情况
		QueryTpAccount(ctx context.Context, userId int) (ThirdPartyAccount []entity.ThirdPartyAccount, err error)
		// 创建微信账号
		CreateWxAccountIfNotExists(ctx context.Context, wxUser entity.WxUser) error
	}
)

var (
	localTpAccount ITpAccount
)

func TpAccount() ITpAccount {
	if localTpAccount == nil {
		panic("implement not found for interface ITpAccount, forgot register?")
	}
	return localTpAccount
}

func RegisterTpAccount(i ITpAccount) {
	localTpAccount = i
}
