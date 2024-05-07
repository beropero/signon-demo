// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeRecord is the golang structure of table code_record for DAO operations like Where/Data.
type CodeRecord struct {
	g.Meta   `orm:"table:code_record, do:true"`
	Id       interface{} //
	Account  interface{} // 账号
	Code     interface{} // 验证码
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
