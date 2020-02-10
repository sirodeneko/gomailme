package service

import (
	"gomailme/model"
	"gomailme/serializer"
	"gopkg.in/gomail.v2"
	"strings"
)

// SendService 设置默认用户
type SendService struct {
	Body        string `form:"body" json:"body"`
	To			string`form:"to" json:"to"`
}

// Send 发送邮件的服务
func (service *SendService)Send() *serializer.Response {
	master:=model.Master{}
	err:=model.DB.First(&master).Error
	if err!=nil{
		return &serializer.Response{
			Code: 40000,
			Msg:  "查找发件人失败",
			Error:err.Error(),
		}
	}
	//QQ 邮箱
	//POP3 服务器地址：qq.com（端口：995）
	//SMTP 服务器地址：smtp.qq.com（端口：465/587）
	//
	//163 邮箱：
	//POP3 服务器地址：pop.163.com（端口：110）
	//SMTP 服务器地址：smtp.163.com（端口：25）
	//
	//126 邮箱：
	//POP3 服务器地址：pop.126.com（端口：110）
	//SMTP 服务器地址：smtp.126.com（端口：25）
	mailConn :=map[string]string{
		"user":master.User,
		"pass": master.Pass,
		"host":"smtp.qq.com",
	}


	m := gomail.NewMessage()
	m.SetAddressHeader("From",mailConn["user"],"机器人")

	s:=bodyTos(service.To)//将多个收件人转化为字符串
	var toSomeboy []string
	for _,item:=range s{
		toSomeboy=append(toSomeboy,m.FormatAddress(item,"master") )
	}
	m.SetHeader("To", toSomeboy...)
	m.SetHeader("Subject","由机器自动发送")
	m.SetBody("text/html",service.Body)
	d:=gomail.NewDialer(mailConn["host"],465,mailConn["user"],mailConn["pass"])
	if err :=d.DialAndSend(m);err==nil{
		return &serializer.Response{
			Code: 40001,
			Msg:  "发送成功",
		}
	}else{
		return &serializer.Response{
			Code: 40000,
			Msg:  "发送失败",
			Error:err.Error(),
		}
	}
}

func bodyTos(body string) []string{
	s:=strings.Split(body,",")
	return s
}
