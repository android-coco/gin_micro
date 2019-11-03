package redis

import (
	"gin_micro/util"
	"testing"
	"time"
)

func TestGetRedisDb(t *testing.T) {
	redis := GetRedisDb()
	err := redis.Set(util.RedisKeyToken, "token", 100*time.Second).Err()
	if err != nil {
		t.Log(err)
		t.Error(err)
	}

}
