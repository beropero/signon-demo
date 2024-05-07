package email_test

import (
	"signon_demo/utility/config"
	"signon_demo/utility/email"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestEmail(t*testing.T){
	ctx := gctx.New()
	code, err  := email.EmailSendCode(ctx,config.Email().Username)
	
	t.Logf("code:%s",code)
	if err != nil {
		t.Fatal(err)
	}
	err = email.SendEmail(ctx,config.Email().Username,"hello","你好")
	if err != nil {
		t.Fatal(err)
	}
}