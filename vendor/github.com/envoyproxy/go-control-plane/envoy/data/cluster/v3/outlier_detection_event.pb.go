// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/cluster/v3/outlier_detection_event.proto

package envoy_data_cluster_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OutlierEjectionType int32

const (
	OutlierEjectionType_CONSECUTIVE_5XX                  OutlierEjectionType = 0
	OutlierEjectionType_CONSECUTIVE_GATEWAY_FAILURE      OutlierEjectionType = 1
	OutlierEjectionType_SUCCESS_RATE                     OutlierEjectionType = 2
	OutlierEjectionType_CONSECUTIVE_LOCAL_ORIGIN_FAILURE OutlierEjectionType = 3
	OutlierEjectionType_SUCCESS_RATE_LOCAL_ORIGIN        OutlierEjectionType = 4
	OutlierEjectionType_FAILURE_PERCENTAGE               OutlierEjectionType = 5
	OutlierEjectionType_FAILURE_PERCENTAGE_LOCAL_ORIGIN  OutlierEjectionType = 6
)

var OutlierEjectionType_name = map[int32]string{
	0: "CONSECUTIVE_5XX",
	1: "CONSECUTIVE_GATEWAY_FAILURE",
	2: "SUCCESS_RATE",
	3: "CONSECUTIVE_LOCAL_ORIGIN_FAILURE",
	4: "SUCCESS_RATE_LOCAL_ORIGIN",
	5: "FAILURE_PERCENTAGE",
	6: "FAILURE_PERCENTAGE_LOCAL_ORIGIN",
}

var OutlierEjectionType_value = map[string]int32{
	"CONSECUTIVE_5XX":                  0,
	"CONSECUTIVE_GATEWAY_FAILURE":      1,
	"SUCCESS_RATE":                     2,
	"CONSECUTIVE_LOCAL_ORIGIN_FAILURE": 3,
	"SUCCESS_RATE_LOCAL_ORIGIN":        4,
	"FAILURE_PERCENTAGE":               5,
	"FAILURE_PERCENTAGE_LOCAL_ORIGIN":  6,
}

func (x OutlierEjectionType) String() string {
	return proto.EnumName(OutlierEjectionType_name, int32(x))
}

func (OutlierEjectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b64fc638e24bdcb, []int{0}
}

type Action int32

const (
	Action_EJECT   Action = 0
	Action_UNEJECT Action = 1
)

var Action_name = map[int32]string{
	0: "EJECT",
	1: "UNEJECT",
}

var Action_value = map[string]int32{
	"EJECT":   0,
	"UNEJECT": 1,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b64fc638e24bdcb, []int{1}
}

type OutlierDetectionEvent struct {
	Type                OutlierEjectionType   `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.data.cluster.v3.OutlierEjectionType" json:"type,omitempty"`
	Timestamp           *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SecsSinceLastAction *wrappers.UInt64Value `protobuf:"bytes,3,opt,name=secs_since_last_action,json=secsSinceLastAction,proto3" json:"secs_since_last_action,omitempty"`
	ClusterName         string                `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	UpstreamUrl         string                `protobuf:"bytes,5,opt,name=upstream_url,json=upstreamUrl,proto3" json:"upstream_url,omitempty"`
	Action              Action                `protobuf:"varint,6,opt,name=action,proto3,enum=envoy.data.cluster.v3.Action" json:"action,omitempty"`
	NumEjections        uint32                `protobuf:"varint,7,opt,name=num_ejections,json=numEjections,proto3" json:"num_ejections,omitempty"`
	Enforced            bool                  `protobuf:"varint,8,opt,name=enforced,proto3" json:"enforced,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*OutlierDetectionEvent_EjectSuccessRateEvent
	//	*OutlierDetectionEvent_EjectConsecutiveEvent
	//	*OutlierDetectionEvent_EjectFailurePercentageEvent
	Event                isOutlierDetectionEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *OutlierDetectionEvent) Reset()         { *m = OutlierDetectionEvent{} }
