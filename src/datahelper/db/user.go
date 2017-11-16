package db

import (
	"utils/function"
	"github.com/go-redis/redis"
	"time"
	"fmt"
)

func UserLoginErrCnt(uid uint32) (cnt int64, err error) {
	cnt, err = rediscache.Incr(function.MakeKey(CACHE_USER_LOGIN_ERROR, uid)).Result()
	if err != redis.Nil&&err != nil {
		return
	}  else {
		cnt++
		err=rediscache.Set(function.MakeKey(CACHE_USER_LOGIN_ERROR, uid),cnt,time.Minute * 10).Err()
	}
	fmt.Println(cnt)
	return
}
func SetUserForbid(uid uint32) (err error) {
	err=rediscache.Set(function.MakeKey(CACHE_USER_LOGIN_FORBID, uid),true,time.Minute * 10).Err()
	return
}