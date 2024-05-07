package email

import (
	"context"
	"fmt"
	"signon_demo/utility/config"
	"signon_demo/utility/vcode"

	"gopkg.in/gomail.v2"
)

// 发送邮箱验证码
func EmailSendCode(ctx context.Context, email string) (code string, err error) {
	// 生成6位随机验证码
	code = vcode.GenerateRandomCode(6)
	m := gomail.NewMessage()
	m.SetHeader("From", config.Email().Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码")
	msg := fmt.Sprintf("你的验证码为: %s", code)
	m.SetBody("text/html", msg)
	d := gomail.NewDialer(config.Email().Host, config.Email().Port, config.Email().Username, config.Email().Password)
	err = d.DialAndSend(m)
	return
}

// 发送通用邮件
func SendEmail(ctx context.Context, email string, subject string, msg string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Email().Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)
	d := gomail.NewDialer(config.Email().Host, config.Email().Port, config.Email().Username, config.Email().Password)
	err := d.DialAndSend(m)
	return err
}
