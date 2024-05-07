package config_test

import (
	"signon_demo/utility/config"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestConfig(t *testing.T) {
	ctx := gctx.New()
	cfg := config.Cfg{}
	g.Cfg().MustGet(ctx, "utility").Scan(&cfg)
	g.Dump(cfg)
}
