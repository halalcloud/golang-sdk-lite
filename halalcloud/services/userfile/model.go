package userfile

import "github.com/halalcloud/golang-sdk-lite/halalcloud/model"

type File struct {
	Identity        string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Parent          string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Name            string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Path            string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	MimeType        string `protobuf:"bytes,6,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Size            string `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Type            string `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	CreateTs        string `protobuf:"varint,9,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	UpdateTs        string `protobuf:"varint,10,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	DeleteTs        string `protobuf:"varint,11,opt,name=delete_ts,json=deleteTs,proto3" json:"delete_ts,omitempty"`
	Deleted         bool   `protobuf:"varint,12,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Dir             bool   `protobuf:"varint,13,opt,name=dir,proto3" json:"dir,omitempty"`
	Hidden          bool   `protobuf:"varint,14,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Locked          bool   `protobuf:"varint,15,opt,name=locked,proto3" json:"locked,omitempty"`
	Shared          bool   `protobuf:"varint,16,opt,name=shared,proto3" json:"shared,omitempty"`
	Starred         bool   `protobuf:"varint,17,opt,name=starred,proto3" json:"starred,omitempty"`
	Trashed         bool   `protobuf:"varint,18,opt,name=trashed,proto3" json:"trashed,omitempty"`
	LockedAt        string `protobuf:"varint,19,opt,name=locked_at,json=lockedAt,proto3" json:"locked_at,omitempty"`
	LockedBy        string `protobuf:"bytes,20,opt,name=locked_by,json=lockedBy,proto3" json:"locked_by,omitempty"`
	SharedAt        string `protobuf:"varint,21,opt,name=shared_at,json=sharedAt,proto3" json:"shared_at,omitempty"`
	Flag            string `protobuf:"varint,22,opt,name=flag,proto3" json:"flag,omitempty"`
	Unique          string `protobuf:"bytes,23,opt,name=unique,proto3" json:"unique,omitempty"`
	ContentIdentity string `protobuf:"bytes,24,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
	Label           string `protobuf:"varint,25,opt,name=label,proto3" json:"label,omitempty"`
	StoreType       string `protobuf:"varint,26,opt,name=store_type,json=storeType,proto3" json:"store_type,omitempty"`
	Version         string `protobuf:"varint,27,opt,name=version,proto3" json:"version,omitempty"`
	Files           string `protobuf:"varint,28,opt,name=files,proto3" json:"files,omitempty"`
	Direcotries     string `protobuf:"varint,29,opt,name=direcotries,proto3" json:"direcotries,omitempty"`
	Nodes           string `protobuf:"varint,30,opt,name=nodes,proto3" json:"nodes,omitempty"`
	SortName        string `protobuf:"bytes,31,opt,name=sort_name,json=sortName,proto3" json:"sort_name,omitempty"`
}

type FileListRequest struct {
	Parent   *File                  `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Filter   *File                  `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	ListInfo *model.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

type FileListResponse struct {
	Files    []*File                `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	ListInfo *model.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

type ParseFileSliceResponse struct {
	ContentIdentity      string               `protobuf:"bytes,1,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
	MetaNodes            []string             `protobuf:"bytes,2,rep,name=meta_nodes,json=metaNodes,proto3" json:"meta_nodes,omitempty"`
	RawNodes             []string             `protobuf:"bytes,3,rep,name=raw_nodes,json=rawNodes,proto3" json:"raw_nodes,omitempty"`
	FileSize             string               `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Sizes                []*SliceSize         `protobuf:"bytes,5,rep,name=sizes,proto3" json:"sizes,omitempty"`
	Sha1                 string               `protobuf:"bytes,6,opt,name=sha1,proto3" json:"sha1,omitempty"`
	WcsEtag              string               `protobuf:"bytes,7,opt,name=wcs_etag,json=wcsEtag,proto3" json:"wcs_etag,omitempty"`
	Name                 string               `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string               `protobuf:"bytes,9,opt,name=path,proto3" json:"path,omitempty"`
	StoreType            string               `protobuf:"varint,10,opt,name=store_type,json=storeType,proto3" json:"store_type,omitempty"`
	RawDownloadAddresses []*SliceDownloadInfo `protobuf:"bytes,11,rep,name=raw_download_addresses,json=rawDownloadAddresses,proto3" json:"raw_download_addresses,omitempty"`
}

type SliceSize struct {
	StartIndex string `protobuf:"varint,1,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	EndIndex   string `protobuf:"varint,2,opt,name=end_index,json=endIndex,proto3" json:"end_index,omitempty"`
	Size       string `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

type SliceDownloadInfo struct {
	Identity        string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	DownloadAddress string `protobuf:"bytes,2,opt,name=download_address,json=downloadAddress,proto3" json:"download_address,omitempty"`
	DownloadToken   string `protobuf:"bytes,3,opt,name=download_token,json=downloadToken,proto3" json:"download_token,omitempty"`
	Encrypt         int32  `protobuf:"varint,4,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	StoreType       string `protobuf:"varint,5,opt,name=store_type,json=storeType,proto3" json:"store_type,omitempty"`
}

type SliceDownloadAddressRequest struct {
	Identity []string `protobuf:"bytes,1,rep,name=identity,proto3" json:"identity,omitempty"`
	Version  int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Filename string   `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Flag     string   `protobuf:"varint,4,opt,name=flag,proto3" json:"flag,omitempty"`
	Parse    bool     `protobuf:"varint,5,opt,name=parse,proto3" json:"parse,omitempty"`
}

type SliceDownloadAddressResponse struct {
	Addresses      []*SliceDownloadInfo `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	ExpireAt       string               `protobuf:"varint,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	RequestAddress string               `protobuf:"bytes,3,opt,name=request_address,json=requestAddress,proto3" json:"request_address,omitempty"`
	Version        int32                `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

type UploadTask struct {
	Created       bool   `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	UploadAddress string `protobuf:"bytes,2,opt,name=upload_address,json=uploadAddress,proto3" json:"upload_address,omitempty"`
	Task          string `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	UploadToken   string `protobuf:"bytes,4,opt,name=upload_token,json=uploadToken,proto3" json:"upload_token,omitempty"`
	UploadKey     string `protobuf:"bytes,5,opt,name=upload_key,json=uploadKey,proto3" json:"upload_key,omitempty"`
	AccessKey     string `protobuf:"bytes,10,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey     string `protobuf:"bytes,11,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Bucket        string `protobuf:"bytes,12,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Region        string `protobuf:"bytes,13,opt,name=region,proto3" json:"region,omitempty"`
	Endpoint      string `protobuf:"bytes,14,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Key           string `protobuf:"bytes,15,opt,name=key,proto3" json:"key,omitempty"`
	BlockSize     string `protobuf:"varint,16,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	BlockHashType string `protobuf:"varint,17,opt,name=block_hash_type,json=blockHashType,proto3" json:"block_hash_type,omitempty"`
	BlockVersion  int32  `protobuf:"varint,18,opt,name=block_version,json=blockVersion,proto3" json:"block_version,omitempty"`
	BlockCodec    string `protobuf:"varint,19,opt,name=block_codec,json=blockCodec,proto3" json:"block_codec,omitempty"`
}

type BatchOperationResponse struct {
	Task     string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Status   int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Affected string `protobuf:"varint,3,opt,name=affected,proto3" json:"affected,omitempty"`
	CreateTs string `protobuf:"varint,4,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Progress string `protobuf:"varint,5,opt,name=progress,proto3" json:"progress,omitempty"`
}

type BatchOperationRequest struct {
	Source    []*File `protobuf:"bytes,1,rep,name=source,proto3" json:"source,omitempty"`
	Dest      *File   `protobuf:"bytes,2,opt,name=dest,proto3" json:"dest,omitempty"`
	Operation int32   `protobuf:"varint,3,opt,name=operation,proto3" json:"operation,omitempty"`
	Flag      int32   `protobuf:"varint,4,opt,name=flag,proto3" json:"flag,omitempty"`
}

type FileDownloadAddressResponse struct {
	Addresses      []*SliceDownloadInfo `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	ExpireAt       string               `protobuf:"varint,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	RequestAddress string               `protobuf:"bytes,3,opt,name=request_address,json=requestAddress,proto3" json:"request_address,omitempty"`
	Version        int32                `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Encrypt        int32                `protobuf:"varint,5,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	StoreType      string               `protobuf:"varint,6,opt,name=store_type,json=storeType,proto3" json:"store_type,omitempty"`
}
