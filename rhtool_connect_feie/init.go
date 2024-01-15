package printer

import (
	"crypto/sha1"
	"encoding/hex"
)

type FeieParam struct {
	User string
	Ukey string
	Furl string
}

func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1 := hex.EncodeToString(s[:])
	return strsha1
}

type PrinterStatusResp struct {
	Msg                string `json:"msg"`
	Ret                int    `json:"ret"`
	Data               string `json:"data"`
	ServerExecutedTime int    `json:"serverExecutedTime"`
}

type PrinterStatusRespAdd struct {
	Msg                string `json:"msg"`
	Ret                int    `json:"ret"`
	Data               Resp   `json:"data"`
	ServerExecutedTime int    `json:"serverExecutedTime"`
}

type Resp struct {
	OK      []string `json:"ok"`
	No      []string `json:"no"`
	NoGuide []Guide  `json:"noGuide"`
}

type Guide struct {
	Sn     string `json:"sn"`
	ImgUrl string `json:"imgUrl"`
}
