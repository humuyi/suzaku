// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: friend/friend.proto

package pb_friend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	pb_com "suzaku/pkg/proto/pb_com"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                     // 用户ID
	TargetUserId string `protobuf:"bytes,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"` // 添加好友用户ID
}

func (x *AddFriendReq) Reset() {
	*x = AddFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendReq) ProtoMessage() {}

func (x *AddFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_friend_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendReq.ProtoReflect.Descriptor instead.
func (*AddFriendReq) Descriptor() ([]byte, []int) {
	return file_friend_friend_proto_rawDescGZIP(), []int{0}
}

func (x *AddFriendReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddFriendReq) GetTargetUserId() string {
	if x != nil {
		return x.TargetUserId
	}
	return ""
}

type AddFriendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *pb_com.CommonResp `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
}

func (x *AddFriendResp) Reset() {
	*x = AddFriendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendResp) ProtoMessage() {}

func (x *AddFriendResp) ProtoReflect() protoreflect.Message {
	mi := &file_friend_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendResp.ProtoReflect.Descriptor instead.
func (*AddFriendResp) Descriptor() ([]byte, []int) {
	return file_friend_friend_proto_rawDescGZIP(), []int{1}
}

func (x *AddFriendResp) GetCommon() *pb_com.CommonResp {
	if x != nil {
		return x.Common
	}
	return nil
}

var File_friend_friend_proto protoreflect.FileDescriptor

var file_friend_friend_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x62, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x1a, 0x13, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x32, 0x48, 0x0a, 0x06, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x14, 0x5a, 0x12, 0x2e,
	0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x3b, 0x70, 0x62, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_friend_proto_rawDescOnce sync.Once
	file_friend_friend_proto_rawDescData = file_friend_friend_proto_rawDesc
)

func file_friend_friend_proto_rawDescGZIP() []byte {
	file_friend_friend_proto_rawDescOnce.Do(func() {
		file_friend_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_friend_proto_rawDescData)
	})
	return file_friend_friend_proto_rawDescData
}

var file_friend_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_friend_friend_proto_goTypes = []interface{}{
	(*AddFriendReq)(nil),      // 0: pb_friend.AddFriendReq
	(*AddFriendResp)(nil),     // 1: pb_friend.AddFriendResp
	(*pb_com.CommonResp)(nil), // 2: pb_com.CommonResp
}
var file_friend_friend_proto_depIdxs = []int32{
	2, // 0: pb_friend.AddFriendResp.common:type_name -> pb_com.CommonResp
	0, // 1: pb_friend.Friend.AddFriend:input_type -> pb_friend.AddFriendReq
	1, // 2: pb_friend.Friend.AddFriend:output_type -> pb_friend.AddFriendResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_friend_proto_init() }
func file_friend_friend_proto_init() {
	if File_friend_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friend_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendReq); i {
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
		file_friend_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendResp); i {
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
			RawDescriptor: file_friend_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_friend_proto_goTypes,
		DependencyIndexes: file_friend_friend_proto_depIdxs,
		MessageInfos:      file_friend_friend_proto_msgTypes,
	}.Build()
	File_friend_friend_proto = out.File
	file_friend_friend_proto_rawDesc = nil
	file_friend_friend_proto_goTypes = nil
	file_friend_friend_proto_depIdxs = nil
}
