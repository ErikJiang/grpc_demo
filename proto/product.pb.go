// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/product.proto

package proto

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

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserDesc string `protobuf:"bytes,3,opt,name=user_desc,json=userDesc,proto3" json:"user_desc,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserResponse) GetUserDesc() string {
	if x != nil {
		return x.UserDesc
	}
	return ""
}

var File_proto_product_proto protoreflect.FileDescriptor

var file_proto_product_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x32, 0x45, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x69,
	0x6b, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_proto_rawDescOnce sync.Once
	file_proto_product_proto_rawDescData = file_proto_product_proto_rawDesc
)

func file_proto_product_proto_rawDescGZIP() []byte {
	file_proto_product_proto_rawDescOnce.Do(func() {
		file_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_proto_rawDescData)
	})
	return file_proto_product_proto_rawDescData
}

var file_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_product_proto_goTypes = []interface{}{
	(*UserRequest)(nil),  // 0: proto.UserRequest
	(*UserResponse)(nil), // 1: proto.UserResponse
}
var file_proto_product_proto_depIdxs = []int32{
	0, // 0: proto.UserService.GetUserInfo:input_type -> proto.UserRequest
	1, // 1: proto.UserService.GetUserInfo:output_type -> proto.UserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_product_proto_init() }
func file_proto_product_proto_init() {
	if File_proto_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_proto_goTypes,
		DependencyIndexes: file_proto_product_proto_depIdxs,
		MessageInfos:      file_proto_product_proto_msgTypes,
	}.Build()
	File_proto_product_proto = out.File
	file_proto_product_proto_rawDesc = nil
	file_proto_product_proto_goTypes = nil
	file_proto_product_proto_depIdxs = nil
}