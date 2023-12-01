// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: synchronization/session.proto

package synchronization

import (
	url "github.com/mutagen-io/mutagen/pkg/url"
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

// Session represents a synchronization session configuration and persistent
// state. It is mutable within the context of the daemon, so it should be
// accessed and modified in a synchronized fashion. Outside of the daemon (e.g.
// when returned via the API), it should be considered immutable.
type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier is the (unique) session identifier. It is static. It cannot be
	// empty.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Version is the session version. It is static.
	Version Version `protobuf:"varint,2,opt,name=version,proto3,enum=synchronization.Version" json:"version,omitempty"`
	// CreationTime is the creation time of the session. It is static. It cannot
	// be nil.
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	// CreatingVersionMajor is the major version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionMajor uint32 `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	// CreatingVersionMinor is the minor version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionMinor uint32 `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	// CreatingVersionPatch is the patch version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionPatch uint32 `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	// Alpha is the alpha endpoint URL. It is static. It cannot be nil.
	Alpha *url.URL `protobuf:"bytes,7,opt,name=alpha,proto3" json:"alpha,omitempty"`
	// Beta is the beta endpoint URL. It is static. It cannot be nil.
	Beta *url.URL `protobuf:"bytes,8,opt,name=beta,proto3" json:"beta,omitempty"`
	// Configuration is the flattened session configuration. It is static. It
	// cannot be nil.
	Configuration *Configuration `protobuf:"bytes,9,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// ConfigurationAlpha are the alpha-specific session configuration
	// overrides. It is static. It may be nil for existing sessions loaded from
	// disk, but it is not considered valid unless non-nil, so it should be
	// replaced with an empty default value in-memory if a nil on-disk value is
	// detected.
	ConfigurationAlpha *Configuration `protobuf:"bytes,11,opt,name=configurationAlpha,proto3" json:"configurationAlpha,omitempty"`
	// ConfigurationBeta are the beta-specific session configuration overrides.
	// It is static. It may be nil for existing sessions loaded from disk, but
	// it is not considered valid unless non-nil, so it should be replaced with
	// an empty default value in-memory if a nil on-disk value is detected.
	ConfigurationBeta *Configuration `protobuf:"bytes,12,opt,name=configurationBeta,proto3" json:"configurationBeta,omitempty"`
	// Name is a user-friendly name for the session. It may be empty and is not
	// guaranteed to be unique across all sessions. It is only used as a simpler
	// handle for specifying sessions. It is static.
	Name string `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the session labels. They are static.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not the session is marked as paused.
	Paused bool `protobuf:"varint,10,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_synchronization_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Session) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_Invalid
}

func (x *Session) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *Session) GetCreatingVersionMajor() uint32 {
	if x != nil {
		return x.CreatingVersionMajor
	}
	return 0
}

func (x *Session) GetCreatingVersionMinor() uint32 {
	if x != nil {
		return x.CreatingVersionMinor
	}
	return 0
}

func (x *Session) GetCreatingVersionPatch() uint32 {
	if x != nil {
		return x.CreatingVersionPatch
	}
	return 0
}

func (x *Session) GetAlpha() *url.URL {
	if x != nil {
		return x.Alpha
	}
	return nil
}

func (x *Session) GetBeta() *url.URL {
	if x != nil {
		return x.Beta
	}
	return nil
}

func (x *Session) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *Session) GetConfigurationAlpha() *Configuration {
	if x != nil {
		return x.ConfigurationAlpha
	}
	return nil
}

func (x *Session) GetConfigurationBeta() *Configuration {
	if x != nil {
		return x.ConfigurationBeta
	}
	return nil
}

func (x *Session) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Session) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Session) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

var File_synchronization_session_proto protoreflect.FileDescriptor

var file_synchronization_session_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x72, 0x6c, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x06, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x14,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x1e, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x12, 0x1c, 0x0a, 0x04, 0x62, 0x65, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x62, 0x65, 0x74, 0x61, 0x12, 0x44,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0x12, 0x4c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f,
	0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_session_proto_rawDescOnce sync.Once
	file_synchronization_session_proto_rawDescData = file_synchronization_session_proto_rawDesc
)

func file_synchronization_session_proto_rawDescGZIP() []byte {
	file_synchronization_session_proto_rawDescOnce.Do(func() {
		file_synchronization_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_session_proto_rawDescData)
	})
	return file_synchronization_session_proto_rawDescData
}

var file_synchronization_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_synchronization_session_proto_goTypes = []interface{}{
	(*Session)(nil),               // 0: synchronization.Session
	nil,                           // 1: synchronization.Session.LabelsEntry
	(Version)(0),                  // 2: synchronization.Version
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*url.URL)(nil),               // 4: url.URL
	(*Configuration)(nil),         // 5: synchronization.Configuration
}
var file_synchronization_session_proto_depIdxs = []int32{
	2, // 0: synchronization.Session.version:type_name -> synchronization.Version
	3, // 1: synchronization.Session.creationTime:type_name -> google.protobuf.Timestamp
	4, // 2: synchronization.Session.alpha:type_name -> url.URL
	4, // 3: synchronization.Session.beta:type_name -> url.URL
	5, // 4: synchronization.Session.configuration:type_name -> synchronization.Configuration
	5, // 5: synchronization.Session.configurationAlpha:type_name -> synchronization.Configuration
	5, // 6: synchronization.Session.configurationBeta:type_name -> synchronization.Configuration
	1, // 7: synchronization.Session.labels:type_name -> synchronization.Session.LabelsEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_synchronization_session_proto_init() }
func file_synchronization_session_proto_init() {
	if File_synchronization_session_proto != nil {
		return
	}
	file_synchronization_configuration_proto_init()
	file_synchronization_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synchronization_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_synchronization_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_session_proto_goTypes,
		DependencyIndexes: file_synchronization_session_proto_depIdxs,
		MessageInfos:      file_synchronization_session_proto_msgTypes,
	}.Build()
	File_synchronization_session_proto = out.File
	file_synchronization_session_proto_rawDesc = nil
	file_synchronization_session_proto_goTypes = nil
	file_synchronization_session_proto_depIdxs = nil
}
