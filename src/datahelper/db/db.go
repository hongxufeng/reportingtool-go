package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"
	"utils/config"
)

var mysqlmain *sql.DB
var rediscache *redis.Client

const (
	CACHE_USER_DETAIL = "W_Redis_Cache_User_Detail_Byte"
	CACHE_USER_LOGIN_ERROR="W_Redis_Cache_User_Login_Error_Cnt"
	CACHE_USER_LOGIN_FORBID="W_Redis_Cache_User_Login_Forbid_Bool"
)

func Init(config config.Config) (err error) {
	mysqlmain, err = sql.Open("mysql", config.Mysql)
	if err != nil {
		return  // Just for example purpose. You should use proper error handling instead of panic
	}
	err = mysqlmain.Ping()
	if err != nil {
		return // proper error handling instead of panic in your app
	}

	opt, err := redis.ParseURL(config.Redis)
	if err != nil {
		return
	}
	// Create client as usually.
	rediscache= redis.NewClient(opt)

	_, err = rediscache.Ping().Result()
	if err != nil {
		return
	}
	return
}