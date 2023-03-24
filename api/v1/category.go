package v1

import (
    "blog/model"
    "blog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
    var category model.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        // 返回错误信息
        // gin.H封装了生成json数据的工具
        c.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})	// 400
        return
    }
    // 添加用户前先判断分类名是否存在
    code := model.CheckCategory(category.Name)
    if code == errmsg.SUCCESS {
        // 将分类数据写入数据库
        model.CreateCate(&category)
    }
    // 给前端返回add操作的结果
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "data": category,
        "message": errmsg.GetErrMsg(code),
        })
}

// 查询分类列表(分页查询)
func GetCategoryList(c *gin.Context) {
    pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))   // Query返回的是string类型，需要将其强转为int类型
    pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))	  // 当前端不传数据时默认为“1”

    categoryList := model.SelectCategoryList(pageSize, pageNum)
    code := errmsg.SUCCESS
    // 给前端返回操作结果
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "data": categoryList,
        "message": errmsg.GetErrMsg(code),
        })
}

// 编辑分类
func EditCategory(c *gin.Context) {
    var category model.Category
    // 获取分类id，将string转为int
    id, _ := strconv.Atoi(c.Param("id"))
    if err := c.ShouldBindJSON(&category); err != nil {
        // 返回错误信息
        // gin.H封装了生成json数据的工具
        c.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})	// 400
        return
    }
    // 修改分类名前也要先判断分类名是否存在（也可以使用钩子函数将此判断放入model.BeforeSave()中）
    code := model.CheckCategory(category.Name)
    if code == errmsg.SUCCESS {
        model.UpdateCategory(id, &category)
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

// 删除分类
func DelCategory(c *gin.Context) {
    // 将string转为int
    id, _ := strconv.Atoi(c.Param("id"))

    code := model.DeleteCategory(id)
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
        })
}