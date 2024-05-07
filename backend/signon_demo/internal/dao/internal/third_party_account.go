// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThirdPartyAccountDao is the data access object for table third_party_account.
type ThirdPartyAccountDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns ThirdPartyAccountColumns // columns contains all the column names of Table for convenient usage.
}

// ThirdPartyAccountColumns defines and stores column names for table third_party_account.
type ThirdPartyAccountColumns struct {
	Id          string //
	UserId      string // 用户id
	Name        string // 平台名
	AccountInfo string // 第三方账号信息
	CreateAt    string //
	UpdateAt    string //
	DeleteAt    string //
}

// thirdPartyAccountColumns holds the columns for table third_party_account.
var thirdPartyAccountColumns = ThirdPartyAccountColumns{
	Id:          "id",
	UserId:      "user_id",
	Name:        "name",
	AccountInfo: "account_info",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
	DeleteAt:    "delete_at",
}

// NewThirdPartyAccountDao creates and returns a new DAO object for table data access.
func NewThirdPartyAccountDao() *ThirdPartyAccountDao {
	return &ThirdPartyAccountDao{
		group:   "default",
		table:   "third_party_account",
		columns: thirdPartyAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ThirdPartyAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ThirdPartyAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ThirdPartyAccountDao) Columns() ThirdPartyAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ThirdPartyAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ThirdPartyAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ThirdPartyAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
