package wechat

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type Wechat struct {
	Appid     string `json:"appid"`
	AppSecret string `json:"app_secret"`

	_redis *redis.Client

	// redis存储的key前缀
	Prefix string `json:"prefix"`

	log *zap.Logger
}

type WechatServices interface {
	CgiGetJsapiTicket(AccessToken string) ([]byte, error)
	CgiGetAccessToken() ([]byte, error)
	SnsGetAccessToken(code string) (ResultForAuthorizationCode, error)
	SnsGetUserInfo(access_token, openid string) (ResultForUserInfo, error)
	CgiGetUserUnionID(access_token, openid string) (ResultForUnionInfo, error)
	SnsGetAndSetJsApiTicket() (JsApiTicket, error)
	SnsGetAndSetAccessToken() (BodyAccessToken, error)
}

func NewWechat(Appid, AppSecret, Prefix string, _redis *redis.Client, log *zap.Logger) WechatServices {
	svc := new(Wechat)
	svc.Appid = Appid
	svc.AppSecret = AppSecret
	svc.Prefix = Prefix
	svc.log = log
	svc._redis = _redis
	return svc
}
