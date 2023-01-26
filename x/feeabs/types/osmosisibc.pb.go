// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/osmosisibc.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FeeabsIbcPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*FeeabsIbcPacketData_NoData
	//	*FeeabsIbcPacketData_IbcSwapAmountInRoute
	//	*FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData
	Packet isFeeabsIbcPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *FeeabsIbcPacketData) Reset()         { *m = FeeabsIbcPacketData{} }
func (m *FeeabsIbcPacketData) String() string { return proto.CompactTextString(m) }
func (*FeeabsIbcPacketData) ProtoMessage()    {}
func (*FeeabsIbcPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7e956854d14a746, []int{0}
}
func (m *FeeabsIbcPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeabsIbcPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeabsIbcPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeabsIbcPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeabsIbcPacketData.Merge(m, src)
}
func (m *FeeabsIbcPacketData) XXX_Size() int {
	return m.Size()
}
func (m *FeeabsIbcPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeabsIbcPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_FeeabsIbcPacketData proto.InternalMessageInfo

type isFeeabsIbcPacketData_Packet interface {
	isFeeabsIbcPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FeeabsIbcPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type FeeabsIbcPacketData_IbcSwapAmountInRoute struct {
	IbcSwapAmountInRoute *SwapAmountInRoute `protobuf:"bytes,2,opt,name=ibcSwapAmountInRoute,proto3,oneof" json:"ibcSwapAmountInRoute,omitempty"`
}
type FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData struct {
	IbcOsmosisQuerySpotPriceRequestPacketData *OsmosisQuerySpotPriceRequestPacketData `protobuf:"bytes,3,opt,name=ibcOsmosisQuerySpotPriceRequestPacketData,proto3,oneof" json:"ibcOsmosisQuerySpotPriceRequestPacketData,omitempty"`
}

func (*FeeabsIbcPacketData_NoData) isFeeabsIbcPacketData_Packet()               {}
func (*FeeabsIbcPacketData_IbcSwapAmountInRoute) isFeeabsIbcPacketData_Packet() {}
func (*FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData) isFeeabsIbcPacketData_Packet() {
}

func (m *FeeabsIbcPacketData) GetPacket() isFeeabsIbcPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *FeeabsIbcPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*FeeabsIbcPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *FeeabsIbcPacketData) GetIbcSwapAmountInRoute() *SwapAmountInRoute {
	if x, ok := m.GetPacket().(*FeeabsIbcPacketData_IbcSwapAmountInRoute); ok {
		return x.IbcSwapAmountInRoute
	}
	return nil
}

func (m *FeeabsIbcPacketData) GetIbcOsmosisQuerySpotPriceRequestPacketData() *OsmosisQuerySpotPriceRequestPacketData {
	if x, ok := m.GetPacket().(*FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData); ok {
		return x.IbcOsmosisQuerySpotPriceRequestPacketData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeeabsIbcPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeeabsIbcPacketData_NoData)(nil),
		(*FeeabsIbcPacketData_IbcSwapAmountInRoute)(nil),
		(*FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7e956854d14a746, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// TODO : add more details for this message
type SwapAmountInRoute struct {
	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty"`
}

func (m *SwapAmountInRoute) Reset()         { *m = SwapAmountInRoute{} }
func (m *SwapAmountInRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountInRoute) ProtoMessage()    {}
func (*SwapAmountInRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7e956854d14a746, []int{2}
}
func (m *SwapAmountInRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountInRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountInRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountInRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountInRoute.Merge(m, src)
}
func (m *SwapAmountInRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountInRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountInRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountInRoute proto.InternalMessageInfo

func (m *SwapAmountInRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *SwapAmountInRoute) GetTokenOutDenom() string {
	if m != nil {
		return m.TokenOutDenom
	}
	return ""
}

// TODO : add more details for this message
type OsmosisQueryRequestPacketData struct {
	PoolId  uint64              `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenIn string              `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	Routes  []SwapAmountInRoute `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes"`
}

func (m *OsmosisQueryRequestPacketData) Reset()         { *m = OsmosisQueryRequestPacketData{} }
func (m *OsmosisQueryRequestPacketData) String() string { return proto.CompactTextString(m) }
func (*OsmosisQueryRequestPacketData) ProtoMessage()    {}
func (*OsmosisQueryRequestPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7e956854d14a746, []int{3}
}
func (m *OsmosisQueryRequestPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmosisQueryRequestPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmosisQueryRequestPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmosisQueryRequestPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmosisQueryRequestPacketData.Merge(m, src)
}
func (m *OsmosisQueryRequestPacketData) XXX_Size() int {
	return m.Size()
}
func (m *OsmosisQueryRequestPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmosisQueryRequestPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_OsmosisQueryRequestPacketData proto.InternalMessageInfo

func (m *OsmosisQueryRequestPacketData) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *OsmosisQueryRequestPacketData) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *OsmosisQueryRequestPacketData) GetRoutes() []SwapAmountInRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type OsmosisQuerySpotPriceRequestPacketData struct {
	PoolId          uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	BaseAssetDenom  string `protobuf:"bytes,2,opt,name=base_asset_denom,json=baseAssetDenom,proto3" json:"base_asset_denom,omitempty"`
	QuoteAssetDenom string `protobuf:"bytes,3,opt,name=quote_asset_denom,json=quoteAssetDenom,proto3" json:"quote_asset_denom,omitempty"`
}

func (m *OsmosisQuerySpotPriceRequestPacketData) Reset() {
	*m = OsmosisQuerySpotPriceRequestPacketData{}
}
func (m *OsmosisQuerySpotPriceRequestPacketData) String() string { return proto.CompactTextString(m) }
func (*OsmosisQuerySpotPriceRequestPacketData) ProtoMessage()    {}
func (*OsmosisQuerySpotPriceRequestPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7e956854d14a746, []int{4}
}
func (m *OsmosisQuerySpotPriceRequestPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmosisQuerySpotPriceRequestPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmosisQuerySpotPriceRequestPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmosisQuerySpotPriceRequestPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmosisQuerySpotPriceRequestPacketData.Merge(m, src)
}
func (m *OsmosisQuerySpotPriceRequestPacketData) XXX_Size() int {
	return m.Size()
}
func (m *OsmosisQuerySpotPriceRequestPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmosisQuerySpotPriceRequestPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_OsmosisQuerySpotPriceRequestPacketData proto.InternalMessageInfo

func (m *OsmosisQuerySpotPriceRequestPacketData) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *OsmosisQuerySpotPriceRequestPacketData) GetBaseAssetDenom() string {
	if m != nil {
		return m.BaseAssetDenom
	}
	return ""
}

func (m *OsmosisQuerySpotPriceRequestPacketData) GetQuoteAssetDenom() string {
	if m != nil {
		return m.QuoteAssetDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*FeeabsIbcPacketData)(nil), "feeabstraction.absfee.v1beta1.FeeabsIbcPacketData")
	proto.RegisterType((*NoData)(nil), "feeabstraction.absfee.v1beta1.NoData")
	proto.RegisterType((*SwapAmountInRoute)(nil), "feeabstraction.absfee.v1beta1.SwapAmountInRoute")
	proto.RegisterType((*OsmosisQueryRequestPacketData)(nil), "feeabstraction.absfee.v1beta1.OsmosisQueryRequestPacketData")
	proto.RegisterType((*OsmosisQuerySpotPriceRequestPacketData)(nil), "feeabstraction.absfee.v1beta1.OsmosisQuerySpotPriceRequestPacketData")
}

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/osmosisibc.proto", fileDescriptor_d7e956854d14a746)
}

var fileDescriptor_d7e956854d14a746 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x9b, 0xca, 0x0d, 0x8b, 0xa0, 0xd4, 0x54, 0x22, 0x54, 0xd4, 0x54, 0x96, 0xa8, 0x0a,
	0x12, 0x5e, 0x02, 0x27, 0x4e, 0xa8, 0x51, 0x41, 0xc9, 0xa5, 0x2d, 0x0e, 0x27, 0x2e, 0xd1, 0xae,
	0xb3, 0x31, 0xab, 0xc6, 0x3b, 0xae, 0x77, 0x5c, 0xe8, 0x5f, 0x70, 0xe3, 0x03, 0x10, 0xff, 0xd2,
	0x63, 0x8f, 0x9c, 0x10, 0x4a, 0x7e, 0x80, 0x4f, 0x40, 0xbb, 0xb6, 0x54, 0x42, 0x28, 0x8d, 0x7a,
	0xdb, 0x79, 0xfb, 0xe6, 0xbd, 0x99, 0xb7, 0x5a, 0x12, 0x8d, 0x84, 0x60, 0x5c, 0x63, 0xc1, 0x12,
	0x94, 0xa0, 0x28, 0xe3, 0x7a, 0x24, 0x04, 0x3d, 0x69, 0x73, 0x81, 0xac, 0x4d, 0x41, 0x67, 0xa0,
	0xa5, 0x96, 0x3c, 0x89, 0xf2, 0x02, 0x10, 0xfc, 0xcd, 0x59, 0x7e, 0x54, 0xf1, 0xa3, 0x9a, 0xbf,
	0xb1, 0x9e, 0x42, 0x0a, 0x96, 0x49, 0xcd, 0xa9, 0x6a, 0xda, 0x78, 0x90, 0x02, 0xa4, 0x63, 0x41,
	0x59, 0x2e, 0x29, 0x53, 0x0a, 0x90, 0x99, 0x5e, 0x5d, 0xdd, 0x86, 0xbf, 0x96, 0xc8, 0xdd, 0x37,
	0x56, 0xb5, 0xc7, 0x93, 0x43, 0x96, 0x1c, 0x09, 0xdc, 0x63, 0xc8, 0xfc, 0x57, 0xc4, 0x53, 0x60,
	0x4e, 0x2d, 0x77, 0xcb, 0xdd, 0xb9, 0xf9, 0xfc, 0x51, 0xf4, 0x5f, 0xef, 0x68, 0xdf, 0x92, 0xbb,
	0x4e, 0x5c, 0xb7, 0xf9, 0x23, 0xb2, 0x2e, 0x79, 0xd2, 0xff, 0xc8, 0xf2, 0xdd, 0x0c, 0x4a, 0x85,
	0x3d, 0x15, 0x43, 0x89, 0xa2, 0xb5, 0x64, 0xe5, 0x9e, 0x5d, 0x21, 0x37, 0xd7, 0xd7, 0x75, 0xe2,
	0x7f, 0xea, 0xf9, 0xdf, 0x5c, 0xf2, 0x58, 0xf2, 0xe4, 0xa0, 0xca, 0xea, 0x6d, 0x29, 0x8a, 0xd3,
	0x7e, 0x0e, 0x78, 0x58, 0xc8, 0x44, 0xc4, 0xe2, 0xb8, 0x14, 0x1a, 0x2f, 0xd6, 0x6a, 0x35, 0xac,
	0xfb, 0xeb, 0x2b, 0xdc, 0x17, 0x13, 0xeb, 0x3a, 0xf1, 0xe2, 0xce, 0x9d, 0x26, 0xf1, 0x72, 0x5b,
	0x85, 0x4d, 0xe2, 0x55, 0x69, 0x85, 0xef, 0xc8, 0xda, 0xfc, 0x42, 0xf7, 0xc8, 0x4a, 0x0e, 0x30,
	0x1e, 0xc8, 0xa1, 0x8d, 0x7e, 0x39, 0xf6, 0x4c, 0xd9, 0x1b, 0xfa, 0xdb, 0x64, 0x15, 0xe1, 0x48,
	0xa8, 0x01, 0x94, 0x38, 0x18, 0x0a, 0x05, 0x99, 0x0d, 0xf3, 0x46, 0x7c, 0xcb, 0xc2, 0x07, 0x25,
	0xee, 0x19, 0x30, 0xfc, 0xea, 0x92, 0xcd, 0x3f, 0x87, 0x9a, 0x9b, 0xe5, 0x72, 0x8b, 0xfb, 0xa4,
	0x59, 0x59, 0x48, 0x55, 0x6b, 0xaf, 0xd8, 0xba, 0xa7, 0xfc, 0x7d, 0xe2, 0x15, 0x66, 0x3e, 0xdd,
	0x6a, 0x6c, 0x35, 0xae, 0xf3, 0x82, 0x9d, 0xe5, 0xb3, 0x1f, 0x0f, 0x9d, 0xb8, 0x56, 0x09, 0xbf,
	0xb8, 0x64, 0x7b, 0xb1, 0xe8, 0x2e, 0x1f, 0x77, 0x87, 0xdc, 0xe1, 0x4c, 0x8b, 0x01, 0xd3, 0x5a,
	0xcc, 0x46, 0x72, 0xdb, 0xe0, 0xbb, 0x06, 0xb6, 0x99, 0xf8, 0x4f, 0xc8, 0xda, 0x71, 0x09, 0x38,
	0x4b, 0x6d, 0x58, 0xea, 0xaa, 0xbd, 0xb8, 0xe0, 0x76, 0xfa, 0x67, 0x93, 0xc0, 0x3d, 0x9f, 0x04,
	0xee, 0xcf, 0x49, 0xe0, 0x7e, 0x9e, 0x06, 0xce, 0xf9, 0x34, 0x70, 0xbe, 0x4f, 0x03, 0xe7, 0xfd,
	0xcb, 0x54, 0xe2, 0x87, 0x92, 0x47, 0x09, 0x64, 0x54, 0x81, 0xd9, 0x9a, 0x8d, 0x9f, 0x8e, 0x19,
	0xd7, 0xf4, 0xaf, 0x8f, 0x7c, 0xd2, 0xa6, 0x9f, 0x6a, 0x8c, 0xe2, 0x69, 0x2e, 0x34, 0xf7, 0xec,
	0x77, 0x7b, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x53, 0x42, 0x1b, 0xf3, 0x03, 0x00, 0x00,
}

