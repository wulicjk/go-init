package router

import "github.com/gin-gonic/gin"

func InitApiRouter(engine *gin.Engine) {
	//路由分组
	user := engine.Group("/user")
	{
		user.GET("/messages", controller.GetMessagesV2)
		user.POST("/message", middleware.Ipblack, controller.SendMessageV2)
	}
	engine.GET("/user", controller.GetCaptcha)
	engine.POST("/check", controller.LoginCheckPass)
}
