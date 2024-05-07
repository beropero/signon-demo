// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CodeRecordDao is the data access object for table code_record.
type CodeRecordDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CodeRecordColumns // columns contains all the column names of Table for convenient usage.
}

// CodeRecordColumns defines and stores column names for table code_record.
type CodeRecordColumns struct {
	Id       string //
	Account  string // 账号
	Code     string // 验证码
	CreateAt string //
	UpdateAt string //
	DeleteAt string //
}

// codeRecordColumns holds the columns for table code_record.
var codeRecordColumns = CodeRecordColumns{
	Id:       "id",
	Account:  "account",
	Code:     "code",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}

// NewCodeRecordDao creates and returns a new DAO object for table data access.
func NewCodeRecordDao() *CodeRecordDao {
	return &CodeRecordDao{
		group:   "default",
		table:   "code_record",
		columns: codeRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CodeRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CodeRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CodeRecordDao) Columns() CodeRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CodeRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CodeRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CodeRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
