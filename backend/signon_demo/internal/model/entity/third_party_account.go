// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThirdPartyAccount is the golang structure for table third_party_account.
type ThirdPartyAccount struct {
	Id          int         `json:"id"          orm:"id"           ` //
	UserId      int         `json:"userId"      orm:"user_id"      ` // 用户id
	Name        string      `json:"name"        orm:"name"         ` // 平台名
	AccountInfo string      `json:"accountInfo" orm:"account_info" ` // 第三方账号信息
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    ` //
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    ` //
	DeleteAt    *gtime.Time `json:"deleteAt"    orm:"delete_at"    ` //
}
