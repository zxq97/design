// protoc --go_out=. --go-grpc_out=. fuxi/api/service/v1/fuxi.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.0
// source: fuxi/api/service/v1/fuxi.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetGenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid int64 `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
}

func (x *GetGenerateResponse) Reset() {
	*x = GetGenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fuxi_api_service_v1_fuxi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenerateResponse) ProtoMessage() {}

func (x *GetGenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fuxi_api_service_v1_fuxi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenerateResponse.ProtoReflect.Descriptor instead.
func (*GetGenerateResponse) Descriptor() ([]byte, []int) {
	return file_fuxi_api_service_v1_fuxi_proto_rawDescGZIP(), []int{0}
}

func (x *GetGenerateResponse) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

var File_fuxi_api_service_v1_fuxi_proto protoreflect.FileDescriptor

var file_fuxi_api_service_v1_fuxi_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x75, 0x78, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x78, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x32, 0x4a, 0x0a, 0x04, 0x46, 0x75,
	0x78, 0x69, 0x12, 0x42, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x66, 0x75, 0x78, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fuxi_api_service_v1_fuxi_proto_rawDescOnce sync.Once
	file_fuxi_api_service_v1_fuxi_proto_rawDescData = file_fuxi_api_service_v1_fuxi_proto_rawDesc
)

func file_fuxi_api_service_v1_fuxi_proto_rawDescGZIP() []byte {
	file_fuxi_api_service_v1_fuxi_proto_rawDescOnce.Do(func() {
		file_fuxi_api_service_v1_fuxi_proto_rawDescData = protoimpl.X.CompressGZIP(file_fuxi_api_service_v1_fuxi_proto_rawDescData)
	})
	return file_fuxi_api_service_v1_fuxi_proto_rawDescData
}

var file_fuxi_api_service_v1_fuxi_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fuxi_api_service_v1_fuxi_proto_goTypes = []interface{}{
	(*GetGenerateResponse)(nil), // 0: v1.GetGenerateResponse
	(*emptypb.Empty)(nil),       // 1: google.protobuf.Empty
}
var file_fuxi_api_service_v1_fuxi_proto_depIdxs = []int32{
	1, // 0: v1.Fuxi.GetGenerateID:input_type -> google.protobuf.Empty
	0, // 1: v1.Fuxi.GetGenerateID:output_type -> v1.GetGenerateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fuxi_api_service_v1_fuxi_proto_init() }
func file_fuxi_api_service_v1_fuxi_proto_init() {
	if File_fuxi_api_service_v1_fuxi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fuxi_api_service_v1_fuxi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGenerateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fuxi_api_service_v1_fuxi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fuxi_api_service_v1_fuxi_proto_goTypes,
		DependencyIndexes: file_fuxi_api_service_v1_fuxi_proto_depIdxs,
		MessageInfos:      file_fuxi_api_service_v1_fuxi_proto_msgTypes,
	}.Build()
	File_fuxi_api_service_v1_fuxi_proto = out.File
	file_fuxi_api_service_v1_fuxi_proto_rawDesc = nil
	file_fuxi_api_service_v1_fuxi_proto_goTypes = nil
	file_fuxi_api_service_v1_fuxi_proto_depIdxs = nil
}
