package model

import (
	"gorm.io/gorm"
	"time"
)

type ShortUrl struct {
	Abbreviation string    `json:"abbreviation"`
	Url          string    `json:"url"`
	Manual       bool      `json:"manual"`
	CreatedAt    time.Time `json:"created_at"`
}

type Menu struct {
	ID        int       `json:"id"`
	Url       string    `json:"url"`
	Name      string    `json:"name"`
	IsHeader  bool      `json:"is_header"`
	CreatedAt time.Time `json:"created_at"`
}

type Blacklist struct {
	gorm.Model
	Pattern string `json:"pattern"`
	Manual  bool   `json:"manual"`
}

func (ShortUrl) TableName() string {
	return "shorturl"
}

func (Menu) TableName() string {
	return "menu"
}

func (Blacklist) TableName() string {
	return "blacklist"
}
