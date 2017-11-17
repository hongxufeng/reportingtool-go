package db

import (
	"utils/function"
	"github.com/go-redis/redis"
	"model"
)

func UserLoginErrCnt(uid uint32) (cnt int64, err error) {
	cnt, err = RedisCache.Incr(function.MakeKey(CACHE_USER_LOGIN_ERROR, uid)).Result()
	if err != redis.Nil&&err != nil {
		return
	}  else {
		//cnt++
		err=RedisCache.Set(function.MakeKey(CACHE_USER_LOGIN_ERROR, uid),cnt,model.User_Forfid_Expiration_Duration).Err()
	}
	//fmt.Println(cnt)
	return
}
func SetUserForbid(uid uint32) (err error) {
	err=RedisCache.Set(function.MakeKey(CACHE_USER_LOGIN_FORBID, uid),true,model.User_Forfid_Expiration_Duration).Err()
	return
}
func CheckUserForbid(uid uint32) (forbid bool, err error) {
	//验证
	_,err=RedisCache.Get(function.MakeKey(CACHE_USER_LOGIN_FORBID, uid)).Result()
	if err == redis.Nil {
		forbid=false
	} else if err != nil {
		forbid=true
	}
	return
}
func CheckUserState(uid uint32) (state bool, e error) {
	//验证
	state=true
	return
}