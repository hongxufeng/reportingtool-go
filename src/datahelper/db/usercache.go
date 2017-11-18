package db

import (
	"model"
	"utils/function"
	"github.com/go-redis/redis"
	"encoding/json"
	"fmt"
	"errors"
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
		//查数据库，设置redis
		return SetUserInfoCache(uid)
	} else if e != nil {
		return nil,e
	} else {
		if len(info)>0{
			var rm map[string]interface{}
			e := json.Unmarshal([]byte(info), &rm)
			if e != nil {
				//查数据库，设置redis
				return SetUserInfoCache(uid)
			}
			if e := function.MapToStruct(rm, &userinfo); e != nil {
				//查数据库，设置redis
				return SetUserInfoCache(uid)
			}
		}else {
			//查数据库，设置redis
			return SetUserInfoCache(uid)
		}
	}
	return
}

func SetUserInfoCache(uid uint32) (userinfo *UserInfo, err error) {
	userinfo,err=GetUserInfobyDB(uid)
	if err!=nil {
		return
	}
	bts, err:= json.Marshal(userinfo)
	if err != nil {
		fmt.Println("json.Marshal error " + err.Error())
		return
	}
	err=RedisCache.Set(function.MakeKey(CACHE_USER_INFO, uid),string(bts),model.User_Info_Persistence_Duration).Err()
	return
}

func GetUserInfobyDB(uid uint32) (userinfo *UserInfo, err error) {
	query:="SELECT uid,username,nickname,password,salt,state,avatar,user_agent FROM w_user_list WHERE uid=?"
	result, e := MysqlMain.Query(query, uid)
	if e != nil {
		return
	}
	defer result.Close()
	if result.Next() {
		err= result.Scan(&userinfo.Uid,&userinfo.UserName,&userinfo.NickName,&userinfo.Password,&userinfo.Salt,&userinfo.Avatar,&userinfo.UserAgent)
	}else {
		err=errors.New("您输入的用户未找到呢！")
	}
	return
}