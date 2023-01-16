// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/portfolio/portfolio.proto

package types

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SpotPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Price float64              `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SpotPrice) Reset() {
	*x = SpotPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotPrice) ProtoMessage() {}

func (x *SpotPrice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotPrice.ProtoReflect.Descriptor instead.
func (*SpotPrice) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{0}
}

func (x *SpotPrice) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *SpotPrice) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PriceHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName   string       `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	GeckoId     string       `protobuf:"bytes,2,opt,name=geckoId,proto3" json:"geckoId,omitempty"`
	BaseDenomId string       `protobuf:"bytes,3,opt,name=baseDenomId,proto3" json:"baseDenomId,omitempty"`
	Prices      []*SpotPrice `protobuf:"bytes,4,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *PriceHistory) Reset() {
	*x = PriceHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceHistory) ProtoMessage() {}

func (x *PriceHistory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceHistory.ProtoReflect.Descriptor instead.
func (*PriceHistory) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{1}
}

func (x *PriceHistory) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *PriceHistory) GetGeckoId() string {
	if x != nil {
		return x.GeckoId
	}
	return ""
}

func (x *PriceHistory) GetBaseDenomId() string {
	if x != nil {
		return x.BaseDenomId
	}
	return ""
}

func (x *PriceHistory) GetPrices() []*SpotPrice {
	if x != nil {
		return x.Prices
	}
	return nil
}

type AmountHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount []uint32 `protobuf:"varint,2,rep,packed,name=amount,proto3" json:"amount,omitempty"`
	//
	//Time recorded is when the amount is updated and added to the list on manual user input
	UpdateTimes []*timestamp.Timestamp `protobuf:"bytes,3,rep,name=update_times,json=updateTimes,proto3" json:"update_times,omitempty"`
}

func (x *AmountHistory) Reset() {
	*x = AmountHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmountHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmountHistory) ProtoMessage() {}

func (x *AmountHistory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmountHistory.ProtoReflect.Descriptor instead.
func (*AmountHistory) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{2}
}

func (x *AmountHistory) GetAmount() []uint32 {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *AmountHistory) GetUpdateTimes() []*timestamp.Timestamp {
	if x != nil {
		return x.UpdateTimes
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GeckoId string         `protobuf:"bytes,2,opt,name=geckoId,proto3" json:"geckoId,omitempty"`
	Amounts *AmountHistory `protobuf:"bytes,3,opt,name=amounts,proto3" json:"amounts,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Token) GetGeckoId() string {
	if x != nil {
		return x.GeckoId
	}
	return ""
}

func (x *Token) GetAmounts() *AmountHistory {
	if x != nil {
		return x.Amounts
	}
	return nil
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addr        string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Tokens      []*Token             `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{4}
}

func (x *Chain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chain) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Chain) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Chain) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Chains []*Chain `protobuf:"bytes,3,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{5}
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetChains() []*Chain {
	if x != nil {
		return x.Chains
	}
	return nil
}

type Portfolio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TrackedTokens []*PriceHistory `protobuf:"bytes,2,rep,name=trackedTokens,proto3" json:"trackedTokens,omitempty"`
	Accounts      []*Account      `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Portfolio) Reset() {
	*x = Portfolio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_portfolio_portfolio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Portfolio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Portfolio) ProtoMessage() {}

func (x *Portfolio) ProtoReflect() protoreflect.Message {
	mi := &file_proto_portfolio_portfolio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Portfolio.ProtoReflect.Descriptor instead.
func (*Portfolio) Descriptor() ([]byte, []int) {
	return file_proto_portfolio_portfolio_proto_rawDescGZIP(), []int{6}
}

func (x *Portfolio) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Portfolio) GetTrackedTokens() []*PriceHistory {
	if x != nil {
		return x.TrackedTokens
	}
	return nil
}

func (x *Portfolio) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_proto_portfolio_portfolio_proto protoreflect.FileDescriptor

var file_proto_portfolio_portfolio_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x09, 0x53, 0x70, 0x6f,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x92, 0x01, 0x0a,
	0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x65, 0x63, 0x6b, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65,
	0x63, 0x6b, 0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x6f, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x66, 0x0a, 0x0d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x22, 0x94, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x86, 0x01, 0x0a,
	0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x78, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_portfolio_portfolio_proto_rawDescOnce sync.Once
	file_proto_portfolio_portfolio_proto_rawDescData = file_proto_portfolio_portfolio_proto_rawDesc
)

func file_proto_portfolio_portfolio_proto_rawDescGZIP() []byte {
	file_proto_portfolio_portfolio_proto_rawDescOnce.Do(func() {
		file_proto_portfolio_portfolio_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_portfolio_portfolio_proto_rawDescData)
	})
	return file_proto_portfolio_portfolio_proto_rawDescData
}

var file_proto_portfolio_portfolio_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_portfolio_portfolio_proto_goTypes = []interface{}{
	(*SpotPrice)(nil),           // 0: types.SpotPrice
	(*PriceHistory)(nil),        // 1: types.PriceHistory
	(*AmountHistory)(nil),       // 2: types.AmountHistory
	(*Token)(nil),               // 3: types.Token
	(*Chain)(nil),               // 4: types.Chain
	(*Account)(nil),             // 5: types.Account
	(*Portfolio)(nil),           // 6: types.Portfolio
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_proto_portfolio_portfolio_proto_depIdxs = []int32{
	7, // 0: types.SpotPrice.time:type_name -> google.protobuf.Timestamp
	0, // 1: types.PriceHistory.prices:type_name -> types.SpotPrice
	7, // 2: types.AmountHistory.update_times:type_name -> google.protobuf.Timestamp
	2, // 3: types.Token.amounts:type_name -> types.AmountHistory
	3, // 4: types.Chain.tokens:type_name -> types.Token
	7, // 5: types.Chain.last_updated:type_name -> google.protobuf.Timestamp
	4, // 6: types.Account.chains:type_name -> types.Chain
	1, // 7: types.Portfolio.trackedTokens:type_name -> types.PriceHistory
	5, // 8: types.Portfolio.accounts:type_name -> types.Account
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_proto_portfolio_portfolio_proto_init() }
func file_proto_portfolio_portfolio_proto_init() {
	if File_proto_portfolio_portfolio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_portfolio_portfolio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotPrice); i {
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
		file_proto_portfolio_portfolio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceHistory); i {
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
		file_proto_portfolio_portfolio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmountHistory); i {
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
		file_proto_portfolio_portfolio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_proto_portfolio_portfolio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain); i {
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
		file_proto_portfolio_portfolio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_proto_portfolio_portfolio_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Portfolio); i {
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
			RawDescriptor: file_proto_portfolio_portfolio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_portfolio_portfolio_proto_goTypes,
		DependencyIndexes: file_proto_portfolio_portfolio_proto_depIdxs,
		MessageInfos:      file_proto_portfolio_portfolio_proto_msgTypes,
	}.Build()
	File_proto_portfolio_portfolio_proto = out.File
	file_proto_portfolio_portfolio_proto_rawDesc = nil
	file_proto_portfolio_portfolio_proto_goTypes = nil
	file_proto_portfolio_portfolio_proto_depIdxs = nil
}
