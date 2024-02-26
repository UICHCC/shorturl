/**
 * @filename: util/initialize.go
 * @description: 配置文件相关
 */

package service

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/django/v3"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var cfg *Config
var once sync.Once

func (addr Address) String() string {
	return fmt.Sprintf("%s:%d", addr.Host, addr.Port)
}

// InitConfig Initialize configuration file
func initConfig() {
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
}

// InitDB Initialize database connection
func initDB() (*gorm.DB, error) {
	formatStr := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	db := cfg.Database
	dsn := fmt.Sprintf(formatStr, db.Username, db.Password, db.Host, db.Port, db.Name)
	Conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return Conn, nil
}

// CloseDB Close database connection
func closeDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// InitRedis Initialize redis connection
func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.String(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// InitApp Initialize application
func InitApp() *fiber.App {
	initConfig()
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(compress.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowMethods: "POST,OPTIONS",
	}))
	initRedis()
	return app
}
