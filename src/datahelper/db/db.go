package db

import (
	"database/sql"
	"github.com/go-redis/redis"
	"utils/config"
)

var mysqlmain *sql.DB
var rediscache *redis.Client

func Init(config config.Config) (err error) {
	mysqlmain, err = sql.Open("mysql", config.Mysql)
	if err != nil {
		return  // Just for example purpose. You should use proper error handling instead of panic
	}

	opt, err := redis.ParseURL(config.Redis)
	if err != nil {
		return
	}
	// Create client as usually.
	rediscache= redis.NewClient(opt)
	return
}