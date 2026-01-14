package main

import (
	"context"
	"log"
	"time"

	"encoding/json"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/config"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/services/userfile"
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
	userFileSvc := userfile.NewUserFileService(client)
	req := &userfile.ListRecentUpdatedFilesRequest{
		Parent: &userfile.File{
			Path: "/",
		},
		StartTs: time.Now().Add(-7 * 24 * time.Hour).Unix(), // 最近7天更新的文件
	}

	jsonData, _ := json.MarshalIndent(req, "", "  ")
	log.Println(string(jsonData))

	userResp, err := userFileSvc.ListRecentUpdatedFiles(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to list user files: %v", err)
	}
	jsonData, _ = json.MarshalIndent(userResp, "", "  ")
	log.Println(string(jsonData))
}
