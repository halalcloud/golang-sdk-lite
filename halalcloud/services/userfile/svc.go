package userfile

import (
	"context"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
)

type UserFileService struct {
	client *apiclient.Client
}

func NewUserFileService(client *apiclient.Client) *UserFileService {
	return &UserFileService{
		client: client,
	}
}

func (s *UserFileService) List(ctx context.Context, req *FileListRequest) (*FileListResponse, error) {
	data := &FileListResponse{}
	err := s.client.Post(ctx, "/v6/userfile/list", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *UserFileService) Get(ctx context.Context, req *File) (*File, error) {
	data := &File{}
	err := s.client.Post(ctx, "/v6/userfile/get", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// rpc ParseFileSlice (File) returns (ParseFileSliceResponse) {
func (s *UserFileService) ParseFileSlice(ctx context.Context, req *File) (*ParseFileSliceResponse, error) {
	data := &ParseFileSliceResponse{}
	err := s.client.Post(ctx, "/v6/userfile/parse_file_slice", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *UserFileService) GetSliceDownloadAddress(ctx context.Context, req *SliceDownloadAddressRequest) (*SliceDownloadAddressResponse, error) {
	data := &SliceDownloadAddressResponse{}
	err := s.client.Post(ctx, "/v6/userfile/get_slice_download_address", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// rpc Create (File) returns (File) {
func (s *UserFileService) Create(ctx context.Context, req *File) (*File, error) {
	data := &File{}
	err := s.client.Post(ctx, "/v6/userfile/create", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// rpc Create (File) returns (File) {
func (s *UserFileService) CreateUploadTask(ctx context.Context, req *File) (*UploadTask, error) {
	data := &UploadTask{}
	err := s.client.Post(ctx, "/v6/userfile/create_upload_task", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
