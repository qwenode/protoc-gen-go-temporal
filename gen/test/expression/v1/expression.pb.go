// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: test/expression/v1/expression.proto

package expressionv1

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

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestVal    string                 `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	IntField      int32                  `protobuf:"varint,3,opt,name=int_field,json=intField,proto3" json:"int_field,omitempty"`
	BoolField     bool                   `protobuf:"varint,4,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	BytesField    []byte                 `protobuf:"bytes,5,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	DoubleField   float64                `protobuf:"fixed64,6,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	OuterSingle   *Request_OuterNested   `protobuf:"bytes,7,opt,name=outer_single,json=outerSingle,proto3" json:"outer_single,omitempty"`
	OuterList     []*Request_OuterNested `protobuf:"bytes,8,rep,name=outer_list,json=outerList,proto3" json:"outer_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_test_expression_v1_expression_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_expression_v1_expression_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_test_expression_v1_expression_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetIntField() int32 {
	if x != nil {
		return x.IntField
	}
	return 0
}

func (x *Request) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *Request) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *Request) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *Request) GetOuterSingle() *Request_OuterNested {
	if x != nil {
		return x.OuterSingle
	}
	return nil
}

func (x *Request) GetOuterList() []*Request_OuterNested {
	if x != nil {
		return x.OuterList
	}
	return nil
}

type Request_OuterNested struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	Foo           string                             `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	InnerSingle   *Request_OuterNested_InnerNested   `protobuf:"bytes,2,opt,name=inner_single,json=innerSingle,proto3" json:"inner_single,omitempty"`
	InnerList     []*Request_OuterNested_InnerNested `protobuf:"bytes,3,rep,name=inner_list,json=innerList,proto3" json:"inner_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request_OuterNested) Reset() {
	*x = Request_OuterNested{}
	mi := &file_test_expression_v1_expression_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request_OuterNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_OuterNested) ProtoMessage() {}

func (x *Request_OuterNested) ProtoReflect() protoreflect.Message {
	mi := &file_test_expression_v1_expression_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_OuterNested.ProtoReflect.Descriptor instead.
func (*Request_OuterNested) Descriptor() ([]byte, []int) {
	return file_test_expression_v1_expression_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_OuterNested) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *Request_OuterNested) GetInnerSingle() *Request_OuterNested_InnerNested {
	if x != nil {
		return x.InnerSingle
	}
	return nil
}

func (x *Request_OuterNested) GetInnerList() []*Request_OuterNested_InnerNested {
	if x != nil {
		return x.InnerList
	}
	return nil
}

type Request_OuterNested_InnerNested struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bar           string                 `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request_OuterNested_InnerNested) Reset() {
	*x = Request_OuterNested_InnerNested{}
	mi := &file_test_expression_v1_expression_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request_OuterNested_InnerNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_OuterNested_InnerNested) ProtoMessage() {}

func (x *Request_OuterNested_InnerNested) ProtoReflect() protoreflect.Message {
	mi := &file_test_expression_v1_expression_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_OuterNested_InnerNested.ProtoReflect.Descriptor instead.
func (*Request_OuterNested_InnerNested) Descriptor() ([]byte, []int) {
	return file_test_expression_v1_expression_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Request_OuterNested_InnerNested) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

var File_test_expression_v1_expression_proto protoreflect.FileDescriptor

var file_test_expression_v1_expression_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xed, 0x04, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x56, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x0b, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x84,
	0x02, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f,
	0x12, 0x62, 0x0a, 0x0c, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x0b, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x5e, 0x0a, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1f, 0x0a, 0x0b, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x61, 0x72, 0x42, 0xa1, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64,
	0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02,
	0x04, 0x54, 0x56, 0x54, 0x45, 0xaa, 0x02, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x56, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_test_expression_v1_expression_proto_rawDescOnce sync.Once
	file_test_expression_v1_expression_proto_rawDescData []byte
)

func file_test_expression_v1_expression_proto_rawDescGZIP() []byte {
	file_test_expression_v1_expression_proto_rawDescOnce.Do(func() {
		file_test_expression_v1_expression_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_test_expression_v1_expression_proto_rawDesc), len(file_test_expression_v1_expression_proto_rawDesc)))
	})
	return file_test_expression_v1_expression_proto_rawDescData
}

var file_test_expression_v1_expression_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_expression_v1_expression_proto_goTypes = []any{
	(*Request)(nil),                         // 0: temporal.v1.test.expression.v1.Request
	(*Request_OuterNested)(nil),             // 1: temporal.v1.test.expression.v1.Request.OuterNested
	(*Request_OuterNested_InnerNested)(nil), // 2: temporal.v1.test.expression.v1.Request.OuterNested.InnerNested
}
var file_test_expression_v1_expression_proto_depIdxs = []int32{
	1, // 0: temporal.v1.test.expression.v1.Request.outer_single:type_name -> temporal.v1.test.expression.v1.Request.OuterNested
	1, // 1: temporal.v1.test.expression.v1.Request.outer_list:type_name -> temporal.v1.test.expression.v1.Request.OuterNested
	2, // 2: temporal.v1.test.expression.v1.Request.OuterNested.inner_single:type_name -> temporal.v1.test.expression.v1.Request.OuterNested.InnerNested
	2, // 3: temporal.v1.test.expression.v1.Request.OuterNested.inner_list:type_name -> temporal.v1.test.expression.v1.Request.OuterNested.InnerNested
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_test_expression_v1_expression_proto_init() }
func file_test_expression_v1_expression_proto_init() {
	if File_test_expression_v1_expression_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_test_expression_v1_expression_proto_rawDesc), len(file_test_expression_v1_expression_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_expression_v1_expression_proto_goTypes,
		DependencyIndexes: file_test_expression_v1_expression_proto_depIdxs,
		MessageInfos:      file_test_expression_v1_expression_proto_msgTypes,
	}.Build()
	File_test_expression_v1_expression_proto = out.File
	file_test_expression_v1_expression_proto_goTypes = nil
	file_test_expression_v1_expression_proto_depIdxs = nil
}
