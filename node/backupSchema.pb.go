// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: backupSchema.proto

package node

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha          []byte                 `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastModified *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	Size         uint64                 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"` // size is in KB
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupSchema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_backupSchema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_backupSchema_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *File) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type BackupSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *BackupSchema) Reset() {
	*x = BackupSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backupSchema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupSchema) ProtoMessage() {}

func (x *BackupSchema) ProtoReflect() protoreflect.Message {
	mi := &file_backupSchema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupSchema.ProtoReflect.Descriptor instead.
func (*BackupSchema) Descriptor() ([]byte, []int) {
	return file_backupSchema_proto_rawDescGZIP(), []int{1}
}

func (x *BackupSchema) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_backupSchema_proto protoreflect.FileDescriptor

var file_backupSchema_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x68, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2b, 0x0a, 0x0c, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backupSchema_proto_rawDescOnce sync.Once
	file_backupSchema_proto_rawDescData = file_backupSchema_proto_rawDesc
)

func file_backupSchema_proto_rawDescGZIP() []byte {
	file_backupSchema_proto_rawDescOnce.Do(func() {
		file_backupSchema_proto_rawDescData = protoimpl.X.CompressGZIP(file_backupSchema_proto_rawDescData)
	})
	return file_backupSchema_proto_rawDescData
}

var file_backupSchema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backupSchema_proto_goTypes = []interface{}{
	(*File)(nil),                  // 0: File
	(*BackupSchema)(nil),          // 1: BackupSchema
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_backupSchema_proto_depIdxs = []int32{
	2, // 0: File.last_modified:type_name -> google.protobuf.Timestamp
	0, // 1: BackupSchema.files:type_name -> File
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_backupSchema_proto_init() }
func file_backupSchema_proto_init() {
	if File_backupSchema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backupSchema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_backupSchema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupSchema); i {
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
			RawDescriptor: file_backupSchema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backupSchema_proto_goTypes,
		DependencyIndexes: file_backupSchema_proto_depIdxs,
		MessageInfos:      file_backupSchema_proto_msgTypes,
	}.Build()
	File_backupSchema_proto = out.File
	file_backupSchema_proto_rawDesc = nil
	file_backupSchema_proto_goTypes = nil
	file_backupSchema_proto_depIdxs = nil
}