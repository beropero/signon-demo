// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure for table user_log.
type UserLog struct {
	Id       int         `json:"id"       orm:"id"        ` //
	UserId   int         `json:"userId"   orm:"user_id"   ` // 用户id
	Type     string      `json:"type"     orm:"type"      ` // 登陆类型
	Ip       string      `json:"ip"       orm:"ip"        ` // ip地址
	Device   string      `json:"device"   orm:"device"    ` // 设备
	Location int         `json:"location" orm:"location"  ` // 坐标
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` //
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` //
}
