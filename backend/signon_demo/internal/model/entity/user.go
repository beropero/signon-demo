// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `json:"id"       orm:"id"        ` //
	Username string      `json:"username" orm:"username"  ` // 用户名
	Nickname string      `json:"nickname" orm:"nickname"  ` // 昵称
	Password string      `json:"password" orm:"password"  ` // 密码
	Mobile   string      `json:"mobile"   orm:"mobile"    ` // 手机号
	Email    string      `json:"email"    orm:"email"     ` // 邮箱
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` //
}
