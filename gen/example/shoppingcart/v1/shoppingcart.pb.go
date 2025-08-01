// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: example/shoppingcart/v1/shoppingcart.proto

package shoppingcartv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
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

type UpdateCartAction int32

const (
	UpdateCartAction_UPDATE_CART_ACTION_UNSPECIFIED UpdateCartAction = 0
	UpdateCartAction_UPDATE_CART_ACTION_ADD         UpdateCartAction = 1
	UpdateCartAction_UPDATE_CART_ACTION_REMOVE      UpdateCartAction = 2
)

// Enum value maps for UpdateCartAction.
var (
	UpdateCartAction_name = map[int32]string{
		0: "UPDATE_CART_ACTION_UNSPECIFIED",
		1: "UPDATE_CART_ACTION_ADD",
		2: "UPDATE_CART_ACTION_REMOVE",
	}
	UpdateCartAction_value = map[string]int32{
		"UPDATE_CART_ACTION_UNSPECIFIED": 0,
		"UPDATE_CART_ACTION_ADD":         1,
		"UPDATE_CART_ACTION_REMOVE":      2,
	}
)

func (x UpdateCartAction) Enum() *UpdateCartAction {
	p := new(UpdateCartAction)
	*p = x
	return p
}

func (x UpdateCartAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateCartAction) Descriptor() protoreflect.EnumDescriptor {
	return file_example_shoppingcart_v1_shoppingcart_proto_enumTypes[0].Descriptor()
}

func (UpdateCartAction) Type() protoreflect.EnumType {
	return &file_example_shoppingcart_v1_shoppingcart_proto_enumTypes[0]
}

func (x UpdateCartAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateCartAction.Descriptor instead.
func (UpdateCartAction) EnumDescriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{0}
}

type CartState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         map[string]int32       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"` // item_id -> quantity
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartState) Reset() {
	*x = CartState{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartState) ProtoMessage() {}

func (x *CartState) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartState.ProtoReflect.Descriptor instead.
func (*CartState) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{0}
}

func (x *CartState) GetItems() map[string]int32 {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckoutInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckoutInput) Reset() {
	*x = CheckoutInput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckoutInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutInput) ProtoMessage() {}

func (x *CheckoutInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutInput.ProtoReflect.Descriptor instead.
func (*CheckoutInput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{1}
}

type DescribeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeInput) Reset() {
	*x = DescribeInput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeInput) ProtoMessage() {}

func (x *DescribeInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeInput.ProtoReflect.Descriptor instead.
func (*DescribeInput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{2}
}

type DescribeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *CartState             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeOutput) Reset() {
	*x = DescribeOutput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeOutput) ProtoMessage() {}

func (x *DescribeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeOutput.ProtoReflect.Descriptor instead.
func (*DescribeOutput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeOutput) GetCart() *CartState {
	if x != nil {
		return x.Cart
	}
	return nil
}

type ShoppingCartInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *CartState             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShoppingCartInput) Reset() {
	*x = ShoppingCartInput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShoppingCartInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingCartInput) ProtoMessage() {}

func (x *ShoppingCartInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingCartInput.ProtoReflect.Descriptor instead.
func (*ShoppingCartInput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{4}
}

func (x *ShoppingCartInput) GetCart() *CartState {
	if x != nil {
		return x.Cart
	}
	return nil
}

type ShoppingCartOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *CartState             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShoppingCartOutput) Reset() {
	*x = ShoppingCartOutput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShoppingCartOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingCartOutput) ProtoMessage() {}

func (x *ShoppingCartOutput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingCartOutput.ProtoReflect.Descriptor instead.
func (*ShoppingCartOutput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{5}
}

func (x *ShoppingCartOutput) GetCart() *CartState {
	if x != nil {
		return x.Cart
	}
	return nil
}

type UpdateCartInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        UpdateCartAction       `protobuf:"varint,1,opt,name=action,proto3,enum=example.shoppingcart.v1.UpdateCartAction" json:"action,omitempty"`
	ItemId        string                 `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"` // item_id to add or remove
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCartInput) Reset() {
	*x = UpdateCartInput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCartInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartInput) ProtoMessage() {}

