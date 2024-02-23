package controller

import (
	"errors"
	"github.com/DRJ31/shorturl-go/model"
	"github.com/DRJ31/shorturl-go/service"
	"github.com/DRJ31/shorturl-go/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func Home(c *fiber.Ctx) error {
	cfg := util.GetConfig()
	return c.Render("index", fiber.Map{
		"siteKey": cfg.Captcha.Key,
	})
}

func Generate(c *fiber.Ctx) error {
	var gr model.GenerateReq
	cfg := util.GetConfig()
	if err := c.BodyParser(&gr); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid Parameters"})
	}

	// 检查长网址
	longUrl := util.B64Decode(gr.LongUrl)
	log.Println(longUrl)
	if len(longUrl) == 0 {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid URL"})
	}

	// 验证hCaptcha
	skip := false // 是否跳过验证
	for _, u := range cfg.Whitelist {
		if strings.Contains(longUrl, u) {
			skip = true
			break
		}
	}
	if !skip {
		err := util.VerifyCode(gr.Token)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusForbidden)
			return c.JSON(fiber.Map{"message": "Token is expired"})
		}
	}

	var short string
	for {
		short = util.NextUrl()
		err := service.InsertUrl(short, longUrl)
		if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
			continue
		} else if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{"message": "Failed to generate short url"})
		}
		break
	}

	if skip {
		return c.JSON(fiber.Map{
			"Code":     1,
			"ShortUrl": c.BaseURL() + "/" + short,
		})
	}

	return c.JSON(fiber.Map{
		"code": http.StatusOK,
		"url":  c.BaseURL() + "/" + short,
	})
}

func Redirect(c *fiber.Ctx) error {
	short := c.Params("short")
	url, err := service.GetUrl(short)
	if err != nil {
		log.Println(err)
		return c.Redirect("/404")
	}
	return c.Redirect(url)
}
