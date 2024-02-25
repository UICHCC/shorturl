package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pquerna/otp/totp"
	"io"
	"net/http"
	"strings"
)

// VerifyCode Verify hCaptcha status
func VerifyCode(token string) error {
	url := "https://api.hcaptcha.com/siteverify"
	data := fmt.Sprintf("secret=%v&response=%v", cfg.Captcha.Secret, token)
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	var cr CaptchaResponse
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
	return totp.Validate(code, cfg.Otp.Secret)
}
