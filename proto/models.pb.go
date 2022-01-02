// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: models.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Address  string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Company  string                 `protobuf:"bytes,4,opt,name=company,proto3" json:"company,omitempty"`
	Score    int32                  `protobuf:"varint,5,opt,name=score,proto3" json:"score,omitempty"`
	Bank     string                 `protobuf:"bytes,6,opt,name=bank,proto3" json:"bank,omitempty"`
	Balance  int32                  `protobuf:"varint,7,opt,name=balance,proto3" json:"balance,omitempty"`
	Updated  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *UserDetail) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserDetail) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDetail) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserDetail) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *UserDetail) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UserDetail) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *UserDetail) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *UserDetail) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type TeamDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    int32                  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	TeamName  string                 `protobuf:"bytes,2,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	MemberNum int32                  `protobuf:"varint,3,opt,name=member_num,json=memberNum,proto3" json:"member_num,omitempty"`
	Business  string                 `protobuf:"bytes,4,opt,name=business,proto3" json:"business,omitempty"`
	Created   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *TeamDetail) Reset() {
	*x = TeamDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamDetail) ProtoMessage() {}

func (x *TeamDetail) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamDetail.ProtoReflect.Descriptor instead.
func (*TeamDetail) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *TeamDetail) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamDetail) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *TeamDetail) GetMemberNum() int32 {
	if x != nil {
		return x.MemberNum
	}
	return 0
}

func (x *TeamDetail) GetBusiness() string {
	if x != nil {
		return x.Business
	}
	return ""
}

func (x *TeamDetail) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x0a, 0x54, 0x65,
	0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72,
	0x69, 0x6b, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_proto_goTypes = []interface{}{
	(*UserDetail)(nil),            // 0: proto.UserDetail
	(*TeamDetail)(nil),            // 1: proto.TeamDetail
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_models_proto_depIdxs = []int32{
	2, // 0: proto.UserDetail.updated:type_name -> google.protobuf.Timestamp
	2, // 1: proto.TeamDetail.created:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamDetail); i {
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
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}