func (m *OutlierDetectionEvent) String() string { return proto.CompactTextString(m) }
func (*OutlierDetectionEvent) ProtoMessage()    {}
func (*OutlierDetectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fc638e24bdcb, []int{0}
}

func (m *OutlierDetectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetectionEvent.Unmarshal(m, b)
}
func (m *OutlierDetectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetectionEvent.Marshal(b, m, deterministic)
}
func (m *OutlierDetectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetectionEvent.Merge(m, src)
}
func (m *OutlierDetectionEvent) XXX_Size() int {
	return xxx_messageInfo_OutlierDetectionEvent.Size(m)
}
func (m *OutlierDetectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetectionEvent proto.InternalMessageInfo

func (m *OutlierDetectionEvent) GetType() OutlierEjectionType {
	if m != nil {
		return m.Type
	}
	return OutlierEjectionType_CONSECUTIVE_5XX
}

func (m *OutlierDetectionEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *OutlierDetectionEvent) GetSecsSinceLastAction() *wrappers.UInt64Value {
	if m != nil {
		return m.SecsSinceLastAction
	}
	return nil
}

func (m *OutlierDetectionEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *OutlierDetectionEvent) GetUpstreamUrl() string {
	if m != nil {
		return m.UpstreamUrl
	}
	return ""
}

func (m *OutlierDetectionEvent) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_EJECT
}

func (m *OutlierDetectionEvent) GetNumEjections() uint32 {
	if m != nil {
		return m.NumEjections
	}
	return 0
}

func (m *OutlierDetectionEvent) GetEnforced() bool {
	if m != nil {
		return m.Enforced
	}
	return false
}

type isOutlierDetectionEvent_Event interface {
	isOutlierDetectionEvent_Event()
}

type OutlierDetectionEvent_EjectSuccessRateEvent struct {
	EjectSuccessRateEvent *OutlierEjectSuccessRate `protobuf:"bytes,9,opt,name=eject_success_rate_event,json=ejectSuccessRateEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectConsecutiveEvent struct {
	EjectConsecutiveEvent *OutlierEjectConsecutive `protobuf:"bytes,10,opt,name=eject_consecutive_event,json=ejectConsecutiveEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectFailurePercentageEvent struct {
	EjectFailurePercentageEvent *OutlierEjectFailurePercentage `protobuf:"bytes,11,opt,name=eject_failure_percentage_event,json=ejectFailurePercentageEvent,proto3,oneof"`
}

func (*OutlierDetectionEvent_EjectSuccessRateEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectConsecutiveEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectFailurePercentageEvent) isOutlierDetectionEvent_Event() {}

func (m *OutlierDetectionEvent) GetEvent() isOutlierDetectionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectSuccessRateEvent() *OutlierEjectSuccessRate {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectSuccessRateEvent); ok {
		return x.EjectSuccessRateEvent
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectConsecutiveEvent() *OutlierEjectConsecutive {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectConsecutiveEvent); ok {
		return x.EjectConsecutiveEvent
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectFailurePercentageEvent() *OutlierEjectFailurePercentage {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectFailurePercentageEvent); ok {
		return x.EjectFailurePercentageEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutlierDetectionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutlierDetectionEvent_EjectSuccessRateEvent)(nil),
		(*OutlierDetectionEvent_EjectConsecutiveEvent)(nil),
		(*OutlierDetectionEvent_EjectFailurePercentageEvent)(nil),
	}
}

type OutlierEjectSuccessRate struct {
	HostSuccessRate                     uint32   `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	ClusterAverageSuccessRate           uint32   `protobuf:"varint,2,opt,name=cluster_average_success_rate,json=clusterAverageSuccessRate,proto3" json:"cluster_average_success_rate,omitempty"`
	ClusterSuccessRateEjectionThreshold uint32   `protobuf:"varint,3,opt,name=cluster_success_rate_ejection_threshold,json=clusterSuccessRateEjectionThreshold,proto3" json:"cluster_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *OutlierEjectSuccessRate) Reset()         { *m = OutlierEjectSuccessRate{} }
func (m *OutlierEjectSuccessRate) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectSuccessRate) ProtoMessage()    {}
func (*OutlierEjectSuccessRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fc638e24bdcb, []int{1}
}

