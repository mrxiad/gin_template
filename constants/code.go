package constants

type ResponseCodeType int

// 状态码
const (
	ResponseCodeSuccess        ResponseCodeType = 1   //成功相应
	ResponseCodeBodyIsNull     ResponseCodeType = 2   //body为空
	ResponseCodeJsonParsError  ResponseCodeType = 3   //json解析错误
	ResponseCodeLessParamError ResponseCodeType = 4   //参数不足
	ResponseCodeServerError    ResponseCodeType = 5   //服务器错误
	ResponseCodeParamError     ResponseCodeType = 6   //参数错误
	ResponseFail               ResponseCodeType = 7   //失败
	ResponseHaveDataNoAllowDel ResponseCodeType = 8   //有数据不允许删除
	ResponseCodeAuthError      ResponseCodeType = 403 //权限校验失败
)

// 根据状态码返回对应的描述
func (p ResponseCodeType) String() string {
	switch p {
	case ResponseCodeSuccess:
		return "success"
	case ResponseCodeBodyIsNull:
		return "body is null"
	case ResponseCodeJsonParsError:
		return "json parse error"
	case ResponseCodeLessParamError:
		return "less param"
	case ResponseCodeServerError:
		return "server error"
	case ResponseCodeParamError:
		return "param error"
	case ResponseFail:
		return "fail"
	case ResponseHaveDataNoAllowDel:
		return "have data, not allow delete"
	case ResponseCodeAuthError:
		return "check auth fail"
	default:
		return "unknown code desc"
	}
}
