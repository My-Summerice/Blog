package v1

import (
	"net/http"
	"blog/model"
	"blog/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)


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
		/* 在user.go中已经将加密操作写入了gorm自带的钩子函数beforeSave()，会在将数据存入数据库前自动执行钩子函数，所以此处即可省略
            // 在将数据写入数据库前对用户密码进行加密
            user.PassWord = model.ScryptPwd(user.PassWord)
        */
		// 将用户数据写入数据库
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

// 查询用户列表(分页查询)
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))   // Query返回的是string类型，需要将其强转为int类型
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))	  // 当前端不传数据时默认为“1”

	userList := model.SelectUserList(pageSize, pageNum)
	code := errmsg.SUCCESS
    // 给前端返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": userList,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户(只修改普通信息，密码单另做接口修改)
func EditUser(c *gin.Context) {
    var user model.User
    // 获取用户id，将string转为int
    id, _ := strconv.Atoi(c.Param("id"))
    if err := c.ShouldBindJSON(&user); err != nil {
        // 返回错误信息
        // gin.H封装了生成json数据的工具
        c.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})	// 400
        return
    }
    // 修改用户名前也要先判断用户名是否存在（也可以使用钩子函数将此判断放入model.BeforeSave()中）
    code := model.CheckUser(user.Username)
    if code == errmsg.SUCCESS {
        model.UpdateUser(id, &user)
    }
    if code == errmsg.ERROR_USERNAME_USED {
        // 该中间件执行完毕之后后续的中间件将不再执行
        c.Abort()
    }
    // 给前端返回操作结果
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
    })
}

// 删除用户
func DelUser(c *gin.Context) {
    // 将string转为int
    id, _ := strconv.Atoi(c.Param("id"))

    code := model.DeleteUser(id)
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
    })
}