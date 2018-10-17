package wechat

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

//GetToken 获取token 同时返回过期时间
type GetToken func(appid, secret string) (string, time.Time)

//GetToken 获取token
func (w *Wechat) GetToken() string {
	if w.token == "" || !w.refreshTokenTime.After(time.Now()) {
		token, refreshTokenTime := w.tokenFunc(w.appID, w.appSecret)
		w.token = token
		w.refreshTokenTime = refreshTokenTime
	}
	return w.token
}

//defaultToken 默认获取token的方法
func (w *Wechat) defaultToken(appid, secret string) (string, time.Time) {
	w.lock.Lock()
	defer w.lock.Unlock()

	URL := fmt.Sprintf(w.host+TokenAPI, appid, secret)
	retryCount := 0
retry:
	if retryCount > MaxRetryCount {
		return "", time.Now()
	}
	resp, err := client.Get(URL)
	if err != nil {
		log.Printf("[ERR][WECHAT][GET TOKEN]:%s\n", err.Error())
		time.Sleep(5 * time.Second)
		goto retry
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERR][WECHAT][GET TOKEN]:%s\n", err.Error())
		time.Sleep(5 * time.Second)
		goto retry
	}
	tokenResp := TokenResp{}
	if err := json.Unmarshal(data, &tokenResp); err != nil {
		log.Printf("[ERR][WECHAT][GET TOKEN]:%s\n", err.Error())
		time.Sleep(5 * time.Second)
		goto retry
	}

	if tokenResp.AccessToken == "" {
		log.Printf("[ERR][WECHAT][GET TOKEN]:%s, URL:%s\n", string(data), URL)
		return "", time.Now()
	}

	return tokenResp.AccessToken, time.Now().Add(time.Duration(tokenResp.ExpiresIn) - 5*time.Minute)
}

type (
	//TokenResp token api 返回值
	TokenResp struct {
		ErrCode     int    `json:"errcode,omitempty"`
		ErrMsg      string `json:"errmsg,omitempty"`
		AccessToken string `json:"access_token,omitempty"`
		ExpiresIn   int    `json:"expires_in,omitempty"`
	}
)

//SignEndpoint 微信公众号 明文模式/URL认证 签名
func SignEndpoint(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

//SignEncryptedMessage 微信公众号/企业号 密文模式消息签名
func SignEncryptedMessage(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce)+len(encryptedMsg))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

func sign(args ...string) (signature string) {
	sorter := sort.StringSlice{}
	for _, arg := range args {
		sorter = append(sorter, arg)
	}
	sorter.Sort()
	s := strings.Join(sorter, "")
	hashsum := sha1.Sum([]byte(s))
	return hex.EncodeToString(hashsum[:])
}

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
