package utils

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Sms struct {
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SignName        string `json:"sign_name"`
	TemplateCode    string `json:"template_code"`
}

func SendOss(mobile, code string, s *Sms) error {
	ossConfig := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(s.AccessKeyID, s.AccessKeySecret)

	ossClient, err := dysmsapi.NewClientWithOptions("cn-hangzhou", ossConfig, credential)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = s.SignName
	request.TemplateCode = s.TemplateCode
	request.TemplateParam = "{\"code\":" + code + "}"

	response, err := ossClient.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}
