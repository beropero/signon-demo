// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IJwtStorage interface {
		// 保存jwt令牌
		GetJwtAndSave(ctx context.Context, userid int) (token string, err error)
		PullTokenFromDB()
	}
)

var (
	localJwtStorage IJwtStorage
)

func JwtStorage() IJwtStorage {
	if localJwtStorage == nil {
		panic("implement not found for interface IJwtStorage, forgot register?")
	}
	return localJwtStorage
}

func RegisterJwtStorage(i IJwtStorage) {
	localJwtStorage = i
}
