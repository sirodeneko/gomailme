package api

import (
	"gomailme/serializer"
	"gomailme/service"

	"github.com/gin-gonic/gin"
)

// Set 设置默认发件人
func Set(c *gin.Context) {
	var service service.SetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Set()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// Send 发送邮件接口
func Send(c *gin.Context) {
	var service service.SendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Send()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetSend 发送邮件
func GetSend(c *gin.Context){
	var service service.SendService

	service.Body=c.Param("body")
	service.To=c.Param("to")


	if service.Body!=""&&service.To!="" {
		res := service.Send()
		c.JSON(200, res)
	} else {

		c.JSON(200, &serializer.Response{
			Code: 40000,
			Msg:  "参数存在空值",
		})
	}
}