func (m *OutlierEjectSuccessRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectSuccessRate.Unmarshal(m, b)
}
func (m *OutlierEjectSuccessRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectSuccessRate.Marshal(b, m, deterministic)
}
func (m *OutlierEjectSuccessRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectSuccessRate.Merge(m, src)
}
func (m *OutlierEjectSuccessRate) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectSuccessRate.Size(m)
}
func (m *OutlierEjectSuccessRate) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectSuccessRate.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectSuccessRate proto.InternalMessageInfo

func (m *OutlierEjectSuccessRate) GetHostSuccessRate() uint32 {
	if m != nil {
		return m.HostSuccessRate
	}
	return 0
}

func (m *OutlierEjectSuccessRate) GetClusterAverageSuccessRate() uint32 {
	if m != nil {
		return m.ClusterAverageSuccessRate
	}
	return 0
}

func (m *OutlierEjectSuccessRate) GetClusterSuccessRateEjectionThreshold() uint32 {
	if m != nil {
		return m.ClusterSuccessRateEjectionThreshold
	}
	return 0
}

type OutlierEjectConsecutive struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutlierEjectConsecutive) Reset()         { *m = OutlierEjectConsecutive{} }
func (m *OutlierEjectConsecutive) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectConsecutive) ProtoMessage()    {}
func (*OutlierEjectConsecutive) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fc638e24bdcb, []int{2}
}

func (m *OutlierEjectConsecutive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectConsecutive.Unmarshal(m, b)
}
func (m *OutlierEjectConsecutive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectConsecutive.Marshal(b, m, deterministic)
}
func (m *OutlierEjectConsecutive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectConsecutive.Merge(m, src)
}
func (m *OutlierEjectConsecutive) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectConsecutive.Size(m)
}
func (m *OutlierEjectConsecutive) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectConsecutive.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectConsecutive proto.InternalMessageInfo

