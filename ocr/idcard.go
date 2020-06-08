package ocr

import (
	"encoding/json"
	"fmt"

	"github.com/yellbuy/wechat/context"
	"github.com/yellbuy/wechat/util"
)

const (
	idcardOcrURL = "https://api.weixin.qq.com/cv/ocr/idcard"
)

//Menu struct
type IdCard struct {
	*context.Context
}

//ResMedia 响应结果
type ResOcrIdCard struct {
	util.CommonError
	Type        string `json:"type"`
	Name        string `json:"name"`
	Id          string `json:"id"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

//NewImg 实例
func NewIdCard(context *context.Context) *IdCard {
	idCard := new(IdCard)
	idCard.Context = context
	return idCard
}

// 图片检查
func (idCard *IdCard) Ocr(imgUrl string) (*ResOcrIdCard, error) {
	accessToken, err := idCard.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s??type=MODE&img_url=%s&access_token=%s", idcardOcrURL, imgUrl, accessToken)

	response, err := util.PostJSON(uri, nil)
	if err != nil {
		return nil, err
	}
	resOcrIdCard := new(ResOcrIdCard)
	err = json.Unmarshal(response, resOcrIdCard)
	if err != nil {
		return resOcrIdCard, err
	}
	if resOcrIdCard.ErrCode != 0 {
		err = fmt.Errorf("MediaCheckAsync Error , errcode=%d , errmsg=%s", resOcrIdCard.ErrCode, resOcrIdCard.ErrMsg)
		return resOcrIdCard, err
	}
	return resOcrIdCard, nil
}
