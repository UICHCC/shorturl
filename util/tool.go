package util

import (
	"encoding/base64"
	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2"
)

var node *snowflake.Node

func InitSnowflake() {
	node, _ = snowflake.NewNode(1)
}

func NextUrl() string {
	id := node.Generate()
	return id.Base58()
}

func B64Decode(b64Str string) string {
	result, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return ""
	}
	return string(result)
}

func GetIP(c *fiber.Ctx) string {
	return string(c.Request().Header.Peek("X-Real-IP"))
}
