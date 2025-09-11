package main

import (
	"context"
	"log"
	"time"

	"encoding/json"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/config"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/services/oauth"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/services/user"
)

func main() {
	// 创建带有token认证的客户端
	cf := config.NewLocalFileConfigStore("/tmp/halal_config.json")
	clientID, _ := cf.GetConfig("client_id")
	clientSecret, _ := cf.GetConfig("client_secret")
	client := apiclient.NewClient(
		nil,
		"openapi.2dland.cn",
		//"05597f894253300bdb8a96706f9626b1ade5fd93bd24a7b6891ef427dfff665d",
		clientID,
		clientSecret,
		config.NewMapConfigStore(),
		apiclient.WithTimeout(15*time.Second),
		// 初始token信息和刷新端点,
	)
	oauthSvc := oauth.NewOAuthService(client)
	resp, err := oauthSvc.DeviceCodeAuthorize(context.Background(), &oauth.AuthorizeRequest{
		ClientId: "AK_5f02f0889301fd7be1ac972c11bf3e7d",
		Device:   "sdk-test-device/1.0",
	})
	if err != nil {
		log.Fatalf("Failed to get device code: %v", err)
	}
	log.Printf("Device Code: %s, User Code: %s, Verification URI: %s, Expires In: %d seconds",
		resp.DeviceCode, resp.UserCode, resp.VerificationUri, resp.ExpiresIn)
	for {
		if resp.Interval <= 0 {
			resp.Interval = 5
		}
		time.Sleep(time.Duration(resp.Interval) * time.Second)
		osp, err := oauthSvc.GetDeviceCodeState(context.Background(), &oauth.DeviceCodeAuthorizeState{
			// UserCode:   resp.UserCode,
			DeviceCode: resp.DeviceCode,
		})
		if err != nil {
			log.Printf("Error checking device code state: %v", err)
			return
		}
		log.Printf("Device Code State: %s UserCode: %s, DeviceCode: %s, ExpiresIn: %d seconds, Login: %t", osp.Status,
			osp.UserCode, osp.DeviceCode, osp.ExpiresIn, osp.Login)
		if osp.Status == "AUTHORIZATION_SUCCESS" {
			log.Println("User has logged in successfully, proceeding to get token...")
			accessToken := osp.AccessToken
			refreshToken := osp.RefreshToken
			log.Printf("Access Token: %s, Refresh Token: %s", accessToken, refreshToken)
			client.SetToken(accessToken, refreshToken, osp.ExpiresIn)
			break
		}
	}

	// 调用另一个API，此时token已经是有效的
	userSvc := user.NewUserService(client)
	userResp, err := userSvc.UserCenterUri(context.Background(), &user.UserCenterUriRequest{})
	if err != nil {
		log.Fatalf("Failed to get user info: %v", err)
	}
	jsonData, _ := json.MarshalIndent(userResp, "", "  ")
	log.Println(string(jsonData))
	log.Println("OAuth flow completed successfully.")
	log.Println("You can now use the access token to make authorized requests to the API.")

}
