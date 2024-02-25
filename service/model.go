/**
 * @filename: util/model.go
 * @description: util相关数据结构
 */

package service

import (
	"time"
)

// Address 服务器IP和端口信息
type Address struct {
	Host string `json:"host"` // IP地址
	Port int    `json:"port"` // 端口号
}

// Database 数据库相关配置
type Database struct {
	Address
	Username string `json:"username"` // 数据库用户名
	Password string `json:"password"` // 数据库密码
	Name     string `json:"name"`     // 数据库名
}

// Captcha hCaptcha相关配置
type Captcha struct {
	Key    string `json:"key"`    // hCaptcha网站Key
	Secret string `json:"secret"` // hCaptcha对应Secret
}

// CaptchaResponse hCaptcha返回内容
type CaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Credit      bool      `json:"credit"`
}

type OTP struct {
	Email  string `json:"email"`
	Secret string `json:"secret"`
}

// Config 配置文件Struct
type Config struct {
	Server       Address  `json:"server"`       // Host and port of website
	Database     Database `json:"database"`     // Configuration of database
	Captcha      Captcha  `json:"captcha"`      // Information of hCaptcha
	Whitelist    []string `json:"whitelist"`    // Urls that do not need to verify
	AllowOrigins string   `json:"allowOrigins"` // Cross origin
	Otp          OTP      `json:"otp"`          // Information on otp settings
	Redis        Address  `json:"redis"`        // Host and port of redis server
	Blacklist    []string `json:"blacklist"`    // Rules on filtering urls
}
