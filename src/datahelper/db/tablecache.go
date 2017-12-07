package db

import (
	"model"
	"utils/function"

	"github.com/go-redis/redis"
)

//Persistence用于判断是否更新SelectorBarCache
func SetSelectorCachePersistence(table string, selectorfeild string) (err error) {
	err = RedisCache.Set(function.MakeKey(CACHE_TABLE_SELECTOR, table, selectorfeild), true, model.User_Info_Persistence_Duration).Err()
	return
}

func GetSelectorCachePersistence(table string, selectorfeild string) (being bool, err error) {
	_, err = RedisCache.Get(function.MakeKey(CACHE_TABLE_SELECTOR, table, selectorfeild)).Result()
	if err == redis.Nil {
		being = false
	} else if err != nil {
		being = false
	} else {
		being = true
	}
	return
}

func HSetSelectorBarCache(table string, selectorfeild string, key string, value string) (err error) {
	err = RedisCache.HSet(function.MakeKey(CACHE_TABLE_SELECTOR_BAR, table, selectorfeild), key, value).Err()
	return
}

func HGetSelectorBarCache(table string, selectorfeild string) (selectordata map[string]string, err error) {
	selectordata, _ = RedisCache.HGetAll(function.MakeKey(CACHE_TABLE_SELECTOR_BAR, table, selectorfeild)).Result()
	return
}
