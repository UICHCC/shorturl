package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

const (
	LONG_EXPIRE     = 12 * time.Hour
	SHORT_EXPIRE    = time.Minute
	SHORT_PREFIX    = "alias_"
	HEADER_MENU_KEY = "header_menu_key"
	MENU_KEY        = "menu_key"
)

func SetKey(k string, v interface{}, expire time.Duration) error {
	return rdb.Set(ctx, k, v, expire).Err()
}

func GetKey(k string) *redis.StringCmd {
	return rdb.Get(ctx, k)
}
