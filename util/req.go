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
	blacklist, _ := service.GetBlacklist(false)
	for _, record := range blacklist {
		reg := regexp.MustCompile(record.Pattern)
		if reg.MatchString(strings.ToLower(url)) {
			log.Println(reg.FindString(strings.ToLower(url)), record.Pattern)
			valid = false
			return valid
		}
	}
	blockRules, _ := parseBlockRules()
	for _, rule := range blockRules {
		reg, err := regexp.Compile(rule)
		if err != nil {
			continue
		}
		if reg.MatchString(strings.ToLower(url)) {
			log.Println(reg.FindString(strings.ToLower(url)), rule)
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

// parseBlockRules Parse the block rule of rule list
func parseBlockRules() ([]string, error) {
	listCached, err := service.GetBlacklistExtraCache()
	if err == nil {
		return listCached, err
	}

	// Get Result
	url := service.GetBlacklistUrl()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resultByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	results := strings.Split(string(resultByte), "\n")

	// Parse rules
	rules := make([]string, 0, len(results))
	regMain := regexp.MustCompile(`^\|\|[-a-z0-9A-Z./*]+\^?\$(all|doc)`)
	regEnd := regexp.MustCompile(`[\^|]$`)
	regEx := regexp.MustCompile(`/\$doc`)
	regDoc := regexp.MustCompile(`\^?\$doc`)
	regBan := regexp.MustCompile(`^!`)
	regStart := regexp.MustCompile(`^\|\|`)
	for _, result := range results {
		var s string
		var end int
		const START_INDEX = 2
		fsm := regMain.FindString(result)
		fsx := regEx.FindString(result)
		fsd := regDoc.FindString(result)
		fsb := regBan.FindString(result)
		if len(fsb) > 0 {
			continue
		}
		if len(fsm) > 0 {
			s = strings.Split(fsm, "$")[0]
			if regEnd.MatchString(s) {
				end = len(s) - 1
			} else {
				end = len(s)
			}
			rules = append(rules, s[START_INDEX:end])
		} else if len(fsx) > 0 {
			s = strings.Split(result, `/$doc`)[0]
			if strings.Contains(s, "?=") {
				continue
			}
			if strings.Contains(s, "||") {
				s = strings.ReplaceAll(s, "||", "")
			}
			rules = append(rules, s[1:])
		} else if len(fsd) > 0 {
			s = strings.Split(result, "$")[0]
			start := 0
			// Handling special case
			if strings.Contains(s, "?") {
				s = strings.ReplaceAll(s, "?", `\?`)
			}
			if s[:1] == "*" {
				start = 1
			}
			// Final check
			if regEnd.MatchString(s) {
				end = len(s) - 1
			} else {
				end = len(s)
			}
			if regStart.MatchString(s) {
				start = START_INDEX
			}
			rules = append(rules, s[start:end])
		}
	}
	err = service.SetBlacklistExtra(rules)
	if err != nil {
		return nil, err
	}
	return rules, err
}
