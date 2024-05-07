package ucontext

import "github.com/gogf/gf/v2/frame/g"

const (
	CtxKey = "uctx"
)

type Ucontext struct {
	User *User
	Data g.Map
}

type User struct {
	UserId   int
	UserName string
}