type OutlierEjectFailurePercentage struct {
	HostSuccessRate      uint32   `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutlierEjectFailurePercentage) Reset()         { *m = OutlierEjectFailurePercentage{} }
func (m *OutlierEjectFailurePercentage) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectFailurePercentage) ProtoMessage()    {}
func (*OutlierEjectFailurePercentage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b64fc638e24bdcb, []int{3}
}

func (m *OutlierEjectFailurePercentage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectFailurePercentage.Unmarshal(m, b)
}
func (m *OutlierEjectFailurePercentage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectFailurePercentage.Marshal(b, m, deterministic)
}
func (m *OutlierEjectFailurePercentage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectFailurePercentage.Merge(m, src)
}
func (m *OutlierEjectFailurePercentage) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectFailurePercentage.Size(m)
}
func (m *OutlierEjectFailurePercentage) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectFailurePercentage.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectFailurePercentage proto.InternalMessageInfo

func (m *OutlierEjectFailurePercentage) GetHostSuccessRate() uint32 {
	if m != nil {
		return m.HostSuccessRate
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.data.cluster.v3.OutlierEjectionType", OutlierEjectionType_name, OutlierEjectionType_value)
	proto.RegisterEnum("envoy.data.cluster.v3.Action", Action_name, Action_value)
	proto.RegisterType((*OutlierDetectionEvent)(nil), "envoy.data.cluster.v3.OutlierDetectionEvent")
	proto.RegisterType((*OutlierEjectSuccessRate)(nil), "envoy.data.cluster.v3.OutlierEjectSuccessRate")
	proto.RegisterType((*OutlierEjectConsecutive)(nil), "envoy.data.cluster.v3.OutlierEjectConsecutive")
	proto.RegisterType((*OutlierEjectFailurePercentage)(nil), "envoy.data.cluster.v3.OutlierEjectFailurePercentage")
}

func init() {
	proto.RegisterFile("envoy/data/cluster/v3/outlier_detection_event.proto", fileDescriptor_5b64fc638e24bdcb)
}

var fileDescriptor_5b64fc638e24bdcb = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xaf, 0xdd, 0x36, 0x6d, 0x5f, 0x5b, 0xd6, 0x4c, 0xd5, 0xad, 0x37, 0xbb, 0x6d, 0x43, 0x8b,
	0x44, 0x94, 0x83, 0x8d, 0x92, 0x02, 0x4b, 0x2f, 0xab, 0x24, 0x78, 0xdb, 0x40, 0x95, 0x16, 0x27,
	0x59, 0x96, 0x03, 0x1a, 0xcd, 0x3a, 0xaf, 0x8d, 0x91, 0x63, 0x5b, 0x33, 0xe3, 0x40, 0xc5, 0x05,
	0x71, 0xe2, 0x0b, 0x70, 0xe1, 0x63, 0x70, 0xe4, 0x8e, 0xc4, 0x95, 0x4f, 0x80, 0xf8, 0x16, 0x68,
	0x4f, 0xc8, 0xff, 0xb2, 0x4e, 0x9b, 0x85, 0xb2, 0xb7, 0x64, 0xde, 0xef, 0xf7, 0x7e, 0xef, 0xfd,
	0xe6, 0xbd, 0x31, 0x34, 0xd0, 0x9f, 0x04, 0xd7, 0xe6, 0x90, 0x49, 0x66, 0x3a, 0x5e, 0x24, 0x24,
	0x72, 0x73, 0xd2, 0x30, 0x83, 0x48, 0x7a, 0x2e, 0x72, 0x3a, 0x44, 0x89, 0x8e, 0x74, 0x03, 0x9f,
	0xe2, 0x04, 0x7d, 0x69, 0x84, 0x3c, 0x90, 0x01, 0xd9, 0x4e, 0x48, 0x46, 0x4c, 0x32, 0x32, 0x92,
	0x31, 0x69, 0x94, 0xf7, 0xaf, 0x82, 0xe0, 0xca, 0x43, 0x33, 0x01, 0xbd, 0x88, 0x2e, 0x4d, 0xe9,
	0x8e, 0x51, 0x48, 0x36, 0x0e, 0x53, 0x5e, 0x79, 0xef, 0x26, 0xe0, 0x1b, 0xce, 0xc2, 0x10, 0xb9,
	0xc8, 0xe2, 0xbb, 0xd1, 0x30, 0x64, 0x26, 0xf3, 0xfd, 0x40, 0xb2, 0x58, 0x54, 0x98, 0x42, 0x32,
	0x19, 0xe5, 0xe1, 0x77, 0x6e, 0x85, 0x27, 0xc8, 0x85, 0x1b, 0xf8, 0xae, 0x7f, 0x95, 0x41, 0x76,
	0x26, 0xcc, 0x73, 0x87, 0x4c, 0xa2, 0x99, 0xff, 0x48, 0x03, 0x07, 0x7f, 0x95, 0x60, 0xfb, 0x3c,
	0x6d, 0xea, 0x93, 0xbc, 0x27, 0x2b, 0x6e, 0x89, 0x9c, 0xc2, 0x92, 0xbc, 0x0e, 0x51, 0x57, 0x2a,
	0x4a, 0xf5, 0xad, 0x7a, 0xcd, 0x98, 0xdb, 0x9b, 0x91, 0x71, 0xad, 0xaf, 0x53, 0x6a, 0xff, 0x3a,
	0xc4, 0xd6, 0xea, 0xcb, 0xd6, 0xf2, 0x0f, 0x8a, 0xaa, 0x29, 0x76, 0x92, 0x81, 0x3c, 0x86, 0xb5,
	0x69, 0xc7, 0xba, 0x5a, 0x51, 0xaa, 0xeb, 0xf5, 0xb2, 0x91, 0xb6, 0x6c, 0xe4, 0x2d, 0x1b, 0xfd,
	0x1c, 0x61, 0xbf, 0x02, 0x93, 0xcf, 0xe1, 0xbe, 0x40, 0x47, 0x50, 0xe1, 0xfa, 0x0e, 0x52, 0x8f,
	0x09, 0x49, 0x59, 0xa2, 0xa3, 0x2f, 0x26, 0x69, 0x1e, 0xdd, 0x4a, 0x33, 0xe8, 0xf8, 0xf2, 0xc3,
	0xa3, 0x67, 0xcc, 0x8b, 0xd0, 0xde, 0x8a, 0xb9, 0xbd, 0x98, 0x7a, 0xc6, 0x84, 0x6c, 0x26, 0x44,
	0x52, 0x83, 0x8d, 0xac, 0x7c, 0xea, 0xb3, 0x31, 0xea, 0x4b, 0x15, 0xa5, 0xba, 0xd6, 0x5a, 0x79,
	0xd9, 0x5a, 0xe2, 0x6a, 0x45, 0xb1, 0xd7, 0xb3, 0x60, 0x97, 0x8d, 0x31, 0xc6, 0x46, 0xa1, 0x90,
	0x1c, 0xd9, 0x98, 0x46, 0xdc, 0xd3, 0x97, 0x6f, 0x60, 0xf3, 0xe0, 0x80, 0x7b, 0xe4, 0x09, 0x94,
	0xb2, 0xd2, 0x4a, 0x89, 0x61, 0xbb, 0xaf, 0x31, 0x2c, 0x2d, 0xa3, 0xe0, 0x51, 0x46, 0x23, 0x87,
	0xb0, 0xe9, 0x47, 0x63, 0x8a, 0x99, 0x93, 0x42, 0x5f, 0xa9, 0x28, 0xd5, 0x4d, 0x7b, 0xc3, 0x8f,
	0xc6, 0xb9, 0xbb, 0x82, 0x94, 0x61, 0x15, 0xfd, 0xcb, 0x80, 0x3b, 0x38, 0xd4, 0x57, 0x2b, 0x4a,
	0x75, 0xd5, 0x9e, 0xfe, 0x27, 0x2e, 0xe8, 0x09, 0x99, 0x8a, 0xc8, 0x71, 0x50, 0x08, 0xca, 0x99,
	0xc4, 0x74, 0x3e, 0xf5, 0xb5, 0xc4, 0x2e, 0xe3, 0x0e, 0x97, 0xd8, 0x4b, 0xc9, 0x36, 0x93, 0x78,
	0xba, 0x60, 0x6f, 0xe3, 0x8d, 0xb3, 0x74, 0x36, 0x46, 0xb0, 0x93, 0x4a, 0x39, 0x81, 0x2f, 0xd0,
	0x89, 0xa4, 0x3b, 0xc9, 0x95, 0xe0, 0xce, 0x4a, 0xed, 0x57, 0xdc, 0xa9, 0x52, 0xe1, 0x2c, 0x55,
	0xfa, 0x0e, 0xf6, 0x52, 0xa5, 0x4b, 0xe6, 0x7a, 0x11, 0x47, 0x1a, 0x22, 0x77, 0xd0, 0x97, 0xec,
	0x2a, 0x17, 0x5c, 0x4f, 0x04, 0x8f, 0xee, 0x20, 0xf8, 0x34, 0x4d, 0x71, 0x31, 0xcd, 0x70, 0xba,
	0x60, 0x3f, 0xc4, 0xb9, 0x91, 0x44, 0xfc, 0xf8, 0xa3, 0x9f, 0x7f, 0xfb, 0x71, 0xaf, 0x0e, 0xef,
	0xcf, 0x4b, 0x5d, 0x67, 0x5e, 0x38, 0x62, 0xc6, 0xdc, 0xdd, 0x69, 0x6d, 0xc0, 0x72, 0x52, 0x1c,
	0x59, 0xfc, 0xbb, 0xa5, 0x1c, 0xfc, 0xa2, 0xc2, 0xce, 0x6b, 0x2c, 0x26, 0x0d, 0x78, 0x7b, 0x14,
	0x88, 0xd9, 0x3b, 0x4b, 0x56, 0x6e, 0x33, 0x99, 0xb3, 0x9a, 0xaa, 0x0f, 0xed, 0x7b, 0x31, 0xa2,
	0x48, 0x3a, 0x85, 0x47, 0xf9, 0x0c, 0xb3, 0x09, 0xf2, 0xd8, 0x8b, 0x19, 0xbe, 0x3a, 0xcb, 0x7f,
	0x90, 0x81, 0x9b, 0x29, 0xb6, 0x98, 0xe9, 0x2b, 0x78, 0x2f, 0xcf, 0x34, 0x3b, 0x35, 0xd9, 0xc4,
	0x51, 0x39, 0xe2, 0x28, 0x46, 0x81, 0x37, 0x4c, 0x36, 0xae, 0x90, 0xf4, 0x30, 0xe3, 0x15, 0x67,
	0x23, 0x7f, 0x04, 0x72, 0xce, 0xf1, 0xc7, 0xb1, 0x81, 0x47, 0x50, 0xff, 0x6f, 0x03, 0x6f, 0x1a,
	0x73, 0xd0, 0x9f, 0xf5, 0xac, 0x30, 0x18, 0xff, 0x37, 0x6b, 0x81, 0x7a, 0xf0, 0x93, 0x02, 0xbb,
	0xff, 0x3a, 0x12, 0x6f, 0x74, 0x21, 0xc7, 0x4f, 0xe2, 0x8a, 0x8e, 0xe1, 0xf1, 0x1d, 0x2b, 0xba,
	0xa5, 0x5a, 0xfb, 0x53, 0x81, 0xad, 0x39, 0x4f, 0x29, 0xd9, 0x82, 0x7b, 0xed, 0xf3, 0x6e, 0xcf,
	0x6a, 0x0f, 0xfa, 0x9d, 0x67, 0x16, 0xfd, 0xe0, 0xf9, 0x73, 0x6d, 0x81, 0xec, 0xc3, 0xc3, 0xe2,
	0xe1, 0x49, 0xb3, 0x6f, 0x7d, 0xd1, 0xfc, 0x92, 0x3e, 0x6d, 0x76, 0xce, 0x06, 0xb6, 0xa5, 0x29,
	0x44, 0x83, 0x8d, 0xde, 0xa0, 0xdd, 0xb6, 0x7a, 0x3d, 0x6a, 0x37, 0xfb, 0x96, 0xa6, 0x92, 0x77,
	0xa1, 0x52, 0xa4, 0x9c, 0x9d, 0xb7, 0x9b, 0x67, 0xf4, 0xdc, 0xee, 0x9c, 0x74, 0xba, 0x53, 0xde,
	0x22, 0xd9, 0x85, 0x07, 0x45, 0xde, 0x0c, 0x4c, 0x5b, 0x22, 0xf7, 0x81, 0x64, 0x58, 0x7a, 0x61,
	0xd9, 0x6d, 0xab, 0xdb, 0x6f, 0x9e, 0x58, 0xda, 0x32, 0x39, 0x84, 0xfd, 0xdb, 0xe7, 0xb3, 0xe4,
	0x52, 0xad, 0x02, 0xa5, 0xec, 0x05, 0x5e, 0x83, 0x65, 0xeb, 0x53, 0xab, 0xdd, 0xd7, 0x16, 0xc8,
	0x3a, 0xac, 0x0c, 0xba, 0xe9, 0x1f, 0xa5, 0xf5, 0xd9, 0xaf, 0xdf, 0xff, 0xfe, 0x47, 0x49, 0xd5,
	0x54, 0x38, 0x74, 0x83, 0x74, 0x9d, 0x43, 0x1e, 0x7c, 0x7b, 0x3d, 0x7f, 0xb3, 0x5b, 0xe5, 0xb9,
	0xab, 0x77, 0x11, 0x7f, 0x08, 0x2e, 0x94, 0x17, 0xa5, 0xe4, 0x8b, 0xd0, 0xf8, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xa1, 0xa0, 0x15, 0x66, 0xc8, 0x07, 0x00, 0x00,
}