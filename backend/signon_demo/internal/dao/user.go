// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"signon_demo/internal/dao/internal"
	"signon_demo/internal/model/entity"
)

// internalUserDao is internal type for wrapping internal DAO implements.
type internalUserDao = *internal.UserDao

// userDao is the data access object for table user.
// You can define custom methods on it to extend its functionality as you wish.
type userDao struct {
	internalUserDao
}

var (
	// User is globally public accessible object for table user operations.
	User = userDao{
		internal.NewUserDao(),
	}
)

// Fill with you ideas below.
// 添加用户
func AddUser(ctx context.Context, user entity.User) (userId int64, err error) {
	userId, err = User.Ctx(ctx).Data(user).OmitEmpty().InsertAndGetId()
	return
}

// 查询用户
func GetUser(ctx context.Context, user entity.User) (res entity.User, err error) {
	User.Ctx(ctx).Where("id=? or mobile=? or email=? or username=?", user.Id, user.Mobile, user.Email, user.Username).Scan(&res)
	return
}

// 用户是否存在
func IsExistUser(ctx context.Context, user entity.User) (exists bool) {
	count, _ := User.Ctx(ctx).Where("mobile=? or email=? or username=?",user.Mobile, user.Email, user.Username).Count()
	return count != 0
}
