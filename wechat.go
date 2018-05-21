package wechat

type Wechat struct {

	appID string
	appSecret string
	host string
	Token string
}

func NewWechat(appid, appSecret, host string) *Wechat {
	if host == "" {
		host = "https://api.weixin.qq.com"
	}
	return &Wechat{
		appID: appid,
		appSecret: appSecret,
		host: host,
	}
}

