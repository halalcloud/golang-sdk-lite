package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
)

func main() {
	// 创建带有token认证的客户端
	client := apiclient.NewClient(
		"openapi.2dland.cn",
		"05597f894253300bdb8a96706f9626b1ade5fd93bd24a7b6891ef427dfff665d",
		"AK_5f02f0889301fd7be1ac972c11bf3e7d",
		"AppSecretForTest_KSMWNDBAJ2hs__AS",
		apiclient.WithTimeout(15*time.Second),
		// 初始token信息和刷新端点,
	)

	// 调用需要鉴权的API
	data := &json.RawMessage{}
	err := client.Post(context.Background(), "/v6/debug/sign", map[string]string{"a": "b"}, []byte(`{"name":"test"}`), data)
	if err != nil {
		log.Fatalf("获取用户资料失败: %v", err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("序列化响应数据失败: %v", err)
	}

	fmt.Println(string(jsonData))

	// 调用另一个API，此时token已经是有效的

}
