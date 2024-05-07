// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// JwtStorage is the golang structure for table jwt_storage.
type JwtStorage struct {
	Id       int         `json:"id"       orm:"id"        ` //
	UserId   int         `json:"userId"   orm:"user_id"   ` // 用户id
	Token    string      `json:"token"    orm:"token"     ` // jwt令牌
	Logout   int         `json:"logout"   orm:"logout"    ` // 是否注销
	ExpireAt *gtime.Time `json:"expireAt" orm:"expire_at" ` // 失效时间
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` //
}
