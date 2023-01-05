package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func (w *Wechat) SnsGetAndSetJsApiTicket() (JsApiTicket, error) {

	result := JsApiTicket{}

	value, err := w._redis.Get(context.Background(), w.Prefix+RedisKeyWechatJsApiTicket).Result()
	if err != nil && err != redis.Nil {
		return JsApiTicket{}, err
	}

	if len(value) > 5 {
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			w.log.Error("bt err ", zap.Error(err))
			return JsApiTicket{}, err
		}
		w.log.Info("打印jsapi", zap.Any("data", value), zap.Any("result", result))
		return result, nil
	} else {
		accessToken, err := w.SnsGetAndSetAccessToken()
		if err != nil {
			return result, err
		}

		body, err := w.CgiGetJsapiTicket(accessToken.AccessToken)
		if err != nil {
			return result, err
		}

		json.Unmarshal(body, &result)
		if result.Errcode != 0 {
			// 获取出错了
			return result, errors.New(result.Errmsg)
		}

		// TODO: 写入redis
		err = w._redis.Set(context.Background(), w.Prefix+RedisKeyWechatJsApiTicket, string(body), time.Second*(result.ExpiresIn-10)).Err()
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func (w *Wechat) SnsGetAndSetAccessToken() (BodyAccessToken, error) {
	result := BodyAccessToken{}

	value, err := w._redis.Get(context.Background(), w.Prefix+RedisKeyWechatAccessToken).Result()
	if err != nil && err != redis.Nil {
		return BodyAccessToken{}, err
	}

	if len(value) > 5 {
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			w.log.Error("bt err ", zap.Error(err))
			return BodyAccessToken{}, err
		}
		w.log.Info("打印astoken", zap.Any("data", value), zap.Any("result", result))
		return result, nil
	} else {
		// 定期从微信拿取accesstoken
		body, err := w.CgiGetAccessToken()
		if err != nil {
			return result, err
		}

		json.Unmarshal(body, &result)

		if result.Errcode != 0 {
			// 获取出错
			return result, errors.New(result.Errmsg)
		}

		// TODO: 写入redis
		err = w._redis.Set(context.Background(), w.Prefix+RedisKeyWechatAccessToken, string(body), time.Second*(result.ExpiresIn-10)).Err()
		if err != nil {
			return result, err
		}

		w.log.Info("写入成功", zap.Any("result.AccessToken", result))
	}
	return result, nil
}
