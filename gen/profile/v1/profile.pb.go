// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: profile/v1/profile.proto

package profilev1

import (
	v1 "github.com/ole-larsen/plutonium/gen/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserForm struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Csrf          string                 `protobuf:"bytes,1,opt,name=csrf,proto3" json:"csrf,omitempty"`
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Address       string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Gravatar      string                 `protobuf:"bytes,6,opt,name=gravatar,proto3" json:"gravatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserForm) Reset() {
	*x = UserForm{}
	mi := &file_profile_v1_profile_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserForm) ProtoMessage() {}

func (x *UserForm) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserForm.ProtoReflect.Descriptor instead.
func (*UserForm) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UserForm) GetCsrf() string {
	if x != nil {
		return x.Csrf
	}
	return ""
}

func (x *UserForm) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserForm) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserForm) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserForm) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserForm) GetGravatar() string {
	if x != nil {
		return x.Gravatar
	}
	return ""
}

type PatchUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Body          *UserForm              `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatchUserRequest) Reset() {
	*x = PatchUserRequest{}
	mi := &file_profile_v1_profile_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUserRequest) ProtoMessage() {}

func (x *PatchUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUserRequest.ProtoReflect.Descriptor instead.
func (*PatchUserRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{1}
}

func (x *PatchUserRequest) GetBody() *UserForm {
	if x != nil {
		return x.Body
	}
	return nil
}

type PatchUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Response:
	//
	//	*PatchUserResponse_User
	//	*PatchUserResponse_Error
	Response      isPatchUserResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatchUserResponse) Reset() {
	*x = PatchUserResponse{}
	mi := &file_profile_v1_profile_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUserResponse) ProtoMessage() {}

func (x *PatchUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUserResponse.ProtoReflect.Descriptor instead.
func (*PatchUserResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{2}
}

func (x *PatchUserResponse) GetResponse() isPatchUserResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *PatchUserResponse) GetUser() *v1.PublicUser {
	if x != nil {
		if x, ok := x.Response.(*PatchUserResponse_User); ok {
			return x.User
		}
	}
	return nil
}

func (x *PatchUserResponse) GetError() *emptypb.Empty {
	if x != nil {
		if x, ok := x.Response.(*PatchUserResponse_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isPatchUserResponse_Response interface {
	isPatchUserResponse_Response()
}

type PatchUserResponse_User struct {
	User *v1.PublicUser `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type PatchUserResponse_Error struct {
	Error *emptypb.Empty `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*PatchUserResponse_User) isPatchUserResponse_Response() {}

func (*PatchUserResponse_Error) isPatchUserResponse_Response() {}

var File_profile_v1_profile_proto protoreflect.FileDescriptor

var file_profile_v1_profile_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x73, 0x72, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x73, 0x72, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x22, 0x3c, 0x0a, 0x10, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x7c, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5a,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x48, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa1, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x65, 0x2d, 0x6c, 0x61,
	0x72, 0x73, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x6e, 0x69, 0x75, 0x6d, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_profile_v1_profile_proto_rawDescOnce sync.Once
	file_profile_v1_profile_proto_rawDescData []byte
)

func file_profile_v1_profile_proto_rawDescGZIP() []byte {
	file_profile_v1_profile_proto_rawDescOnce.Do(func() {
		file_profile_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_profile_v1_profile_proto_rawDesc), len(file_profile_v1_profile_proto_rawDesc)))
	})
	return file_profile_v1_profile_proto_rawDescData
}

var file_profile_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_profile_v1_profile_proto_goTypes = []any{
	(*UserForm)(nil),          // 0: profile.v1.UserForm
	(*PatchUserRequest)(nil),  // 1: profile.v1.PatchUserRequest
	(*PatchUserResponse)(nil), // 2: profile.v1.PatchUserResponse
	(*v1.PublicUser)(nil),     // 3: common.v1.PublicUser
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_profile_v1_profile_proto_depIdxs = []int32{
	0, // 0: profile.v1.PatchUserRequest.body:type_name -> profile.v1.UserForm
	3, // 1: profile.v1.PatchUserResponse.user:type_name -> common.v1.PublicUser
	4, // 2: profile.v1.PatchUserResponse.error:type_name -> google.protobuf.Empty
	1, // 3: profile.v1.ProfileService.PatchUser:input_type -> profile.v1.PatchUserRequest
	2, // 4: profile.v1.ProfileService.PatchUser:output_type -> profile.v1.PatchUserResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_profile_v1_profile_proto_init() }
func file_profile_v1_profile_proto_init() {
	if File_profile_v1_profile_proto != nil {
		return
	}
	file_profile_v1_profile_proto_msgTypes[2].OneofWrappers = []any{
		(*PatchUserResponse_User)(nil),
		(*PatchUserResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_profile_v1_profile_proto_rawDesc), len(file_profile_v1_profile_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_v1_profile_proto_goTypes,
		DependencyIndexes: file_profile_v1_profile_proto_depIdxs,
		MessageInfos:      file_profile_v1_profile_proto_msgTypes,
	}.Build()
	File_profile_v1_profile_proto = out.File
	file_profile_v1_profile_proto_goTypes = nil
	file_profile_v1_profile_proto_depIdxs = nil
}
