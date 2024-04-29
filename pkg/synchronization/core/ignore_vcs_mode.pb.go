// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: synchronization/core/ignore_vcs_mode.proto

package core

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

// IgnoreVCSMode specifies the mode for ignoring VCS directories.
type IgnoreVCSMode int32

const (
	// IgnoreVCSMode_IgnoreVCSModeDefault represents an unspecified VCS ignore
	// mode. It is not valid for use with Scan. It should be converted to one of
	// the following values based on the desired default behavior.
	IgnoreVCSMode_IgnoreVCSModeDefault IgnoreVCSMode = 0
	// IgnoreVCSMode_IgnoreVCSModeIgnore indicates that VCS directories should
	// be ignored.
	IgnoreVCSMode_IgnoreVCSModeIgnore IgnoreVCSMode = 1
	// IgnoreVCSMode_IgnoreVCSModePropagate indicates that VCS directories
	// should be propagated.
	IgnoreVCSMode_IgnoreVCSModePropagate IgnoreVCSMode = 2
)

// Enum value maps for IgnoreVCSMode.
var (
	IgnoreVCSMode_name = map[int32]string{
		0: "IgnoreVCSModeDefault",
		1: "IgnoreVCSModeIgnore",
		2: "IgnoreVCSModePropagate",
	}
	IgnoreVCSMode_value = map[string]int32{
		"IgnoreVCSModeDefault":   0,
		"IgnoreVCSModeIgnore":    1,
		"IgnoreVCSModePropagate": 2,
	}
)

func (x IgnoreVCSMode) Enum() *IgnoreVCSMode {
	p := new(IgnoreVCSMode)
	*p = x
	return p
}

func (x IgnoreVCSMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IgnoreVCSMode) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_core_ignore_vcs_mode_proto_enumTypes[0].Descriptor()
}

func (IgnoreVCSMode) Type() protoreflect.EnumType {
	return &file_synchronization_core_ignore_vcs_mode_proto_enumTypes[0]
}

func (x IgnoreVCSMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IgnoreVCSMode.Descriptor instead.
func (IgnoreVCSMode) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_core_ignore_vcs_mode_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_core_ignore_vcs_mode_proto protoreflect.FileDescriptor

var file_synchronization_core_ignore_vcs_mode_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x63,
	0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x72, 0x65, 0x2a, 0x5e, 0x0a, 0x0d, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x43, 0x53, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x43, 0x53,
	0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x43, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x56, 0x43, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65,
	0x10, 0x02, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_ignore_vcs_mode_proto_rawDescOnce sync.Once
	file_synchronization_core_ignore_vcs_mode_proto_rawDescData = file_synchronization_core_ignore_vcs_mode_proto_rawDesc
)

func file_synchronization_core_ignore_vcs_mode_proto_rawDescGZIP() []byte {
	file_synchronization_core_ignore_vcs_mode_proto_rawDescOnce.Do(func() {
		file_synchronization_core_ignore_vcs_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_ignore_vcs_mode_proto_rawDescData)
	})
	return file_synchronization_core_ignore_vcs_mode_proto_rawDescData
}

var file_synchronization_core_ignore_vcs_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_core_ignore_vcs_mode_proto_goTypes = []interface{}{
	(IgnoreVCSMode)(0), // 0: core.IgnoreVCSMode
}
var file_synchronization_core_ignore_vcs_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_core_ignore_vcs_mode_proto_init() }
func file_synchronization_core_ignore_vcs_mode_proto_init() {
	if File_synchronization_core_ignore_vcs_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_core_ignore_vcs_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_ignore_vcs_mode_proto_goTypes,
		DependencyIndexes: file_synchronization_core_ignore_vcs_mode_proto_depIdxs,
		EnumInfos:         file_synchronization_core_ignore_vcs_mode_proto_enumTypes,
	}.Build()
	File_synchronization_core_ignore_vcs_mode_proto = out.File
	file_synchronization_core_ignore_vcs_mode_proto_rawDesc = nil
	file_synchronization_core_ignore_vcs_mode_proto_goTypes = nil
	file_synchronization_core_ignore_vcs_mode_proto_depIdxs = nil
}
