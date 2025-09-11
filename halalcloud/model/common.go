package model

type ScanListRequest struct {
	Token   string         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Limit   string         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy []*OrderByInfo `protobuf:"bytes,3,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Version int32          `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

type OrderByInfo struct {
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Asc   bool   `protobuf:"varint,2,opt,name=asc,proto3" json:"asc,omitempty"`
}
