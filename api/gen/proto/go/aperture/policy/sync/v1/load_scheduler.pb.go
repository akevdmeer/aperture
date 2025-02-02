// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aperture/policy/sync/v1/load_scheduler.proto

package syncv1

import (
	v1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
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

type LoadSchedulerWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommonAttributes
	CommonAttributes *CommonAttributes `protobuf:"bytes,1,opt,name=common_attributes,json=commonAttributes,proto3" json:"common_attributes,omitempty"`
	// Load Scheduler
	LoadScheduler *v1.LoadScheduler `protobuf:"bytes,2,opt,name=load_scheduler,json=loadScheduler,proto3" json:"load_scheduler,omitempty"`
}

func (x *LoadSchedulerWrapper) Reset() {
	*x = LoadSchedulerWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadSchedulerWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadSchedulerWrapper) ProtoMessage() {}

func (x *LoadSchedulerWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadSchedulerWrapper.ProtoReflect.Descriptor instead.
func (*LoadSchedulerWrapper) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_load_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *LoadSchedulerWrapper) GetCommonAttributes() *CommonAttributes {
	if x != nil {
		return x.CommonAttributes
	}
	return nil
}

func (x *LoadSchedulerWrapper) GetLoadScheduler() *v1.LoadScheduler {
	if x != nil {
		return x.LoadScheduler
	}
	return nil
}

type LoadDecisionWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommonAttributes
	CommonAttributes *CommonAttributes `protobuf:"bytes,1,opt,name=common_attributes,json=commonAttributes,proto3" json:"common_attributes,omitempty"`
	// Load Decision
	LoadDecision *LoadDecision `protobuf:"bytes,2,opt,name=load_decision,json=loadDecision,proto3" json:"load_decision,omitempty"`
}

func (x *LoadDecisionWrapper) Reset() {
	*x = LoadDecisionWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadDecisionWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadDecisionWrapper) ProtoMessage() {}

func (x *LoadDecisionWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadDecisionWrapper.ProtoReflect.Descriptor instead.
func (*LoadDecisionWrapper) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_load_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *LoadDecisionWrapper) GetCommonAttributes() *CommonAttributes {
	if x != nil {
		return x.CommonAttributes
	}
	return nil
}

func (x *LoadDecisionWrapper) GetLoadDecision() *LoadDecision {
	if x != nil {
		return x.LoadDecision
	}
	return nil
}

