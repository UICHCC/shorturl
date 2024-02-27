package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DRJ31/shorturl-go/service"
	"github.com/pquerna/otp/totp"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// VerifyCode Verify hCaptcha status
func VerifyCode(token string) error {
	url := "https://api.hcaptcha.com/siteverify"
	captchaInfo := service.GetCaptchaInfo()
	data := fmt.Sprintf("secret=%v&response=%v", captchaInfo.Secret, token)
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	var cr service.CaptchaResponse
	err = json.NewDecoder(res.Body).Decode(&cr)
	if err != nil {
		return err
	}
	if !cr.Success {
		return errors.New("verify failed")
	}
	return nil
}

// GetWallpaperUrl Get Url for microsoft wallpaper
func GetWallpaperUrl() string {
	const UNSPLASH_URL = "https://source.unsplash.com/random"
	url := "https://api.drjchn.com/api/wallpaper/url"
	res, err := http.Get(url)
	if err != nil {
		return UNSPLASH_URL
	}

	defer res.Body.Close()
	resultByte, err := io.ReadAll(res.Body)
	if err != nil {
		return UNSPLASH_URL
	}
	return string(resultByte)
}

// ValidateTotp Totp validation
func ValidateTotp(code string) bool {
	otpInfo := service.GetOtpInfo()
	return totp.Validate(code, otpInfo.Secret)
}

// ValidateUrl Check if the url is valid
func ValidateUrl(url string) bool {
	valid := true
	// Check url from blacklist
	blacklist, _ := service.GetBlacklist()
	for _, record := range blacklist {
		reg := regexp.MustCompile(record.Pattern)
		if reg.MatchString(url) {
			log.Println(reg.FindString(url), record.Pattern)
			valid = false
			return valid
		}
	}

	// Check if the url has redirect
	urlReg := regexp.MustCompile("://[-a-zA-Z0-9.]+/")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(url)
	if err != nil || res.StatusCode != 200 {
		blackUrl := urlReg.FindString(url)
		_ = service.InsertBlacklistRecord(blackUrl, false)
		return false
	}
	return true
}
