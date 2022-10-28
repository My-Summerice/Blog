package routers

import (
	"net/http"
	"blog/utils"
	"blog/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/list", v1.GetUserList)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DelUser)
		// 文章模块的路由接口

		// 分类模块的路由接口

	}

	r.Run(utils.HttpPort)
}