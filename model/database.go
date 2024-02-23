package model

import (
	"fmt"
	"github.com/DRJ31/shorturl-go/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ShortUrl struct {
	Abbreviation string    `json:"abbreviation"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
}

func (ShortUrl) TableName() string {
	return "shorturl"
}

func Init() (*gorm.DB, error) {
	formatStr := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	cf := util.GetConfig()
	dsn := fmt.Sprintf(formatStr, cf.Database.Username, cf.Database.Password, cf.Database.Host, cf.Database.Port, cf.Database.Name)
	Conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return Conn, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
