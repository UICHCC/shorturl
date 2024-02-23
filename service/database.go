package service

import (
	"errors"
	"github.com/DRJ31/shorturl-go/model"
	"github.com/DRJ31/shorturl-go/util"
	"gorm.io/gorm"
)

func GetUrl(short string) (string, error) {
	db, err := util.InitDB()
	if err != nil {
		return "", err
	}
	defer util.CloseDB(db)

	var u model.ShortUrl
	result := db.First(&u, "abbreviation = ?", short)
	if result.Error != nil {
		return "", err
	}
	return u.Url, nil
}

func InsertUrl(short, url string) error {
	db, err := util.InitDB()
	if err != nil {
		return err
	}
	defer util.CloseDB(db)

	var uCheck, uInsert model.ShortUrl
	result := db.First(&uCheck, "abbreviation = ?", short)
	if result.Error == nil {
		return gorm.ErrDuplicatedKey
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		uInsert.Abbreviation = short
		uInsert.Url = url
		result = db.Create(&uInsert)
		return result.Error
	}
	return result.Error
}
