package v1

import (
    "blog/model"
    "blog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
    var article model.Article
    if err := c.ShouldBindJSON(&article); err != nil {
        // 返回错误信息
        // gin.H封装了生成json数据的工具
        c.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})	// 400
        return
    }
    // 向数据库中添加文章
    code := model.CreateArticle(&article)
    // 给前端返回add操作的结果
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "data": article,
        "message": errmsg.GetErrMsg(code),
        })
}

// todo 查询单个文章信息
func GetArticle(c *gin.Context) {
    
}

// 查询文章列表(分页查询)
func GetArticleList(c *gin.Context) {
    pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))   // Query返回的是string类型，需要将其强转为int类型
    pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))	  // 当前端不传数据时默认为“1”

    articleList := model.SelectArticleList(pageSize, pageNum)
    code := errmsg.SUCCESS
    // 给前端返回操作结果
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "data": articleList,
        "message": errmsg.GetErrMsg(code),
        })
}

// todo：查询分类下的所有文章
func GetCateArticleList(c *gin.Context) {

}

// 编辑文章
func EditArticle(c *gin.Context) {
    var article model.Article
    // 获取cid，将string转为int
    cid, _ := strconv.Atoi(c.Param("cid"))
    if err := c.ShouldBindJSON(&article); err != nil {
        // 返回错误信息
        // gin.H封装了生成json数据的工具
        c.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})	// 400
        return
    }
    code := model.UpdateArticle(cid, &article)
    // 给前端返回操作结果
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
        })
}

// 删除文章
func DelArticle(c *gin.Context) {
    // 将string转为int
    cid, _ := strconv.Atoi(c.Param("cid"))

    code := model.DeleteArticle(cid)
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "message": errmsg.GetErrMsg(code),
        })
}