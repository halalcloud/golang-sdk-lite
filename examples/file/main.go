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
	req := &userfile.File{
		// Name: "新建文件夹",
		Path: "/0v0000/新建文件夹/1.pdf",
		Size: "10240",
	}

	jsonData, _ := json.MarshalIndent(req, "", "  ")
	log.Println(string(jsonData))

	userResp, err := userFileSvc.CreateUploadTask(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create upload task: %v", err)
	}
	jsonData, _ = json.MarshalIndent(userResp, "", "  ")
	log.Println(string(jsonData))
}
