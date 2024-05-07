package jwtstorage

import (
	"context"
	"fmt"
	"signon_demo/internal/dao"
	"signon_demo/internal/model/entity"
	"signon_demo/internal/service"
	"signon_demo/utility/config"
	"signon_demo/utility/jwt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sJwtStorage struct{}

func init() {
	service.RegisterJwtStorage(&sJwtStorage{})
}

// 生成并保存jwt令牌
func (s *sJwtStorage) GetJwtAndSave(ctx context.Context, userid int) (token string, err error) {
	// 生成jwt令牌
	token, err = jwt.JwtEncode(ctx, userid)
	if err != nil {
		return
	}
	// Redis存储
	_, err = g.Redis().Do(ctx, "SETEX", fmt.Sprintf("jwt.token.%d", userid), 60*60*24, token)
	if err != nil {
		return "", err
	}
	// 数据库存储
	err = dao.AddJwt(ctx, entity.JwtStorage{
		UserId:   userid,
		Token:    token,
		ExpireAt: gconv.GTime(time.Now().Add(time.Duration(config.Jwt().Duration) * time.Hour)),
	})
	return
}
// 拉取Jwt令牌到Redis
func (s *sJwtStorage) PullTokenFromDB() {
	ctx := gctx.New()
	tokens := []entity.JwtStorage{}
	model := dao.JwtStorage.Ctx(ctx)
	// 数据库查询
	if err := model.Where("expire_at>? and logout=0", time.Now()).Scan(&tokens); err != nil {
		panic(err)
	}
	for _, token := range tokens {
		k := fmt.Sprintf("jwt.token.%d", token.UserId)
		v := token.Token
		exp := token.ExpireAt.Sub(gtime.Now()).Seconds()
		// 存入Redis
		if _, err := g.Redis().Do(ctx, "SETEX", k, int(exp) , v); err != nil {
			panic(err)
		}
	}
}
