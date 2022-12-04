// Package errorx
/*
@Coding : utf-8
@time : 2022/12/2 22:28
@Author : yizhigopher
@Software : GoLand
*/
package errorx

var codeMap = make(map[string]string)

func init() {

	codeMap[NotLoggedIn] = "未登录"
	codeMap[UnauthorizedUserId] = "非法的用户Id"
	codeMap[Unauthorized] = "未授权"
	codeMap[OperationFailure] = "操作失败"
	codeMap[CommonError] = "服务器一般错误"

	// 通用错误
	codeMap[OK] = "Success"
	codeMap[NotData] = "没有数据"
	codeMap[DataExist] = "数据已存在"
	codeMap[DataError] = "数据错误"

	// 网络级错误
	codeMap[ParameterIllegal] = "参数不合法"
	codeMap[RequestOverDue] = "请求已过期"
	codeMap[LoginError] = "登录已过期"
	codeMap[AccessDenied] = "拒绝访问"
	codeMap[RoutingNotExist] = "路由不存在"
	codeMap[PasswordError] = "密码错误"
	codeMap[RequestError] = "非法访问"
	codeMap[IPError] = "IP受限"

	// 系统级错误
	codeMap[InternalError] = "系统错误"
	codeMap[DBError] = "数据库错误"
	codeMap[ThirdError] = "第三方系统错误"
	codeMap[IOError] = "IO错误"
	codeMap[UnKnownError] = "未知错误"
}

func Code2Msg(errcode string) string {
	if msg, ok := codeMap[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode string) bool {
	if _, ok := codeMap[errcode]; ok {
		return true
	} else {
		return false
	}
}