func (m *FeeabsIbcPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeabsIbcPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeabsIbcPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *FeeabsIbcPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeabsIbcPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOsmosisibc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *FeeabsIbcPacketData_IbcSwapAmountInRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeabsIbcPacketData_IbcSwapAmountInRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IbcSwapAmountInRoute != nil {
		{
			size, err := m.IbcSwapAmountInRoute.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOsmosisibc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IbcOsmosisQuerySpotPriceRequestPacketData != nil {
		{
			size, err := m.IbcOsmosisQuerySpotPriceRequestPacketData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOsmosisibc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SwapAmountInRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountInRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountInRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOutDenom) > 0 {
		i -= len(m.TokenOutDenom)
		copy(dAtA[i:], m.TokenOutDenom)
		i = encodeVarintOsmosisibc(dAtA, i, uint64(len(m.TokenOutDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosisibc(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OsmosisQueryRequestPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmosisQueryRequestPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmosisQueryRequestPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Routes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOsmosisibc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintOsmosisibc(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosisibc(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OsmosisQuerySpotPriceRequestPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmosisQuerySpotPriceRequestPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmosisQuerySpotPriceRequestPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QuoteAssetDenom) > 0 {
		i -= len(m.QuoteAssetDenom)
		copy(dAtA[i:], m.QuoteAssetDenom)
		i = encodeVarintOsmosisibc(dAtA, i, uint64(len(m.QuoteAssetDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAssetDenom) > 0 {
		i -= len(m.BaseAssetDenom)
		copy(dAtA[i:], m.BaseAssetDenom)
		i = encodeVarintOsmosisibc(dAtA, i, uint64(len(m.BaseAssetDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosisibc(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOsmosisibc(dAtA []byte, offset int, v uint64) int {
	offset -= sovOsmosisibc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeabsIbcPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *FeeabsIbcPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	return n
}
func (m *FeeabsIbcPacketData_IbcSwapAmountInRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcSwapAmountInRoute != nil {
		l = m.IbcSwapAmountInRoute.Size()
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	return n
}
func (m *FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcOsmosisQuerySpotPriceRequestPacketData != nil {
		l = m.IbcOsmosisQuerySpotPriceRequestPacketData.Size()
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SwapAmountInRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosisibc(uint64(m.PoolId))
	}
	l = len(m.TokenOutDenom)
	if l > 0 {
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	return n
}

func (m *OsmosisQueryRequestPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosisibc(uint64(m.PoolId))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovOsmosisibc(uint64(l))
		}
	}
	return n
}

func (m *OsmosisQuerySpotPriceRequestPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosisibc(uint64(m.PoolId))
	}
	l = len(m.BaseAssetDenom)
	if l > 0 {
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	l = len(m.QuoteAssetDenom)
	if l > 0 {
		n += 1 + l + sovOsmosisibc(uint64(l))
	}
	return n
}

func sovOsmosisibc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOsmosisibc(x uint64) (n int) {
	return sovOsmosisibc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeabsIbcPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosisibc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeeabsIbcPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeabsIbcPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &FeeabsIbcPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcSwapAmountInRoute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SwapAmountInRoute{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &FeeabsIbcPacketData_IbcSwapAmountInRoute{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcOsmosisQuerySpotPriceRequestPacketData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OsmosisQuerySpotPriceRequestPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &FeeabsIbcPacketData_IbcOsmosisQuerySpotPriceRequestPacketData{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosisibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosisibc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosisibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SwapAmountInRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosisibc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SwapAmountInRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountInRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOutDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosisibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OsmosisQueryRequestPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosisibc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OsmosisQueryRequestPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmosisQueryRequestPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, SwapAmountInRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosisibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OsmosisQuerySpotPriceRequestPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosisibc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OsmosisQuerySpotPriceRequestPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmosisQuerySpotPriceRequestPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosisibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosisibc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOsmosisibc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOsmosisibc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOsmosisibc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOsmosisibc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOsmosisibc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOsmosisibc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOsmosisibc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOsmosisibc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOsmosisibc = fmt.Errorf("proto: unexpected end of group")
)