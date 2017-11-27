package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"
	"utils/config"
	"github.com/hongxufeng/fileLogger"
)

var MysqlMain *sql.DB
var RedisCache *redis.Client
var DBLog *fileLogger.FileLogger

const (
	CACHE_USER_INFO = "W_Redis_Cache_User_Info_Byte"
	CACHE_USER_LOGIN_ERROR="W_Redis_Cache_User_Login_Error_Cnt"
	CACHE_USER_LOGIN_FORBID="W_Redis_Cache_User_Login_Forbid_Bool"
	CACHE_TABLE_SELECTOR_BAR="W_Redis_Cache_Table_Selector_Bar_String"
)

func Init(config *config.Config) (err error) {
	DBLog=fileLogger.NewDefaultLogger(config.LogDir, "DB_LOG.log")
	DBLog.SetPrefix("[DB] ")
	MysqlMain, err = sql.Open("mysql", config.Mysql)
	if err != nil {
		return  // Just for example purpose. You should use proper error handling instead of panic
	}
	err = MysqlMain.Ping()
	if err != nil {
		return // proper error handling instead of panic in your app
	}

	opt, err := redis.ParseURL(config.Redis)
	if err != nil {
		return
	}
	// Create client as usually.
	RedisCache= redis.NewClient(opt)

	_, err = RedisCache.Ping().Result()
	if err != nil {
		return
	}
	return
}