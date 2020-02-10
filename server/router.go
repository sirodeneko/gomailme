package server

import (
	"os"
	"gomailme/api"
	"gomailme/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v2:=r.Group("/mail")
	{
		// 设置默认发送人
		v2.POST("set",api.Set)

		// 发送消息
		v2.POST("send",api.Send)

		// 发送消息
		v2.GET("send/:to/:body",api.GetSend)

		//// 定时消息（待完成）
		//v2.POST("setTimeMsg",api.SetTimeMsg)
		//
		//// 更新定时消息
		//v2.POST("updateTimeMsg",api.UpdateTimeMsg)
		//
		//// 删除定时消息
		//v2.DELETE("deleteTimeMsg",api.DeleteTimeMsg)
	}

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
