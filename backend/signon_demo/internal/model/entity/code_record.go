// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeRecord is the golang structure for table code_record.
type CodeRecord struct {
	Id       int         `json:"id"       orm:"id"        ` //
	Account  string      `json:"account"  orm:"account"   ` // 账号
	Code     string      `json:"code"     orm:"code"      ` // 验证码
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` //
}
