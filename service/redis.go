package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

const (
	LONG_EXPIRE         = 12 * time.Hour
	SHORT_EXPIRE        = time.Minute
	SHORT_PREFIX        = "alias_"
	HEADER_MENU_KEY     = "header_menu_key"
	MENU_KEY            = "menu_key"
	BLACKLIST_KEY       = "blacklist"
	BLACKLIST_EXTRA_KEY = "blacklistExtra"
)

func SetKey(k string, v interface{}, expire time.Duration) error {
	return rdb.Set(ctx, k, v, expire).Err()
}

func GetKey(k string) *redis.StringCmd {
	return rdb.Get(ctx, k)
}

func GetBlacklistExtraCache() ([]string, error) {
	b, err := GetKey(BLACKLIST_EXTRA_KEY).Bytes()
	if err != nil {
		return nil, err
	}

	var results []string
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, redis.Nil
	}
	return results, nil
}

func SetBlacklistExtra(l []string) error {
	b, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return SetKey(BLACKLIST_EXTRA_KEY, b, LONG_EXPIRE)
}
