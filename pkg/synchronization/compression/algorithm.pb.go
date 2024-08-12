// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: synchronization/compression/algorithm.proto

package compression

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

// Algorithm specifies a compression algorithm.
type Algorithm int32

const (
	// Algorithm_AlgorithmDefault represents an unspecified compression
	// algorithm. It should be converted to one of the following values based on
	// the desired default behavior.
	Algorithm_AlgorithmDefault Algorithm = 0
	// Algorithm_AlgorithmNone specifies that no compression should be used.
	Algorithm_AlgorithmNone Algorithm = 1
	// Algorithm_AlgorithmDeflate specifies that DEFLATE compression should be
	// used.
	Algorithm_AlgorithmDeflate Algorithm = 2
	// Algorithm_AlgorithmZstandard specifies that Zstandard compression should
	// be used.
	Algorithm_AlgorithmZstandard Algorithm = 3
)

// Enum value maps for Algorithm.
var (
	Algorithm_name = map[int32]string{
		0: "AlgorithmDefault",
		1: "AlgorithmNone",
		2: "AlgorithmDeflate",
		3: "AlgorithmZstandard",
	}
	Algorithm_value = map[string]int32{
		"AlgorithmDefault":   0,
		"AlgorithmNone":      1,
		"AlgorithmDeflate":   2,
		"AlgorithmZstandard": 3,
	}
)

func (x Algorithm) Enum() *Algorithm {
	p := new(Algorithm)
	*p = x
	return p
}

func (x Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_compression_algorithm_proto_enumTypes[0].Descriptor()
}

func (Algorithm) Type() protoreflect.EnumType {
	return &file_synchronization_compression_algorithm_proto_enumTypes[0]
}

func (x Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Algorithm.Descriptor instead.
func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_compression_algorithm_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_compression_algorithm_proto protoreflect.FileDescriptor

var file_synchronization_compression_algorithm_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x62, 0x0a, 0x09, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x44, 0x65, 0x66,
	0x6c, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x5a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x10, 0x03, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74,
	0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_compression_algorithm_proto_rawDescOnce sync.Once
	file_synchronization_compression_algorithm_proto_rawDescData = file_synchronization_compression_algorithm_proto_rawDesc
)

func file_synchronization_compression_algorithm_proto_rawDescGZIP() []byte {
	file_synchronization_compression_algorithm_proto_rawDescOnce.Do(func() {
		file_synchronization_compression_algorithm_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_compression_algorithm_proto_rawDescData)
	})
	return file_synchronization_compression_algorithm_proto_rawDescData
}

var file_synchronization_compression_algorithm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_compression_algorithm_proto_goTypes = []interface{}{
	(Algorithm)(0), // 0: compression.Algorithm
}
var file_synchronization_compression_algorithm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_compression_algorithm_proto_init() }
func file_synchronization_compression_algorithm_proto_init() {
	if File_synchronization_compression_algorithm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_compression_algorithm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_compression_algorithm_proto_goTypes,
		DependencyIndexes: file_synchronization_compression_algorithm_proto_depIdxs,
		EnumInfos:         file_synchronization_compression_algorithm_proto_enumTypes,
	}.Build()
	File_synchronization_compression_algorithm_proto = out.File
	file_synchronization_compression_algorithm_proto_rawDesc = nil
	file_synchronization_compression_algorithm_proto_goTypes = nil
	file_synchronization_compression_algorithm_proto_depIdxs = nil
}