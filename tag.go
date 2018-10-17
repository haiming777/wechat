package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

//Tag 用户标签管理
type Tag struct {
	ResultError
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Count int    `json:"count,omitempty"`
}

//TagCreate 创建用户标签  标签名（30个字符以内）
func (wx *Wechat) TagCreate(name string, id *int) error {
	URL := wx.getURL(TagCreateAPI, wx.GetToken())
	reqBody := fmt.Sprintf(`{"tag":{"name":"%s"}}`, name)

	resp, err := client.Post(URL, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return err
	}

	result := struct {
		ResultError
		Tag struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"tag,omitempty"`
	}{}

	if err := getResultError(resp, &result); err != nil {
		return err
	}
	*id = result.Tag.ID
	return nil
}

//TagDelete 删除标签
//请注意，当某个标签下的粉丝超过10w时，后台不可直接删除标签。
// 此时，开发者可以对该标签下的openid列表，先进行取消标签的操作，直到粉丝数不超过10w后，才可直接删除该标签。
func (wx *Wechat) TagDelete(id int) error {
	URL := wx.getURL(TagDeleteAPI, wx.GetToken())
	reqBody := fmt.Sprintf(`{"tag":{"id":%d}}`, id)

	resp, err := client.Post(URL, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	return getResultError(resp, &ResultError{})
}

//TagList 获取公众号已创建的标签
func (wx *Wechat) TagList(tags *[]*Tag) error {
	URL := wx.getURL(TagListAPI, wx.GetToken())

	resp, err := client.Get(URL)
	if err != nil {
		return err
	}

	result := struct {
		ResultError
		Tags *[]*Tag `json:"tags,omitempty"`
	}{
		Tags: tags,
	}

	return getResultError(resp, &result)
}

//TagUpdate 更新Tag
func (wx *Wechat) TagUpdate(id int, name string) error {
	URL := wx.getURL(TagUpdateAPI, wx.GetToken())
	reqBody := fmt.Sprintf(`{"tag":{"id":%d,"name":"%s"}}`, id, name)

	resp, err := client.Post(URL, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	return getResultError(resp, &ResultError{})
}

//BatchTagging 批量为用户打标签
func (wx *Wechat) BatchTagging(id int, openids ...string) error {
	URL := wx.getURL(BatchTaggingAPI, wx.GetToken())
	// reqBody := fmt.Sprintf(`{"tag":{"id":%d,"name":"%s"}}`, id, name)
	reqBody := struct {
		OpenidList []string `json:"openid_list,omitempty"`
		TagID      int      `json:"tagid,omitempty"`
	}{
		OpenidList: openids,
		TagID:      id,
	}

	data, _ := json.Marshal(&reqBody)

	resp, err := client.Post(URL, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	return getResultError(resp, &ResultError{})
}

// BatchUnTagging 批量为用户取消标签
func (wx *Wechat) BatchUnTagging(id int, openids ...string) error {
	URL := wx.getURL(BatchUnTaggingAPI, wx.GetToken())
	// reqBody := fmt.Sprintf(`{"tag":{"id":%d,"name":"%s"}}`, id, name)
	reqBody := struct {
		OpenidList []string `json:"openid_list,omitempty"`
		TagID      int      `json:"tagid,omitempty"`
	}{
		OpenidList: openids,
		TagID:      id,
	}

	data, _ := json.Marshal(&reqBody)

	resp, err := client.Post(URL, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	return getResultError(resp, &ResultError{})
}

//GetUserTag 获取用户身上的标签列表
func (wx *Wechat) GetUserTag(openid string, tagids *[]string) error {
	URL := wx.getURL(GetUserTagAPI, wx.GetToken())

	resp, err := client.Get(URL)
	if err != nil {
		return err
	}

	result := struct {
		ResultError
		TagIDList *[]string `json:"tagid_list,omitempty"`
	}{
		TagIDList: tagids,
	}

	return getResultError(resp, &result)
}
