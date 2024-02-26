package model

type GenerateDto struct {
	LongUrl string `json:"longUrl"`
	Token   string `json:"token"`
}

type ManualDto struct {
	Short  string `json:"short"`
	Origin string `json:"origin"`
	Code   string `json:"code"`
}
