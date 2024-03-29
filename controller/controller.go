package controller

import (
	"errors"
	"fmt"
	"github.com/DRJ31/shorturl-go/model"
	"github.com/DRJ31/shorturl-go/service"
	"github.com/DRJ31/shorturl-go/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"siteKey": service.GetCaptchaInfo().Key,
		"year":    time.Now().Year(),
	})
}

func ManualPage(c *fiber.Ctx) error {
	return c.Render("manual", fiber.Map{
		"year": time.Now().Year(),
	})
}

func Manual(c *fiber.Ctx) error {
	var mr model.ManualDto
	if err := c.BodyParser(&mr); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid Parameters"})
	}

	if !util.ValidateTotp(mr.Code) {
		c.Status(http.StatusForbidden)
		return c.JSON(fiber.Map{"message": "Invalid Code"})
	}

	if !util.CheckAlias(mr.Short) {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid short url!"})
	}

	_, err := service.GetUrl(mr.Short)
	log.Println("Get url:", err)
	if err == nil {
		c.Status(http.StatusForbidden)
		return c.JSON(fiber.Map{"message": "The alias is existed. Please contact administrator."})
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		longUrl := util.B64Decode(mr.Origin)
		err = service.InsertUrl(mr.Short, longUrl, true)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{"message": "Failed to add record."})
		}
		return c.JSON(fiber.Map{"message": fmt.Sprintf("You have added %v/%v", c.BaseURL(), mr.Short)})
	}
	c.Status(http.StatusInternalServerError)
	return c.JSON(fiber.Map{"message": "Operation failed."})
}

func Generate(c *fiber.Ctx) error {
	var gr model.GenerateDto
	if err := c.BodyParser(&gr); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid Parameters"})
	}

	// 检查长网址
	longUrl := util.B64Decode(gr.LongUrl)
	log.Println(longUrl)
	if len(longUrl) == 0 {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "Invalid URL"})
	}

	// 验证hCaptcha
	skip := false // 是否跳过验证
	whitelist := service.GetWhitelist()
	for _, u := range whitelist {
		if strings.Contains(longUrl, u) {
			skip = true
			break
		}
	}
	if !skip {
		if !util.ValidateUrl(longUrl) {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{"message": "Failed to pass URL check."})
		}
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
		err := service.InsertUrl(short, longUrl, false)
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
	if err != nil || url == "invalid" {
		log.Println(err)
		return c.Redirect("/404")
	}
	return c.Redirect(url)
}

func PageNotFound(c *fiber.Ctx) error {
	var headerMenu, menu []model.Menu
	var err error
	headerMenu, err = service.GetMenu(true)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "Failed to get header menu."})
	}
	menu, err = service.GetMenu(false)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "Failed to get menu."})
	}
	return c.Render("404", fiber.Map{
		"headers": headerMenu,
		"menus":   menu,
		"year":    time.Now().Year(),
		"img":     util.GetWallpaperUrl(),
	})
}
