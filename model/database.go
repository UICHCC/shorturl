package model

import (
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
