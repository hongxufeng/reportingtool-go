package usercache

import "errors"

var ERROR_USER_NOT_FOUND = errors.New("用户未找到！")

type UserDetail struct {
	Uid       uint32  `json:"uid"`
	Salt      uint32  `json:"salt"`
	Password  string  `json:"password"`
	Avatar    string  `json:"avatar"`
}
func GetUserDetail(uid uint32) (detail *UserDetail, e error) {
	if uid==331805370{
		detail=&UserDetail{uid,148360,"3dfaf9b0fc31457f7d068946181201f3","https://avatars1.githubusercontent.com/u/20455425?s=40&v=4"}
	}else {
		e=ERROR_USER_NOT_FOUND
	}
	return
}