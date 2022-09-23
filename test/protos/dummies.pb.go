// Copyright 2015 gRPC authors.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: test/protos/dummies.proto

package protos

import (
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

type ErrorDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category      string            `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Code          string            `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	CorrelationId string            `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Status        string            `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Message       string            `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Cause         string            `protobuf:"bytes,6,opt,name=cause,proto3" json:"cause,omitempty"`
	StackTrace    string            `protobuf:"bytes,7,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Details       map[string]string `protobuf:"bytes,8,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ErrorDescription) Reset() {
	*x = ErrorDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDescription) ProtoMessage() {}

func (x *ErrorDescription) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDescription.ProtoReflect.Descriptor instead.
func (*ErrorDescription) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDescription) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ErrorDescription) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ErrorDescription) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ErrorDescription) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ErrorDescription) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorDescription) GetCause() string {
	if x != nil {
		return x.Cause
	}
	return ""
}

func (x *ErrorDescription) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *ErrorDescription) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

type PagingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skip  int64 `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take  int64 `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	Total bool  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PagingParams) Reset() {
	*x = PagingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingParams) ProtoMessage() {}

func (x *PagingParams) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingParams.ProtoReflect.Descriptor instead.
func (*PagingParams) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{1}
}

func (x *PagingParams) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *PagingParams) GetTake() int64 {
	if x != nil {
		return x.Take
	}
	return 0
}

func (x *PagingParams) GetTotal() bool {
	if x != nil {
		return x.Total
	}
	return false
}

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{2}
}

func (x *Dummy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dummy) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Dummy) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DummiesPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*Dummy `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DummiesPage) Reset() {
	*x = DummiesPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummiesPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummiesPage) ProtoMessage() {}

func (x *DummiesPage) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummiesPage.ProtoReflect.Descriptor instead.
func (*DummiesPage) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{3}
}

func (x *DummiesPage) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DummiesPage) GetData() []*Dummy {
	if x != nil {
		return x.Data
	}
	return nil
}

// The request message containing the page request.
type DummiesPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string            `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Filter        map[string]string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Paging        *PagingParams     `protobuf:"bytes,3,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *DummiesPageRequest) Reset() {
	*x = DummiesPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummiesPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummiesPageRequest) ProtoMessage() {}

func (x *DummiesPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummiesPageRequest.ProtoReflect.Descriptor instead.
func (*DummiesPageRequest) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{4}
}

func (x *DummiesPageRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DummiesPageRequest) GetFilter() map[string]string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *DummiesPageRequest) GetPaging() *PagingParams {
	if x != nil {
		return x.Paging
	}
	return nil
}

// The request message containing the object id request.
type DummyIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	DummyId       string `protobuf:"bytes,2,opt,name=dummy_id,json=dummyId,proto3" json:"dummy_id,omitempty"`
}

func (x *DummyIdRequest) Reset() {
	*x = DummyIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyIdRequest) ProtoMessage() {}

func (x *DummyIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyIdRequest.ProtoReflect.Descriptor instead.
func (*DummyIdRequest) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{5}
}

func (x *DummyIdRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DummyIdRequest) GetDummyId() string {
	if x != nil {
		return x.DummyId
	}
	return ""
}

// The request message containing the object value request.
type DummyObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Dummy         *Dummy `protobuf:"bytes,2,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *DummyObjectRequest) Reset() {
	*x = DummyObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_dummies_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyObjectRequest) ProtoMessage() {}

func (x *DummyObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_dummies_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyObjectRequest.ProtoReflect.Descriptor instead.
func (*DummyObjectRequest) Descriptor() ([]byte, []int) {
	return file_test_protos_dummies_proto_rawDescGZIP(), []int{6}
}

func (x *DummyObjectRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DummyObjectRequest) GetDummy() *Dummy {
	if x != nil {
		return x.Dummy
	}
	return nil
}

var File_test_protos_dummies_proto protoreflect.FileDescriptor

var file_test_protos_dummies_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x75,
	0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x75, 0x6d,
	0x6d, 0x69, 0x65, 0x73, 0x22, 0xd0, 0x02, 0x0a, 0x10, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x75, 0x6d,
	0x6d, 0x69, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x61, 0x6b, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x0b, 0x44, 0x75,
	0x6d, 0x6d, 0x69, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xe6, 0x01, 0x0a, 0x12, 0x44, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d,
	0x69, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x0e,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x49, 0x64,
	0x22, 0x61, 0x0a, 0x12, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64,
	0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x05, 0x64, 0x75,
	0x6d, 0x6d, 0x79, 0x32, 0xca, 0x02, 0x0a, 0x07, 0x44, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x12,
	0x42, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x12, 0x1b,
	0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75,
	0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x50, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x17, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73,
	0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x12, 0x1b, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x12, 0x1b, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f,
	0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x17, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x00,
	0x42, 0x3b, 0x0a, 0x19, 0x70, 0x69, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x42, 0x0c, 0x44,
	0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x08, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_protos_dummies_proto_rawDescOnce sync.Once
	file_test_protos_dummies_proto_rawDescData = file_test_protos_dummies_proto_rawDesc
)

