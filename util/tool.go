package util

import (
	"encoding/base64"
	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2"
	"log"
	"regexp"
)

var snowFlakeNode *snowflake.Node

// InitSnowflake Initialize snowflake node
func InitSnowflake() {
	snowFlakeNode, _ = snowflake.NewNode(1)
}

// NextUrl Next short url
func NextUrl() string {
	id := snowFlakeNode.Generate()
	return id.Base58()
}

// B64Decode Decode base64 string
func B64Decode(b64Str string) string {
	result, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return ""
	}
	return string(result)
}

// GetIP Get IP address from X-Real-IP header
func GetIP(c *fiber.Ctx) string {
	return string(c.Request().Header.Peek("X-Real-IP"))
}

// CheckAlias Check if the short url is valid
func CheckAlias(short string) bool {
	reg := regexp.MustCompile("[0-9a-z]{1,10}")
	s := reg.FindString(short)
	return s == short
}

// CheckUrl Check if the url is valid
func CheckUrl(url string, patterns []string) bool {
	valid := true
	for _, pat := range patterns {
		reg := regexp.MustCompile(pat)
		if reg.MatchString(url) {
			log.Println(reg.FindString(url))
			valid = false
			return valid
		}
	}
	return valid
}
