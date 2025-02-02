// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: frontend/v1/collectible.proto

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

type MarketplaceCollectible struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Attributes    *MarketplaceCollectibleAttributes `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	unknownFields protoimpl.UnknownFields
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *MarketplaceCollectible) Reset() {
	*x = MarketplaceCollectible{}
	mi := &file_frontend_v1_collectible_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketplaceCollectible) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketplaceCollectible) ProtoMessage() {}

func (x *MarketplaceCollectible) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_collectible_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketplaceCollectible.ProtoReflect.Descriptor instead.
func (*MarketplaceCollectible) Descriptor() ([]byte, []int) {
	return file_frontend_v1_collectible_proto_rawDescGZIP(), []int{0}
}

func (x *MarketplaceCollectible) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MarketplaceCollectible) GetAttributes() *MarketplaceCollectibleAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type MarketplaceCollectibleAttributes struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	Creator       *PublicUser                     `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	Owner         *PublicUser                     `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Details       *MarketplaceCollectibleDetails  `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
	Metadata      *MarketplaceCollectibleMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Uri           string                          `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	TokenIds      []int64                         `protobuf:"varint,2,rep,packed,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	CollectionId  int64 `protobuf:"varint,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *MarketplaceCollectibleAttributes) Reset() {
	*x = MarketplaceCollectibleAttributes{}
	mi := &file_frontend_v1_collectible_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketplaceCollectibleAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketplaceCollectibleAttributes) ProtoMessage() {}

