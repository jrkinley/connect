// Copyright 2025 Redpanda Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: redpanda/runtime/v1alpha1/message.proto

package runtimepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// `NullValue` is a representation of a null value.
type NullValue int32

const (
	NullValue_NULL_VALUE NullValue = 0
)

// Enum value maps for NullValue.
var (
	NullValue_name = map[int32]string{
		0: "NULL_VALUE",
	}
	NullValue_value = map[string]int32{
		"NULL_VALUE": 0,
	}
)

func (x NullValue) Enum() *NullValue {
	p := new(NullValue)
	*p = x
	return p
}

func (x NullValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NullValue) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_runtime_v1alpha1_message_proto_enumTypes[0].Descriptor()
}

func (NullValue) Type() protoreflect.EnumType {
	return &file_redpanda_runtime_v1alpha1_message_proto_enumTypes[0]
}

func (x NullValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NullValue.Descriptor instead.
func (NullValue) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{0}
}

// `StructValue` represents a struct value which can be used to represent a
// structured data value.
type StructValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        map[string]*Value      `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StructValue) Reset() {
	*x = StructValue{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StructValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructValue) ProtoMessage() {}

func (x *StructValue) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructValue.ProtoReflect.Descriptor instead.
func (*StructValue) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{0}
}

func (x *StructValue) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

// `ListValue` represents a list value which can be used to represent a list of
// values.
type ListValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []*Value               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{1}
}

func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

// `Value` represents a dynamically typed value which can be used to represent
// a value within a Redpanda Connect pipeline.
type Value struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Kind:
	//
	//	*Value_NullValue
	//	*Value_StringValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_BoolValue
	//	*Value_TimestampValue
	//	*Value_BytesValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind          isValue_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetKind() isValue_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *Value) GetNullValue() NullValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_NullValue); ok {
			return x.NullValue
		}
	}
	return NullValue_NULL_VALUE
}

func (x *Value) GetStringValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *Value) GetIntegerValue() int64 {
	if x != nil {
		if x, ok := x.Kind.(*Value_IntegerValue); ok {
			return x.IntegerValue
		}
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x != nil {
		if x, ok := x.Kind.(*Value_DoubleValue); ok {
			return x.DoubleValue
		}
	}
	return 0
}

func (x *Value) GetBoolValue() bool {
	if x != nil {
		if x, ok := x.Kind.(*Value_BoolValue); ok {
			return x.BoolValue
		}
	}
	return false
}

func (x *Value) GetTimestampValue() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.Kind.(*Value_TimestampValue); ok {
			return x.TimestampValue
		}
	}
	return nil
}

func (x *Value) GetBytesValue() []byte {
	if x != nil {
		if x, ok := x.Kind.(*Value_BytesValue); ok {
			return x.BytesValue
		}
	}
	return nil
}

func (x *Value) GetStructValue() *StructValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_StructValue); ok {
			return x.StructValue
		}
	}
	return nil
}

func (x *Value) GetListValue() *ListValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_ListValue); ok {
			return x.ListValue
		}
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=redpanda.runtime.v1alpha1.NullValue,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,3,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	TimestampValue *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_StructValue struct {
	StructValue *StructValue `protobuf:"bytes,8,opt,name=struct_value,json=structValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,9,opt,name=list_value,json=listValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_IntegerValue) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_TimestampValue) isValue_Kind() {}

func (*Value_BytesValue) isValue_Kind() {}

