package main

import (
	_ "signon_demo/internal/packed"

	_ "signon_demo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"signon_demo/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
