package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

//Error 微信错误接口
type Error interface {
	Error() error
}

//ResultError 返回值错误
type ResultError struct {
	ErrMsg  string `json:"errmsg,omitempty"`
	ErrCode int    `json:"errcode,omitempty"`
}

//Error 返回错误
func (c ResultError) Error() error {
	if c.ErrCode != 0 {
		return fmt.Errorf("errcode:%d,errmsg:%s", c.ErrCode, c.ErrMsg)
	}
	return nil
}

//getResultError 获取微信返回内容是否错误
func getResultError(resp *http.Response, result Error) error {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("[Read Response Body]:" + err.Error())
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	return result.Error()
}
