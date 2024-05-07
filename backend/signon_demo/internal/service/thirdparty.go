// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IThirdParty interface{}
)

var (
	localThirdParty IThirdParty
)

func ThirdParty() IThirdParty {
	if localThirdParty == nil {
		panic("implement not found for interface IThirdParty, forgot register?")
	}
	return localThirdParty
}

func RegisterThirdParty(i IThirdParty) {
	localThirdParty = i
}
