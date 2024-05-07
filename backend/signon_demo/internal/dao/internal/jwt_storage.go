// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// JwtStorageDao is the data access object for table jwt_storage.
type JwtStorageDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns JwtStorageColumns // columns contains all the column names of Table for convenient usage.
}

// JwtStorageColumns defines and stores column names for table jwt_storage.
type JwtStorageColumns struct {
	Id       string //
	UserId   string // 用户id
	Token    string // jwt令牌
	Logout   string // 是否注销
	ExpireAt string // 失效时间
	CreateAt string //
	UpdateAt string //
	DeleteAt string //
}

// jwtStorageColumns holds the columns for table jwt_storage.
var jwtStorageColumns = JwtStorageColumns{
	Id:       "id",
	UserId:   "user_id",
	Token:    "token",
	Logout:   "logout",
	ExpireAt: "expire_at",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}

// NewJwtStorageDao creates and returns a new DAO object for table data access.
func NewJwtStorageDao() *JwtStorageDao {
	return &JwtStorageDao{
		group:   "default",
		table:   "jwt_storage",
		columns: jwtStorageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *JwtStorageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *JwtStorageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *JwtStorageDao) Columns() JwtStorageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *JwtStorageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *JwtStorageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *JwtStorageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
