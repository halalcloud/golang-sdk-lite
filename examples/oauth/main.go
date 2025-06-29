package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/config"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/services/oauth"
)

func main() {
	// 创建带有token认证的客户端
	client := apiclient.NewClient(
		"openapi.2dland.cn",
		//"05597f894253300bdb8a96706f9626b1ade5fd93bd24a7b6891ef427dfff665d",
		"AK_5f02f0889301fd7be1ac972c11bf3e7d",
		"AppSecretForTest_KSMWNDBAJ2hs__AS",
		config.NewMapConfigStore(),
		apiclient.WithTimeout(15*time.Second),
		// 初始token信息和刷新端点,
	)
	oauthSvc := oauth.NewOAuthService(client)
	resp, err := oauthSvc.GetToken(context.Background(), &oauth.TokenRequest{
		Code:         "0197af06482573aa83dd2f59a9d94ab7",
		CodeVerifier: "pUoVClMY6SedhDbSpLubVUdVnX1p4qyE",
	})
	if err != nil {
		log.Fatalf("获取授权失败: %v", err)
	}

	jsonData, _ := json.Marshal(resp)
	log.Printf("获取授权成功: %s", string(jsonData))

	// 调用另一个API，此时token已经是有效的

}
