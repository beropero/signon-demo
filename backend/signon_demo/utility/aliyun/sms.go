package aliyun

import (
	"encoding/json"
	"signon_demo/utility/config"
	"signon_demo/utility/vcode"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendSMS(mobile string) (code string, err error) {
	// 生成6位随机Code
	code = vcode.GenerateRandomCode(6)
	//参数一：连接的节点地址
	//参数二：AccessKey ID
	//参数三：AccessKey Secret
	client, err := dysmsapi.NewClientWithAccessKey(config.Aliyun().RegionId, config.Aliyun().AccessKeyId, config.Aliyun().AccessKeySecret)
	if err != nil {
		return "", err
	}
	
	request := dysmsapi.CreateSendSmsRequest()          //创建请求
	request.Scheme = config.Aliyun().Scheme             //请求协议
	request.PhoneNumbers = mobile                       //接收短信的手机号码
	request.SignName = config.Aliyun().SignName         //短信签名名称
	request.TemplateCode = config.Aliyun().TemplateCode //短信模板ID
	par, err := json.Marshal(map[string]interface{}{    //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": code,
	})
	if err != nil {
		return "", err
	}
	request.TemplateParam = string(par) //将短信模板参数传入短信模板
	_, err = client.SendSms(request)    //调用阿里云API发送信息
	return
}
