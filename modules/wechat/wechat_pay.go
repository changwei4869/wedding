package wechat

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/changwei4869/wedding/model"
)

// 生成支付二维码

// GenerateNonceStr generates a random nonce string
func GenerateNonceStr(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// SignWeChatRequest signs the WeChat request
func SignWeChatRequest(params map[string]string, apiKey string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var signStrings string
	for _, k := range keys {
		signStrings += fmt.Sprintf("%s=%s&", k, params[k])
	}
	signStrings += "key=" + apiKey

	hasher := md5.New()
	hasher.Write([]byte(signStrings))
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}

// CreateWeChatOrder creates a WeChat order and returns the code URL for QR generation
func CreateWeChatOrder(config model.WeChatPayConfig) (string, error) {
	nonceStr := GenerateNonceStr(32)
	params := map[string]string{
		"appid":            config.AppID,
		"mch_id":           config.MchID,
		"nonce_str":        nonceStr,
		"body":             "Test Product",
		"out_trade_no":     fmt.Sprintf("%d", time.Now().Unix()),
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
		"notify_url":       config.NotifyURL,
		"trade_type":       "NATIVE",
	}
	sign := SignWeChatRequest(params, config.APIKey)
	orderRequest := model.WeChatOrderRequest{
		AppID:          config.AppID,
		MchID:          config.MchID,
		NonceStr:       nonceStr,
		Sign:           sign,
		Body:           "Test Product",
		OutTradeNo:     params["out_trade_no"],
		TotalFee:       1,
		SpbillCreateIP: "127.0.0.1",
		NotifyURL:      config.NotifyURL,
		TradeType:      "NATIVE",
	}

	output, err := xml.Marshal(orderRequest)
	if err != nil {
		return "", err
	}

	response, err := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", "application/xml; charset=utf-8", strings.NewReader(string(output)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var orderResponse model.WeChatOrderResponse
	err = xml.Unmarshal(body, &orderResponse)
	if err != nil {
		return "", err
	}
	if orderResponse.ReturnCode != "SUCCESS" || orderResponse.ResultCode != "SUCCESS" {
		return "", fmt.Errorf("error from WeChat: %s", orderResponse.ReturnMsg)
	}

	return orderResponse.CodeURL, nil
}

// 处理支付结果
func HandleWeChatPaymentNotification(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var notification model.WeChatOrderResponse
	if err := xml.Unmarshal(body, &notification); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Verify the request is from WeChat by checking the signature, etc.

	// Process the notification, update order status in the database, etc.

	// Respond to WeChat to acknowledge receipt of the notification
	response := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}
	respXML, _ := xml.Marshal(response)
	w.Header().Set("Content-Type", "application/xml")
	w.Write(respXML)
}