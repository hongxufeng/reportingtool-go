package service

import "fmt"

const (
	ERR_NOERR     = 0    //没有错误
	ERR_UNKNOWN   = 1001 //未知错误
	ERR_INTERNAL  = 1002 //内部错误
	ERR_MYSQL     = 1003 //mysql错误
	ERR_REDIS     = 1004 //redis错误
	ERR_NOT_FOUND = 1005 //未找到

	ERR_INVALID_PARAM         = 2001 //请求参数错误
	ERR_INVALID_FORMAT        = 2002 //格式错误
	ERR_ENCRYPT_ERROR         = 2003 //加密错误
	ERR_INVALID_REQUEST       = 2004 //不合法的请求
	ERR_VCODE_ERROR           = 2005 //验证码错误
	ERR_VERIFY_FAIL           = 2006 //验证失败
	ERR_VCODE_TIMEOUT         = 2007 //验证码超时
	ERR_INVALID_USER          = 2008 //用户验证不通过
	RR_STATUS_DENIED		    =2009  //用户状态关闭
	ERR_PERMISSION_DENIED     = 2010 //权限不足
	ERR_LOGIN_FREQUENT        =2011   //登录频繁
	ERR_USER_NOT_FOUND        =2012   //用户不存在

	ERR_XML_ATTRIBUTE_LACK    =3001   //XML配置属性缺失
)
type Error struct {
	Code uint
	Show string //客户端显示的内容
}

func NewError(ecode uint, show ...string) (err Error) {
	if len(show) > 0 {
		err = Error{ecode, show[0]}
	} else {
		switch ecode {
		case ERR_INVALID_PARAM:
			err = Error{ecode, "参数错误"}
		case ERR_INVALID_REQUEST:
			err = Error{ecode, "不合法的请求"}
		case ERR_MYSQL, ERR_REDIS:
			err = Error{ecode, "数据库错误"}
		default:
			err = Error{ecode, "内部错误"}
		}
	}
	return
}
func (e Error) Error() (re string) {
	return fmt.Sprintf("ecode=%v, show=%v", e.Code, e.Show)
}