func (x *MarketplaceCollectibleAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_collectible_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketplaceCollectibleAttributes.ProtoReflect.Descriptor instead.
func (*MarketplaceCollectibleAttributes) Descriptor() ([]byte, []int) {
	return file_frontend_v1_collectible_proto_rawDescGZIP(), []int{1}
}

func (x *MarketplaceCollectibleAttributes) GetCollectionId() int64 {
	if x != nil {
		return x.CollectionId
	}
	return 0
}

func (x *MarketplaceCollectibleAttributes) GetTokenIds() []int64 {
	if x != nil {
		return x.TokenIds
	}
	return nil
}

func (x *MarketplaceCollectibleAttributes) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *MarketplaceCollectibleAttributes) GetCreator() *PublicUser {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *MarketplaceCollectibleAttributes) GetOwner() *PublicUser {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *MarketplaceCollectibleAttributes) GetDetails() *MarketplaceCollectibleDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *MarketplaceCollectibleAttributes) GetMetadata() *MarketplaceCollectibleMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MarketplaceCollectibleDetails struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	StartPrice      string                 `protobuf:"bytes,15,opt,name=start_price,json=startPrice,proto3" json:"start_price,omitempty"`
	Total           string                 `protobuf:"bytes,19,opt,name=total,proto3" json:"total,omitempty"`
	PriceWei        string                 `protobuf:"bytes,11,opt,name=price_wei,json=priceWei,proto3" json:"price_wei,omitempty"`
	Collection      string                 `protobuf:"bytes,4,opt,name=collection,proto3" json:"collection,omitempty"`
	ReservePriceWei string                 `protobuf:"bytes,14,opt,name=reserve_price_wei,json=reservePriceWei,proto3" json:"reserve_price_wei,omitempty"`
	Fee             string                 `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
	FeeWei          string                 `protobuf:"bytes,7,opt,name=fee_wei,json=feeWei,proto3" json:"fee_wei,omitempty"`
	TotalWei        string                 `protobuf:"bytes,20,opt,name=total_wei,json=totalWei,proto3" json:"total_wei,omitempty"`
	ReservePrice    string                 `protobuf:"bytes,13,opt,name=reserve_price,json=reservePrice,proto3" json:"reserve_price,omitempty"`
	Price           string                 `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	Address         string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Tags            string                 `protobuf:"bytes,18,opt,name=tags,proto3" json:"tags,omitempty"`
	StartPriceWei   string                 `protobuf:"bytes,16,opt,name=start_price_wei,json=startPriceWei,proto3" json:"start_price_wei,omitempty"`
	unknownFields   protoimpl.UnknownFields
	StartTime       int64 `protobuf:"varint,17,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime         int64 `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Quantity        int64 `protobuf:"varint,12,opt,name=quantity,proto3" json:"quantity,omitempty"`
	sizeCache       protoimpl.SizeCache
	IsStarted       bool `protobuf:"varint,9,opt,name=is_started,json=isStarted,proto3" json:"is_started,omitempty"`
	Auction         bool `protobuf:"varint,2,opt,name=auction,proto3" json:"auction,omitempty"`
	Cancelled       bool `protobuf:"varint,3,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	Fulfilled       bool `protobuf:"varint,8,opt,name=fulfilled,proto3" json:"fulfilled,omitempty"`
}

func (x *MarketplaceCollectibleDetails) Reset() {
	*x = MarketplaceCollectibleDetails{}
	mi := &file_frontend_v1_collectible_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketplaceCollectibleDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketplaceCollectibleDetails) ProtoMessage() {}

func (x *MarketplaceCollectibleDetails) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_collectible_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketplaceCollectibleDetails.ProtoReflect.Descriptor instead.
func (*MarketplaceCollectibleDetails) Descriptor() ([]byte, []int) {
	return file_frontend_v1_collectible_proto_rawDescGZIP(), []int{2}
}

func (x *MarketplaceCollectibleDetails) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetAuction() bool {
	if x != nil {
		return x.Auction
	}
	return false
}

func (x *MarketplaceCollectibleDetails) GetCancelled() bool {
	if x != nil {
		return x.Cancelled
	}
	return false
}

func (x *MarketplaceCollectibleDetails) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *MarketplaceCollectibleDetails) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetFeeWei() string {
	if x != nil {
		return x.FeeWei
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetFulfilled() bool {
	if x != nil {
		return x.Fulfilled
	}
	return false
}

func (x *MarketplaceCollectibleDetails) GetIsStarted() bool {
	if x != nil {
		return x.IsStarted
	}
	return false
}

func (x *MarketplaceCollectibleDetails) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetPriceWei() string {
	if x != nil {
		return x.PriceWei
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *MarketplaceCollectibleDetails) GetReservePrice() string {
	if x != nil {
		return x.ReservePrice
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetReservePriceWei() string {
	if x != nil {
		return x.ReservePriceWei
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetStartPrice() string {
	if x != nil {
		return x.StartPrice
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetStartPriceWei() string {
	if x != nil {
		return x.StartPriceWei
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *MarketplaceCollectibleDetails) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *MarketplaceCollectibleDetails) GetTotalWei() string {
	if x != nil {
		return x.TotalWei
	}
	return ""
}

type MarketplaceCollectibleMetadata struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExternalUrl     string                 `protobuf:"bytes,2,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	AnimationUrl    string                 `protobuf:"bytes,3,opt,name=animation_url,json=animationUrl,proto3" json:"animation_url,omitempty"`
	BackgroundColor string                 `protobuf:"bytes,4,opt,name=background_color,json=backgroundColor,proto3" json:"background_color,omitempty"`
	Description     string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	YoutubeUrl      string                 `protobuf:"bytes,6,opt,name=youtube_url,json=youtubeUrl,proto3" json:"youtube_url,omitempty"`
	ImageUrl        string                 `protobuf:"bytes,7,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Attributes      []*MetadataAttributes  `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MarketplaceCollectibleMetadata) Reset() {
	*x = MarketplaceCollectibleMetadata{}
	mi := &file_frontend_v1_collectible_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketplaceCollectibleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketplaceCollectibleMetadata) ProtoMessage() {}

