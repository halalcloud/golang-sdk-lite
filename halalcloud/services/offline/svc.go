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
