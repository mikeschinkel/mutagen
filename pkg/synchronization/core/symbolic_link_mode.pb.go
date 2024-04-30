// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: synchronization/core/symbolic_link_mode.proto

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

// SymbolicLinkMode specifies the mode for handling symbolic links.
type SymbolicLinkMode int32

const (
	// SymbolicLinkMode_SymbolicLinkModeDefault represents an unspecified
	// symbolic link mode. It is not valid for use with Scan or Transition. It
	// should be converted to one of the following values based on the desired
	// default behavior.
	SymbolicLinkMode_SymbolicLinkModeDefault SymbolicLinkMode = 0
	// SymbolicLinkMode_SymbolicLinkModeIgnore specifies that all symbolic links
	// should be ignored.
	SymbolicLinkMode_SymbolicLinkModeIgnore SymbolicLinkMode = 1
	// SymbolicLinkMode_SymbolicLinkModePortable specifies that only portable
	// symbolic links should be synchronized. Any absolute symbolic links or
	// symbolic links which are otherwise non-portable will be treate as
	// problematic content.
	SymbolicLinkMode_SymbolicLinkModePortable SymbolicLinkMode = 2
	// SymbolicLinkMode_SymbolicLinkModePOSIXRaw specifies that symbolic links
	// should be propagated in their raw form. It is only valid on POSIX systems
	// and only makes sense in the context of POSIX-to-POSIX synchronization.
	SymbolicLinkMode_SymbolicLinkModePOSIXRaw SymbolicLinkMode = 3
)

// Enum value maps for SymbolicLinkMode.
var (
	SymbolicLinkMode_name = map[int32]string{
		0: "SymbolicLinkModeDefault",
		1: "SymbolicLinkModeIgnore",
		2: "SymbolicLinkModePortable",
		3: "SymbolicLinkModePOSIXRaw",
	}
	SymbolicLinkMode_value = map[string]int32{
		"SymbolicLinkModeDefault":  0,
		"SymbolicLinkModeIgnore":   1,
		"SymbolicLinkModePortable": 2,
		"SymbolicLinkModePOSIXRaw": 3,
	}
)

func (x SymbolicLinkMode) Enum() *SymbolicLinkMode {
	p := new(SymbolicLinkMode)
	*p = x
	return p
}

func (x SymbolicLinkMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SymbolicLinkMode) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_core_symbolic_link_mode_proto_enumTypes[0].Descriptor()
}

func (SymbolicLinkMode) Type() protoreflect.EnumType {
	return &file_synchronization_core_symbolic_link_mode_proto_enumTypes[0]
}

func (x SymbolicLinkMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SymbolicLinkMode.Descriptor instead.
func (SymbolicLinkMode) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_core_symbolic_link_mode_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_core_symbolic_link_mode_proto protoreflect.FileDescriptor

var file_synchronization_core_symbolic_link_mode_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x87, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x4c,
	0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x02, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e,
	0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x4f, 0x53, 0x49, 0x58, 0x52, 0x61, 0x77, 0x10, 0x03, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75,
	0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_synchronization_core_symbolic_link_mode_proto_rawDescOnce sync.Once
	file_synchronization_core_symbolic_link_mode_proto_rawDescData = file_synchronization_core_symbolic_link_mode_proto_rawDesc
)

func file_synchronization_core_symbolic_link_mode_proto_rawDescGZIP() []byte {
	file_synchronization_core_symbolic_link_mode_proto_rawDescOnce.Do(func() {
		file_synchronization_core_symbolic_link_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_symbolic_link_mode_proto_rawDescData)
	})
	return file_synchronization_core_symbolic_link_mode_proto_rawDescData
}

var file_synchronization_core_symbolic_link_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_core_symbolic_link_mode_proto_goTypes = []interface{}{
	(SymbolicLinkMode)(0), // 0: core.SymbolicLinkMode
}
var file_synchronization_core_symbolic_link_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_core_symbolic_link_mode_proto_init() }
func file_synchronization_core_symbolic_link_mode_proto_init() {
	if File_synchronization_core_symbolic_link_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_core_symbolic_link_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_symbolic_link_mode_proto_goTypes,
		DependencyIndexes: file_synchronization_core_symbolic_link_mode_proto_depIdxs,
		EnumInfos:         file_synchronization_core_symbolic_link_mode_proto_enumTypes,
	}.Build()
	File_synchronization_core_symbolic_link_mode_proto = out.File
	file_synchronization_core_symbolic_link_mode_proto_rawDesc = nil
	file_synchronization_core_symbolic_link_mode_proto_goTypes = nil
	file_synchronization_core_symbolic_link_mode_proto_depIdxs = nil
}
