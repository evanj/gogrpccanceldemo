// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: proto/echo.proto

package echopb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerAction int32

const (
	ServerAction_UNSPECIFIED                      ServerAction = 0
	ServerAction_RETURN_CONTEXT_DEADLINE_EXCEEDED ServerAction = 1
	ServerAction_RETURN_CONTEXT_CANCELED          ServerAction = 2
)

// Enum value maps for ServerAction.
var (
	ServerAction_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "RETURN_CONTEXT_DEADLINE_EXCEEDED",
		2: "RETURN_CONTEXT_CANCELED",
	}
	ServerAction_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"RETURN_CONTEXT_DEADLINE_EXCEEDED": 1,
		"RETURN_CONTEXT_CANCELED":          2,
	}
)

func (x ServerAction) Enum() *ServerAction {
	p := new(ServerAction)
	*p = x
	return p
}

func (x ServerAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerAction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_echo_proto_enumTypes[0].Descriptor()
}

func (ServerAction) Type() protoreflect.EnumType {
	return &file_proto_echo_proto_enumTypes[0]
}

func (x ServerAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerAction.Descriptor instead.
func (ServerAction) EnumDescriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{0}
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input       string               `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	ServerSleep *durationpb.Duration `protobuf:"bytes,2,opt,name=server_sleep,json=serverSleep,proto3" json:"server_sleep,omitempty"`
	Action      ServerAction         `protobuf:"varint,3,opt,name=action,proto3,enum=echopb.ServerAction" json:"action,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[0]
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
	return file_proto_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *EchoRequest) GetServerSleep() *durationpb.Duration {
	if x != nil {
		return x.ServerSleep
	}
	return nil
}

func (x *EchoRequest) GetAction() ServerAction {
	if x != nil {
		return x.Action
	}
	return ServerAction_UNSPECIFIED
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[1]
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
	return file_proto_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_proto_echo_proto protoreflect.FileDescriptor

var file_proto_echo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x12, 0x2c,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0c,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2a, 0x62, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52,
	0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0x3b, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f,
	0x12, 0x33, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x13, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70,
	0x62, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_echo_proto_rawDescOnce sync.Once
	file_proto_echo_proto_rawDescData = file_proto_echo_proto_rawDesc
)

func file_proto_echo_proto_rawDescGZIP() []byte {
	file_proto_echo_proto_rawDescOnce.Do(func() {
		file_proto_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_echo_proto_rawDescData)
	})
	return file_proto_echo_proto_rawDescData
}

var file_proto_echo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_echo_proto_goTypes = []interface{}{
	(ServerAction)(0),           // 0: echopb.ServerAction
	(*EchoRequest)(nil),         // 1: echopb.EchoRequest
	(*EchoResponse)(nil),        // 2: echopb.EchoResponse
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_proto_echo_proto_depIdxs = []int32{
	3, // 0: echopb.EchoRequest.server_sleep:type_name -> google.protobuf.Duration
	0, // 1: echopb.EchoRequest.action:type_name -> echopb.ServerAction
	1, // 2: echopb.Echo.Echo:input_type -> echopb.EchoRequest
	2, // 3: echopb.Echo.Echo:output_type -> echopb.EchoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_echo_proto_init() }
func file_proto_echo_proto_init() {
	if File_proto_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_echo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_echo_proto_goTypes,
		DependencyIndexes: file_proto_echo_proto_depIdxs,
		EnumInfos:         file_proto_echo_proto_enumTypes,
		MessageInfos:      file_proto_echo_proto_msgTypes,
	}.Build()
	File_proto_echo_proto = out.File
	file_proto_echo_proto_rawDesc = nil
	file_proto_echo_proto_goTypes = nil
	file_proto_echo_proto_depIdxs = nil
}