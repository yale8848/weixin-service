// Create by Yale 2019/6/14 14:10
package do

import (
	"encoding/json"
	"io/ioutil"
)

type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Result
}

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
type XCXQRCodeParams struct {
	Path  string `json:"path"`
	Width int    `json:"width"`
}
type XCXCode2SessionResult struct {
	Result
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}
type XCXCodeParams struct {
	Path      string `json:"path"`
	Width     int    `json:"width"`
	AutoColor bool   `json:"auto_color"`
	LineColor *RGB   `json:"line_color"`
	IsHyaline bool   `json:"is_hyaline"`
}

type XCXCodeUnlimitedParams struct {
	Path      string `json:"path"`
	Width     int    `json:"width"`
	AutoColor bool   `json:"auto_color"`
	LineColor *RGB   `json:"line_color"`
	IsHyaline bool   `json:"is_hyaline"`
	Scene     string `json:"scene"`
}

type XCXInfo struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func (xi *XCXInfo) Parse() {
	b, er := ioutil.ReadFile("xcxinfo.json")
	if er != nil {
		panic(er)
	}
	_ = json.Unmarshal(b, xi)
}
