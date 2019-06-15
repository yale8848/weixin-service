// Create by Yale 2019/6/14 16:53
package xcx

import (
	"github.com/yale8848/weixin-service/common/do"
	"image/jpeg"
	"os"
	"testing"
)

func service() XCX {
	info := do.XCXInfo{}
	info.Parse()

	return NewXCXService(info.Appid, info.Secret)
}

func TestXcxService_Code2Session(t *testing.T) {

	x := service()
	x.Code2Session("code")

}
func GetToken(x XCX) *do.AccessToken {
	to, err := x.GetAccessToken()
	if err != nil {
		panic(err)
	}
	return to
}

func TestXcxService_GetXCXQRCode(t *testing.T) {
	xcx := service()
	to := GetToken(xcx)

	param := do.XCXQRCodeParams{Path: "page/index/index", Width: 800}
	img, err := xcx.GetXCXQRCode(to.AccessToken, param)
	if err != nil {
		panic(err)
	}
	f, _ := os.Create("test.jpg")
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		panic(err)
	}
}
func TestXcxService_GetXCXCode(t *testing.T) {

	xcx := service()
	to := GetToken(xcx)

	rgb := do.RGB{R: 255, G: 0, B: 0}
	param := do.XCXCodeParams{Path: "page/index/index", Width: 800, AutoColor: true, LineColor: &rgb, IsHyaline: true}
	img, err := xcx.GetXCXCode(to.AccessToken, param)
	if err != nil {
		panic(err)
	}
	f, _ := os.Create("test.jpg")
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		panic(err)
	}

}

func TestXcxService_GetXCXCodeUnlimited(t *testing.T) {
	xcx := service()
	to := GetToken(xcx)

	rgb := do.RGB{R: 255, G: 0, B: 0}
	param := do.XCXCodeUnlimitedParams{"page/index/index", 800, true, &rgb, true, "aaa"}

	img, err := xcx.GetXCXCodeUnlimited(to.AccessToken, param)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("test.jpg")
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		panic(err)
	}
}
