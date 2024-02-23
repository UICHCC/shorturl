package util

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
)

func NextUrl() string {
	id := snowFlakeNode.Generate()
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

func ValidateTotp(code string) bool {
	return totp.Validate(code, cfg.Otp.Secret)
}
