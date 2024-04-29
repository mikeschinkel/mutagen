// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: service/forwarding/forwarding.proto

package forwarding

import (
	forwarding "github.com/mutagen-io/mutagen/pkg/forwarding"
	selection "github.com/mutagen-io/mutagen/pkg/selection"
	url "github.com/mutagen-io/mutagen/pkg/url"
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

// CreationSpecification contains the metadata required for a new session.
type CreationSpecification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source is the source endpoint URL for the session.
	Source *url.URL `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Destination is the destination endpoint URL for the session.
	Destination *url.URL `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Configuration is the base session configuration. It is the result of
	// merging the global configuration (unless disabled), any manually
	// specified configuration file, and any command line configuration
	// parameters.
	Configuration *forwarding.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// ConfigurationSource is the source-specific session configuration. It is
	// determined based on command line configuration parameters.
	ConfigurationSource *forwarding.Configuration `protobuf:"bytes,4,opt,name=configurationSource,proto3" json:"configurationSource,omitempty"`
	// ConfigurationDestination is the destination-specific session
	// configuration. It is determined based on command line configuration
	// parameters.
	ConfigurationDestination *forwarding.Configuration `protobuf:"bytes,5,opt,name=configurationDestination,proto3" json:"configurationDestination,omitempty"`
	// Name is the name for the session object.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the labels for the session object.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not to create the session pre-paused.
	Paused bool `protobuf:"varint,8,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *CreationSpecification) Reset() {
	*x = CreationSpecification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreationSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreationSpecification) ProtoMessage() {}

func (x *CreationSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreationSpecification.ProtoReflect.Descriptor instead.
func (*CreationSpecification) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{0}
}

func (x *CreationSpecification) GetSource() *url.URL {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CreationSpecification) GetDestination() *url.URL {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *CreationSpecification) GetConfiguration() *forwarding.Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *CreationSpecification) GetConfigurationSource() *forwarding.Configuration {
	if x != nil {
		return x.ConfigurationSource
	}
	return nil
}

func (x *CreationSpecification) GetConfigurationDestination() *forwarding.Configuration {
	if x != nil {
		return x.ConfigurationDestination
	}
	return nil
}

func (x *CreationSpecification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreationSpecification) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreationSpecification) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

// CreateRequest encodes a request for session creation.
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompter is the prompter identifier to use for creating sessions.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Specification is the creation specification.
	Specification *CreationSpecification `protobuf:"bytes,2,opt,name=specification,proto3" json:"specification,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetPrompter() string {
	if x != nil {
		return x.Prompter
	}
	return ""
}

func (x *CreateRequest) GetSpecification() *CreationSpecification {
	if x != nil {
		return x.Specification
	}
	return nil
}

// CreateResponse encodes a session creation response.
type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session is the resulting session identifier.
	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

// ListRequest encodes a request for session metadata.
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selection is the session selection criteria.
	Selection *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	// PreviousStateIndex is the previously seen state index. 0 may be provided
	// to force an immediate state listing.
	PreviousStateIndex uint64 `protobuf:"varint,2,opt,name=previousStateIndex,proto3" json:"previousStateIndex,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetSelection() *selection.Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

func (x *ListRequest) GetPreviousStateIndex() uint64 {
	if x != nil {
		return x.PreviousStateIndex
	}
	return 0
}

// ListResponse encodes session metadata.
type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// StateIndex is the state index associated with the session metadata.
	StateIndex uint64 `protobuf:"varint,1,opt,name=stateIndex,proto3" json:"stateIndex,omitempty"`
	// SessionStates are the session metadata states.
	SessionStates []*forwarding.State `protobuf:"bytes,2,rep,name=sessionStates,proto3" json:"sessionStates,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetStateIndex() uint64 {
	if x != nil {
		return x.StateIndex
	}
	return 0
}

func (x *ListResponse) GetSessionStates() []*forwarding.State {
	if x != nil {
		return x.SessionStates
	}
	return nil
}

