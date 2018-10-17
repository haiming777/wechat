package wechat

//微信API
const (
	//TokenAPI token api
	TokenAPI = "/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	//TagCreateAPI 标签创建API
	TagCreateAPI = "/cgi-bin/tags/create?access_token=%s"
	//TagListAPI 获取公众号已创建的标签
	TagListAPI = "/cgi-bin/tags/get?access_token=%s"
	//TagUpdateAPI  编辑标签
	TagUpdateAPI = "/cgi-bin/tags/update?access_token=%s"
	//TagDeleteAPI 删除标签
	TagDeleteAPI = "/cgi-bin/tags/delete?access_token=%s"
	// BatchTaggingAPI 批量为用户打标签
	BatchTaggingAPI = "/cgi-bin/tags/members/batchtagging?access_token=%s"
	//BatchUnTaggingAPI 批量为用户取消标签
	BatchUnTaggingAPI = "/cgi-bin/tags/members/batchuntagging?access_token=%s"
	//GetUserTagAPI 获取用户身上的标签列表
	GetUserTagAPI = "/cgi-bin/tags/getidlist?access_token=%s"
	//TagUserListAPI 获取标签下粉丝列表
	TagUserListAPI = "/cgi-bin/user/tag/get?access_token=%s"
)

const (
	// MaxRetryCount 默认最大重试次数
	MaxRetryCount = 3
)
