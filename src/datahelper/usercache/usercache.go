package usercache

import "errors"

var ERROR_USER_NOT_FOUND = errors.New("用户未找到！")

type UserDetail struct {
	Uid       uint32  `json:"uid"`
	Salt      uint32  `json:"salt"`
	Password  string  `json:"password"`
	Avatar    string  `json:"avatar"`
	UserAgent string  `json:"user_agent"`
}
func GetUserDetail(uid uint32) (detail *UserDetail, e error) {
	if uid==331805370{
		detail=&UserDetail{uid,148360,"3dfaf9b0fc31457f7d068946181201f3","assets/img/avatar/W.jpg","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"}
	}else {
		e=ERROR_USER_NOT_FOUND
	}
	return
}