// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: cilium/api/network_filter.proto

package cilium

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NetworkFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the proxylib to be opened
	Proxylib string `protobuf:"bytes,1,opt,name=proxylib,proto3" json:"proxylib,omitempty"`
	// Transparent set of parameters provided for proxylib initialization
	ProxylibParams map[string]string `protobuf:"bytes,2,rep,name=proxylib_params,json=proxylibParams,proto3" json:"proxylib_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Path to the unix domain socket for the cilium access log.
	AccessLogPath string `protobuf:"bytes,5,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
}

func (x *NetworkFilter) Reset() {
	*x = NetworkFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_network_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkFilter) ProtoMessage() {}

func (x *NetworkFilter) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_network_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkFilter.ProtoReflect.Descriptor instead.
func (*NetworkFilter) Descriptor() ([]byte, []int) {
	return file_cilium_api_network_filter_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkFilter) GetProxylib() string {
	if x != nil {
		return x.Proxylib
	}
	return ""
}

func (x *NetworkFilter) GetProxylibParams() map[string]string {
	if x != nil {
		return x.ProxylibParams
	}
	return nil
}

func (x *NetworkFilter) GetAccessLogPath() string {
	if x != nil {
		return x.AccessLogPath
	}
	return ""
}

var File_cilium_api_network_filter_proto protoreflect.FileDescriptor

var file_cilium_api_network_filter_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0xea, 0x01, 0x0a, 0x0d, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x6c, 0x69, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x6c, 0x69, 0x62, 0x12, 0x52, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x6c, 0x69, 0x62, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x6c, 0x69, 0x62,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x6c, 0x69, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x50,
	0x61, 0x74, 0x68, 0x1a, 0x41, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x6c, 0x69, 0x62, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x3b,
	0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cilium_api_network_filter_proto_rawDescOnce sync.Once
	file_cilium_api_network_filter_proto_rawDescData = file_cilium_api_network_filter_proto_rawDesc
)

func file_cilium_api_network_filter_proto_rawDescGZIP() []byte {
	file_cilium_api_network_filter_proto_rawDescOnce.Do(func() {
		file_cilium_api_network_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_cilium_api_network_filter_proto_rawDescData)
	})
	return file_cilium_api_network_filter_proto_rawDescData
}

var file_cilium_api_network_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cilium_api_network_filter_proto_goTypes = []interface{}{
	(*NetworkFilter)(nil), // 0: cilium.NetworkFilter
	nil,                   // 1: cilium.NetworkFilter.ProxylibParamsEntry
}
var file_cilium_api_network_filter_proto_depIdxs = []int32{
	1, // 0: cilium.NetworkFilter.proxylib_params:type_name -> cilium.NetworkFilter.ProxylibParamsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cilium_api_network_filter_proto_init() }
func file_cilium_api_network_filter_proto_init() {
	if File_cilium_api_network_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cilium_api_network_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkFilter); i {
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
			RawDescriptor: file_cilium_api_network_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cilium_api_network_filter_proto_goTypes,
		DependencyIndexes: file_cilium_api_network_filter_proto_depIdxs,
		MessageInfos:      file_cilium_api_network_filter_proto_msgTypes,
	}.Build()
	File_cilium_api_network_filter_proto = out.File
	file_cilium_api_network_filter_proto_rawDesc = nil
	file_cilium_api_network_filter_proto_goTypes = nil
	file_cilium_api_network_filter_proto_depIdxs = nil
}
