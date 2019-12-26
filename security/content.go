package security

import (
	"encoding/json"
	"fmt"

	"github.com/yellbuy/wechat/context"
	"github.com/yellbuy/wechat/util"
)

const (
	imgSecCheckURL   = "https://api.weixin.qq.com/wxa/img_sec_check"
	midiaSecCheckURL = "https://api.weixin.qq.com/wxa/media_check_async"
	msgSecCheckURL   = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

//Menu struct
type Content struct {
	*context.Context
}

//ButtonNew 图文消息菜单
type Media struct {
	ContentType string `json:"contentType"`
	Value       []byte `json:"value"`
}

//reqImg  请求图片数据
type reqImg struct {
	Media *Media `json:"media"`
}

type ResImg struct {
	util.CommonError
}

//ResMedia 响应结果
type ResMedia struct {
	util.CommonError
	TraceId string `json:"trace_id"`
}

//NewImg 实例
func NewContent(context *context.Context) *Content {
	content := new(Content)
	content.Context = context
	return content
}

// 图片检查
func (content *Content) ImgSecCheck(media *Media) error {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", imgSecCheckURL, accessToken)
	reqImg := &reqImg{
		Media: media,
	}

	response, err := util.PostJSON(uri, reqImg)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "ImgSecCheck")
}

// 异步校验图片/音频是否含有违法违规内容。
// media_type	number		是	1:音频;2:图片
func (content *Content) MediaCheckAsync(mediaUrl string, mediaType uint8) (resMedia ResMedia, err error) {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", imgSecCheckURL, accessToken)
	var req struct {
		MediaUrl  string `json:"mediaUrl"`
		MediaType uint8  `json:"mediaType"`
	}
	req.MediaUrl = mediaUrl
	req.MediaType = mediaType

	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &resMedia)
	if err != nil {
		return
	}
	if resMedia.ErrCode != 0 {
		err = fmt.Errorf("MediaCheckAsync Error , errcode=%d , errmsg=%s", resMedia.ErrCode, resMedia.ErrMsg)
		return
	}
	return
}

// 异步校验图片/音频是否含有违法违规内容。
// media_type	number		是	1:音频;2:图片
func (content *Content) MsgSecCheck(msg string) error {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", msgSecCheckURL, accessToken)
	var req struct {
		Content string `json:"content"`
	}
	req.Content = msg

	response, err := util.PostJSON(uri, req)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "MsgSecCheck")
}