func (*Value_StructValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

// An error in the context of a data pipeline.
type Error struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The error message. If non empty, then the error to be "valid" and
	// if empty the error is ignored as if a success (due to proto3 empty
	// semantics).
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Additional error details for specific Redpanda Connect behavior.
	// If one of these fields is set, then message must be non-empty.
	//
	// Types that are valid to be assigned to Detail:
	//
	//	*Error_Backoff
	//	*Error_NotConnected_
	//	*Error_EndOfInput_
	Detail        isError_Detail `protobuf_oneof:"detail"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetail() isError_Detail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *Error) GetBackoff() *durationpb.Duration {
	if x != nil {
		if x, ok := x.Detail.(*Error_Backoff); ok {
			return x.Backoff
		}
	}
	return nil
}

func (x *Error) GetNotConnected() *Error_NotConnected {
	if x != nil {
		if x, ok := x.Detail.(*Error_NotConnected_); ok {
			return x.NotConnected
		}
	}
	return nil
}

func (x *Error) GetEndOfInput() *Error_EndOfInput {
	if x != nil {
		if x, ok := x.Detail.(*Error_EndOfInput_); ok {
			return x.EndOfInput
		}
	}
	return nil
}

type isError_Detail interface {
	isError_Detail()
}

type Error_Backoff struct {
	// BackOff is an error that plugins can optionally wrap another error with which instructs upstream components to wait for a specified period of time before retrying the errored call.
	//
	// Only suppported by Connect methods in the Input and Output services.
	Backoff *durationpb.Duration `protobuf:"bytes,2,opt,name=backoff,proto3,oneof"`
}

type Error_NotConnected_ struct {
	NotConnected *Error_NotConnected `protobuf:"bytes,3,opt,name=not_connected,json=notConnected,proto3,oneof"`
}

type Error_EndOfInput_ struct {
	EndOfInput *Error_EndOfInput `protobuf:"bytes,4,opt,name=end_of_input,json=endOfInput,proto3,oneof"`
}

func (*Error_Backoff) isError_Detail() {}

func (*Error_NotConnected_) isError_Detail() {}

func (*Error_EndOfInput_) isError_Detail() {}

// Message represents a piece of data or an event that flows through the runtime.
type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*Message_Bytes
	//	*Message_Structured
	Payload       isMessage_Payload `protobuf_oneof:"payload"`
	Metadata      *StructValue      `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Error         *Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetPayload() isMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Message) GetBytes() []byte {
	if x != nil {
		if x, ok := x.Payload.(*Message_Bytes); ok {
			return x.Bytes
		}
	}
	return nil
}

func (x *Message) GetStructured() *Value {
	if x != nil {
		if x, ok := x.Payload.(*Message_Structured); ok {
			return x.Structured
		}
	}
	return nil
}

func (x *Message) GetMetadata() *StructValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Message) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Bytes struct {
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3,oneof"`
}

type Message_Structured struct {
	Structured *Value `protobuf:"bytes,2,opt,name=structured,proto3,oneof"`
}

func (*Message_Bytes) isMessage_Payload() {}

func (*Message_Structured) isMessage_Payload() {}

type MessageBatch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageBatch) Reset() {
	*x = MessageBatch{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBatch) ProtoMessage() {}

func (x *MessageBatch) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBatch.ProtoReflect.Descriptor instead.
func (*MessageBatch) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{5}
}

func (x *MessageBatch) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// NotConnected is returned by inputs and outputs when their Read or
// Write methods are called and the connection that they maintain is lost.
// This error prompts the upstream component to call Connect until the
// connection is re-established.
type Error_NotConnected struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error_NotConnected) Reset() {
	*x = Error_NotConnected{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error_NotConnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error_NotConnected) ProtoMessage() {}

func (x *Error_NotConnected) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error_NotConnected.ProtoReflect.Descriptor instead.
func (*Error_NotConnected) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{3, 0}
}

// EndOfInput is returned by inputs that have exhausted their source of
// data to the point where subsequent Read calls will be ineffective. This
// error prompts the upstream component to gracefully terminate the
// pipeline.
type Error_EndOfInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error_EndOfInput) Reset() {
	*x = Error_EndOfInput{}
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error_EndOfInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error_EndOfInput) ProtoMessage() {}

func (x *Error_EndOfInput) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_runtime_v1alpha1_message_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error_EndOfInput.ProtoReflect.Descriptor instead.
func (*Error_EndOfInput) Descriptor() ([]byte, []int) {
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP(), []int{3, 1}
}

var File_redpanda_runtime_v1alpha1_message_proto protoreflect.FileDescriptor

