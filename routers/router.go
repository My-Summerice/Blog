package routers

import (
	_"net/http"
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
        router.POST("article/add", v1.AddArticle)
        router.GET("article/:cid", v1.GetArticle)
        router.GET("article/list", v1.GetArticleList)
        router.GET("article/list/cate", v1.GetCateArticleList)
        router.PUT("article/:cid", v1.EditArticle)
        router.DELETE("article/:cid", v1.DelArticle)
		// 分类模块的路由接口
        router.POST("category/add", v1.AddCategory)
        router.GET("category/list", v1.GetCategoryList)
        router.PUT("category/:id", v1.EditCategory)
        router.DELETE("category/:id", v1.DelCategory)
	}

	r.Run(utils.HttpPort)
}