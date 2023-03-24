package model

import (
	"blog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
    _"net"

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

// 钩子函数将在数据存入数据库前自动被调用
func (u *User) BeforeSave() {
    u.PassWord = ScryptPwd(u.PassWord)
}
// 使用scrypt算法对用户密码进行加密
func ScryptPwd(password string) string {
	keyLen := 10
	salt := []byte{10, 3, 52, 137, 90, 2, 255, 71}
	bytePwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(bytePwd)
}

// 添加用户
func CreateUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR		// 500
	}
	return errmsg.SUCCESS		// 200
}

// 查询用户列表
func SelectUserList(pageSize int, pageNum int) []User {
	var userList []User
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&userList).Error	// 分页查询
	if err != nil {	// 因为返回的值是List类型,可以为空，所以不需要考虑查询出空列表的问题(&& err != gorm.ErrRecordNotFound)
		return nil
	}
	return userList
}

// 更新用户信息
func UpdateUser(id int, u_info *User) int {
    var maps = make(map[string]interface{})
    maps["username"] = u_info.Username
    maps["role"] = u_info.Role
    err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error  //此处若直接使用struct更新将不会更新其中为零值的项，而role可能为0,所以使用map进行数据更新
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
    var user User
	err := db.Where("id = ? ", id).Delete(&user).Error     // 软删除：数据库中保存的数据并不会被真正的删除，只是无法通过正常的方法访问到
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}