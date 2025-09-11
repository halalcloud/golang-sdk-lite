package user

import (
	"context"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
)

type UserService struct {
	client *apiclient.Client
}

func NewUserService(client *apiclient.Client) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) Get(ctx context.Context, req *User) (*User, error) {
	data := &User{}
	err := s.client.Post(ctx, "/v6/user/get", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *UserService) GetStatisticsAndQuota(ctx context.Context) (*UserStatisticsAndQuota, error) {
	req := &map[string]any{}
	data := &UserStatisticsAndQuota{}
	err := s.client.Post(ctx, "/v6/user/get_statistics_and_quota", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *UserService) UserCenterUri(ctx context.Context, req *UserCenterUriRequest) (*UserCenterUriResponse, error) {
	data := &UserCenterUriResponse{}
	err := s.client.Post(ctx, "/v6/user/user_center_uri", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
