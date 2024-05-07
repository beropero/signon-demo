// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// JwtStorage is the golang structure of table jwt_storage for DAO operations like Where/Data.
type JwtStorage struct {
	g.Meta   `orm:"table:jwt_storage, do:true"`
	Id       interface{} //
	UserId   interface{} // 用户id
	Token    interface{} // jwt令牌
	Logout   interface{} // 是否注销
	ExpireAt *gtime.Time // 失效时间
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
