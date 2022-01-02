// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: team.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TeamLevel int32

const (
	TeamLevel_High   TeamLevel = 0
	TeamLevel_Middle TeamLevel = 1
	TeamLevel_Low    TeamLevel = 2
)

// Enum value maps for TeamLevel.
var (
	TeamLevel_name = map[int32]string{
		0: "High",
		1: "Middle",
		2: "Low",
	}
	TeamLevel_value = map[string]int32{
		"High":   0,
		"Middle": 1,
		"Low":    2,
	}
)

func (x TeamLevel) Enum() *TeamLevel {
	p := new(TeamLevel)
	*p = x
	return p
}

func (x TeamLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TeamLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_team_proto_enumTypes[0].Descriptor()
}

func (TeamLevel) Type() protoreflect.EnumType {
	return &file_team_proto_enumTypes[0]
}

func (x TeamLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TeamLevel.Descriptor instead.
func (TeamLevel) EnumDescriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{0}
}

type TeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int32 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *TeamRequest) Reset() {
	*x = TeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamRequest) ProtoMessage() {}

func (x *TeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamRequest.ProtoReflect.Descriptor instead.
func (*TeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{0}
}

func (x *TeamRequest) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type TeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId   int32     `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	TeamName string    `protobuf:"bytes,2,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	TeamDesc string    `protobuf:"bytes,3,opt,name=team_desc,json=teamDesc,proto3" json:"team_desc,omitempty"`
	Level    TeamLevel `protobuf:"varint,4,opt,name=level,proto3,enum=proto.TeamLevel" json:"level,omitempty"`
}

func (x *TeamResponse) Reset() {
	*x = TeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamResponse) ProtoMessage() {}

func (x *TeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamResponse.ProtoReflect.Descriptor instead.
func (*TeamResponse) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{1}
}

func (x *TeamResponse) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamResponse) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *TeamResponse) GetTeamDesc() string {
	if x != nil {
		return x.TeamDesc
	}
	return ""
}

func (x *TeamResponse) GetLevel() TeamLevel {
	if x != nil {
		return x.Level
	}
	return TeamLevel_High
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *TeamDetail `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetTeam() *TeamDetail {
	if x != nil {
		return x.Team
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{3}
}

func (x *AddResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_team_proto protoreflect.FileDescriptor

var file_team_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x26, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x12, 0x26, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x33, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x3f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2a, 0x0a, 0x09, 0x54, 0x65, 0x61,
	0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x4c, 0x6f, 0x77, 0x10, 0x02, 0x32, 0xab, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f,
	0x7b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x54, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x3a, 0x04, 0x74,
	0x65, 0x61, 0x6d, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x72, 0x69, 0x6b, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_team_proto_rawDescOnce sync.Once
	file_team_proto_rawDescData = file_team_proto_rawDesc
)

func file_team_proto_rawDescGZIP() []byte {
	file_team_proto_rawDescOnce.Do(func() {
		file_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_proto_rawDescData)
	})
	return file_team_proto_rawDescData
}

var file_team_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_team_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_team_proto_goTypes = []interface{}{
	(TeamLevel)(0),       // 0: proto.TeamLevel
	(*TeamRequest)(nil),  // 1: proto.TeamRequest
	(*TeamResponse)(nil), // 2: proto.TeamResponse
	(*AddRequest)(nil),   // 3: proto.AddRequest
	(*AddResponse)(nil),  // 4: proto.AddResponse
	(*TeamDetail)(nil),   // 5: proto.TeamDetail
}
var file_team_proto_depIdxs = []int32{
	0, // 0: proto.TeamResponse.level:type_name -> proto.TeamLevel
	5, // 1: proto.AddRequest.team:type_name -> proto.TeamDetail
	1, // 2: proto.TeamService.GetTeamInfo:input_type -> proto.TeamRequest
	3, // 3: proto.TeamService.AddTeam:input_type -> proto.AddRequest
	2, // 4: proto.TeamService.GetTeamInfo:output_type -> proto.TeamResponse
	4, // 5: proto.TeamService.AddTeam:output_type -> proto.AddResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_team_proto_init() }
func file_team_proto_init() {
	if File_team_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamRequest); i {
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
		file_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamResponse); i {
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
		file_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
			RawDescriptor: file_team_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_team_proto_goTypes,
		DependencyIndexes: file_team_proto_depIdxs,
		EnumInfos:         file_team_proto_enumTypes,
		MessageInfos:      file_team_proto_msgTypes,
	}.Build()
	File_team_proto = out.File
	file_team_proto_rawDesc = nil
	file_team_proto_goTypes = nil
	file_team_proto_depIdxs = nil
}