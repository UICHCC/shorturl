/**
 * @filename: util/config.go
 * @description: 配置文件相关
 */

package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

var cfg *Config
var once sync.Once

// String 把Address转换成网址
func (addr Address) String() string {
	return fmt.Sprintf("%s:%d", addr.Host, addr.Port)
}

// GetConfig 读取配置文件
func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		jsonFile, err := os.Open("config.json")
		if err != nil {
			log.Fatal("[Error] config.json 配置文件不存在")
		}
		defer jsonFile.Close()

		err = json.NewDecoder(jsonFile).Decode(cfg)
		if err != nil {
			log.Fatal("[Error] 配置文件解析失败")
		}
	})
	return cfg
}
