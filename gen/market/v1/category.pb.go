// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: market/v1/category.proto

package marketv1

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

type PublicCategory struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Id            int64                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Attributes    *PublicCategoryAttributes `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicCategory) Reset() {
	*x = PublicCategory{}
	mi := &file_market_v1_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCategory) ProtoMessage() {}

func (x *PublicCategory) ProtoReflect() protoreflect.Message {
	mi := &file_market_v1_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCategory.ProtoReflect.Descriptor instead.
func (*PublicCategory) Descriptor() ([]byte, []int) {
	return file_market_v1_category_proto_rawDescGZIP(), []int{0}
}

func (x *PublicCategory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PublicCategory) GetAttributes() *PublicCategoryAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type PublicCategoryAttributes struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Title         string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Slug          string                   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Description   string                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Content       string                   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Image         *v1.PublicFile           `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Collections   []*MarketplaceCollection `protobuf:"bytes,6,rep,name=collections,proto3" json:"collections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicCategoryAttributes) Reset() {
	*x = PublicCategoryAttributes{}
	mi := &file_market_v1_category_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicCategoryAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCategoryAttributes) ProtoMessage() {}

func (x *PublicCategoryAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_market_v1_category_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCategoryAttributes.ProtoReflect.Descriptor instead.
func (*PublicCategoryAttributes) Descriptor() ([]byte, []int) {
	return file_market_v1_category_proto_rawDescGZIP(), []int{1}
}

func (x *PublicCategoryAttributes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicCategoryAttributes) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *PublicCategoryAttributes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PublicCategoryAttributes) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PublicCategoryAttributes) GetImage() *v1.PublicFile {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *PublicCategoryAttributes) GetCollections() []*MarketplaceCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type CategoriesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Provider      string                 `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoriesRequest) Reset() {
	*x = CategoriesRequest{}
	mi := &file_market_v1_category_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesRequest) ProtoMessage() {}

func (x *CategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_v1_category_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesRequest.ProtoReflect.Descriptor instead.
func (*CategoriesRequest) Descriptor() ([]byte, []int) {
	return file_market_v1_category_proto_rawDescGZIP(), []int{2}
}

func (x *CategoriesRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type SuccessCategories struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*PublicCategory      `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuccessCategories) Reset() {
	*x = SuccessCategories{}
	mi := &file_market_v1_category_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessCategories) ProtoMessage() {}

func (x *SuccessCategories) ProtoReflect() protoreflect.Message {
	mi := &file_market_v1_category_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessCategories.ProtoReflect.Descriptor instead.
func (*SuccessCategories) Descriptor() ([]byte, []int) {
	return file_market_v1_category_proto_rawDescGZIP(), []int{3}
}

func (x *SuccessCategories) GetCategories() []*PublicCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CategoriesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Response:
	//
	//	*CategoriesResponse_Data
	//	*CategoriesResponse_Error
	Response      isCategoriesResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoriesResponse) Reset() {
	*x = CategoriesResponse{}
	mi := &file_market_v1_category_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesResponse) ProtoMessage() {}

func (x *CategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_market_v1_category_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesResponse.ProtoReflect.Descriptor instead.
func (*CategoriesResponse) Descriptor() ([]byte, []int) {
	return file_market_v1_category_proto_rawDescGZIP(), []int{4}
}

func (x *CategoriesResponse) GetResponse() isCategoriesResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *CategoriesResponse) GetData() *SuccessCategories {
	if x != nil {
		if x, ok := x.Response.(*CategoriesResponse_Data); ok {
			return x.Data
		}
	}
	return nil
}

func (x *CategoriesResponse) GetError() *emptypb.Empty {
	if x != nil {
		if x, ok := x.Response.(*CategoriesResponse_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isCategoriesResponse_Response interface {
	isCategoriesResponse_Response()
}

type CategoriesResponse_Data struct {
	Data *SuccessCategories `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type CategoriesResponse_Error struct {
	Error *emptypb.Empty `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*CategoriesResponse_Data) isCategoriesResponse_Response() {}

func (*CategoriesResponse_Error) isCategoriesResponse_Response() {}

var File_market_v1_category_proto protoreflect.FileDescriptor

var file_market_v1_category_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x18,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2b,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x2f, 0x0a, 0x11, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x22, 0x4e, 0x0a, 0x11, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x84, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9b, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x65, 0x2d, 0x6c, 0x61, 0x72, 0x73, 0x65,
	0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x6e, 0x69, 0x75, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_market_v1_category_proto_rawDescOnce sync.Once
	file_market_v1_category_proto_rawDescData []byte
)

func file_market_v1_category_proto_rawDescGZIP() []byte {
	file_market_v1_category_proto_rawDescOnce.Do(func() {
		file_market_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_market_v1_category_proto_rawDesc), len(file_market_v1_category_proto_rawDesc)))
	})
	return file_market_v1_category_proto_rawDescData
}

var file_market_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_market_v1_category_proto_goTypes = []any{
	(*PublicCategory)(nil),           // 0: market.v1.PublicCategory
	(*PublicCategoryAttributes)(nil), // 1: market.v1.PublicCategoryAttributes
	(*CategoriesRequest)(nil),        // 2: market.v1.CategoriesRequest
	(*SuccessCategories)(nil),        // 3: market.v1.SuccessCategories
	(*CategoriesResponse)(nil),       // 4: market.v1.CategoriesResponse
	(*v1.PublicFile)(nil),            // 5: common.v1.PublicFile
	(*MarketplaceCollection)(nil),    // 6: market.v1.MarketplaceCollection
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_market_v1_category_proto_depIdxs = []int32{
	1, // 0: market.v1.PublicCategory.attributes:type_name -> market.v1.PublicCategoryAttributes
	5, // 1: market.v1.PublicCategoryAttributes.image:type_name -> common.v1.PublicFile
	6, // 2: market.v1.PublicCategoryAttributes.collections:type_name -> market.v1.MarketplaceCollection
	0, // 3: market.v1.SuccessCategories.categories:type_name -> market.v1.PublicCategory
	3, // 4: market.v1.CategoriesResponse.data:type_name -> market.v1.SuccessCategories
	7, // 5: market.v1.CategoriesResponse.error:type_name -> google.protobuf.Empty
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_market_v1_category_proto_init() }
func file_market_v1_category_proto_init() {
	if File_market_v1_category_proto != nil {
		return
	}
	file_market_v1_collection_proto_init()
	file_market_v1_category_proto_msgTypes[4].OneofWrappers = []any{
		(*CategoriesResponse_Data)(nil),
		(*CategoriesResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_market_v1_category_proto_rawDesc), len(file_market_v1_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_market_v1_category_proto_goTypes,
		DependencyIndexes: file_market_v1_category_proto_depIdxs,
		MessageInfos:      file_market_v1_category_proto_msgTypes,
	}.Build()
	File_market_v1_category_proto = out.File
	file_market_v1_category_proto_goTypes = nil
	file_market_v1_category_proto_depIdxs = nil
}
