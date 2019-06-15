// Create by Yale 2019/6/14 11:43
package wxhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yale8848/weixin-service/common/wxerrs"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Map2Values(mp map[string]string) url.Values {
	value := make(url.Values)
	for k, v := range mp {
		value.Set(k, v)
	}
	return value
}
func Values2Map(value url.Values) map[string]string {
	mp := make(map[string]string)
	for k, v := range value {
		mp[k] = v[0]
	}
	return mp
}
func GetParams(u string, params map[string]string) {

}
func Get(u string) (ret []byte, err error) {
	var (
		res *http.Response
	)
	res, err = http.Get(u)
	err = errMsg(res, err)
	if err != nil {
		return
	}
	ret, err = ioutil.ReadAll(res.Body)
	return
}
func GetRes(u string) (ret *http.Response, err error) {
	ret, err = http.Get(u)
	err = errMsg(ret, err)
	if err != nil {
		return
	}
	return
}
func JsonRes(u string, param interface{}) (*http.Response, error) {

	if param == nil {
		return nil, wxerrs.ErrObjNil
	}
	b, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	bt := bytes.Buffer{}
	bt.Write(b)

	fmt.Println(bt.String())
	res, er := http.Post(u, "application/json", &bt)
	er = errMsg(res, er)
	if er != nil {
		return nil, er
	}
	return res, nil
}
func errMsg(res *http.Response, err error) error {
	if err != nil {
		return err
	}
	if res == nil {
		return wxerrs.ErrObjNil
	}

	if res.StatusCode == 200 {
		return nil
	}
	if res.ContentLength > 0 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("StatusCode = %d, Body: %s", res.StatusCode, string(b)))
	}
	return errors.New(fmt.Sprintf("StatusCode = %v", res.StatusCode))
}
func PostForm(u string, data url.Values) (*http.Response, error) {
	res, err := http.PostForm(u, data)
	err = errMsg(res, err)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func GetJsonObj(u string, obj interface{}) (err error) {
	var (
		retByte []byte
	)
	if obj == nil {
		return wxerrs.ErrObjNil
	}
	retByte, err = Get(u)
	if err != nil {
		return
	}
	err = json.Unmarshal(retByte, obj)
	if err != nil {
		return
	}
	return nil
}
