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
// source: experience.proto

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

type UserExperience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CumulativeLayoutShift float64          `protobuf:"fixed64,1,opt,name=cumulative_layout_shift,json=cumulativeLayoutShift,proto3" json:"cumulative_layout_shift,omitempty"`
	FirstInputDelay       float64          `protobuf:"fixed64,2,opt,name=first_input_delay,json=firstInputDelay,proto3" json:"first_input_delay,omitempty"`
	TotalBlockingTime     float64          `protobuf:"fixed64,3,opt,name=total_blocking_time,json=totalBlockingTime,proto3" json:"total_blocking_time,omitempty"`
	LongTask              *LongtaskMetrics `protobuf:"bytes,4,opt,name=long_task,json=longTask,proto3" json:"long_task,omitempty"`
}

func (x *UserExperience) Reset() {
	*x = UserExperience{}
	mi := &file_experience_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserExperience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserExperience) ProtoMessage() {}

func (x *UserExperience) ProtoReflect() protoreflect.Message {
	mi := &file_experience_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserExperience.ProtoReflect.Descriptor instead.
func (*UserExperience) Descriptor() ([]byte, []int) {
	return file_experience_proto_rawDescGZIP(), []int{0}
}

func (x *UserExperience) GetCumulativeLayoutShift() float64 {
	if x != nil {
		return x.CumulativeLayoutShift
	}
	return 0
}

func (x *UserExperience) GetFirstInputDelay() float64 {
	if x != nil {
		return x.FirstInputDelay
	}
	return 0
}

func (x *UserExperience) GetTotalBlockingTime() float64 {
	if x != nil {
		return x.TotalBlockingTime
	}
	return 0
}

func (x *UserExperience) GetLongTask() *LongtaskMetrics {
	if x != nil {
		return x.LongTask
	}
	return nil
}

type LongtaskMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Sum   float64 `protobuf:"fixed64,2,opt,name=sum,proto3" json:"sum,omitempty"`
	Max   float64 `protobuf:"fixed64,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *LongtaskMetrics) Reset() {
	*x = LongtaskMetrics{}
	mi := &file_experience_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LongtaskMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongtaskMetrics) ProtoMessage() {}

func (x *LongtaskMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_experience_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongtaskMetrics.ProtoReflect.Descriptor instead.
func (*LongtaskMetrics) Descriptor() ([]byte, []int) {
	return file_experience_proto_rawDescGZIP(), []int{1}
}

func (x *LongtaskMetrics) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *LongtaskMetrics) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *LongtaskMetrics) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_experience_proto protoreflect.FileDescriptor

var file_experience_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e,
	0x76, 0x31, 0x22, 0xe2, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x68, 0x69, 0x66, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x53, 0x68, 0x69, 0x66, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x6e, 0x67, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x08, 0x6c,
	0x6f, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x4b, 0x0a, 0x0f, 0x4c, 0x6f, 0x6e, 0x67, 0x74,
	0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x73,
	0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_experience_proto_rawDescOnce sync.Once
	file_experience_proto_rawDescData = file_experience_proto_rawDesc
)

func file_experience_proto_rawDescGZIP() []byte {
	file_experience_proto_rawDescOnce.Do(func() {
		file_experience_proto_rawDescData = protoimpl.X.CompressGZIP(file_experience_proto_rawDescData)
	})
	return file_experience_proto_rawDescData
}

var file_experience_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_experience_proto_goTypes = []any{
	(*UserExperience)(nil),  // 0: elastic.apm.v1.UserExperience
	(*LongtaskMetrics)(nil), // 1: elastic.apm.v1.LongtaskMetrics
}
var file_experience_proto_depIdxs = []int32{
	1, // 0: elastic.apm.v1.UserExperience.long_task:type_name -> elastic.apm.v1.LongtaskMetrics
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_experience_proto_init() }
func file_experience_proto_init() {
	if File_experience_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_experience_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_experience_proto_goTypes,
		DependencyIndexes: file_experience_proto_depIdxs,
		MessageInfos:      file_experience_proto_msgTypes,
	}.Build()
	File_experience_proto = out.File
	file_experience_proto_rawDesc = nil
	file_experience_proto_goTypes = nil
	file_experience_proto_depIdxs = nil
}
