package db

import (
	"utils/function"
	"github.com/go-redis/redis"
	"model"
	"database/sql"
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
func GetUidbyName(name string) (uid uint32,err error) {
	uid=0
	query := "SELECT uid FROM w_user_list WHERE username=?"
	result, err := MysqlMain.Query(query, name)
	if err!=nil {
		return
	}
	defer result.Close()
	if result.Next(){
		err = result.Scan(&uid)
	}
	return
}
func Query(query string) (*sql.Rows, error){
	//还是不要设置redis
	return MysqlMain.Query(query)

}