const file_redpanda_runtime_v1alpha1_message_proto_rawDesc = "" +
	"\n" +
	"'redpanda/runtime/v1alpha1/message.proto\x12\x19redpanda.runtime.v1alpha1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"\xb6\x01\n" +
	"\vStructValue\x12J\n" +
	"\x06fields\x18\x01 \x03(\v22.redpanda.runtime.v1alpha1.StructValue.FieldsEntryR\x06fields\x1a[\n" +
	"\vFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x126\n" +
	"\x05value\x18\x02 \x01(\v2 .redpanda.runtime.v1alpha1.ValueR\x05value:\x028\x01\"E\n" +
	"\tListValue\x128\n" +
	"\x06values\x18\x01 \x03(\v2 .redpanda.runtime.v1alpha1.ValueR\x06values\"\xe6\x03\n" +
	"\x05Value\x12E\n" +
	"\n" +
	"null_value\x18\x01 \x01(\x0e2$.redpanda.runtime.v1alpha1.NullValueH\x00R\tnullValue\x12#\n" +
	"\fstring_value\x18\x02 \x01(\tH\x00R\vstringValue\x12%\n" +
	"\rinteger_value\x18\x03 \x01(\x03H\x00R\fintegerValue\x12#\n" +
	"\fdouble_value\x18\x04 \x01(\x01H\x00R\vdoubleValue\x12\x1f\n" +
	"\n" +
	"bool_value\x18\x05 \x01(\bH\x00R\tboolValue\x12E\n" +
	"\x0ftimestamp_value\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampH\x00R\x0etimestampValue\x12!\n" +
	"\vbytes_value\x18\a \x01(\fH\x00R\n" +
	"bytesValue\x12K\n" +
	"\fstruct_value\x18\b \x01(\v2&.redpanda.runtime.v1alpha1.StructValueH\x00R\vstructValue\x12E\n" +
	"\n" +
	"list_value\x18\t \x01(\v2$.redpanda.runtime.v1alpha1.ListValueH\x00R\tlistValueB\x06\n" +
	"\x04kind\"\xa7\x02\n" +
	"\x05Error\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x125\n" +
	"\abackoff\x18\x02 \x01(\v2\x19.google.protobuf.DurationH\x00R\abackoff\x12T\n" +
	"\rnot_connected\x18\x03 \x01(\v2-.redpanda.runtime.v1alpha1.Error.NotConnectedH\x00R\fnotConnected\x12O\n" +
	"\fend_of_input\x18\x04 \x01(\v2+.redpanda.runtime.v1alpha1.Error.EndOfInputH\x00R\n" +
	"endOfInput\x1a\x0e\n" +
	"\fNotConnected\x1a\f\n" +
	"\n" +
	"EndOfInputB\b\n" +
	"\x06detail\"\xec\x01\n" +
	"\aMessage\x12\x16\n" +
	"\x05bytes\x18\x01 \x01(\fH\x00R\x05bytes\x12B\n" +
	"\n" +
	"structured\x18\x02 \x01(\v2 .redpanda.runtime.v1alpha1.ValueH\x00R\n" +
	"structured\x12B\n" +
	"\bmetadata\x18\x03 \x01(\v2&.redpanda.runtime.v1alpha1.StructValueR\bmetadata\x126\n" +
	"\x05error\x18\x04 \x01(\v2 .redpanda.runtime.v1alpha1.ErrorR\x05errorB\t\n" +
	"\apayload\"N\n" +
	"\fMessageBatch\x12>\n" +
	"\bmessages\x18\x01 \x03(\v2\".redpanda.runtime.v1alpha1.MessageR\bmessages*\x1b\n" +
	"\tNullValue\x12\x0e\n" +
	"\n" +
	"NULL_VALUE\x10\x00BBZ@github.com/redpanda-data/connect/v4/internal/rpcplugin/runtimepbb\x06proto3"

var (
	file_redpanda_runtime_v1alpha1_message_proto_rawDescOnce sync.Once
	file_redpanda_runtime_v1alpha1_message_proto_rawDescData []byte
)

func file_redpanda_runtime_v1alpha1_message_proto_rawDescGZIP() []byte {
	file_redpanda_runtime_v1alpha1_message_proto_rawDescOnce.Do(func() {
		file_redpanda_runtime_v1alpha1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_redpanda_runtime_v1alpha1_message_proto_rawDesc), len(file_redpanda_runtime_v1alpha1_message_proto_rawDesc)))
	})
	return file_redpanda_runtime_v1alpha1_message_proto_rawDescData
}

