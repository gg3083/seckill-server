package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

const (
	DefaultId = "id:generator:default"
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

//得到cache的名字
func getIdCacheName(idType string) string {
	if idType == "" {
		return DefaultId
	}
	return "id:generator:" + idType
}

//得到一个id
func GetOneId(idType string) (int64, error) {
	key := getIdCacheName(idType)
	luaId := redis.NewScript(`
local id_key = KEYS[1]
local current = redis.call('get',id_key)
if current == false then
    redis.call('set',id_key,1)
    return '1'
end
--redis.log(redis.LOG_NOTICE,' current:'..current..':')
local result =  tonumber(current)+1
--redis.log(redis.LOG_NOTICE,' result:'..result..':')
redis.call('set',id_key,result)
return tostring(result)
	`)
	var ctx = context.Background()
	n, err := luaId.Run(ctx, rdb, []string{key}, 2).Result()

	if err != nil {
		return -1, err
	} else {
		var ret string = n.(string)
		retint, err := strconv.ParseInt(ret, 10, 64)
		if err == nil {
			return retint, err
		} else {
			return -1, err
		}
	}
}
