package aliyun_test

import (
	"signon_demo/utility/aliyun"
	"testing"

)

func TestSms(t *testing.T){
	code, err:= aliyun.SendSMS("13542107807") 
	t.Logf("code: %s",code)
	if err != nil {
		t.Fatal(err)
	}
}