type LoadDecision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TickInfo              *TickInfo         `protobuf:"bytes,1,opt,name=tick_info,json=tickInfo,proto3" json:"tick_info,omitempty"`
	LoadMultiplier        float64           `protobuf:"fixed64,2,opt,name=load_multiplier,json=loadMultiplier,proto3" json:"load_multiplier,omitempty"`
	PassThrough           bool              `protobuf:"varint,3,opt,name=pass_through,json=passThrough,proto3" json:"pass_through,omitempty"`
	TokensByWorkloadIndex map[string]uint64 `protobuf:"bytes,4,rep,name=tokens_by_workload_index,json=tokensByWorkloadIndex,proto3" json:"tokens_by_workload_index,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *LoadDecision) Reset() {
	*x = LoadDecision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadDecision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadDecision) ProtoMessage() {}

func (x *LoadDecision) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadDecision.ProtoReflect.Descriptor instead.
func (*LoadDecision) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_load_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *LoadDecision) GetTickInfo() *TickInfo {
	if x != nil {
		return x.TickInfo
	}
	return nil
}

func (x *LoadDecision) GetLoadMultiplier() float64 {
	if x != nil {
		return x.LoadMultiplier
	}
	return 0
}

func (x *LoadDecision) GetPassThrough() bool {
	if x != nil {
		return x.PassThrough
	}
	return false
}

func (x *LoadDecision) GetTokensByWorkloadIndex() map[string]uint64 {
	if x != nil {
		return x.TokensByWorkloadIndex
	}
	return nil
}

var File_aperture_policy_sync_v1_load_scheduler_proto protoreflect.FileDescriptor

var file_aperture_policy_sync_v1_load_scheduler_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x69, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x14,
	0x4c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x52, 0x0d, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22,
	0xb9, 0x01, 0x0a, 0x13, 0x4c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x10, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x0d, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xdf, 0x02, 0x0a, 0x0c,
	0x4c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x09,
	0x74, 0x69, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x68,
	0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x61, 0x73,
	0x73, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x12, 0x79, 0x0a, 0x18, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x1a, 0x48, 0x0a, 0x1a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x79, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x96, 0x02,
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61, 0x70,
	0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x79, 0x6e, 0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x50, 0x53, 0xaa,
	0x02, 0x17, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x41, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x53, 0x79, 0x6e, 0x63,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x53, 0x79, 0x6e, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x53, 0x79,
	0x6e, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aperture_policy_sync_v1_load_scheduler_proto_rawDescOnce sync.Once
	file_aperture_policy_sync_v1_load_scheduler_proto_rawDescData = file_aperture_policy_sync_v1_load_scheduler_proto_rawDesc
)

func file_aperture_policy_sync_v1_load_scheduler_proto_rawDescGZIP() []byte {
	file_aperture_policy_sync_v1_load_scheduler_proto_rawDescOnce.Do(func() {
		file_aperture_policy_sync_v1_load_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_policy_sync_v1_load_scheduler_proto_rawDescData)
	})
	return file_aperture_policy_sync_v1_load_scheduler_proto_rawDescData
}

var file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aperture_policy_sync_v1_load_scheduler_proto_goTypes = []interface{}{
	(*LoadSchedulerWrapper)(nil), // 0: aperture.policy.sync.v1.LoadSchedulerWrapper
	(*LoadDecisionWrapper)(nil),  // 1: aperture.policy.sync.v1.LoadDecisionWrapper
	(*LoadDecision)(nil),         // 2: aperture.policy.sync.v1.LoadDecision
	nil,                          // 3: aperture.policy.sync.v1.LoadDecision.TokensByWorkloadIndexEntry
	(*CommonAttributes)(nil),     // 4: aperture.policy.sync.v1.CommonAttributes
	(*v1.LoadScheduler)(nil),     // 5: aperture.policy.language.v1.LoadScheduler
	(*TickInfo)(nil),             // 6: aperture.policy.sync.v1.TickInfo
}
var file_aperture_policy_sync_v1_load_scheduler_proto_depIdxs = []int32{
	4, // 0: aperture.policy.sync.v1.LoadSchedulerWrapper.common_attributes:type_name -> aperture.policy.sync.v1.CommonAttributes
	5, // 1: aperture.policy.sync.v1.LoadSchedulerWrapper.load_scheduler:type_name -> aperture.policy.language.v1.LoadScheduler
	4, // 2: aperture.policy.sync.v1.LoadDecisionWrapper.common_attributes:type_name -> aperture.policy.sync.v1.CommonAttributes
	2, // 3: aperture.policy.sync.v1.LoadDecisionWrapper.load_decision:type_name -> aperture.policy.sync.v1.LoadDecision
	6, // 4: aperture.policy.sync.v1.LoadDecision.tick_info:type_name -> aperture.policy.sync.v1.TickInfo
	3, // 5: aperture.policy.sync.v1.LoadDecision.tokens_by_workload_index:type_name -> aperture.policy.sync.v1.LoadDecision.TokensByWorkloadIndexEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_aperture_policy_sync_v1_load_scheduler_proto_init() }
func file_aperture_policy_sync_v1_load_scheduler_proto_init() {
	if File_aperture_policy_sync_v1_load_scheduler_proto != nil {
		return
	}
	file_aperture_policy_sync_v1_common_attributes_proto_init()
	file_aperture_policy_sync_v1_tick_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadSchedulerWrapper); i {
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
		file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadDecisionWrapper); i {
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
		file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadDecision); i {
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
			RawDescriptor: file_aperture_policy_sync_v1_load_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aperture_policy_sync_v1_load_scheduler_proto_goTypes,
		DependencyIndexes: file_aperture_policy_sync_v1_load_scheduler_proto_depIdxs,
		MessageInfos:      file_aperture_policy_sync_v1_load_scheduler_proto_msgTypes,
	}.Build()
	File_aperture_policy_sync_v1_load_scheduler_proto = out.File
	file_aperture_policy_sync_v1_load_scheduler_proto_rawDesc = nil
	file_aperture_policy_sync_v1_load_scheduler_proto_goTypes = nil
	file_aperture_policy_sync_v1_load_scheduler_proto_depIdxs = nil
}
