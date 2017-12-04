package db

import (
	"utils/function"
	"model"
	"github.com/go-redis/redis"
)

func SetSelectorBarCacheDuration(table string,selectorfeild string)  (err error){
	err=RedisCache.Expire(function.MakeKey(CACHE_TABLE_SELECTOR_BAR,table, selectorfeild),model.User_Info_Persistence_Duration).Err()
	return
}

func HSetSelectorBarCache(table string,selectorfeild string,key string,value string) (err error) {
	err=RedisCache.HSet(function.MakeKey(CACHE_TABLE_SELECTOR_BAR,table, selectorfeild),key,value).Err()
	return
}

func HGetSelectorBarCache(table string,selectorfeild string)  (being bool,selectordata map[string]string){
	selectordata,err:=RedisCache.HGetAll(function.MakeKey(CACHE_TABLE_SELECTOR_BAR,table, selectorfeild)).Result()
	if err == redis.Nil {
		being=false
	} else if err != nil {
		being=false
	} else if len(selectordata)==0{
		being=false
	}else {
		being=true
	}
	return
}
