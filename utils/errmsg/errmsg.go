package errmsg

import (

)

const (
	SUCCESS = 200
	ERROR = 500

	// code = 100... 用户模块的错误
	ERROR_USERNAME_USED = 1001		
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_NOT_EXIST = 1004
	ERROR_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_WRONG = 1006

	// code = 200... 文章模块的错误

	// code = 300... 分类模块的错误

)

// 构建一个错误字典
var codemsg = map[int]string{

}

// 构建一个输出错误信息的函数
func GetErrMsg(code int) string {
	return codemsg[code]
}
