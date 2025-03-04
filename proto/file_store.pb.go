// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: file_store.proto

package file_store

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

type FileData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileData) Reset() {
	*x = FileData{}
	mi := &file_file_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileData) ProtoMessage() {}

func (x *FileData) ProtoReflect() protoreflect.Message {
	mi := &file_file_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileData.ProtoReflect.Descriptor instead.
func (*FileData) Descriptor() ([]byte, []int) {
	return file_file_store_proto_rawDescGZIP(), []int{0}
}

func (x *FileData) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *FileData) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type FileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	mi := &file_file_store_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_store_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_file_store_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_file_store_proto protoreflect.FileDescriptor

var file_file_store_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x3c,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x4c, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x3c, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x0e,
	0x5a, 0x0c, 0x2f, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_file_store_proto_rawDescOnce sync.Once
	file_file_store_proto_rawDescData []byte
)

func file_file_store_proto_rawDescGZIP() []byte {
	file_file_store_proto_rawDescOnce.Do(func() {
		file_file_store_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_file_store_proto_rawDesc), len(file_file_store_proto_rawDesc)))
	})
	return file_file_store_proto_rawDescData
}

var file_file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_file_store_proto_goTypes = []any{
	(*FileData)(nil),     // 0: file_store.FileData
	(*FileResponse)(nil), // 1: file_store.FileResponse
}
var file_file_store_proto_depIdxs = []int32{
	0, // 0: file_store.FileTransfer.SendFile:input_type -> file_store.FileData
	1, // 1: file_store.FileTransfer.SendFile:output_type -> file_store.FileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_store_proto_init() }
func file_file_store_proto_init() {
	if File_file_store_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_file_store_proto_rawDesc), len(file_file_store_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_store_proto_goTypes,
		DependencyIndexes: file_file_store_proto_depIdxs,
		MessageInfos:      file_file_store_proto_msgTypes,
	}.Build()
	File_file_store_proto = out.File
	file_file_store_proto_goTypes = nil
	file_file_store_proto_depIdxs = nil
}