func file_test_protos_dummies_proto_rawDescGZIP() []byte {
	file_test_protos_dummies_proto_rawDescOnce.Do(func() {
		file_test_protos_dummies_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_protos_dummies_proto_rawDescData)
	})
	return file_test_protos_dummies_proto_rawDescData
}

var file_test_protos_dummies_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_test_protos_dummies_proto_goTypes = []interface{}{
	(*ErrorDescription)(nil),   // 0: dummies.ErrorDescription
	(*PagingParams)(nil),       // 1: dummies.PagingParams
	(*Dummy)(nil),              // 2: dummies.Dummy
	(*DummiesPage)(nil),        // 3: dummies.DummiesPage
	(*DummiesPageRequest)(nil), // 4: dummies.DummiesPageRequest
	(*DummyIdRequest)(nil),     // 5: dummies.DummyIdRequest
	(*DummyObjectRequest)(nil), // 6: dummies.DummyObjectRequest
	nil,                        // 7: dummies.ErrorDescription.DetailsEntry
	nil,                        // 8: dummies.DummiesPageRequest.FilterEntry
}
var file_test_protos_dummies_proto_depIdxs = []int32{
	7,  // 0: dummies.ErrorDescription.details:type_name -> dummies.ErrorDescription.DetailsEntry
	2,  // 1: dummies.DummiesPage.data:type_name -> dummies.Dummy
	8,  // 2: dummies.DummiesPageRequest.filter:type_name -> dummies.DummiesPageRequest.FilterEntry
	1,  // 3: dummies.DummiesPageRequest.paging:type_name -> dummies.PagingParams
	2,  // 4: dummies.DummyObjectRequest.dummy:type_name -> dummies.Dummy
	4,  // 5: dummies.Dummies.get_dummies:input_type -> dummies.DummiesPageRequest
	5,  // 6: dummies.Dummies.get_dummy_by_id:input_type -> dummies.DummyIdRequest
	6,  // 7: dummies.Dummies.create_dummy:input_type -> dummies.DummyObjectRequest
	6,  // 8: dummies.Dummies.update_dummy:input_type -> dummies.DummyObjectRequest
	5,  // 9: dummies.Dummies.delete_dummy_by_id:input_type -> dummies.DummyIdRequest
	3,  // 10: dummies.Dummies.get_dummies:output_type -> dummies.DummiesPage
	2,  // 11: dummies.Dummies.get_dummy_by_id:output_type -> dummies.Dummy
	2,  // 12: dummies.Dummies.create_dummy:output_type -> dummies.Dummy
	2,  // 13: dummies.Dummies.update_dummy:output_type -> dummies.Dummy
	2,  // 14: dummies.Dummies.delete_dummy_by_id:output_type -> dummies.Dummy
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_test_protos_dummies_proto_init() }
func file_test_protos_dummies_proto_init() {
	if File_test_protos_dummies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_protos_dummies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDescription); i {
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
		file_test_protos_dummies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingParams); i {
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
		file_test_protos_dummies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dummy); i {
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
		file_test_protos_dummies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummiesPage); i {
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
		file_test_protos_dummies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummiesPageRequest); i {
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
		file_test_protos_dummies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyIdRequest); i {
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
		file_test_protos_dummies_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyObjectRequest); i {
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
			RawDescriptor: file_test_protos_dummies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_protos_dummies_proto_goTypes,
		DependencyIndexes: file_test_protos_dummies_proto_depIdxs,
		MessageInfos:      file_test_protos_dummies_proto_msgTypes,
	}.Build()
	File_test_protos_dummies_proto = out.File
	file_test_protos_dummies_proto_rawDesc = nil
	file_test_protos_dummies_proto_goTypes = nil
	file_test_protos_dummies_proto_depIdxs = nil
}
