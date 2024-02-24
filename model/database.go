package model

import (
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

func (ShortUrl) TableName() string {
	return "shorturl"
}

func (Menu) TableName() string {
	return "menu"
}
