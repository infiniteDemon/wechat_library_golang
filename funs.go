package wechat

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

// 获取JsApiTicket
func (w *Wechat) CgiGetJsapiTicket(AccessToken string) ([]byte, error) {
	getInfoLink := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", AccessToken)
	a := ClientGet(getInfoLink)
	body, err := a()
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 获取ACCESS_TOKEN
func (w *Wechat) CgiGetAccessToken() ([]byte, error) {
	getInfoLink := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", w.Appid, w.AppSecret)
	a := ClientGet(getInfoLink)
	body, err := a()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (w *Wechat) SnsGetAccessToken(code string) (ResultForAuthorizationCode, error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		w.Appid,
		w.AppSecret,
		code,
	)

	resultData := ResultForAuthorizationCode{}

	a := ClientGet(uri)
	body, err := a()
	if err != nil {
		w.log.Error("request err ", zap.Error(err))
		return resultData, err
	}

	json.Unmarshal(body, &resultData)
	return resultData, nil
}

func (w *Wechat) SnsGetUserInfo(access_token, openid string) (ResultForUserInfo, error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		access_token,
		openid,
	)

	resultData := ResultForUserInfo{}

	a := ClientGet(uri)
	body, err := a()
	if err != nil {
		w.log.Error("request err ", zap.Error(err))
		return resultData, err
	}

	json.Unmarshal(body, &resultData)
	return resultData, nil
}

func (w *Wechat) CgiGetUserUnionID(access_token, openid string) (ResultForUnionInfo, error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN",
		access_token,
		openid,
	)

	resultData := ResultForUnionInfo{}

	a := ClientGet(uri)
	body, err := a()
	if err != nil {
		w.log.Error("request err ", zap.Error(err))
		return resultData, err
	}

	json.Unmarshal(body, &resultData)
	return resultData, nil
}
