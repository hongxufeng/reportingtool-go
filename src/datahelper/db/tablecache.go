package db

import (
	"utils/function"
	"model"
	"github.com/go-redis/redis"
)

func SetSelectorBarCache(selectorfeild string,selectorhtml string) (err error) {
	err=RedisCache.Set(function.MakeKey(CACHE_TABLE_SELECTOR_BAR, selectorfeild),string(selectorhtml),model.User_Info_Persistence_Duration).Err()
	return
}

func GetSelectorBarCache(selectorfeild string)  (being bool,html string){
	html,err:=RedisCache.Get(function.MakeKey(CACHE_USER_LOGIN_FORBID, selectorfeild)).Result()
	if err == redis.Nil {
		being=false
	} else if err != nil {
		being=false
	} else {
		being=true
	}
	return
}
