package db

import (
	"model"
	"utils/function"
	"github.com/go-redis/redis"
	"encoding/json"
)

type UserInfo struct {
	Uid       uint32  `json:"uid"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	State     bool    `json:"state"`
	Salt      uint32  `json:"salt"`
	Password  string  `json:"password"`
	Avatar    string  `json:"avatar"`
	UserAgent string  `json:"user_agent"`
	CacheUpdateTime int64 `json:"cache_update_time"` //上次缓存更新时间
}
func GetUserInfo(uid uint32) (userinfo *UserInfo, err error) {
	userinfo=new(UserInfo)
	if uid==model.User_W_Uid{
		userinfo.Uid=model.User_W_Uid
		userinfo.UserName=model.User_W_UserName
		userinfo.NickName=model.User_W_NickName
		userinfo.Salt=model.User_W_Salt
		userinfo.Password=model.User_W_Password
		userinfo.Avatar=model.User_W_Avatar
		userinfo.UserAgent=model.User_W_UserAgent
	}
	info,e:=RedisCache.Get(function.MakeKey(CACHE_USER_INFO, uid)).Result()
	if e == redis.Nil {
		//查数据库
		//return nil,errors.New("您输入的用户并未找到呢！")
	} else if e != nil {
		return nil,e
	} else {
		if len(info)>0{
			var rm map[string]interface{}
			e := json.Unmarshal([]byte(info), &rm)
			if e != nil {
				//查数据库
			}
			if e := function.MapToStruct(rm, &userinfo); e != nil {
				//查数据库
			}
		}else {
			//查数据库
		}
	}
	return
}