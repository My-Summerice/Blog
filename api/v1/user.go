package v1

import (
	"net/http"
	"blog/model"
	"blog/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// // 查询用户名是否存在	// 已在User.go中实现
// func UserExist(c *gin.Context) {
// 
// }

// 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})	// 400
		return
	}
	// 添加用户前先判断用户名是否存在
	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	}
	// 给前端返回add操作的结果
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": user,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUser(c *gin.Context) {

}

// 查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))	// Query返回的是string类型，需要将其强转为int类型
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))		// 当前端不传数据时默认为“1”

	userList := model.SelectUserList(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": userList,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DelUser(c *gin.Context) {

}