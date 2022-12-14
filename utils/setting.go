package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode string      
	HttpPort string      

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}

	LoadServer(file)
	LoadData(file)
}

// 读取server相关的参数
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")			// 去取section为“server”、key为“AppMode”的value,若没取到则获得默认值“debug”
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

// 读取data相关的参数
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString(":3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("summerice")
	DbName = file.Section("database").Key("DbName").MustString("blog")
}