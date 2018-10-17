# Readme

## 微信域名

通用域名(api.weixin.qq.com)，使用该域名将访问官方指定就近的接入点；  
上海域名(sh.api.weixin.qq.com)，使用该域名将访问上海的接入点；  
深圳域名(sz.api.weixin.qq.com)，使用该域名将访问深圳的接入点；  
香港域名(hk.api.weixin.qq.com)，使用该域名将访问香港的接入点。  

## 微信token验证
```go
//VerifySignature 校验微信token
// demo: verfiy?signature=b88ac16ed67b5080cd6c5833302f030614cda808&echostr=8705336224975363871&timestamp=1538399775&nonce=694203948
func VerifySignature(token, signature, timestamp, nonce, echostr string) (string, error) {

	if signature == "" || timestamp == "" || nonce == "" {

		return "", errors.New("invalid parameters")
	}

	signatureCalculated := SignEndpoint(token, timestamp, nonce)
	if signature != signatureCalculated {
		return "", errors.New("invalid verifySignature")
	}

	return echostr, nil
}

```
使用案例(token: wxhellobintoptoken):
```
//VerifySignature 微信校验token
func VerifySignature(w http.ResponseWriter, r *http.Request) {
	echo, err := wechat.VerifySignature("wxhellobintoptoken",
		r.URL.Query().Get("signature"), r.URL.Query().Get("timestamp"),
		r.URL.Query().Get("nonce"), r.URL.Query().Get("echostr"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(echo))
}

```
