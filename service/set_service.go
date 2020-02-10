package service

import (
	"gomailme/model"
	"gomailme/serializer"
)

// SetService 设置默认用户
type SetService struct {
	User        string `form:"user" json:"user"`
	Pass  		string`form:"pass" json:"pass"`
}

// Set 设置默认地址的服务
func (service *SetService) Set() *serializer.Response {
	master:=model.Master{}
	count := 0
	model.DB.Model(&model.Master{}).Count(&count)
	if count > 0 {
		model.DB.First(&master)
		master.User=service.User
		master.Pass=service.Pass
		err:=model.DB.Save(&master).Error
		if err==nil{
			return &serializer.Response{
				Code: 40001,
				Msg:  "更改成功",
			}
		}else{
			return &serializer.Response{
				Code: 40000,
				Msg:  "失败",
				Error:err.Error(),
			}
		}

	}else{
		master=model.Master{
			User: service.User,
			Pass:service.Pass,
		}
		err:=model.DB.Create(&master).Error
		if model.DB.NewRecord(master)==false {
			return &serializer.Response{
				Code: 40001,
				Msg:  "创建成功",
			}
		}else{
			return &serializer.Response{
				Code: 40000,
				Msg:  "创建失败",
				Error:err.Error(),
			}
		}

	}

}
