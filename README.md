## Golang微信后台接口

### Install

`go get github.com/yale8848/weixin-service@v0.1.1`

### 小程序

 #### demo
 
 ```go

  x:=xcx.NewXCXService("Appid", "Secret")
  x.Code2Session("js code")

```
 #### 接口
 
- 登录
  - Code2Session
- 接口凭证
  - GetAccessToken
- 小程序码
  - GetXCXCode
  - GetXCXQRCode
  - GetXCXCodeUnlimited
