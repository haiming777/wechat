package wechat

import (
	"time"
	"net/url"
)

type Token struct {
	requestURL string
	AccessAoken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

func NewToken(appid, secret, host string) *Token {
	values := url.Values{}
	values.Set("grant_type", "client_credential")
	values.Set("appid", appid)
	values.Set("secret", secret)

	requestURL := host + "/cgi-bin/token?" + values.Encode()

	return &Token{
		requestURL: requestURL,
	}
}


func (t *Token)GetToken(host string) string  {
	if time.Now().Unix() > t.ExpiresIn {

	}

	return t.AccessAoken
}

func (t *Token)  refreshToken() error {
	get(t.requestURL)
	return nil
}




