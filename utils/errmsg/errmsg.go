package errmsg

import ()

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 100... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code = 200... 文章模块的错误

	// code = 300... 分类模块的错误

)

// 构建一个错误码map
var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户不存在！",
	ERROR_TOKEN_NOT_EXIST:  "token不存在",
	ERROR_TOKEN_RUNTIME:    "token已过期",
	ERROR_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
}

// 构建一个输出错误信息的函数
func GetErrMsg(code int) string {
	return codeMsg[code]
}
