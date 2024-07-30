// Copyright 2023 LiveKit, Inc.
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
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: rpc/analytics.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rpc_analytics_proto protoreflect.FileDescriptor

var file_rpc_analytics_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf5, 0x01, 0x0a, 0x18, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x42, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x17, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x44, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4f, 0x0a, 0x14, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x21, 0x5a, 0x1f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_rpc_analytics_proto_goTypes = []any{
	(*livekit.AnalyticsStats)(nil),     // 0: livekit.AnalyticsStats
	(*livekit.AnalyticsEvents)(nil),    // 1: livekit.AnalyticsEvents
	(*livekit.AnalyticsNodeRooms)(nil), // 2: livekit.AnalyticsNodeRooms
	(*emptypb.Empty)(nil),              // 3: google.protobuf.Empty
}
var file_rpc_analytics_proto_depIdxs = []int32{
	0, // 0: livekit.AnalyticsRecorderService.IngestStats:input_type -> livekit.AnalyticsStats
	1, // 1: livekit.AnalyticsRecorderService.IngestEvents:input_type -> livekit.AnalyticsEvents
	2, // 2: livekit.AnalyticsRecorderService.IngestNodeRoomStates:input_type -> livekit.AnalyticsNodeRooms
	3, // 3: livekit.AnalyticsRecorderService.IngestStats:output_type -> google.protobuf.Empty
	3, // 4: livekit.AnalyticsRecorderService.IngestEvents:output_type -> google.protobuf.Empty
	3, // 5: livekit.AnalyticsRecorderService.IngestNodeRoomStates:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_analytics_proto_init() }
func file_rpc_analytics_proto_init() {
	if File_rpc_analytics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_analytics_proto_goTypes,
		DependencyIndexes: file_rpc_analytics_proto_depIdxs,
	}.Build()
	File_rpc_analytics_proto = out.File
	file_rpc_analytics_proto_rawDesc = nil
	file_rpc_analytics_proto_goTypes = nil
	file_rpc_analytics_proto_depIdxs = nil
}
