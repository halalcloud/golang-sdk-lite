package offline

import (
	"context"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
)

type OfflineTaskService struct {
	client *apiclient.Client
}

func NewOfflineTaskService(client *apiclient.Client) *OfflineTaskService {
	return &OfflineTaskService{
		client: client,
	}
}

// Parse 解析离线任务
func (s *OfflineTaskService) Parse(ctx context.Context, req *TaskParseRequest) (*TaskParseResponse, error) {
	data := &TaskParseResponse{}
	err := s.client.Post(ctx, "/v6/offline_task/parse", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// - callbacks 生成流程
// 1. 整理需要回调的URL
// 2. 创建JSON:  {"url": "xxxxxxxxx"} （目前只支持HTTP/HTTPS）
// 3. 序列化之后将其BASE64 （使用标准base64编码，不要使用URL安全的base64编码）
// 4. 将序列化之后的字符串放到callbacks字段中
// 5. 创建离线任务
// rpc Add (UserTask) returns (UserTask) {

func (s *OfflineTaskService) Add(ctx context.Context, req *UserTask) (*UserTask, error) {
	data := &UserTask{}
	err := s.client.Post(ctx, "/v6/offline_task/add", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *OfflineTaskService) List(ctx context.Context, req *OfflineTaskListRequest) (*OfflineTaskListResponse, error) {
	data := &OfflineTaskListResponse{}
	err := s.client.Post(ctx, "/v6/offline_task/list", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// rpc Delete (OfflineTaskDeleteRequest) returns (OfflineTaskDeleteResponse) {

func (s *OfflineTaskService) Delete(ctx context.Context, req *OfflineTaskDeleteRequest) (*OfflineTaskDeleteResponse, error) {
	data := &OfflineTaskDeleteResponse{}
	err := s.client.Post(ctx, "/v6/offline_task/delete", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
