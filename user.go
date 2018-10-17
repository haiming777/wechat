package wechat

import (
	"fmt"
	"strings"
)

//UserList 粉丝列表
type UserList struct {
	ResultError
	Count int `json:"count,omitempty"`
	Data  struct {
		OpenID []string `json:"openid,omitempty"`
	} `json:"data,omitempty"`
	NextOpenID string `json:"next_openid,omitempty"`
}

//TagUserList 获取标签下的粉丝列表
func (wx *Wechat) TagUserList(tagid int, nextOpenID string, list *UserList) error {
	URL := wx.getURL(TagUserListAPI, wx.GetToken())
	reqBody := fmt.Sprintf(`{"tagid":%d,"next_openid":"%s"}`, tagid, nextOpenID)
	if nextOpenID == "" {
		reqBody = fmt.Sprintf(`{"tagid":%d}`, tagid)
	}
	resp, err := client.Post(URL, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return err
	}

	return getResultError(resp, list)
}
