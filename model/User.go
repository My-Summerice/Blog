package model

import (
	"blog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username	string 		`gorm:"type:varchar(20);not null" json:"username"`
	PassWord	string 		`gorm:"type:varchar(20);not null" json:"password"`
	Role 		int 		`gorm:"type:int" json:"role"`
}


// 查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED	// 1001
	}
	return errmsg.SUCCESS		// 200
}

// 添加用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR		// 500
	}
	return errmsg.SUCCESS		// 200
}

// 查询用户列表
func SelectUserList(pageSize int, pageNum int) []User {
	var userList []User
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&userList).Error	// 分页查询
    if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return userList
}

// 编辑用户
func UpdateUser() {

}

// 删除用户
func DeleteUser() {
	
}