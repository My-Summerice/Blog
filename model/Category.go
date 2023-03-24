package model

import (
    "blog/utils/errmsg"
    "github.com/jinzhu/gorm"
)

// 文章类别模型
type Category struct {
    gorm.Model
	Name		string 		`gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) int {
    var category Category
    db.Select("id").Where("name = ?", name).First(&category)
    if category.ID > 0 {
        return errmsg.ERROR_CATENAME_USED	// 2001
    }
    return errmsg.SUCCESS		// 200
}

// 添加分类
func CreateCate(category *Category) int {
    err := db.Create(&category).Error
    if err != nil {
        return errmsg.ERROR		// 500
    }
    return errmsg.SUCCESS		// 200
}

// 查询分类列表
func SelectCategoryList(pageSize int, pageNum int) []Category {
    var categoryList []Category
    err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&categoryList).Error	// 分页查询
    if err != nil {	// 因为返回的值是List类型,可以为空，所以不需要考虑查询出空列表的问题(&& err != gorm.ErrRecordNotFound)
        return nil
    }
    return categoryList
}

// 更新分类信息
func UpdateCategory(id int, c_info *Category) int {
    var maps = make(map[string]interface{})
    maps["name"] = c_info.Name
    err := db.Model(&Category{}).Where("id = ?", id).Updates(maps).Error  //此处若直接使用struct更新将不会更新其中为零值的项
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
    var category Category
    err := db.Where("id = ? ", id).Delete(&category).Error     // 软删除：数据库中保存的数据并不会被真正的删除，只是无法通过正常的方法访问到
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}