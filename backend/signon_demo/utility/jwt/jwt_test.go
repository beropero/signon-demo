package jwt_test

import (
	"signon_demo/utility/jwt"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestJwt(t *testing.T) {
	ctx := gctx.New()
    token, _ := jwt.JwtEncode(ctx, 1111111)
	g.Dump(token)
	res, _ := jwt.JwtDecode(ctx, token)
	g.Dump(res)
}