var file_redpanda_runtime_v1alpha1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_redpanda_runtime_v1alpha1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_redpanda_runtime_v1alpha1_message_proto_goTypes = []any{
	(NullValue)(0),                // 0: redpanda.runtime.v1alpha1.NullValue
	(*StructValue)(nil),           // 1: redpanda.runtime.v1alpha1.StructValue
	(*ListValue)(nil),             // 2: redpanda.runtime.v1alpha1.ListValue
	(*Value)(nil),                 // 3: redpanda.runtime.v1alpha1.Value
	(*Error)(nil),                 // 4: redpanda.runtime.v1alpha1.Error
	(*Message)(nil),               // 5: redpanda.runtime.v1alpha1.Message
	(*MessageBatch)(nil),          // 6: redpanda.runtime.v1alpha1.MessageBatch
	nil,                           // 7: redpanda.runtime.v1alpha1.StructValue.FieldsEntry
	(*Error_NotConnected)(nil),    // 8: redpanda.runtime.v1alpha1.Error.NotConnected
	(*Error_EndOfInput)(nil),      // 9: redpanda.runtime.v1alpha1.Error.EndOfInput
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 11: google.protobuf.Duration
}
var file_redpanda_runtime_v1alpha1_message_proto_depIdxs = []int32{
	7,  // 0: redpanda.runtime.v1alpha1.StructValue.fields:type_name -> redpanda.runtime.v1alpha1.StructValue.FieldsEntry
	3,  // 1: redpanda.runtime.v1alpha1.ListValue.values:type_name -> redpanda.runtime.v1alpha1.Value
	0,  // 2: redpanda.runtime.v1alpha1.Value.null_value:type_name -> redpanda.runtime.v1alpha1.NullValue
	10, // 3: redpanda.runtime.v1alpha1.Value.timestamp_value:type_name -> google.protobuf.Timestamp
	1,  // 4: redpanda.runtime.v1alpha1.Value.struct_value:type_name -> redpanda.runtime.v1alpha1.StructValue
	2,  // 5: redpanda.runtime.v1alpha1.Value.list_value:type_name -> redpanda.runtime.v1alpha1.ListValue
	11, // 6: redpanda.runtime.v1alpha1.Error.backoff:type_name -> google.protobuf.Duration
	8,  // 7: redpanda.runtime.v1alpha1.Error.not_connected:type_name -> redpanda.runtime.v1alpha1.Error.NotConnected
	9,  // 8: redpanda.runtime.v1alpha1.Error.end_of_input:type_name -> redpanda.runtime.v1alpha1.Error.EndOfInput
	3,  // 9: redpanda.runtime.v1alpha1.Message.structured:type_name -> redpanda.runtime.v1alpha1.Value
	1,  // 10: redpanda.runtime.v1alpha1.Message.metadata:type_name -> redpanda.runtime.v1alpha1.StructValue
	4,  // 11: redpanda.runtime.v1alpha1.Message.error:type_name -> redpanda.runtime.v1alpha1.Error
	5,  // 12: redpanda.runtime.v1alpha1.MessageBatch.messages:type_name -> redpanda.runtime.v1alpha1.Message
	3,  // 13: redpanda.runtime.v1alpha1.StructValue.FieldsEntry.value:type_name -> redpanda.runtime.v1alpha1.Value
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_redpanda_runtime_v1alpha1_message_proto_init() }
func file_redpanda_runtime_v1alpha1_message_proto_init() {
	if File_redpanda_runtime_v1alpha1_message_proto != nil {
		return
	}
	file_redpanda_runtime_v1alpha1_message_proto_msgTypes[2].OneofWrappers = []any{
		(*Value_NullValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
	file_redpanda_runtime_v1alpha1_message_proto_msgTypes[3].OneofWrappers = []any{
		(*Error_Backoff)(nil),
		(*Error_NotConnected_)(nil),
		(*Error_EndOfInput_)(nil),
	}
	file_redpanda_runtime_v1alpha1_message_proto_msgTypes[4].OneofWrappers = []any{
		(*Message_Bytes)(nil),
		(*Message_Structured)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_redpanda_runtime_v1alpha1_message_proto_rawDesc), len(file_redpanda_runtime_v1alpha1_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_runtime_v1alpha1_message_proto_goTypes,
		DependencyIndexes: file_redpanda_runtime_v1alpha1_message_proto_depIdxs,
		EnumInfos:         file_redpanda_runtime_v1alpha1_message_proto_enumTypes,
		MessageInfos:      file_redpanda_runtime_v1alpha1_message_proto_msgTypes,
	}.Build()
	File_redpanda_runtime_v1alpha1_message_proto = out.File
	file_redpanda_runtime_v1alpha1_message_proto_goTypes = nil
	file_redpanda_runtime_v1alpha1_message_proto_depIdxs = nil
}
