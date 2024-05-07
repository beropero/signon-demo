package jwt

import (
	"context"
	"signon_demo/utility/config"
	"time"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	jwt "github.com/golang-jwt/jwt/v5"
)

type Data jwt.MapClaims

func JwtEncode(ctx context.Context, userId int) (tokenString string, err error) {
	data := jwt.MapClaims{
		"userid": userId,
		"exp":    jwt.NewNumericDate(time.Now().Add(time.Duration(config.Jwt().Duration) * time.Hour)),
		"jti":    guid.S(),
	}
	key := config.Jwt().Key
	Secret := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err = token.SignedString(Secret)
	return
}

func JwtDecode(ctx context.Context, token string) (userid int, err error) {
	key := config.Jwt().Key
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return 0, err
	}
	// 校验 Claims 对象是否有效
	if !res.Valid {
		return 0, gerror.New("invalid token")
	}
	claims := res.Claims.(jwt.MapClaims)
	userid = gconv.Int(claims["userid"])
	// jwt是否有效
	//search, _ := g.Redis().Do(ctx, "GET", fmt.Sprintf("jwt.token.%d", userid))
	//if search.String() != token {
	//	return 0, gerror.New("invalid token")
	//}
	return
}
