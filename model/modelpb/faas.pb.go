// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: faas.proto

package modelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Faas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ColdStart        *bool  `protobuf:"varint,2,opt,name=cold_start,json=coldStart,proto3,oneof" json:"cold_start,omitempty"`
	Execution        string `protobuf:"bytes,3,opt,name=execution,proto3" json:"execution,omitempty"`
	TriggerType      string `protobuf:"bytes,4,opt,name=trigger_type,json=triggerType,proto3" json:"trigger_type,omitempty"`
	TriggerRequestId string `protobuf:"bytes,5,opt,name=trigger_request_id,json=triggerRequestId,proto3" json:"trigger_request_id,omitempty"`
	Name             string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Version          string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Faas) Reset() {
	*x = Faas{}
	mi := &file_faas_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Faas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Faas) ProtoMessage() {}

func (x *Faas) ProtoReflect() protoreflect.Message {
	mi := &file_faas_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Faas.ProtoReflect.Descriptor instead.
func (*Faas) Descriptor() ([]byte, []int) {
	return file_faas_proto_rawDescGZIP(), []int{0}
}

func (x *Faas) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Faas) GetColdStart() bool {
	if x != nil && x.ColdStart != nil {
		return *x.ColdStart
	}
	return false
}

func (x *Faas) GetExecution() string {
	if x != nil {
		return x.Execution
	}
	return ""
}

func (x *Faas) GetTriggerType() string {
	if x != nil {
		return x.TriggerType
	}
	return ""
}

func (x *Faas) GetTriggerRequestId() string {
	if x != nil {
		return x.TriggerRequestId
	}
	return ""
}

func (x *Faas) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Faas) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_faas_proto protoreflect.FileDescriptor

var file_faas_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6c,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0xe6, 0x01, 0x0a,
	0x04, 0x46, 0x61, 0x61, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6c, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_faas_proto_rawDescOnce sync.Once
	file_faas_proto_rawDescData = file_faas_proto_rawDesc
)

func file_faas_proto_rawDescGZIP() []byte {
	file_faas_proto_rawDescOnce.Do(func() {
		file_faas_proto_rawDescData = protoimpl.X.CompressGZIP(file_faas_proto_rawDescData)
	})
	return file_faas_proto_rawDescData
}

var file_faas_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_faas_proto_goTypes = []any{
	(*Faas)(nil), // 0: elastic.apm.v1.Faas
}
var file_faas_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_faas_proto_init() }
func file_faas_proto_init() {
	if File_faas_proto != nil {
		return
	}
	file_faas_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_faas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_faas_proto_goTypes,
		DependencyIndexes: file_faas_proto_depIdxs,
		MessageInfos:      file_faas_proto_msgTypes,
	}.Build()
	File_faas_proto = out.File
	file_faas_proto_rawDesc = nil
	file_faas_proto_goTypes = nil
	file_faas_proto_depIdxs = nil
}
