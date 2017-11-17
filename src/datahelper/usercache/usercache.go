package usercache

import (
	"errors"
	"model"
)

type UserDetail struct {
	Uid       uint32  `json:"uid"`
	State     bool    `json:"state"`
	Salt      uint32  `json:"salt"`
	Password  string  `json:"password"`
	Avatar    string  `json:"avatar"`
	UserAgent string  `json:"user_agent"`
	Cache_Update_Time int64 `json:"cache_update_time"` //上次缓存更新时间
}
func GetUserDetail(uid uint32) (detail *UserDetail, err error) {
	detail=new(UserDetail)
	if uid==model.User_W_Uid{
		detail.Uid=model.User_W_Uid
		detail.Salt=model.User_W_Salt
		detail.Password=model.User_W_Password
		detail.Avatar=model.User_W_Avatar
		detail.UserAgent=model.User_W_UserAgent
	}else {
		err=errors.New("您输入的用户并未找到呢！")
	}
	return
}