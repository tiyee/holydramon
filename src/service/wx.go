package service

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tiyee/holydramon/src/engine"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const WxAccessTokenUrlFmt = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

const cacheKey = "access_token"

type AccessTokenValue struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type OriginalAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func getAccessToken() (string, error) {
	url := fmt.Sprintf(WxAccessTokenUrlFmt, engine.ImmutableConfig.Wx.AppId, engine.ImmutableConfig.Wx.AppSecret)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return "", err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	if data, err := ioutil.ReadAll(res.Body); err == nil {
		var accessToken OriginalAccessToken
		if err := json.Unmarshal(data, &accessToken); err == nil {
			cacheValue := AccessTokenValue{
				accessToken.AccessToken,
				time.Now().Unix() + 6400,
			}
			if value, err := json.Marshal(cacheValue); err == nil {
				if err := engine.ImmutableComponents.BigCache.Set(cacheKey, value); err != nil {
					log.Println(err)
				}
			}
			return accessToken.AccessToken, nil
		}
		{
			return accessToken.AccessToken, nil
		}
	}
	return "", errors.New("get access_token error")
}
func GetAccessToken() (string, error) {
	if entry, err := engine.ImmutableComponents.BigCache.Get(cacheKey); err == nil {
		var cacheValue AccessTokenValue
		if err := json.Unmarshal(entry, &cacheValue); err == nil {
			if cacheValue.Expire > time.Now().Unix() {
				log.Println("get access_token from cache")
				return cacheValue.Token, nil
			}
		}
	}
	return getAccessToken()
}