func (x *UpdateCartInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartInput.ProtoReflect.Descriptor instead.
func (*UpdateCartInput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCartInput) GetAction() UpdateCartAction {
	if x != nil {
		return x.Action
	}
	return UpdateCartAction_UPDATE_CART_ACTION_UNSPECIFIED
}

func (x *UpdateCartInput) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type UpdateCartOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *CartState             `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCartOutput) Reset() {
	*x = UpdateCartOutput{}
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCartOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartOutput) ProtoMessage() {}

func (x *UpdateCartOutput) ProtoReflect() protoreflect.Message {
	mi := &file_example_shoppingcart_v1_shoppingcart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartOutput.ProtoReflect.Descriptor instead.
func (*UpdateCartOutput) Descriptor() ([]byte, []int) {
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCartOutput) GetCart() *CartState {
	if x != nil {
		return x.Cart
	}
	return nil
}

var File_example_shoppingcart_v1_shoppingcart_proto protoreflect.FileDescriptor

var file_example_shoppingcart_v1_shoppingcart_proto_rawDesc = string([]byte{
	0x0a, 0x2a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x01, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0f, 0x0a, 0x0d, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x0f, 0x0a, 0x0d,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x48, 0x0a,
	0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x36, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x4b, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x63, 0x61, 0x72, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x43, 0x61, 0x72, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x22, 0x6d, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x22, 0x4a, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x2a, 0x71, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x54,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x41, 0x52, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x44, 0x10,
	0x01, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x54,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x02,
	0x32, 0x80, 0x05, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72,
	0x74, 0x12, 0xfc, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x2a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2b,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x43, 0x61, 0x72, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x92, 0x01, 0x8a, 0xc4,
	0x03, 0x8d, 0x01, 0x0a, 0x0a, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x0a, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x1a, 0x18, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x06, 0x42, 0x02, 0x08, 0x0a, 0x48,
	0x03, 0x18, 0x01, 0x28, 0x02, 0x2a, 0x33, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x2f, 0x24, 0x7b, 0x21, 0x20,
	0x6e, 0x61, 0x6e, 0x6f, 0x69, 0x64, 0x28, 0x29, 0x20, 0x7d, 0x72, 0x24, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x61, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x26, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x04, 0x9a,
	0xc4, 0x03, 0x00, 0x12, 0x9c, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x28, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x29, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x39, 0xaa, 0xc4, 0x03, 0x35, 0x0a, 0x31, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x6e, 0x61, 0x6e, 0x6f, 0x69, 0x64, 0x28, 0x29, 0x20, 0x7d,
	0x10, 0x01, 0x12, 0x50, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x26,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04,
	0xa2, 0xc4, 0x03, 0x00, 0x1a, 0x1d, 0x8a, 0xc4, 0x03, 0x19, 0x0a, 0x17, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74,
	0x2d, 0x76, 0x31, 0x42, 0x84, 0x02, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x45, 0x53, 0x58, 0xaa, 0x02, 0x17, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x17, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x63, 0x61, 0x72, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_example_shoppingcart_v1_shoppingcart_proto_rawDescOnce sync.Once
	file_example_shoppingcart_v1_shoppingcart_proto_rawDescData []byte
)

func file_example_shoppingcart_v1_shoppingcart_proto_rawDescGZIP() []byte {
	file_example_shoppingcart_v1_shoppingcart_proto_rawDescOnce.Do(func() {
		file_example_shoppingcart_v1_shoppingcart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_example_shoppingcart_v1_shoppingcart_proto_rawDesc), len(file_example_shoppingcart_v1_shoppingcart_proto_rawDesc)))
	})
	return file_example_shoppingcart_v1_shoppingcart_proto_rawDescData
}

var file_example_shoppingcart_v1_shoppingcart_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_shoppingcart_v1_shoppingcart_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_example_shoppingcart_v1_shoppingcart_proto_goTypes = []any{
	(UpdateCartAction)(0),      // 0: example.shoppingcart.v1.UpdateCartAction
	(*CartState)(nil),          // 1: example.shoppingcart.v1.CartState
	(*CheckoutInput)(nil),      // 2: example.shoppingcart.v1.CheckoutInput
	(*DescribeInput)(nil),      // 3: example.shoppingcart.v1.DescribeInput
	(*DescribeOutput)(nil),     // 4: example.shoppingcart.v1.DescribeOutput
	(*ShoppingCartInput)(nil),  // 5: example.shoppingcart.v1.ShoppingCartInput
	(*ShoppingCartOutput)(nil), // 6: example.shoppingcart.v1.ShoppingCartOutput
	(*UpdateCartInput)(nil),    // 7: example.shoppingcart.v1.UpdateCartInput
	(*UpdateCartOutput)(nil),   // 8: example.shoppingcart.v1.UpdateCartOutput
	nil,                        // 9: example.shoppingcart.v1.CartState.ItemsEntry
	(*emptypb.Empty)(nil),      // 10: google.protobuf.Empty
}
var file_example_shoppingcart_v1_shoppingcart_proto_depIdxs = []int32{
	9,  // 0: example.shoppingcart.v1.CartState.items:type_name -> example.shoppingcart.v1.CartState.ItemsEntry
	1,  // 1: example.shoppingcart.v1.DescribeOutput.cart:type_name -> example.shoppingcart.v1.CartState
	1,  // 2: example.shoppingcart.v1.ShoppingCartInput.cart:type_name -> example.shoppingcart.v1.CartState
	1,  // 3: example.shoppingcart.v1.ShoppingCartOutput.cart:type_name -> example.shoppingcart.v1.CartState
	0,  // 4: example.shoppingcart.v1.UpdateCartInput.action:type_name -> example.shoppingcart.v1.UpdateCartAction
	1,  // 5: example.shoppingcart.v1.UpdateCartOutput.cart:type_name -> example.shoppingcart.v1.CartState
	5,  // 6: example.shoppingcart.v1.ShoppingCart.ShoppingCart:input_type -> example.shoppingcart.v1.ShoppingCartInput
	3,  // 7: example.shoppingcart.v1.ShoppingCart.Describe:input_type -> example.shoppingcart.v1.DescribeInput
	7,  // 8: example.shoppingcart.v1.ShoppingCart.UpdateCart:input_type -> example.shoppingcart.v1.UpdateCartInput
	2,  // 9: example.shoppingcart.v1.ShoppingCart.Checkout:input_type -> example.shoppingcart.v1.CheckoutInput
	6,  // 10: example.shoppingcart.v1.ShoppingCart.ShoppingCart:output_type -> example.shoppingcart.v1.ShoppingCartOutput
	4,  // 11: example.shoppingcart.v1.ShoppingCart.Describe:output_type -> example.shoppingcart.v1.DescribeOutput
	8,  // 12: example.shoppingcart.v1.ShoppingCart.UpdateCart:output_type -> example.shoppingcart.v1.UpdateCartOutput
	10, // 13: example.shoppingcart.v1.ShoppingCart.Checkout:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_example_shoppingcart_v1_shoppingcart_proto_init() }
func file_example_shoppingcart_v1_shoppingcart_proto_init() {
	if File_example_shoppingcart_v1_shoppingcart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_shoppingcart_v1_shoppingcart_proto_rawDesc), len(file_example_shoppingcart_v1_shoppingcart_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_shoppingcart_v1_shoppingcart_proto_goTypes,
		DependencyIndexes: file_example_shoppingcart_v1_shoppingcart_proto_depIdxs,
		EnumInfos:         file_example_shoppingcart_v1_shoppingcart_proto_enumTypes,
		MessageInfos:      file_example_shoppingcart_v1_shoppingcart_proto_msgTypes,
	}.Build()
	File_example_shoppingcart_v1_shoppingcart_proto = out.File
	file_example_shoppingcart_v1_shoppingcart_proto_goTypes = nil
	file_example_shoppingcart_v1_shoppingcart_proto_depIdxs = nil
}
