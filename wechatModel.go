package wechat

import (
	"encoding/xml"
	"time"
)

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WX CODE Response
type WxCodeResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

//{
//"openid":" OPENID",
//"nickname": NICKNAME,
//"sex":"1",
//"province":"PROVINCE",
//"city":"CITY",
//"country":"COUNTRY",
//"headimgurl":"https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
//"privilege":[ "PRIVILEGE1" "PRIVILEGE2"     ],
//"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
//}

// WX INFO Response
type WxInfoResponse struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// WX CODE Error Response
type WxCodeErrorResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// {"access_token":"ACCESS_TOKEN","expires_in":7200}
type BodyAccessToken struct {
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	Errcode     int           `json:"errcode"`
	Errmsg      string        `json:"errmsg"`
}

// jsapi_ticket
// {
// "errcode":0,
// "errmsg":"ok",
// "ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
// "expires_in":7200
// }
type JsApiTicket struct {
	Errcode   int           `json:"errcode"`
	Errmsg    string        `json:"errmsg"`
	Ticket    string        `json:"ticket"`
	ExpiresIn time.Duration `json:"expires_in"`
}

type RequestWxShareConfig struct {
	Uri       string   `json:"uri" validate:"required"`
	JsApiList []string `json:"js_api_list" validate:"required"`
	Debug     bool     `json:"debug"`
}

type WxShareConfig struct {
	AppId     string   `json:"appId"`
	Timestamp int64    `json:"timestamp"`
	NonceStr  string   `json:"nonceStr"`
	Signature string   `json:"signature"`
	JsApiList []string `json:"jsApiList"`
	Debug     bool     `json:"debug"`
}

type ResultForAuthorizationCode struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Openid         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	Unionid        string `json:"unionid"`
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
}

type ResultForUserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	Errcode    int      `json:"errcode"`
	Errmsg     string   `json:"errmsg"`
}

type ResultForUnionInfo struct {
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int    `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	Groupid        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
}
