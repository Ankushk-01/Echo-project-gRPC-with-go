// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.5
// source: Echo/echo.proto

package Echo

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

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EchoRequest string `protobuf:"bytes,1,opt,name=echoRequest,proto3" json:"echoRequest,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_Echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetEchoRequest() string {
	if x != nil {
		return x.EchoRequest
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EchoResponse string `protobuf:"bytes,1,opt,name=echoResponse,proto3" json:"echoResponse,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_Echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetEchoResponse() string {
	if x != nil {
		return x.EchoResponse
	}
	return ""
}

var File_Echo_echo_proto protoreflect.FileDescriptor

var file_Echo_echo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x45, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x3e, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x65, 0x63, 0x68, 0x6f, 0x2f, 0x45, 0x63, 0x68, 0x6f, 0x2f, 0x45, 0x63, 0x68, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Echo_echo_proto_rawDescOnce sync.Once
	file_Echo_echo_proto_rawDescData = file_Echo_echo_proto_rawDesc
)

func file_Echo_echo_proto_rawDescGZIP() []byte {
	file_Echo_echo_proto_rawDescOnce.Do(func() {
		file_Echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_Echo_echo_proto_rawDescData)
	})
	return file_Echo_echo_proto_rawDescData
}

var file_Echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Echo_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),  // 0: Echo.EchoRequest
	(*EchoResponse)(nil), // 1: Echo.EchoResponse
}
var file_Echo_echo_proto_depIdxs = []int32{
	0, // 0: Echo.EchoService.Echo:input_type -> Echo.EchoRequest
	1, // 1: Echo.EchoService.Echo:output_type -> Echo.EchoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Echo_echo_proto_init() }
func file_Echo_echo_proto_init() {
	if File_Echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_Echo_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_Echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Echo_echo_proto_goTypes,
		DependencyIndexes: file_Echo_echo_proto_depIdxs,
		MessageInfos:      file_Echo_echo_proto_msgTypes,
	}.Build()
	File_Echo_echo_proto = out.File
	file_Echo_echo_proto_rawDesc = nil
	file_Echo_echo_proto_goTypes = nil
	file_Echo_echo_proto_depIdxs = nil
}