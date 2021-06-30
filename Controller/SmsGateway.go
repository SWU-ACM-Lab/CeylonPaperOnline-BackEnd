package Controller

import (
	"CeylonPaperOnline-BackEnd/Middleware"
	"encoding/json"
	"fmt"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"io/ioutil"
)

type SmsInfo struct {
	PhoneNumbers string `json:"phone_numbers"`
	SignName string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	TemplateParam string `json:"template_param"`
}

func (config* SmsInfo) LoadConfig (path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		jsonStr := string(data)
		err = json.Unmarshal([]byte(jsonStr), config)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (config* SmsInfo) SetSms(phonenumbers, templateparam string) {
	config.PhoneNumbers = phonenumbers
	config.TemplateParam = templateparam
}

func SendSms(info SmsInfo, config Middleware.SmsApiConfig) {
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AccessKeyId, config.AccessSecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = info.PhoneNumbers
	request.SignName = info.SignName
	request.TemplateCode = info.TemplateCode
	request.TemplateParam = info.TemplateParam

	response, err := client.SendSms(request)
	if err != nil {
		Middleware.Console.Log(err, "Send SMS")
	}
	fmt.Printf("response is %#v\n", response)
}