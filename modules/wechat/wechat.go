package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/changwei4869/wedding/model"
)

// 获取二维码
func GetQRCodeURL(w http.ResponseWriter, r *http.Request) {
	appID := "your_appid"
	redirectURI := "http://yourserver.com/api/wechat/callback"
	state := "some_random_state_string"
	qrCodeURL := fmt.Sprintf(
		"https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=%s#wechat_redirect",
		appID, redirectURI, state,
	)
	fmt.Fprintln(w, qrCodeURL)
}

// 微信回调处理
func WeChatCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	// 验证state的合法性

	// 将code传递给前端，或直接在后端获取access_token和用户信息
	fmt.Fprintf(w, "Received code: %s and state: %s\n", code, state)

	// 这里可以进一步处理：使用code获取access_token和用户信息
}

// 获取用户信息
func GetAccessToken(appID, appSecret, code string) (*model.WeChatAccessTokenResponse, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appID, appSecret, code,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request for access token failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var tokenResponse model.WeChatAccessTokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal access token response: %v", err)
	}
	if tokenResponse.ErrCode != 0 {
		return nil, fmt.Errorf("wechat API error: %v, %s", tokenResponse.ErrCode, tokenResponse.ErrMsg)
	}

	return &tokenResponse, nil
}

func GetUserInfo(accessToken, openID string) (*model.WeChatUserInfo, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		accessToken, openID,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request for user info failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var userInfo model.WeChatUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info response: %v", err)
	}
	if userInfo.ErrCode != 0 {
		return nil, fmt.Errorf("wechat API error: %v, %s", userInfo.ErrCode, userInfo.ErrMsg)
	}

	return &userInfo, nil
}