func (x *MarketplaceCollectibleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_collectible_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketplaceCollectibleMetadata.ProtoReflect.Descriptor instead.
func (*MarketplaceCollectibleMetadata) Descriptor() ([]byte, []int) {
	return file_frontend_v1_collectible_proto_rawDescGZIP(), []int{3}
}

func (x *MarketplaceCollectibleMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetAnimationUrl() string {
	if x != nil {
		return x.AnimationUrl
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetBackgroundColor() string {
	if x != nil {
		return x.BackgroundColor
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetYoutubeUrl() string {
	if x != nil {
		return x.YoutubeUrl
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *MarketplaceCollectibleMetadata) GetAttributes() []*MetadataAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type MetadataAttributes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TraitType     string                 `protobuf:"bytes,1,opt,name=trait_type,json=traitType,proto3" json:"trait_type,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetadataAttributes) Reset() {
	*x = MetadataAttributes{}
	mi := &file_frontend_v1_collectible_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataAttributes) ProtoMessage() {}

func (x *MetadataAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_v1_collectible_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataAttributes.ProtoReflect.Descriptor instead.
func (*MetadataAttributes) Descriptor() ([]byte, []int) {
	return file_frontend_v1_collectible_proto_rawDescGZIP(), []int{4}
}

func (x *MetadataAttributes) GetTraitType() string {
	if x != nil {
		return x.TraitType
	}
	return ""
}

func (x *MetadataAttributes) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_frontend_v1_collectible_proto protoreflect.FileDescriptor

var file_frontend_v1_collectible_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x16, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4d,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xe7, 0x02,
	0x0a, 0x20, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x47,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe3, 0x04, 0x0a, 0x1d, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x62,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x5f,
	0x77, 0x65, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x65, 0x65, 0x57, 0x65,
	0x69, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x65,
	0x69, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x57, 0x65,
	0x69, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x57, 0x65, 0x69, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x77,
	0x65, 0x69, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x57, 0x65, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x22, 0xc8, 0x02,
	0x0a, 0x1e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x29, 0x0a, 0x10,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x79, 0x6f, 0x75,
	0x74, 0x75, 0x62, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0xac, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
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
	file_frontend_v1_collectible_proto_rawDescOnce sync.Once
	file_frontend_v1_collectible_proto_rawDescData []byte
)

func file_frontend_v1_collectible_proto_rawDescGZIP() []byte {
	file_frontend_v1_collectible_proto_rawDescOnce.Do(func() {
		file_frontend_v1_collectible_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_frontend_v1_collectible_proto_rawDesc), len(file_frontend_v1_collectible_proto_rawDesc)))
	})
	return file_frontend_v1_collectible_proto_rawDescData
}

var file_frontend_v1_collectible_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_frontend_v1_collectible_proto_goTypes = []any{
	(*MarketplaceCollectible)(nil),           // 0: frontend.v1.MarketplaceCollectible
	(*MarketplaceCollectibleAttributes)(nil), // 1: frontend.v1.MarketplaceCollectibleAttributes
	(*MarketplaceCollectibleDetails)(nil),    // 2: frontend.v1.MarketplaceCollectibleDetails
	(*MarketplaceCollectibleMetadata)(nil),   // 3: frontend.v1.MarketplaceCollectibleMetadata
	(*MetadataAttributes)(nil),               // 4: frontend.v1.MetadataAttributes
	(*PublicUser)(nil),                       // 5: frontend.v1.PublicUser
}
var file_frontend_v1_collectible_proto_depIdxs = []int32{
	1, // 0: frontend.v1.MarketplaceCollectible.attributes:type_name -> frontend.v1.MarketplaceCollectibleAttributes
	5, // 1: frontend.v1.MarketplaceCollectibleAttributes.creator:type_name -> frontend.v1.PublicUser
	5, // 2: frontend.v1.MarketplaceCollectibleAttributes.owner:type_name -> frontend.v1.PublicUser
	2, // 3: frontend.v1.MarketplaceCollectibleAttributes.details:type_name -> frontend.v1.MarketplaceCollectibleDetails
	3, // 4: frontend.v1.MarketplaceCollectibleAttributes.metadata:type_name -> frontend.v1.MarketplaceCollectibleMetadata
	4, // 5: frontend.v1.MarketplaceCollectibleMetadata.attributes:type_name -> frontend.v1.MetadataAttributes
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_frontend_v1_collectible_proto_init() }
func file_frontend_v1_collectible_proto_init() {
	if File_frontend_v1_collectible_proto != nil {
		return
	}
	file_frontend_v1_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_frontend_v1_collectible_proto_rawDesc), len(file_frontend_v1_collectible_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frontend_v1_collectible_proto_goTypes,
		DependencyIndexes: file_frontend_v1_collectible_proto_depIdxs,
		MessageInfos:      file_frontend_v1_collectible_proto_msgTypes,
	}.Build()
	File_frontend_v1_collectible_proto = out.File
	file_frontend_v1_collectible_proto_goTypes = nil
	file_frontend_v1_collectible_proto_depIdxs = nil
}
