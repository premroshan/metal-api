// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/network.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Network struct {
	Common               *Common               `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	PartitionID          *wrappers.StringValue `protobuf:"bytes,2,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	ProjectID            *wrappers.StringValue `protobuf:"bytes,3,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Labels               map[string]string     `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{0}
}

func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (m *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(m, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetCommon() *Common {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *Network) GetPartitionID() *wrappers.StringValue {
	if m != nil {
		return m.PartitionID
	}
	return nil
}

func (m *Network) GetProjectID() *wrappers.StringValue {
	if m != nil {
		return m.ProjectID
	}
	return nil
}

func (m *Network) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// a network which contains prefixes from which IP addresses can be allocated
type NetworkImmutable struct {
	Prefixes             []string              `protobuf:"bytes,1,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	DestinationPrefixes  []string              `protobuf:"bytes,2,rep,name=destinationPrefixes,proto3" json:"destinationPrefixes,omitempty"`
	Nat                  bool                  `protobuf:"varint,3,opt,name=nat,proto3" json:"nat,omitempty"`
	PrivateSuper         bool                  `protobuf:"varint,4,opt,name=privateSuper,proto3" json:"privateSuper,omitempty"`
	Underlay             bool                  `protobuf:"varint,5,opt,name=underlay,proto3" json:"underlay,omitempty"`
	Vrf                  *wrappers.UInt64Value `protobuf:"bytes,6,opt,name=vrf,proto3" json:"vrf,omitempty"`
	VrfShared            *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=vrfShared,proto3" json:"vrfShared,omitempty"`
	ParentNetworkID      *wrappers.StringValue `protobuf:"bytes,8,opt,name=parentNetworkID,proto3" json:"parentNetworkID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NetworkImmutable) Reset()         { *m = NetworkImmutable{} }
func (m *NetworkImmutable) String() string { return proto.CompactTextString(m) }
func (*NetworkImmutable) ProtoMessage()    {}
func (*NetworkImmutable) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{1}
}

func (m *NetworkImmutable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkImmutable.Unmarshal(m, b)
}
func (m *NetworkImmutable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkImmutable.Marshal(b, m, deterministic)
}
func (m *NetworkImmutable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkImmutable.Merge(m, src)
}
func (m *NetworkImmutable) XXX_Size() int {
	return xxx_messageInfo_NetworkImmutable.Size(m)
}
func (m *NetworkImmutable) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkImmutable.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkImmutable proto.InternalMessageInfo

func (m *NetworkImmutable) GetPrefixes() []string {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

func (m *NetworkImmutable) GetDestinationPrefixes() []string {
	if m != nil {
		return m.DestinationPrefixes
	}
	return nil
}

func (m *NetworkImmutable) GetNat() bool {
	if m != nil {
		return m.Nat
	}
	return false
}

func (m *NetworkImmutable) GetPrivateSuper() bool {
	if m != nil {
		return m.PrivateSuper
	}
	return false
}

func (m *NetworkImmutable) GetUnderlay() bool {
	if m != nil {
		return m.Underlay
	}
	return false
}

func (m *NetworkImmutable) GetVrf() *wrappers.UInt64Value {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *NetworkImmutable) GetVrfShared() *wrappers.BoolValue {
	if m != nil {
		return m.VrfShared
	}
	return nil
}

func (m *NetworkImmutable) GetParentNetworkID() *wrappers.StringValue {
	if m != nil {
		return m.ParentNetworkID
	}
	return nil
}

type NetworkUsage struct {
	AvailableIPs         uint64   `protobuf:"varint,1,opt,name=availableIPs,proto3" json:"availableIPs,omitempty"`
	UsedIPs              uint64   `protobuf:"varint,2,opt,name=usedIPs,proto3" json:"usedIPs,omitempty"`
	AvailablePrefixes    uint64   `protobuf:"varint,3,opt,name=availablePrefixes,proto3" json:"availablePrefixes,omitempty"`
	UsedPrefixes         uint64   `protobuf:"varint,4,opt,name=usedPrefixes,proto3" json:"usedPrefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkUsage) Reset()         { *m = NetworkUsage{} }
