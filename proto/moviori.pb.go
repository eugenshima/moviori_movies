// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: moviori.proto

package moviori_movies

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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moviori_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_moviori_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_moviori_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type FindMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *FindMovieRequest) Reset() {
	*x = FindMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moviori_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMovieRequest) ProtoMessage() {}

func (x *FindMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moviori_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMovieRequest.ProtoReflect.Descriptor instead.
func (*FindMovieRequest) Descriptor() ([]byte, []int) {
	return file_moviori_proto_rawDescGZIP(), []int{1}
}

func (x *FindMovieRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type FindMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieInfo []byte `protobuf:"bytes,1,opt,name=MovieInfo,proto3" json:"MovieInfo,omitempty"`
}

func (x *FindMovieResponse) Reset() {
	*x = FindMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moviori_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMovieResponse) ProtoMessage() {}

func (x *FindMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moviori_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMovieResponse.ProtoReflect.Descriptor instead.
func (*FindMovieResponse) Descriptor() ([]byte, []int) {
	return file_moviori_proto_rawDescGZIP(), []int{2}
}

func (x *FindMovieResponse) GetMovieInfo() []byte {
	if x != nil {
		return x.MovieInfo
	}
	return nil
}

var File_moviori_proto protoreflect.FileDescriptor

var file_moviori_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x76, 0x69, 0x6f, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x17, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x22, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x31, 0x0a, 0x11,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0x45, 0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x69, 0x6f, 0x72, 0x69, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x11, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x75, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6d, 0x61, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x6f, 0x72, 0x69, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moviori_proto_rawDescOnce sync.Once
	file_moviori_proto_rawDescData = file_moviori_proto_rawDesc
)

func file_moviori_proto_rawDescGZIP() []byte {
	file_moviori_proto_rawDescOnce.Do(func() {
		file_moviori_proto_rawDescData = protoimpl.X.CompressGZIP(file_moviori_proto_rawDescData)
	})
	return file_moviori_proto_rawDescData
}

var file_moviori_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_moviori_proto_goTypes = []any{
	(*Movie)(nil),             // 0: Movie
	(*FindMovieRequest)(nil),  // 1: FindMovieRequest
	(*FindMovieResponse)(nil), // 2: FindMovieResponse
}
var file_moviori_proto_depIdxs = []int32{
	1, // 0: Moviori_movies.FindByName:input_type -> FindMovieRequest
	2, // 1: Moviori_movies.FindByName:output_type -> FindMovieResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_moviori_proto_init() }
func file_moviori_proto_init() {
	if File_moviori_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moviori_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Movie); i {
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
		file_moviori_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FindMovieRequest); i {
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
		file_moviori_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FindMovieResponse); i {
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
			RawDescriptor: file_moviori_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moviori_proto_goTypes,
		DependencyIndexes: file_moviori_proto_depIdxs,
		MessageInfos:      file_moviori_proto_msgTypes,
	}.Build()
	File_moviori_proto = out.File
	file_moviori_proto_rawDesc = nil
	file_moviori_proto_goTypes = nil
	file_moviori_proto_depIdxs = nil
}
