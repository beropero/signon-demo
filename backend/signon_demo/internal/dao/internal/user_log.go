// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLogDao is the data access object for table user_log.
type UserLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns UserLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserLogColumns defines and stores column names for table user_log.
type UserLogColumns struct {
	Id       string //
	UserId   string // 用户id
	Type     string // 登陆类型
	Ip       string // ip地址
	Device   string // 设备
	Location string // 坐标
	CreateAt string //
	UpdateAt string //
	DeleteAt string //
}

// userLogColumns holds the columns for table user_log.
var userLogColumns = UserLogColumns{
	Id:       "id",
	UserId:   "user_id",
	Type:     "type",
	Ip:       "ip",
	Device:   "device",
	Location: "location",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}

// NewUserLogDao creates and returns a new DAO object for table data access.
func NewUserLogDao() *UserLogDao {
	return &UserLogDao{
		group:   "default",
		table:   "user_log",
		columns: userLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserLogDao) Columns() UserLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
