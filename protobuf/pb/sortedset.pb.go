// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: protobuf/sortedset.proto

package pb

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

type ZAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Score  int64  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Member string `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *ZAddRequest) Reset() {
	*x = ZAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_sortedset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAddRequest) ProtoMessage() {}

func (x *ZAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_sortedset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAddRequest.ProtoReflect.Descriptor instead.
func (*ZAddRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_sortedset_proto_rawDescGZIP(), []int{0}
}

func (x *ZAddRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ZAddRequest) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ZAddRequest) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

type ZAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ZAddResponse) Reset() {
	*x = ZAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_sortedset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZAddResponse) ProtoMessage() {}

func (x *ZAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_sortedset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZAddResponse.ProtoReflect.Descriptor instead.
func (*ZAddResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_sortedset_proto_rawDescGZIP(), []int{1}
}

type ZRankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Member string `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *ZRankRequest) Reset() {
	*x = ZRankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_sortedset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZRankRequest) ProtoMessage() {}

func (x *ZRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_sortedset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZRankRequest.ProtoReflect.Descriptor instead.
func (*ZRankRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_sortedset_proto_rawDescGZIP(), []int{2}
}

func (x *ZRankRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ZRankRequest) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

type ZRankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank int64 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *ZRankResponse) Reset() {
	*x = ZRankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_sortedset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZRankResponse) ProtoMessage() {}

func (x *ZRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_sortedset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZRankResponse.ProtoReflect.Descriptor instead.
func (*ZRankResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_sortedset_proto_rawDescGZIP(), []int{3}
}

func (x *ZRankResponse) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

var File_protobuf_sortedset_proto protoreflect.FileDescriptor

var file_protobuf_sortedset_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0b, 0x5a, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x5a, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x5a, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0d, 0x5a, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x32, 0x5c, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x5a, 0x41, 0x64, 0x64, 0x12, 0x0c, 0x2e,
	0x5a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x5a, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05,
	0x5a, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x0d, 0x2e, 0x5a, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x5a, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x75, 0x70, 0x61, 0x6e, 0x6f, 0x6a, 0x69, 0x73,
	0x61, 0x6e, 0x2f, 0x6d, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_sortedset_proto_rawDescOnce sync.Once
	file_protobuf_sortedset_proto_rawDescData = file_protobuf_sortedset_proto_rawDesc
)

func file_protobuf_sortedset_proto_rawDescGZIP() []byte {
	file_protobuf_sortedset_proto_rawDescOnce.Do(func() {
		file_protobuf_sortedset_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_sortedset_proto_rawDescData)
	})
	return file_protobuf_sortedset_proto_rawDescData
}

var file_protobuf_sortedset_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_sortedset_proto_goTypes = []interface{}{
	(*ZAddRequest)(nil),   // 0: ZAddRequest
	(*ZAddResponse)(nil),  // 1: ZAddResponse
	(*ZRankRequest)(nil),  // 2: ZRankRequest
	(*ZRankResponse)(nil), // 3: ZRankResponse
}
var file_protobuf_sortedset_proto_depIdxs = []int32{
	0, // 0: SortedSet.ZAdd:input_type -> ZAddRequest
	2, // 1: SortedSet.ZRank:input_type -> ZRankRequest
	1, // 2: SortedSet.ZAdd:output_type -> ZAddResponse
	3, // 3: SortedSet.ZRank:output_type -> ZRankResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_sortedset_proto_init() }
func file_protobuf_sortedset_proto_init() {
	if File_protobuf_sortedset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_sortedset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAddRequest); i {
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
		file_protobuf_sortedset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZAddResponse); i {
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
		file_protobuf_sortedset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZRankRequest); i {
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
		file_protobuf_sortedset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZRankResponse); i {
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
			RawDescriptor: file_protobuf_sortedset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_sortedset_proto_goTypes,
		DependencyIndexes: file_protobuf_sortedset_proto_depIdxs,
		MessageInfos:      file_protobuf_sortedset_proto_msgTypes,
	}.Build()
	File_protobuf_sortedset_proto = out.File
	file_protobuf_sortedset_proto_rawDesc = nil
	file_protobuf_sortedset_proto_goTypes = nil
	file_protobuf_sortedset_proto_depIdxs = nil
}