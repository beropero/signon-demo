// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThirdPartyAccount is the golang structure of table third_party_account for DAO operations like Where/Data.
type ThirdPartyAccount struct {
	g.Meta      `orm:"table:third_party_account, do:true"`
	Id          interface{} //
	UserId      interface{} // 用户id
	Name        interface{} // 平台名
	AccountInfo interface{} // 第三方账号信息
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
	DeleteAt    *gtime.Time //
}
