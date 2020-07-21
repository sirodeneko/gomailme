package service

import (
	"github.com/robfig/cron"
	"gomailme/delayJob"
	"gomailme/model"
	"gomailme/serializer"
	"strconv"
	"time"
)

// SetTimeMsgService
type SetTimeMsgService struct {
	Name  string `form:"name" json:"name" `
	Type  string `form:"type" json:"type" `
	User  string `form:"user" json:"user" `
	Pass  string `form:"pass" json:"pass" `
	Count int    `form:"count" json:"count" binding:"required"` //表示任务的执行次数，-1表示执行无限次，0表示不用执行了
	Cron  string `form:"cron" json:"cron" binding:"required"`   //表示任务的执行规律，满足cron规则
	To    string `form:"to" json:"to" binding:"required"`
	Msg   string `form:"msg" json:"msg" binding:"required"`
	Desc  string `form:"desc" json:"desc" `
}

// Register 用户注册
func (service *SetTimeMsgService) SetTimeMsg() serializer.Response {
	master := model.Master{}
	err := model.DB.First(&master).Error
	if err != nil {
		return serializer.Response{
			Code:  40000,
			Msg:   "查找发件人失败",
			Error: err.Error(),
		}
	}
	if _,err:=cron.Parse(service.Cron);err!=nil{
		return serializer.ParamErr("cron错误", err)
	}
	if service.Name == "" {
		service.Name = "定时任务"
	}
	if service.Type == "" {
		service.Type = model.SendMail
	}

	if service.User == "" {
		service.User = master.User
		service.Pass = master.Pass
	}

	task := model.Task{
		Name:  service.Name,
		Type:  service.Type,
		User:  service.User,
		Pass:  service.Pass,
		Count: service.Count,
		Cron:  service.Cron,
		To:    service.To,
		Msg:   service.Msg,
		Desc:  service.Desc,
	}

	// 存入数据库
	if err := model.DB.Create(&task).Error; err != nil {
		return serializer.ParamErr("数据库保存任务失败", err)
	}

	// 加入定时任务
	Schedule,_:=cron.Parse(task.Cron)
	if err := delayJob.DJ.AddJob(strconv.FormatUint(uint64(task.ID), 10),Schedule.Next(time.Now()).Unix()); err != nil {
		return serializer.DBErr("rides保存失败",err)
	}
	return serializer.Response{
		Code:  200,
		Data:  nil,
		Msg:   "保存成功",
		Error: "",
	}
}
