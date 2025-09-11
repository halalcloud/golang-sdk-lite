package user

type User struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity for user, unique in system
	Type     int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`        // 1: user, 2: admin, 3: super admin
	Status   int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`    // 1: normal, 2: disabled, 3: deleted
	UpdateTs string `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Name     string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Addon    string `protobuf:"bytes,8,opt,name=addon,proto3" json:"addon,omitempty"`
	CreateTs string `protobuf:"varint,9,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Hash     string `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Icon     string `protobuf:"bytes,11,opt,name=icon,proto3" json:"icon,omitempty"`
}

type UserStatisticsAndQuota struct {
	DiskStatisticsQuota        *DiskStatisticsAndQuota        `protobuf:"bytes,1,opt,name=disk_statistics_quota,json=diskStatisticsQuota,proto3" json:"disk_statistics_quota,omitempty"`
	TrafficStatisticsQuota     *TrafficStatisticsAndQuota     `protobuf:"bytes,2,opt,name=traffic_statistics_quota,json=trafficStatisticsQuota,proto3" json:"traffic_statistics_quota,omitempty"`
	OfflineTaskStatisticsQuota *OfflineTaskStatisticsAndQuota `protobuf:"bytes,3,opt,name=offline_task_statistics_quota,json=offlineTaskStatisticsQuota,proto3" json:"offline_task_statistics_quota,omitempty"`
}

type DiskStatisticsAndQuota struct {
	BytesQuota     string `protobuf:"varint,1,opt,name=bytes_quota,json=bytesQuota,proto3" json:"bytes_quota,omitempty"`
	BytesUsed      string `protobuf:"varint,2,opt,name=bytes_used,json=bytesUsed,proto3" json:"bytes_used,omitempty"`
	BytesFree      string `protobuf:"varint,3,opt,name=bytes_free,json=bytesFree,proto3" json:"bytes_free,omitempty"`
	FilesQuota     string `protobuf:"varint,4,opt,name=files_quota,json=filesQuota,proto3" json:"files_quota,omitempty"`
	FilesUsed      string `protobuf:"varint,5,opt,name=files_used,json=filesUsed,proto3" json:"files_used,omitempty"`
	BytesTrashUsed string `protobuf:"varint,6,opt,name=bytes_trash_used,json=bytesTrashUsed,proto3" json:"bytes_trash_used,omitempty"`
}

type TrafficStatisticsAndQuota struct {
	DailyBytesDownloadQuota string `protobuf:"varint,1,opt,name=daily_bytes_download_quota,json=dailyBytesDownloadQuota,proto3" json:"daily_bytes_download_quota,omitempty"`
	DailyBytesUploadQuota   string `protobuf:"varint,2,opt,name=daily_bytes_upload_quota,json=dailyBytesUploadQuota,proto3" json:"daily_bytes_upload_quota,omitempty"`
	BytesDownloadedToday    string `protobuf:"varint,3,opt,name=bytes_downloaded_today,json=bytesDownloadedToday,proto3" json:"bytes_downloaded_today,omitempty"`
	BytesUploadedToday      string `protobuf:"varint,4,opt,name=bytes_uploaded_today,json=bytesUploadedToday,proto3" json:"bytes_uploaded_today,omitempty"`
}

type OfflineTaskStatisticsAndQuota struct {
	DailyTasksQuota    string `protobuf:"varint,1,opt,name=daily_tasks_quota,json=dailyTasksQuota,proto3" json:"daily_tasks_quota,omitempty"`
	TasksCommitedToday string `protobuf:"varint,2,opt,name=tasks_commited_today,json=tasksCommitedToday,proto3" json:"tasks_commited_today,omitempty"`
	DailyBytesQuota    string `protobuf:"varint,3,opt,name=daily_bytes_quota,json=dailyBytesQuota,proto3" json:"daily_bytes_quota,omitempty"`
	BytesCommitedToday string `protobuf:"varint,4,opt,name=bytes_commited_today,json=bytesCommitedToday,proto3" json:"bytes_commited_today,omitempty"`
}

type UserCenterUriRequest struct {
	DisableSsl bool   `protobuf:"varint,1,opt,name=disable_ssl,json=disableSsl,proto3" json:"disable_ssl,omitempty"` // disable ssl, default is false
	Route      string `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`                              // route for user center, default is "/"
}

type UserCenterUriResponse struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"` // user center uri
}
