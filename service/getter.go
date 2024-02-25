package service

func GetCaptchaInfo() Captcha {
	return cfg.Captcha
}

func GetWhitelist() []string {
	return cfg.Whitelist
}

func GetServerInfo() Address {
	return cfg.Server
}

func GetBlacklist() []string {
	return cfg.BlackList
}
