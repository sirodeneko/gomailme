package delayJob

import (
	"errors"
	"github.com/robfig/cron"
	"gomailme/model"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
	"time"
)

const (
	SendMailType string = "sendMail"
)


// 加入任务类型
func (dj *DelayJob) SetMethods() {
	dj.AddJobFun(SendMailType, SendMailFun)
}

// 邮件发送服务实现
func SendMailFun(id string) error {
	var task model.Task
	err:=model.DB.Where("id = ?",id).First(&task).Error
	if err!=nil{
		return errors.New(time.Now().String()+" 任务查询失败"+err.Error())
	}
	// 执行发送邮件服务
	err=SendMail(task.User,task.Pass,task.To,task.Msg)

	// 更新任务描述
	if task.Count!=-1{
		task.Count--
		if task.Count!=0{
			// 更新描述
			model.DB.Model(&task).Where("id = ?", id).Update("count", task.Count)
		}else {
			model.DB.Delete(&task)
			return nil
		}
	}

	Schedule,_:=cron.Parse(task.Cron)
	_=DJ.AddJob(strconv.FormatUint(uint64(task.ID), 10),Schedule.Next(time.Now()).Unix())
	if err != nil {
		return err
	}
	return nil
}
// SendMail 发送邮件
func SendMail(user string, pass string, to string, msg string) error {
	//fmt.Println(time.Now(),"do this send mail job")
	//return nil
	// user string,pass string,to string,msg string
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


	mailConn := map[string]string{
		"user": user,
		"pass": pass,
		"host": "smtp.qq.com",
	}
	m := gomail.NewMessage()
	m.SetAddressHeader("From", mailConn["user"], "机器人")

	s := strings.Split(to, ",") //将多个收件人转化为字符串
	var toSomeboy []string
	for _, item := range s {
		toSomeboy = append(toSomeboy, m.FormatAddress(item, "master"))
	}
	m.SetHeader("To", toSomeboy...)
	m.SetHeader("Subject", "由机器人自动发送")
	m.SetBody("text/html", msg)
	d := gomail.NewDialer(mailConn["host"], 465, mailConn["user"], mailConn["pass"])
	return d.DialAndSend(m)
}