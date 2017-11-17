package usercache

import (
	"errors"
)

var ERROR_USER_NOT_FOUND = errors.New("用户未找到！")

type UserDetail struct {
	Uid       uint32  `json:"uid"`
	State     bool    `json:"state"`
	Salt      uint32  `json:"salt"`
	Password  string  `json:"password"`
	Avatar    string  `json:"avatar"`
	UserAgent string  `json:"user_agent"`
	Cache_Update_Time int64 `json:"cache_update_time"` //上次缓存更新时间
}
func GetUserDetail(uid uint32) (detail *UserDetail, e error) {
	detail=new(UserDetail)
	if uid==331805370{
		detail.Uid=uid
		detail.Salt=148360
		detail.Password="3dfaf9b0fc31457f7d068946181201f3"
		detail.Avatar="assets/img/avatar/W.jpg"
		detail.UserAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome"
	}else {
		e=ERROR_USER_NOT_FOUND
	}
	return
}