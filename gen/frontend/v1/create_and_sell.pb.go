// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: frontend/v1/create_and_sell.proto

package frontendv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PublicCreateAndSellItem struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	Attributes    *PublicCreateAndSellItemAttributes `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	unknownFields protoimpl.UnknownFields
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *PublicCreateAndSellItem) Reset() {
	*x = PublicCreateAndSellItem{}
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicCreateAndSellItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCreateAndSellItem) ProtoMessage() {}

func (x *PublicCreateAndSellItem) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCreateAndSellItem.ProtoReflect.Descriptor instead.
func (*PublicCreateAndSellItem) Descriptor() ([]byte, []int) {
	return file_frontend_v1_create_and_sell_proto_rawDescGZIP(), []int{0}
}

func (x *PublicCreateAndSellItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PublicCreateAndSellItem) GetAttributes() *PublicCreateAndSellItemAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type PublicCreateAndSellItemAttributes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Link          string                 `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image         *PublicFile            `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicCreateAndSellItemAttributes) Reset() {
	*x = PublicCreateAndSellItemAttributes{}
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicCreateAndSellItemAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCreateAndSellItemAttributes) ProtoMessage() {}

func (x *PublicCreateAndSellItemAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCreateAndSellItemAttributes.ProtoReflect.Descriptor instead.
func (*PublicCreateAndSellItemAttributes) Descriptor() ([]byte, []int) {
	return file_frontend_v1_create_and_sell_proto_rawDescGZIP(), []int{1}
}

func (x *PublicCreateAndSellItemAttributes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicCreateAndSellItemAttributes) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *PublicCreateAndSellItemAttributes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PublicCreateAndSellItemAttributes) GetImage() *PublicFile {
	if x != nil {
		return x.Image
	}
	return nil
}

type CreateAndSellItemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Provider      string                 `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAndSellItemsRequest) Reset() {
	*x = CreateAndSellItemsRequest{}
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAndSellItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndSellItemsRequest) ProtoMessage() {}

func (x *CreateAndSellItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAndSellItemsRequest.ProtoReflect.Descriptor instead.
func (*CreateAndSellItemsRequest) Descriptor() ([]byte, []int) {
	return file_frontend_v1_create_and_sell_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAndSellItemsRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type CreateAndSellRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Provider      string                 `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAndSellRequest) Reset() {
	*x = CreateAndSellRequest{}
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAndSellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndSellRequest) ProtoMessage() {}

func (x *CreateAndSellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_create_and_sell_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAndSellRequest.ProtoReflect.Descriptor instead.
func (*CreateAndSellRequest) Descriptor() ([]byte, []int) {
	return file_frontend_v1_create_and_sell_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAndSellRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

var File_frontend_v1_create_and_sell_proto protoreflect.FileDescriptor

var file_frontend_v1_create_and_sell_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x16, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x17, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6c, 0x6c, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x21, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e,
	0x64, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x32, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x42, 0xae, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64,
	0x53, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x65, 0x2d, 0x6c, 0x61, 0x72, 0x73,
	0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x6e, 0x69, 0x75, 0x6d, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x0b,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_frontend_v1_create_and_sell_proto_rawDescOnce sync.Once
	file_frontend_v1_create_and_sell_proto_rawDescData []byte
)

func file_frontend_v1_create_and_sell_proto_rawDescGZIP() []byte {
	file_frontend_v1_create_and_sell_proto_rawDescOnce.Do(func() {
		file_frontend_v1_create_and_sell_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_frontend_v1_create_and_sell_proto_rawDesc), len(file_frontend_v1_create_and_sell_proto_rawDesc)))
	})
	return file_frontend_v1_create_and_sell_proto_rawDescData
}

var file_frontend_v1_create_and_sell_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_frontend_v1_create_and_sell_proto_goTypes = []any{
	(*PublicCreateAndSellItem)(nil),           // 0: frontend.v1.PublicCreateAndSellItem
	(*PublicCreateAndSellItemAttributes)(nil), // 1: frontend.v1.PublicCreateAndSellItemAttributes
	(*CreateAndSellItemsRequest)(nil),         // 2: frontend.v1.CreateAndSellItemsRequest
	(*CreateAndSellRequest)(nil),              // 3: frontend.v1.CreateAndSellRequest
	(*PublicFile)(nil),                        // 4: frontend.v1.PublicFile
}
var file_frontend_v1_create_and_sell_proto_depIdxs = []int32{
	1, // 0: frontend.v1.PublicCreateAndSellItem.attributes:type_name -> frontend.v1.PublicCreateAndSellItemAttributes
	4, // 1: frontend.v1.PublicCreateAndSellItemAttributes.image:type_name -> frontend.v1.PublicFile
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_frontend_v1_create_and_sell_proto_init() }
func file_frontend_v1_create_and_sell_proto_init() {
	if File_frontend_v1_create_and_sell_proto != nil {
		return
	}
	file_frontend_v1_file_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_frontend_v1_create_and_sell_proto_rawDesc), len(file_frontend_v1_create_and_sell_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frontend_v1_create_and_sell_proto_goTypes,
		DependencyIndexes: file_frontend_v1_create_and_sell_proto_depIdxs,
		MessageInfos:      file_frontend_v1_create_and_sell_proto_msgTypes,
	}.Build()
	File_frontend_v1_create_and_sell_proto = out.File
	file_frontend_v1_create_and_sell_proto_goTypes = nil
	file_frontend_v1_create_and_sell_proto_depIdxs = nil
}
