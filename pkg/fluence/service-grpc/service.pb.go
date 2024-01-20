// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: pkg/fluence/service-grpc/service.proto

package service

import (
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

// GroupRequest for a group
type GroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GroupRequest) Reset() {
	*x = GroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRequest) ProtoMessage() {}

func (x *GroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRequest.ProtoReflect.Descriptor instead.
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_fluence_service_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *GroupRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

// GroupResponse
type GroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GroupResponse) Reset() {
	*x = GroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupResponse) ProtoMessage() {}

func (x *GroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupResponse.ProtoReflect.Descriptor instead.
func (*GroupResponse) Descriptor() ([]byte, []int) {
	return file_pkg_fluence_service_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *GroupResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResourceRequest) Reset() {
	*x = ResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRequest) ProtoMessage() {}

func (x *ResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRequest.ProtoReflect.Descriptor instead.
func (*ResourceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_fluence_service_grpc_service_proto_rawDescGZIP(), []int{2}
}

type ResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Graph string `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
}

func (x *ResourceResponse) Reset() {
	*x = ResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceResponse) ProtoMessage() {}

func (x *ResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_fluence_service_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceResponse.ProtoReflect.Descriptor instead.
func (*ResourceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_fluence_service_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceResponse) GetGraph() string {
	if x != nil {
		return x.Graph
	}
	return ""
}

var File_pkg_fluence_service_grpc_service_proto protoreflect.FileDescriptor

var file_pkg_fluence_service_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x24, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x37, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x32, 0xda, 0x01,
	0x0a, 0x15, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_fluence_service_grpc_service_proto_rawDescOnce sync.Once
	file_pkg_fluence_service_grpc_service_proto_rawDescData = file_pkg_fluence_service_grpc_service_proto_rawDesc
)

func file_pkg_fluence_service_grpc_service_proto_rawDescGZIP() []byte {
	file_pkg_fluence_service_grpc_service_proto_rawDescOnce.Do(func() {
		file_pkg_fluence_service_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_fluence_service_grpc_service_proto_rawDescData)
	})
	return file_pkg_fluence_service_grpc_service_proto_rawDescData
}

var file_pkg_fluence_service_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_fluence_service_grpc_service_proto_goTypes = []interface{}{
	(*GroupRequest)(nil),     // 0: service.GroupRequest
	(*GroupResponse)(nil),    // 1: service.GroupResponse
	(*ResourceRequest)(nil),  // 2: service.ResourceRequest
	(*ResourceResponse)(nil), // 3: service.ResourceResponse
}
var file_pkg_fluence_service_grpc_service_proto_depIdxs = []int32{
	0, // 0: service.ExternalPluginService.ListGroups:input_type -> service.GroupRequest
	0, // 1: service.ExternalPluginService.GetGroup:input_type -> service.GroupRequest
	2, // 2: service.ExternalPluginService.GetResources:input_type -> service.ResourceRequest
	1, // 3: service.ExternalPluginService.ListGroups:output_type -> service.GroupResponse
	1, // 4: service.ExternalPluginService.GetGroup:output_type -> service.GroupResponse
	3, // 5: service.ExternalPluginService.GetResources:output_type -> service.ResourceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_fluence_service_grpc_service_proto_init() }
func file_pkg_fluence_service_grpc_service_proto_init() {
	if File_pkg_fluence_service_grpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_fluence_service_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRequest); i {
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
		file_pkg_fluence_service_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupResponse); i {
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
		file_pkg_fluence_service_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRequest); i {
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
		file_pkg_fluence_service_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceResponse); i {
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
			RawDescriptor: file_pkg_fluence_service_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_fluence_service_grpc_service_proto_goTypes,
		DependencyIndexes: file_pkg_fluence_service_grpc_service_proto_depIdxs,
		MessageInfos:      file_pkg_fluence_service_grpc_service_proto_msgTypes,
	}.Build()
	File_pkg_fluence_service_grpc_service_proto = out.File
	file_pkg_fluence_service_grpc_service_proto_rawDesc = nil
	file_pkg_fluence_service_grpc_service_proto_goTypes = nil
	file_pkg_fluence_service_grpc_service_proto_depIdxs = nil
}
