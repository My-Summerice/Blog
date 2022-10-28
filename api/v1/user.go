package v1

import (
	"net/http"
	"blog/model"
	"blog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

// // 查询用户名是否存在
// func UserExist(c *gin.Context) {
// 
// }

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	// 400
		return
	}
	// 添加用户前先判断用户名是否存在
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code)
	})
}

// 查询单个用户
func GetUser(c *gin.Context) {

}

// 查询用户列表
func GetUserList(c *gin.Context) {

}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DelUser(c *gin.Context) {

}