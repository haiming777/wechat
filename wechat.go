package wechat

import (
	"fmt"
	"sync"
	"time"
)

//Wechat 微信账号
type Wechat struct {
	appID     string
	appSecret string
	host      string

	token            string
	refreshTokenTime time.Time
	lock             *sync.RWMutex
	tokenFunc        GetToken
}

//NewWechat 获取新的微信
func NewWechat(appid, appSecret, host string, tokenFunc GetToken) *Wechat {
	if host == "" {
		host = "https://api.weixin.qq.com"
	}

	wx := &Wechat{
		lock:      new(sync.RWMutex),
		appID:     appid,
		appSecret: appSecret,
		host:      host,
	}
	if tokenFunc == nil {
		wx.tokenFunc = wx.defaultToken
	}

	return wx
}

func (wx *Wechat) getURL(url string, parameters ...interface{}) string {
	return fmt.Sprintf(wx.host+url, parameters...)
}
