package service

import (
	"errors"
	"github.com/DRJ31/shorturl-go/model"
	"gorm.io/gorm"
)

func GetUrl(short string) (string, error) {
	db, err := model.Init()
	if err != nil {
		return "", err
	}
	defer model.Close(db)

	var u model.ShortUrl
	result := db.First(&u, "abbreviation = ?", short)
	if result.Error != nil {
		return "", err
	}
	return u.Url, nil
}

func InsertUrl(short, url string) error {
	db, err := model.Init()
	if err != nil {
		return err
	}
	defer model.Close(db)

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