// PauseRequest encodes a request to pause sessions.
type PauseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompter is the prompter to use for status message updates.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Selection is the session selection criteria.
	Selection *selection.Selection `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
}

func (x *PauseRequest) Reset() {
	*x = PauseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseRequest) ProtoMessage() {}

func (x *PauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseRequest.ProtoReflect.Descriptor instead.
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{5}
}

func (x *PauseRequest) GetPrompter() string {
	if x != nil {
		return x.Prompter
	}
	return ""
}

func (x *PauseRequest) GetSelection() *selection.Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

// PauseResponse indicates completion of pause operation(s).
type PauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PauseResponse) Reset() {
	*x = PauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseResponse) ProtoMessage() {}

func (x *PauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseResponse.ProtoReflect.Descriptor instead.
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{6}
}

// ResumeRequest encodes a request to resume sessions.
type ResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompter is the prompter identifier to use for resuming sessions.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Selection is the session selection criteria.
	Selection *selection.Selection `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
}

func (x *ResumeRequest) Reset() {
	*x = ResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeRequest) ProtoMessage() {}

func (x *ResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeRequest.ProtoReflect.Descriptor instead.
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{7}
}

func (x *ResumeRequest) GetPrompter() string {
	if x != nil {
		return x.Prompter
	}
	return ""
}

func (x *ResumeRequest) GetSelection() *selection.Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

// ResumeResponse indicates completion of resume operation(s).
type ResumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResumeResponse) Reset() {
	*x = ResumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeResponse) ProtoMessage() {}

func (x *ResumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeResponse.ProtoReflect.Descriptor instead.
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{8}
}

