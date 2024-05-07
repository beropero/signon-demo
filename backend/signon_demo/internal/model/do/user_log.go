// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure of table user_log for DAO operations like Where/Data.
type UserLog struct {
	g.Meta   `orm:"table:user_log, do:true"`
	Id       interface{} //
	UserId   interface{} // 用户id
	Type     interface{} // 登陆类型
	Ip       interface{} // ip地址
	Device   interface{} // 设备
	Location interface{} // 坐标
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
