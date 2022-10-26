package model

import (
	"github.com/jinzhu/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username	string 		`gorm:"type:varchar(20);not null" json:"username"`
	PassWord	string 		`gorm:"type:varchar(20);not null" json:"password"`
	Role 		int 		`gorm:"type:int" json:"role"`
}