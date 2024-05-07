// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"signon_demo/internal/dao/internal"
	"signon_demo/utility/config"
	"time"
)

// internalCodeRecordDao is internal type for wrapping internal DAO implements.
type internalCodeRecordDao = *internal.CodeRecordDao

// codeRecordDao is the data access object for table code_record.
// You can define custom methods on it to extend its functionality as you wish.
type codeRecordDao struct {
	internalCodeRecordDao
}

var (
	// CodeRecord is globally public accessible object for table code_record operations.
	CodeRecord = codeRecordDao{
		internal.NewCodeRecordDao(),
	}
)

// Fill with you ideas below.
func Allow(ctx context.Context, acc string) (res bool, err error) {
	model := CodeRecord.Ctx(ctx)
	// 一分钟内发送次数
	min, err := model.Where("account=? and create_at>?", acc, time.Now().Add(-time.Minute*1)).Count()
	if err != nil {
		return
	}
	// 一小时内发送次数
	hour, err := model.Where("account=? and create_at>?", acc, time.Now().Add(-time.Hour*1)).Count()
	if err != nil {
		return
	}
	// 一天内发送次数
	day, err := model.Where("account=? and create_at>?", acc, time.Now().Add(-time.Hour*24)).Count()
	if err != nil {
		return
	}
	// 配置信息
	cfg := config.Code()
	res = min < cfg.MinLimit && hour < cfg.HourLimit && day < cfg.DayLimit
	return res, nil
}
