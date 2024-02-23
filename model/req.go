package model

type GenerateReq struct {
	LongUrl string `json:"longUrl"`
	Token   string `json:"token"`
}

type ManualReq struct {
	Short  string `json:"short"`
	Origin string `json:"origin"`
	Code   string `json:"code"`
}
