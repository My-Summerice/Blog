package model

import (
	"fmt"
	"time"
	"blog/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err != nil {
		fmt.Printf("数据库连接失败，请检查参数：", err)
	}

	// 禁用默认表名的复数形式
	db.SingularTable(true)

	// 自动进行数据库模型迁移
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// 数据库连接时间应该小于框架的连接时间，否则容易报错，一般设置为10s
	db.DB().SetConnMaxLifetime(10 * time.Second)

	// 新版gorm(database/sql)中维护了一个连接池（上方已启用），连接池会自动对数据库的连接进行限制，因此不再需要手动关闭数据库
	//db.Close()
}