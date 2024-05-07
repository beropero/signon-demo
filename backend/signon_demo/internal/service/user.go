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
	IUser interface {
		// 用户注册
		UserSignUp(ctx context.Context, user entity.User) (userId int64, err error)
		// 用户登陆校验
		UserLoginVerify(ctx context.Context, user entity.User) (res bool, err error)
		// 获取用户基础信息
		UserInfo(ctx context.Context, user entity.User) (res entity.User, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
