package errorx

type Status int

const (
	StatusErrUserNotFound Status = 1001 + iota
	StatusErrUserExist
	StatusErrUserPwd
	StatusErrUserNoAuth
	StatusErrParam
	StatusErrAdminAuth
	StatusErrNotKnown
	StatusErrSystemBusy
)

func (s Status) Error() string {
	switch s {
	case StatusErrUserNotFound:
		return "该账号还未注册"
	case StatusErrUserExist:
		return "该账号已被注册"
	case StatusErrUserPwd:
		return "用户名或密码错误"
	case StatusErrUserNoAuth:
		return "请登录后查看"
	case StatusErrParam:
		return "系统内部参数转换类型不匹配"
	case StatusErrAdminAuth:
		return "没有管理员权限"
	case StatusErrNotKnown:
		return "发生了未知错误"
	case StatusErrSystemBusy:
		return "系统繁忙, 请稍后再试!"
	default:
		return "系统繁忙, 请稍后再试!"
	}
}
