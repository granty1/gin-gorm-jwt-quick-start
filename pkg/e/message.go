package e

var messageFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	LOGIN_INFO_ERROR:               "用户名或密码错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "登录信息鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "登录信息已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTH_NOT_FOUND:           "登录信息为空",
}

func Message(code int) string {
	val, ok := messageFlags[code]
	if ok {
		return val
	}
	return messageFlags[ERROR]
}