func (m *NetworkUsage) String() string { return proto.CompactTextString(m) }
func (*NetworkUsage) ProtoMessage()    {}
func (*NetworkUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{2}
}

func (m *NetworkUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkUsage.Unmarshal(m, b)
}
func (m *NetworkUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkUsage.Marshal(b, m, deterministic)
}
func (m *NetworkUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkUsage.Merge(m, src)
}
func (m *NetworkUsage) XXX_Size() int {
	return xxx_messageInfo_NetworkUsage.Size(m)
}
func (m *NetworkUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkUsage.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkUsage proto.InternalMessageInfo

func (m *NetworkUsage) GetAvailableIPs() uint64 {
	if m != nil {
		return m.AvailableIPs
	}
	return 0
}

func (m *NetworkUsage) GetUsedIPs() uint64 {
	if m != nil {
		return m.UsedIPs
	}
	return 0
}

func (m *NetworkUsage) GetAvailablePrefixes() uint64 {
	if m != nil {
		return m.AvailablePrefixes
	}
	return 0
}

func (m *NetworkUsage) GetUsedPrefixes() uint64 {
	if m != nil {
		return m.UsedPrefixes
	}
	return 0
}

// NetworkSearchQuery can be used to search networks.
type NetworkSearchQuery struct {
	ID                   *wrappers.StringValue   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 *wrappers.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PartitionID          *wrappers.StringValue   `protobuf:"bytes,3,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	ProjectID            *wrappers.StringValue   `protobuf:"bytes,4,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Prefixes             []*wrappers.StringValue `protobuf:"bytes,5,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	DestinationPrefixes  []*wrappers.StringValue `protobuf:"bytes,6,rep,name=destinationPrefixes,proto3" json:"destinationPrefixes,omitempty"`
	Nat                  *wrappers.BoolValue     `protobuf:"bytes,7,opt,name=nat,proto3" json:"nat,omitempty"`
	PrivateSuper         *wrappers.BoolValue     `protobuf:"bytes,8,opt,name=privateSuper,proto3" json:"privateSuper,omitempty"`
	Underlay             *wrappers.BoolValue     `protobuf:"bytes,9,opt,name=underlay,proto3" json:"underlay,omitempty"`
	Vrf                  *wrappers.UInt64Value   `protobuf:"bytes,10,opt,name=vrf,proto3" json:"vrf,omitempty"`
	VrfShared            *wrappers.BoolValue     `protobuf:"bytes,11,opt,name=vrfShared,proto3" json:"vrfShared,omitempty"`
	ParentNetworkID      *wrappers.StringValue   `protobuf:"bytes,12,opt,name=parentNetworkID,proto3" json:"parentNetworkID,omitempty"`
	Labels               map[string]string       `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NetworkSearchQuery) Reset()         { *m = NetworkSearchQuery{} }
func (m *NetworkSearchQuery) String() string { return proto.CompactTextString(m) }
func (*NetworkSearchQuery) ProtoMessage()    {}
func (*NetworkSearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{3}
}

func (m *NetworkSearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkSearchQuery.Unmarshal(m, b)
}
func (m *NetworkSearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkSearchQuery.Marshal(b, m, deterministic)
}
func (m *NetworkSearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkSearchQuery.Merge(m, src)
}
func (m *NetworkSearchQuery) XXX_Size() int {
	return xxx_messageInfo_NetworkSearchQuery.Size(m)
}
func (m *NetworkSearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkSearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkSearchQuery proto.InternalMessageInfo

func (m *NetworkSearchQuery) GetID() *wrappers.StringValue {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *NetworkSearchQuery) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *NetworkSearchQuery) GetPartitionID() *wrappers.StringValue {
	if m != nil {
		return m.PartitionID
	}
	return nil
}

func (m *NetworkSearchQuery) GetProjectID() *wrappers.StringValue {
	if m != nil {
		return m.ProjectID
	}
	return nil
}

func (m *NetworkSearchQuery) GetPrefixes() []*wrappers.StringValue {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

func (m *NetworkSearchQuery) GetDestinationPrefixes() []*wrappers.StringValue {
	if m != nil {
		return m.DestinationPrefixes
	}
	return nil
}

func (m *NetworkSearchQuery) GetNat() *wrappers.BoolValue {
	if m != nil {
		return m.Nat
	}
	return nil
}

func (m *NetworkSearchQuery) GetPrivateSuper() *wrappers.BoolValue {
	if m != nil {
		return m.PrivateSuper
	}
	return nil
}

func (m *NetworkSearchQuery) GetUnderlay() *wrappers.BoolValue {
	if m != nil {
		return m.Underlay
	}
	return nil
}

func (m *NetworkSearchQuery) GetVrf() *wrappers.UInt64Value {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *NetworkSearchQuery) GetVrfShared() *wrappers.BoolValue {
	if m != nil {
		return m.VrfShared
	}
	return nil
}

func (m *NetworkSearchQuery) GetParentNetworkID() *wrappers.StringValue {
	if m != nil {
		return m.ParentNetworkID
	}
	return nil
}

func (m *NetworkSearchQuery) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type NetworkCreateRequest struct {
	Network              *Network          `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	NetworkImmutable     *NetworkImmutable `protobuf:"bytes,2,opt,name=networkImmutable,proto3" json:"networkImmutable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NetworkCreateRequest) Reset()         { *m = NetworkCreateRequest{} }
func (m *NetworkCreateRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkCreateRequest) ProtoMessage()    {}
func (*NetworkCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{4}
}

func (m *NetworkCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkCreateRequest.Unmarshal(m, b)
}
func (m *NetworkCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkCreateRequest.Marshal(b, m, deterministic)
}
func (m *NetworkCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkCreateRequest.Merge(m, src)
}
func (m *NetworkCreateRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkCreateRequest.Size(m)
}
func (m *NetworkCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkCreateRequest proto.InternalMessageInfo

func (m *NetworkCreateRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *NetworkCreateRequest) GetNetworkImmutable() *NetworkImmutable {
	if m != nil {
		return m.NetworkImmutable
	}
	return nil
}

type NetworkUpdateRequest struct {
	Network              *Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Prefixes             []string `protobuf:"bytes,2,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkUpdateRequest) Reset()         { *m = NetworkUpdateRequest{} }
func (m *NetworkUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkUpdateRequest) ProtoMessage()    {}
func (*NetworkUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{5}
}

func (m *NetworkUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkUpdateRequest.Unmarshal(m, b)
}
func (m *NetworkUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkUpdateRequest.Marshal(b, m, deterministic)
}
func (m *NetworkUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkUpdateRequest.Merge(m, src)
}
func (m *NetworkUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkUpdateRequest.Size(m)
}
func (m *NetworkUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkUpdateRequest proto.InternalMessageInfo

func (m *NetworkUpdateRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *NetworkUpdateRequest) GetPrefixes() []string {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type NetworkFindRequest struct {
	NetworkSearchQuery   *NetworkSearchQuery `protobuf:"bytes,1,opt,name=networkSearchQuery,proto3" json:"networkSearchQuery,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkFindRequest) Reset()         { *m = NetworkFindRequest{} }
func (m *NetworkFindRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkFindRequest) ProtoMessage()    {}
func (*NetworkFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{6}
}

func (m *NetworkFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkFindRequest.Unmarshal(m, b)
}
func (m *NetworkFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkFindRequest.Marshal(b, m, deterministic)
}
func (m *NetworkFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkFindRequest.Merge(m, src)
}
func (m *NetworkFindRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkFindRequest.Size(m)
}
func (m *NetworkFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkFindRequest proto.InternalMessageInfo

func (m *NetworkFindRequest) GetNetworkSearchQuery() *NetworkSearchQuery {
	if m != nil {
		return m.NetworkSearchQuery
	}
	return nil
}

type NetworkResponse struct {
	Network              *Network          `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	NetworkImmutable     *NetworkImmutable `protobuf:"bytes,2,opt,name=networkImmutable,proto3" json:"networkImmutable,omitempty"`
	Usage                *NetworkUsage     `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NetworkResponse) Reset()         { *m = NetworkResponse{} }
func (m *NetworkResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkResponse) ProtoMessage()    {}
func (*NetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{7}
}

func (m *NetworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkResponse.Unmarshal(m, b)
}
func (m *NetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkResponse.Marshal(b, m, deterministic)
}
func (m *NetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkResponse.Merge(m, src)
}
func (m *NetworkResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkResponse.Size(m)
}
func (m *NetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkResponse proto.InternalMessageInfo

func (m *NetworkResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *NetworkResponse) GetNetworkImmutable() *NetworkImmutable {
	if m != nil {
		return m.NetworkImmutable
	}
	return nil
}

func (m *NetworkResponse) GetUsage() *NetworkUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

type NetworkListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkListRequest) Reset()         { *m = NetworkListRequest{} }
func (m *NetworkListRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkListRequest) ProtoMessage()    {}
func (*NetworkListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{8}
}

func (m *NetworkListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkListRequest.Unmarshal(m, b)
}
func (m *NetworkListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkListRequest.Marshal(b, m, deterministic)
}
func (m *NetworkListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkListRequest.Merge(m, src)
}
func (m *NetworkListRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkListRequest.Size(m)
}
func (m *NetworkListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkListRequest proto.InternalMessageInfo

type NetworkListResponse struct {
	Networks             []*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NetworkListResponse) Reset()         { *m = NetworkListResponse{} }
func (m *NetworkListResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkListResponse) ProtoMessage()    {}
func (*NetworkListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ef602c4c85062d, []int{9}
}

func (m *NetworkListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkListResponse.Unmarshal(m, b)
}
func (m *NetworkListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkListResponse.Marshal(b, m, deterministic)
}
func (m *NetworkListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkListResponse.Merge(m, src)
}
func (m *NetworkListResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkListResponse.Size(m)
}
func (m *NetworkListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkListResponse proto.InternalMessageInfo

func (m *NetworkListResponse) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

func init() {
	proto.RegisterType((*Network)(nil), "v1.Network")
	proto.RegisterMapType((map[string]string)(nil), "v1.Network.LabelsEntry")
	proto.RegisterType((*NetworkImmutable)(nil), "v1.NetworkImmutable")
	proto.RegisterType((*NetworkUsage)(nil), "v1.NetworkUsage")
	proto.RegisterType((*NetworkSearchQuery)(nil), "v1.NetworkSearchQuery")
	proto.RegisterMapType((map[string]string)(nil), "v1.NetworkSearchQuery.LabelsEntry")
	proto.RegisterType((*NetworkCreateRequest)(nil), "v1.NetworkCreateRequest")
	proto.RegisterType((*NetworkUpdateRequest)(nil), "v1.NetworkUpdateRequest")
	proto.RegisterType((*NetworkFindRequest)(nil), "v1.NetworkFindRequest")
	proto.RegisterType((*NetworkResponse)(nil), "v1.NetworkResponse")
	proto.RegisterType((*NetworkListRequest)(nil), "v1.NetworkListRequest")
	proto.RegisterType((*NetworkListResponse)(nil), "v1.NetworkListResponse")
}

func init() { proto.RegisterFile("v1/network.proto", fileDescriptor_77ef602c4c85062d) }

var fileDescriptor_77ef602c4c85062d = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xd1, 0x6e, 0xeb, 0x44,
	0x10, 0x55, 0x1c, 0xc7, 0x49, 0x26, 0x85, 0x86, 0x6d, 0x04, 0x56, 0x84, 0x50, 0x65, 0x09, 0xe8,
	0x43, 0xaf, 0x7d, 0x53, 0xd0, 0x6d, 0xe9, 0x43, 0x85, 0xee, 0x0d, 0x95, 0x22, 0x55, 0x55, 0x71,
	0x54, 0x24, 0x10, 0x2f, 0x9b, 0x64, 0x93, 0x9a, 0x38, 0xb6, 0x59, 0xaf, 0x5d, 0xf2, 0xc4, 0x57,
	0xf0, 0xca, 0x03, 0x7f, 0xc3, 0x7f, 0xf0, 0x09, 0x7c, 0x00, 0xda, 0xf5, 0xda, 0x5d, 0x27, 0x69,
	0x9b, 0xb6, 0xe8, 0xbe, 0xed, 0xce, 0xcc, 0xd9, 0x99, 0x3d, 0xb3, 0x67, 0x6c, 0x68, 0xa7, 0x3d,
	0x27, 0x20, 0xec, 0x36, 0xa4, 0x73, 0x3b, 0xa2, 0x21, 0x0b, 0x91, 0x96, 0xf6, 0xba, 0xbb, 0x69,
	0xcf, 0x19, 0x87, 0x8b, 0x45, 0x18, 0x64, 0xc6, 0xee, 0x67, 0xb3, 0x30, 0x9c, 0xf9, 0xc4, 0x11,
	0xbb, 0x51, 0x32, 0x75, 0x6e, 0x29, 0x8e, 0x22, 0x42, 0xe3, 0xcc, 0x6f, 0xfd, 0xa1, 0x41, 0xfd,
	0x32, 0x3b, 0x06, 0x59, 0x60, 0x64, 0x58, 0xb3, 0xb2, 0x5f, 0x39, 0x68, 0x1d, 0x81, 0x9d, 0xf6,
	0xec, 0x77, 0xc2, 0xe2, 0x4a, 0x0f, 0x3a, 0x83, 0x56, 0x84, 0x29, 0xf3, 0x98, 0x17, 0x06, 0x83,
	0xbe, 0xa9, 0x89, 0xc0, 0x4f, 0xed, 0x2c, 0x8b, 0x9d, 0x67, 0xb1, 0x87, 0x8c, 0x7a, 0xc1, 0xec,
	0x07, 0xec, 0x27, 0xc4, 0x55, 0x01, 0xe8, 0x14, 0x9a, 0x11, 0x0d, 0x7f, 0x21, 0x63, 0x36, 0xe8,
	0x9b, 0xd5, 0x2d, 0xd0, 0x77, 0xe1, 0xc8, 0x01, 0xc3, 0xc7, 0x23, 0xe2, 0xc7, 0xa6, 0xbe, 0x5f,
	0x3d, 0x68, 0x1d, 0x7d, 0xc2, 0xeb, 0x93, 0xc5, 0xdb, 0x17, 0xc2, 0xf3, 0x5d, 0xc0, 0xe8, 0xd2,
	0x95, 0x61, 0xdd, 0x6f, 0xa0, 0xa5, 0x98, 0x51, 0x1b, 0xaa, 0x73, 0xb2, 0x14, 0x97, 0x6b, 0xba,
	0x7c, 0x89, 0x3a, 0x50, 0x4b, 0x79, 0x16, 0x71, 0x8f, 0xa6, 0x9b, 0x6d, 0x4e, 0xb5, 0x93, 0x8a,
	0xf5, 0x8f, 0x06, 0x6d, 0x79, 0xf4, 0x60, 0xb1, 0x48, 0x18, 0x1e, 0xf9, 0x04, 0x75, 0xa1, 0x11,
	0x51, 0x32, 0xf5, 0x7e, 0x23, 0xb1, 0x59, 0xd9, 0xaf, 0x1e, 0x34, 0xdd, 0x62, 0x8f, 0x5e, 0xc3,
	0xde, 0x84, 0xc4, 0xcc, 0x0b, 0x30, 0xbf, 0xe9, 0x55, 0x1e, 0xa6, 0x89, 0xb0, 0x4d, 0x2e, 0x5e,
	0x4e, 0x80, 0x99, 0x20, 0xa1, 0xe1, 0xf2, 0x25, 0xb2, 0x60, 0x27, 0xa2, 0x5e, 0x8a, 0x19, 0x19,
	0x26, 0x11, 0xa1, 0xa6, 0x2e, 0x5c, 0x25, 0x1b, 0xaf, 0x21, 0x09, 0x26, 0x84, 0xfa, 0x78, 0x69,
	0xd6, 0x84, 0xbf, 0xd8, 0x23, 0x1b, 0xaa, 0x29, 0x9d, 0x9a, 0xc6, 0x3d, 0xb4, 0x5e, 0x0f, 0x02,
	0xf6, 0xe6, 0xeb, 0x8c, 0x56, 0x1e, 0x88, 0x4e, 0xa0, 0x99, 0xd2, 0xe9, 0xf0, 0x06, 0x53, 0x32,
	0x31, 0xeb, 0x02, 0xd5, 0x5d, 0x43, 0xbd, 0x0d, 0x43, 0x5f, 0xb6, 0xa2, 0x08, 0x46, 0xe7, 0xb0,
	0x1b, 0x61, 0x4a, 0x02, 0x96, 0x73, 0xd4, 0x37, 0x1b, 0x5b, 0x34, 0x73, 0x15, 0x64, 0xfd, 0x59,
	0x81, 0x1d, 0xb9, 0xbb, 0x8e, 0xf1, 0x8c, 0x70, 0x0a, 0x70, 0x8a, 0x3d, 0x9f, 0xf3, 0x3d, 0xb8,
	0x8a, 0x45, 0xb3, 0x74, 0xb7, 0x64, 0x43, 0x26, 0xd4, 0x93, 0x98, 0x4c, 0xb8, 0x5b, 0x13, 0xee,
	0x7c, 0x8b, 0x0e, 0xe1, 0xa3, 0x22, 0xb2, 0x68, 0x41, 0x55, 0xc4, 0xac, 0x3b, 0x78, 0x2e, 0x0e,
	0x2c, 0x02, 0xf5, 0x2c, 0x97, 0x6a, 0xb3, 0xfe, 0x36, 0x00, 0xc9, 0x02, 0x87, 0x04, 0xd3, 0xf1,
	0xcd, 0xf7, 0x09, 0xa1, 0x4b, 0x74, 0x08, 0xda, 0xa0, 0x2f, 0x65, 0xf2, 0xf0, 0x95, 0xb5, 0x41,
	0x1f, 0xbd, 0x06, 0x3d, 0xc0, 0x0b, 0xb2, 0x95, 0x5a, 0x44, 0xe4, 0xaa, 0xcc, 0xaa, 0x2f, 0x92,
	0x99, 0xfe, 0x34, 0x99, 0x9d, 0x28, 0xaf, 0xbc, 0x26, 0x84, 0xf6, 0x30, 0xf4, 0x4e, 0x03, 0x97,
	0x9b, 0x35, 0x60, 0x6c, 0x71, 0xc8, 0x46, 0x85, 0x1c, 0x66, 0x0a, 0x79, 0xfc, 0x65, 0x0a, 0xf5,
	0x9c, 0xad, 0xa8, 0xa7, 0xf1, 0x28, 0xac, 0xac, 0xac, 0x37, 0x8a, 0xb2, 0x9a, 0x8f, 0x62, 0xd7,
	0x54, 0x07, 0xcf, 0x52, 0x5d, 0xeb, 0x85, 0xaa, 0xdb, 0x79, 0x86, 0xea, 0xd0, 0x69, 0x31, 0x48,
	0x3f, 0x10, 0xad, 0xb1, 0x94, 0x41, 0xaa, 0xbc, 0xf2, 0xff, 0x7b, 0xa6, 0xfe, 0x0e, 0x1d, 0x99,
	0xe4, 0x1d, 0x25, 0x98, 0x11, 0x97, 0xfc, 0x9a, 0x90, 0x98, 0xa1, 0xcf, 0xa1, 0x2e, 0xbf, 0x64,
	0x52, 0x51, 0x2d, 0xa5, 0x1e, 0x37, 0xf7, 0xa1, 0x6f, 0xa1, 0x1d, 0xac, 0x4c, 0x64, 0xa9, 0xa8,
	0x8e, 0x12, 0x5f, 0xf8, 0xdc, 0xb5, 0x68, 0xeb, 0xc7, 0xa2, 0x80, 0xeb, 0x68, 0xf2, 0xf4, 0x02,
	0xd4, 0xf1, 0xaf, 0x95, 0xc7, 0xbf, 0xf5, 0x73, 0x31, 0x26, 0xce, 0xbd, 0x60, 0x92, 0x1f, 0x7c,
	0x0e, 0x28, 0x58, 0xa3, 0x55, 0xe6, 0xf8, 0x78, 0x33, 0xe9, 0xee, 0x06, 0x84, 0xf5, 0x57, 0x05,
	0x76, 0xf3, 0x72, 0x48, 0x1c, 0x85, 0x41, 0x4c, 0xde, 0x1b, 0x6b, 0xe8, 0x0b, 0xa8, 0x25, 0x7c,
	0x36, 0xcb, 0x29, 0xd4, 0x56, 0x60, 0x62, 0x66, 0xbb, 0x99, 0xdb, 0xea, 0x14, 0x14, 0x5c, 0x78,
	0x31, 0x93, 0x14, 0x58, 0x67, 0xb0, 0x57, 0xb2, 0xca, 0xea, 0xbf, 0x84, 0x86, 0x4c, 0x94, 0x7d,
	0x4a, 0x57, 0xca, 0x2f, 0x9c, 0x47, 0xff, 0x56, 0xe0, 0xc3, 0x82, 0x25, 0x9a, 0x7a, 0x63, 0x82,
	0x8e, 0xc1, 0xc8, 0x1e, 0x10, 0x32, 0x15, 0x4c, 0xe9, 0x4d, 0x75, 0xf7, 0xd4, 0xd3, 0xf2, 0xa4,
	0xc7, 0x60, 0x64, 0x8d, 0x2f, 0x01, 0x4b, 0x6f, 0xe1, 0x3e, 0xa0, 0xce, 0xdb, 0x8a, 0xd4, 0x9e,
	0x29, 0x7d, 0xee, 0xaa, 0x7f, 0x22, 0xa5, 0x6b, 0x1e, 0x83, 0xce, 0xf7, 0x25, 0xa0, 0xc2, 0xce,
	0xbd, 0xc0, 0xb7, 0xce, 0x4f, 0xaf, 0x66, 0x1e, 0xbb, 0x49, 0x46, 0xf6, 0x38, 0x5c, 0x38, 0x0b,
	0xc2, 0xb0, 0xff, 0x2a, 0x66, 0x78, 0x3c, 0x97, 0x6b, 0x1c, 0x79, 0x4e, 0x34, 0x9f, 0x65, 0xff,
	0x75, 0x4e, 0xda, 0x1b, 0x19, 0x62, 0xf5, 0xd5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xd3,
	0xbe, 0xc0, 0x18, 0x0a, 0x00, 0x00,
}