// TerminateRequest encodes a request to terminate sessions.
type TerminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompter is the prompter to use for status message updates.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Selection is the session selection criteria.
	Selection *selection.Selection `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
}

func (x *TerminateRequest) Reset() {
	*x = TerminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateRequest) ProtoMessage() {}

func (x *TerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateRequest.ProtoReflect.Descriptor instead.
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{9}
}

func (x *TerminateRequest) GetPrompter() string {
	if x != nil {
		return x.Prompter
	}
	return ""
}

func (x *TerminateRequest) GetSelection() *selection.Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

// TerminateResponse indicates completion of termination operation(s).
type TerminateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminateResponse) Reset() {
	*x = TerminateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_forwarding_forwarding_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateResponse) ProtoMessage() {}

func (x *TerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_forwarding_forwarding_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateResponse.ProtoReflect.Descriptor instead.
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return file_service_forwarding_forwarding_proto_rawDescGZIP(), []int{10}
}

var File_service_forwarding_forwarding_proto protoreflect.FileDescriptor

var file_service_forwarding_forwarding_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x1a, 0x19, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x72, 0x6c, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x03, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x75, 0x72, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x13,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x18, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75,
	0x73, 0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x74,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0d, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x71, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x67, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x0c,
	0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d,
	0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x10,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x62, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdb, 0x02, 0x0a, 0x0a, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05, 0x50, 0x61, 0x75, 0x73,
	0x65, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x12, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f,
	0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_forwarding_forwarding_proto_rawDescOnce sync.Once
	file_service_forwarding_forwarding_proto_rawDescData = file_service_forwarding_forwarding_proto_rawDesc
)

func file_service_forwarding_forwarding_proto_rawDescGZIP() []byte {
	file_service_forwarding_forwarding_proto_rawDescOnce.Do(func() {
		file_service_forwarding_forwarding_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_forwarding_forwarding_proto_rawDescData)
	})
	return file_service_forwarding_forwarding_proto_rawDescData
}

var file_service_forwarding_forwarding_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_service_forwarding_forwarding_proto_goTypes = []interface{}{
	(*CreationSpecification)(nil),    // 0: forwarding.CreationSpecification
	(*CreateRequest)(nil),            // 1: forwarding.CreateRequest
	(*CreateResponse)(nil),           // 2: forwarding.CreateResponse
	(*ListRequest)(nil),              // 3: forwarding.ListRequest
	(*ListResponse)(nil),             // 4: forwarding.ListResponse
	(*PauseRequest)(nil),             // 5: forwarding.PauseRequest
	(*PauseResponse)(nil),            // 6: forwarding.PauseResponse
	(*ResumeRequest)(nil),            // 7: forwarding.ResumeRequest
	(*ResumeResponse)(nil),           // 8: forwarding.ResumeResponse
	(*TerminateRequest)(nil),         // 9: forwarding.TerminateRequest
	(*TerminateResponse)(nil),        // 10: forwarding.TerminateResponse
	nil,                              // 11: forwarding.CreationSpecification.LabelsEntry
	(*url.URL)(nil),                  // 12: url.URL
	(*forwarding.Configuration)(nil), // 13: forwarding.Configuration
	(*selection.Selection)(nil),      // 14: selection.Selection
	(*forwarding.State)(nil),         // 15: forwarding.State
}
var file_service_forwarding_forwarding_proto_depIdxs = []int32{
	12, // 0: forwarding.CreationSpecification.source:type_name -> url.URL
	12, // 1: forwarding.CreationSpecification.destination:type_name -> url.URL
	13, // 2: forwarding.CreationSpecification.configuration:type_name -> forwarding.Configuration
	13, // 3: forwarding.CreationSpecification.configurationSource:type_name -> forwarding.Configuration
	13, // 4: forwarding.CreationSpecification.configurationDestination:type_name -> forwarding.Configuration
	11, // 5: forwarding.CreationSpecification.labels:type_name -> forwarding.CreationSpecification.LabelsEntry
	0,  // 6: forwarding.CreateRequest.specification:type_name -> forwarding.CreationSpecification
	14, // 7: forwarding.ListRequest.selection:type_name -> selection.Selection
	15, // 8: forwarding.ListResponse.sessionStates:type_name -> forwarding.State
	14, // 9: forwarding.PauseRequest.selection:type_name -> selection.Selection
	14, // 10: forwarding.ResumeRequest.selection:type_name -> selection.Selection
	14, // 11: forwarding.TerminateRequest.selection:type_name -> selection.Selection
	1,  // 12: forwarding.Forwarding.Create:input_type -> forwarding.CreateRequest
	3,  // 13: forwarding.Forwarding.List:input_type -> forwarding.ListRequest
	5,  // 14: forwarding.Forwarding.Pause:input_type -> forwarding.PauseRequest
	7,  // 15: forwarding.Forwarding.Resume:input_type -> forwarding.ResumeRequest
	9,  // 16: forwarding.Forwarding.Terminate:input_type -> forwarding.TerminateRequest
	2,  // 17: forwarding.Forwarding.Create:output_type -> forwarding.CreateResponse
	4,  // 18: forwarding.Forwarding.List:output_type -> forwarding.ListResponse
	6,  // 19: forwarding.Forwarding.Pause:output_type -> forwarding.PauseResponse
	8,  // 20: forwarding.Forwarding.Resume:output_type -> forwarding.ResumeResponse
	10, // 21: forwarding.Forwarding.Terminate:output_type -> forwarding.TerminateResponse
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_service_forwarding_forwarding_proto_init() }
func file_service_forwarding_forwarding_proto_init() {
	if File_service_forwarding_forwarding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_forwarding_forwarding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreationSpecification); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseRequest); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseResponse); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeRequest); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeResponse); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateRequest); i {
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
		file_service_forwarding_forwarding_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateResponse); i {
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
			RawDescriptor: file_service_forwarding_forwarding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_forwarding_forwarding_proto_goTypes,
		DependencyIndexes: file_service_forwarding_forwarding_proto_depIdxs,
		MessageInfos:      file_service_forwarding_forwarding_proto_msgTypes,
	}.Build()
	File_service_forwarding_forwarding_proto = out.File
	file_service_forwarding_forwarding_proto_rawDesc = nil
	file_service_forwarding_forwarding_proto_goTypes = nil
	file_service_forwarding_forwarding_proto_depIdxs = nil
}
