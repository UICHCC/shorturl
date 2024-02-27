package service

import (
	"encoding/json"
	"errors"
	"github.com/DRJ31/shorturl-go/model"
	"gorm.io/gorm"
	"regexp"
	"time"
)

func GetUrl(short string) (string, error) {
	res, err := GetKey(SHORT_PREFIX + short).Result()
	if err == nil {
		return res, nil
	}

	db, err := initDB()
	if err != nil {
		return "", err
	}
	defer closeDB(db)

	var u model.ShortUrl
	result := db.First(&u, "abbreviation = ?", short)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			_ = SetKey(SHORT_PREFIX+short, "invalid", LONG_EXPIRE)
		}
		return "", result.Error
	}

	var expire time.Duration
	if u.Manual {
		expire = SHORT_EXPIRE
	} else {
		expire = LONG_EXPIRE
	}
	err = SetKey(SHORT_PREFIX+short, u.Url, expire)
	return u.Url, err
}

func InsertUrl(short, url string, manual bool) error {
	db, err := initDB()
	if err != nil {
		return err
	}
	defer closeDB(db)

	var uCheck, uInsert model.ShortUrl
	result := db.First(&uCheck, "abbreviation = ?", short)
	if result.Error == nil {
		return gorm.ErrDuplicatedKey
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		uInsert.Abbreviation = short
		uInsert.Url = url
		result = db.Create(&uInsert)
		if result.Error != nil {
			return result.Error
		}
		var expire time.Duration
		if manual {
			expire = SHORT_EXPIRE
		} else {
			expire = LONG_EXPIRE
		}
		err = SetKey(SHORT_PREFIX+short, url, expire)
		return err
	}
	return result.Error
}

func GetMenu(isHeader bool) ([]model.Menu, error) {
	var menuList []model.Menu
	var k string
	if isHeader {
		k = HEADER_MENU_KEY
	} else {
		k = MENU_KEY
	}
	menuBytes, err := GetKey(k).Bytes()
	if err == nil {
		err = json.Unmarshal(menuBytes, &menuList)
		if err == nil {
			return menuList, nil
		}
	}

	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer closeDB(db)

	result := db.Where("is_header = ?", isHeader).Find(&menuList)
	if result.Error != nil {
		return nil, result.Error
	}
	menuBytes, err = json.Marshal(menuList)
	if err == nil {
		_ = SetKey(k, menuBytes, SHORT_EXPIRE)
	}
	return menuList, nil
}

func GetBlacklist() ([]model.Blacklist, error) {
	var blacklist []model.Blacklist
	blacklistBytes, err := GetKey(BLACKLIST_KEY).Bytes()
	if err == nil {
		err = json.Unmarshal(blacklistBytes, &blacklist)
		if err == nil {
			return blacklist, nil
		}
	}

	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer closeDB(db)

	result := db.Find(&blacklist)
	if result.Error != nil {
		return nil, result.Error
	}
	blacklistBytes, err = json.Marshal(blacklist)
	if err == nil {
		_ = SetKey(BLACKLIST_KEY, blacklistBytes, LONG_EXPIRE)
	}
	return blacklist, nil
}

func InsertBlacklistRecord(pattern string, manual bool) error {
	urlReg := regexp.MustCompile("://[-a-zA-Z0-9.]+/")
	if !urlReg.MatchString(pattern) {
		return nil
	}

	db, err := initDB()
	if err != nil {
		return err
	}
	defer closeDB(db)

	var bCheck, bInsert model.Blacklist
	result := db.First(&bCheck, "pattern = ?", pattern)
	if result.Error == nil {
		return nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		bInsert.Pattern = pattern[:len(pattern)-1]
		bInsert.Manual = manual
		result = db.Create(&bInsert)
		if result.Error != nil {
			return result.Error
		}
		_, err = GetBlacklist()
		return err
	}
	return result.Error
}
