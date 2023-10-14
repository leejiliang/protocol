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
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: rpc/io.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Query an ingress info from an ingress ID or stream key
type GetIngressInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressId string `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	StreamKey string `protobuf:"bytes,2,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
}

func (x *GetIngressInfoRequest) Reset() {
	*x = GetIngressInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngressInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngressInfoRequest) ProtoMessage() {}

func (x *GetIngressInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngressInfoRequest.ProtoReflect.Descriptor instead.
func (*GetIngressInfoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{0}
}

func (x *GetIngressInfoRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *GetIngressInfoRequest) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

type GetIngressInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *livekit.IngressInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Token string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl string               `protobuf:"bytes,3,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
}

func (x *GetIngressInfoResponse) Reset() {
	*x = GetIngressInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngressInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngressInfoResponse) ProtoMessage() {}

func (x *GetIngressInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngressInfoResponse.ProtoReflect.Descriptor instead.
func (*GetIngressInfoResponse) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{1}
}

func (x *GetIngressInfoResponse) GetInfo() *livekit.IngressInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GetIngressInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetIngressInfoResponse) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

// Request to store an update to the ingress state ingress -> service
type UpdateIngressStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressId string                `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	State     *livekit.IngressState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UpdateIngressStateRequest) Reset() {
	*x = UpdateIngressStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIngressStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIngressStateRequest) ProtoMessage() {}

func (x *UpdateIngressStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIngressStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateIngressStateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_io_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateIngressStateRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *UpdateIngressStateRequest) GetState() *livekit.IngressState {
	if x != nil {
		return x.State
	}
	return nil
}

var File_rpc_io_proto protoreflect.FileDescriptor

var file_rpc_io_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4b, 0x65, 0x79, 0x22, 0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x77, 0x73, 0x55, 0x72, 0x6c, 0x22, 0x67, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xe2,
	0x01, 0x0a, 0x06, 0x49, 0x4f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_io_proto_rawDescOnce sync.Once
	file_rpc_io_proto_rawDescData = file_rpc_io_proto_rawDesc
)

func file_rpc_io_proto_rawDescGZIP() []byte {
	file_rpc_io_proto_rawDescOnce.Do(func() {
		file_rpc_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_io_proto_rawDescData)
	})
	return file_rpc_io_proto_rawDescData
}

var file_rpc_io_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_io_proto_goTypes = []interface{}{
	(*GetIngressInfoRequest)(nil),     // 0: rpc.GetIngressInfoRequest
	(*GetIngressInfoResponse)(nil),    // 1: rpc.GetIngressInfoResponse
	(*UpdateIngressStateRequest)(nil), // 2: rpc.UpdateIngressStateRequest
	(*livekit.IngressInfo)(nil),       // 3: livekit.IngressInfo
	(*livekit.IngressState)(nil),      // 4: livekit.IngressState
	(*livekit.EgressInfo)(nil),        // 5: livekit.EgressInfo
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_rpc_io_proto_depIdxs = []int32{
	3, // 0: rpc.GetIngressInfoResponse.info:type_name -> livekit.IngressInfo
	4, // 1: rpc.UpdateIngressStateRequest.state:type_name -> livekit.IngressState
	5, // 2: rpc.IOInfo.UpdateEgressInfo:input_type -> livekit.EgressInfo
	0, // 3: rpc.IOInfo.GetIngressInfo:input_type -> rpc.GetIngressInfoRequest
	2, // 4: rpc.IOInfo.UpdateIngressState:input_type -> rpc.UpdateIngressStateRequest
	6, // 5: rpc.IOInfo.UpdateEgressInfo:output_type -> google.protobuf.Empty
	1, // 6: rpc.IOInfo.GetIngressInfo:output_type -> rpc.GetIngressInfoResponse
	6, // 7: rpc.IOInfo.UpdateIngressState:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_io_proto_init() }
func file_rpc_io_proto_init() {
	if File_rpc_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngressInfoRequest); i {
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
		file_rpc_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngressInfoResponse); i {
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
		file_rpc_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIngressStateRequest); i {
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
			RawDescriptor: file_rpc_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_io_proto_goTypes,
		DependencyIndexes: file_rpc_io_proto_depIdxs,
		MessageInfos:      file_rpc_io_proto_msgTypes,
	}.Build()
	File_rpc_io_proto = out.File
	file_rpc_io_proto_rawDesc = nil
	file_rpc_io_proto_goTypes = nil
	file_rpc_io_proto_depIdxs = nil
}
