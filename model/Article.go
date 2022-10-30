package model

import (
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