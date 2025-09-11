package main

import (
	"context"
	"log"
	"time"

	"encoding/json"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/config"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/services/user"
)

func main() {
	cf := config.NewLocalFileConfigStore("/tmp/halal_config.json")
	clientID, _ := cf.GetConfig("client_id")
	clientSecret, _ := cf.GetConfig("client_secret")
	client := apiclient.NewClient(
		nil,
		"openapi.2dland.cn",
		clientID,
		clientSecret,
		cf,
		apiclient.WithTimeout(15*time.Second),
		// 初始token信息和刷新端点,
	)
	userSvc := user.NewUserService(client)
	userResp, err := userSvc.Get(context.Background(), &user.User{})
	if err != nil {
		log.Fatalf("Failed to get user info: %v", err)
	}
	jsonData, _ := json.MarshalIndent(userResp, "", "  ")
	log.Println(string(jsonData))
}
