package service

import (
	"gomailme/delayJob"
	"gomailme/model"
	"gomailme/serializer"
)

// SendService 设置默认用户
type SendService struct {
	Body string `form:"body" json:"body"`
	To   string `form:"to" json:"to"`
}

// Send 发送邮件的服务
func (service *SendService) Send() *serializer.Response {
	master := model.Master{}
	err := model.DB.First(&master).Error
	if err != nil {
		return &serializer.Response{
			Code:  40000,
			Msg:   "查找发件人失败",
			Error: err.Error(),
		}
	}

	err = delayJob.SendMail(master.User, master.Pass, service.To, service.Body)

	if err == nil {
		return &serializer.Response{
			Code: 200,
			Msg:  "发送成功",
		}
	} else {
		return &serializer.Response{
			Code:  40000,
			Msg:   "发送失败",
			Error: err.Error(),
		}
	}
}


