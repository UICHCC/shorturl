/**
 * @filename: util/model.go
 * @description: util相关数据结构
 */

package util

import "time"

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

// Config 配置文件Struct
type Config struct {
	Server       Address  `json:"server"`       // 当前网站的Host和端口
	Database     Database `json:"database"`     // 数据库相关配置
	Captcha      Captcha  `json:"captcha"`      // 是否限制请求域名
	Whitelist    []string `json:"whitelist"`    // 允许跳过验证的网址
	AllowOrigins string   `json:"allowOrigins"` // 允许跨域的网站
}
