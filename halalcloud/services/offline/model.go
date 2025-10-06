package offline

import "github.com/halalcloud/golang-sdk-lite/halalcloud/model"

type TaskMeta struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Type     int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status   int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Locked   bool   `protobuf:"varint,4,opt,name=locked,proto3" json:"locked,omitempty"`
	UpdateTs int64  `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty,string"`
	File     string `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	CreateTs int64  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty,string"`
	Url      string `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Size     int64  `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty,string"`
	Name     string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// string lock_by = 11;
	Code           int32  `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`
	Message        string `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
	Addon          string `protobuf:"bytes,14,opt,name=addon,proto3" json:"addon,omitempty"`
	Retries        int64  `protobuf:"varint,15,opt,name=retries,proto3" json:"retries,omitempty,string"`
	Progress       int64  `protobuf:"varint,16,opt,name=progress,proto3" json:"progress,omitempty,string"`
	BytesTotal     int64  `protobuf:"varint,17,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty,string"`
	BytesProcessed int64  `protobuf:"varint,18,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty,string"`
	Flag           int32  `protobuf:"varint,19,opt,name=flag,proto3" json:"flag,omitempty"` // int32 retry_status = 20;
}

type TaskFile struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	// string file_identity = 2;
	// string task_identity = 3;
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	UpdateTs int64  `protobuf:"varint,6,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty,string"`
	CreateTs int64  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty,string"`
	Name     string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Size     int64  `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty,string"`
	// string content_identity = 10;
	// int32 code = 11;
	// string message = 12;
	BytesTotal int64 `protobuf:"varint,13,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty,string"`
	Index      int32 `protobuf:"varint,15,opt,name=index,proto3" json:"index,omitempty"`
	Directory  bool  `protobuf:"varint,16,opt,name=directory,proto3" json:"directory,omitempty"`
}

type TaskParseRequest struct {
	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	File     string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	Identity string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Addon    string `protobuf:"bytes,4,opt,name=addon,proto3" json:"addon,omitempty"`
	Flag     int32  `protobuf:"varint,5,opt,name=flag,proto3" json:"flag,omitempty"`
	From     string `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
}

type TaskParseResponse struct {
	Meta      *TaskMeta   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	TaskFiles []*TaskFile `protobuf:"bytes,2,rep,name=task_files,json=taskFiles,proto3" json:"task_files,omitempty"`
}

type UserTask struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Type     int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status   int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// string user_identity = 4;
	UpdateTs       int64    `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty,string"`
	File           string   `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	CreateTs       int64    `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty,string"`
	Url            string   `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Size           int64    `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty,string"`
	Name           string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	TaskIdentity   string   `protobuf:"bytes,11,opt,name=task_identity,json=taskIdentity,proto3" json:"task_identity,omitempty"`
	Code           int32    `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`
	Message        string   `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
	Addon          string   `protobuf:"bytes,14,opt,name=addon,proto3" json:"addon,omitempty"`
	Progress       int64    `protobuf:"varint,16,opt,name=progress,proto3" json:"progress,omitempty,string"`
	BytesTotal     int64    `protobuf:"varint,17,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty,string"`
	BytesProcessed int64    `protobuf:"varint,18,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty,string"`
	Flag           int32    `protobuf:"varint,19,opt,name=flag,proto3" json:"flag,omitempty"`
	SavePath       string   `protobuf:"bytes,20,opt,name=save_path,json=savePath,proto3" json:"save_path,omitempty"`
	Callbacks      []string `protobuf:"bytes,21,rep,name=callbacks,proto3" json:"callbacks,omitempty"`
	IgnoreFiles    []string `protobuf:"bytes,22,rep,name=ignore_files,json=ignoreFiles,proto3" json:"ignore_files,omitempty"`
}

type OfflineTaskListRequest struct {
	ListInfo *model.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

type OfflineTaskListResponse struct {
	Tasks    []*UserTask            `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	ListInfo *model.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

type OfflineTaskDeleteRequest struct {
	Identity    []string `protobuf:"bytes,1,rep,name=identity,proto3" json:"identity,omitempty"`
	DeleteFiles bool     `protobuf:"varint,2,opt,name=delete_files,json=deleteFiles,proto3" json:"delete_files,omitempty"`
}

type OfflineTaskDeleteResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty,string"`
}
