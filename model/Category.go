package model

import (
	"github.com/jinzhu/gorm"
)

// 文章类别模型
type Category struct {
	gorm.Model
	Name		string 		`gorm:"type:varchar(20);not null" json:"name"`
}