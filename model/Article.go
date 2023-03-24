package model

import (
    "blog/utils/errmsg"
    "github.com/jinzhu/gorm"
)

// 文章类型
type Article struct {
	Category	Category
	gorm.Model
	Title 		string		`gorm:"type:varchar(100);not null" json:"title"`
	Cid			int			`gorm:"type:int;not null" json:"cid"`
	Desc		string		`gorm:"type:varchar(200)" json:"desc"`
	Content		string		`gorm:"type:longtext" json:"content"`
	Img			string		`gorm:"type:varchar(100)" json:"img"`
}

// 添加文章
func CreateArticle(article *Article) int {
    err := db.Create(&article).Error
    if err != nil {
        return errmsg.ERROR		// 500
    }
    return errmsg.SUCCESS		// 200
}

// 查询单个文章
//func SelectArtcle(id int) Article {
//
//}

// 查询文章列表
func SelectArticleList(pageSize int, pageNum int) []Article {
    var articleList []Article
    err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articleList).Error	// 分页查询
    if err != nil {	// 因为返回的值是List类型,可以为空，所以不需要考虑查询出空列表的问题(&& err != gorm.ErrRecordNotFound)
        return nil
    }
    return articleList
}

// 查询分类下的文章列表
//func SelectCateArticle(cid int, pageSize int, pageNum int) []Article {
//
//}

// 编辑文章
func UpdateArticle(id int, a_info *Article) int {
    var maps = make(map[string]interface{})
    maps["title"] = a_info.Title
    maps["cid"] = a_info.Cid
    maps["desc"] = a_info.Desc
    maps["content"] = a_info.Content
    maps["img"] = a_info.Img
    err := db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error  //此处若直接使用struct更新将不会更新其中为零值的项
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
    err := db.Where("id = ? ", id).Delete(&Article{}).Error     // 软删除：数据库中保存的数据并不会被真正的删除，只是无法通过正常的方法访问到
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}