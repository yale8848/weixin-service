// Create by Yale 2019/6/14 14:11
package xcx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yale8848/weixin-service/common/do"
	"github.com/yale8848/weixin-service/common/wxhttp"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"strings"
)

type XCX interface {
	GetAccessToken() (*do.AccessToken, error)
	GetXCXCode(accessToken string, params do.XCXCodeParams) (image.Image, error)
	GetXCXQRCode(accessToken string, params do.XCXQRCodeParams) (image.Image, error)
	GetXCXCodeUnlimited(accessToken string, params do.XCXCodeUnlimitedParams) (image.Image, error)
	Code2Session(jsCode string) (*do.XCXCode2SessionResult, error)
}

type xcxService struct {
	appId  string
	secret string
}

func NewXCXService(appId, secret string) XCX {
	return &xcxService{appId: appId, secret: secret}
}

//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (xcxs *xcxService) Code2Session(jsCode string) (*do.XCXCode2SessionResult, error) {
	u := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	res := do.XCXCode2SessionResult{}
	err := wxhttp.GetJsonObj(fmt.Sprintf(u, xcxs.appId, xcxs.secret, jsCode), &res)
	if err != nil {
		return nil, err
	}
	if res.Errcode != 0 {
		return nil, errors.New(res.Errmsg)
	}
	return &res, nil
}
func (xcxs *xcxService) getXCXImage(res *http.Response, err error) (image.Image, error) {
	if err != nil {
		return nil, err
	}

	ct := res.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		rs := do.Result{}
		err = json.Unmarshal(b, &rs)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(rs.Errmsg)
	}
	if ct == "image/jpeg" {
		by, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		img, err := jpeg.Decode(bytes.NewReader(by))
		if err != nil {
			return png.Decode(bytes.NewReader(by))
		}
		return img, err
	}
	if ct == "image/png" {
		return png.Decode(res.Body)
	}
	return nil, errors.New(ct)
}

//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (xcxs *xcxService) GetXCXQRCode(accessToken string, params do.XCXQRCodeParams) (image.Image, error) {

	u := "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=" + accessToken
	res, err := wxhttp.JsonRes(u, &params)
	return xcxs.getXCXImage(res, err)
}

//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (xcxs *xcxService) GetXCXCode(accessToken string, params do.XCXCodeParams) (image.Image, error) {
	u := "https://api.weixin.qq.com/wxa/getwxacode?access_token=" + accessToken

	res, err := wxhttp.JsonRes(u, &params)
	return xcxs.getXCXImage(res, err)

}

//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (xcxs *xcxService) GetXCXCodeUnlimited(accessToken string, params do.XCXCodeUnlimitedParams) (image.Image, error) {
	u := "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=" + accessToken
	res, err := wxhttp.JsonRes(u, &params)
	return xcxs.getXCXImage(res, err)
}

//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (xcxs *xcxService) GetAccessToken() (*do.AccessToken, error) {
	u := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	at := do.AccessToken{}
	err := wxhttp.GetJsonObj(fmt.Sprintf(u, xcxs.appId, xcxs.secret), &at)
	if err != nil {
		return nil, err
	}
	if at.Errcode != 0 {
		return nil, errors.New(at.Errmsg)
	}
	return &at, nil
}
