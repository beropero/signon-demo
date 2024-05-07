package user

import (
	"context"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

// 用户注册
func (s *sUser) UserSignUp(ctx context.Context, user entity.User) (userId int64, err error) {
	exist := dao.IsExistUser(ctx, user)
	if exist {
		return -1, gerror.New("用户已存在")
	}
	userId, err = dao.AddUser(ctx, user)
	return
}

// 用户登陆校验
func (s *sUser) UserLoginVerify(ctx context.Context, user entity.User) (res bool, err error) {
	search, err := dao.GetUser(ctx, user)
	if err != nil {
		return
	}
	if search.Password != user.Password {
		return res, gerror.New("用户不存在或密码错误")
	} 
	res = true
	return
}

// 获取用户基础信息
func (s *sUser) UserInfo(ctx context.Context, user entity.User) (res entity.User, err error) {
	res, err = dao.GetUser(ctx, user)
	return
}
