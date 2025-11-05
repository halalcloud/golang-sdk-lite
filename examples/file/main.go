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
	req := &userfile.FileListRequest{
		// Name: "新建文件夹",
		Parent: &userfile.File{
			Path: "/0v0000",
		},
	}

	jsonData, _ := json.MarshalIndent(req, "", "  ")
	log.Println(string(jsonData))

	userResp, err := userFileSvc.GetDirectDownloadAddress(context.Background(), &userfile.DirectDownloadRequest{
		Path: "/0v0000/zh-cn_windows_server_2025_updated_june_2025_x64_dvd_e743555f.iso",
	})
	if err != nil {
		log.Fatalf("Failed to list user files: %v", err)
	}
	jsonData, _ = json.MarshalIndent(userResp, "", "  ")
	log.Println(string(jsonData))
}
