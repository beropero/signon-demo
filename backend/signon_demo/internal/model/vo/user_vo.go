package vo

// User is the golang structure for table user.
type UserVO struct {
	Username string `json:"username" ` // 用户名
	Nickname string `json:"nickname" ` // 昵称
	Mobile   string `json:"mobile"   ` // 手机号
	Email    string `json:"email"    ` // 邮箱
}

type TpAccountVO struct {
	Name        string        `json:"name"        ` // 平台名
	AccountInfo AccountInfoVO `json:"accountInfo" ` // 第三方账号信息
}

type AccountInfoVO struct {
	NickName string `json:"nickname,omitempty"`
	OpenId   string `json:"openid,omitempty" `
	UnionId  string `json:"unionid,omitempty"`
}
