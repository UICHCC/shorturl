package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

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
