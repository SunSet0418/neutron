// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/limit_order_tranche.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	github_com_neutron_org_neutron_v2_utils_math "github.com/neutron-org/neutron/v2/utils/math"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LimitOrderTrancheKey struct {
	TradePairId           *TradePairID `protobuf:"bytes,1,opt,name=trade_pair_id,json=tradePairId,proto3" json:"trade_pair_id,omitempty"`
	TickIndexTakerToMaker int64        `protobuf:"varint,2,opt,name=tick_index_taker_to_maker,json=tickIndexTakerToMaker,proto3" json:"tick_index_taker_to_maker,omitempty"`
	TrancheKey            string       `protobuf:"bytes,3,opt,name=tranche_key,json=trancheKey,proto3" json:"tranche_key,omitempty"`
}

func (m *LimitOrderTrancheKey) Reset()         { *m = LimitOrderTrancheKey{} }
func (m *LimitOrderTrancheKey) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheKey) ProtoMessage()    {}
func (*LimitOrderTrancheKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c2ded67c80756d1, []int{0}
}
func (m *LimitOrderTrancheKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheKey.Merge(m, src)
}
func (m *LimitOrderTrancheKey) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheKey.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheKey proto.InternalMessageInfo

func (m *LimitOrderTrancheKey) GetTradePairId() *TradePairID {
	if m != nil {
		return m.TradePairId
	}
	return nil
}

func (m *LimitOrderTrancheKey) GetTickIndexTakerToMaker() int64 {
	if m != nil {
		return m.TickIndexTakerToMaker
	}
	return 0
}

func (m *LimitOrderTrancheKey) GetTrancheKey() string {
	if m != nil {
		return m.TrancheKey
	}
	return ""
}

type LimitOrderTranche struct {
	Key                *LimitOrderTrancheKey                  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ReservesMakerDenom github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=reserves_maker_denom,json=reservesMakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserves_maker_denom" yaml:"reserves_maker_denom"`
	ReservesTakerDenom github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=reserves_taker_denom,json=reservesTakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserves_taker_denom" yaml:"reserves_taker_denom"`
	TotalMakerDenom    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=total_maker_denom,json=totalMakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_maker_denom" yaml:"total_maker_denom"`
	TotalTakerDenom    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_taker_denom,json=totalTakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_taker_denom" yaml:"total_taker_denom"`
	// expiration_time is represented as an RFC 3339 formatted date.
	// LimitOrders with expiration_time set are valid as long as blockTime <= expiration_time.
	// JIT orders also use expiration_time to handle deletion, but represent a special case.
	// All JIT orders have an expiration_time of 0001-01-01T00:00:00Z, and an exception is made to
	// still treat these orders as live. Order deletion still functions the
	// same, and the orders will be deleted at the end of the block.
	ExpirationTime    *time.Time                                           `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time,omitempty"`
	PriceTakerToMaker github_com_neutron_org_neutron_v2_utils_math.PrecDec `protobuf:"bytes,7,opt,name=price_taker_to_maker,json=priceTakerToMaker,proto3,customtype=github.com/neutron-org/neutron/v2/utils/math.PrecDec" json:"price_taker_to_maker" yaml:"price_taker_to_maker"`
}

func (m *LimitOrderTranche) Reset()         { *m = LimitOrderTranche{} }
func (m *LimitOrderTranche) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTranche) ProtoMessage()    {}
func (*LimitOrderTranche) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c2ded67c80756d1, []int{1}
}
func (m *LimitOrderTranche) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTranche) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTranche.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTranche) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTranche.Merge(m, src)
}
func (m *LimitOrderTranche) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTranche) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTranche.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTranche proto.InternalMessageInfo

func (m *LimitOrderTranche) GetKey() *LimitOrderTrancheKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LimitOrderTranche) GetExpirationTime() *time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitOrderTrancheKey)(nil), "neutron.dex.LimitOrderTrancheKey")
	proto.RegisterType((*LimitOrderTranche)(nil), "neutron.dex.LimitOrderTranche")
}

func init() {
	proto.RegisterFile("neutron/dex/limit_order_tranche.proto", fileDescriptor_8c2ded67c80756d1)
}

var fileDescriptor_8c2ded67c80756d1 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x8b, 0xd3, 0x4e,
	0x14, 0xc7, 0x3b, 0xbf, 0xfe, 0xdc, 0x75, 0xa7, 0xe8, 0xb2, 0xa1, 0x42, 0x76, 0x85, 0xa4, 0x06,
	0x94, 0x5e, 0x9a, 0xc0, 0xae, 0x07, 0x11, 0x4f, 0xa5, 0x20, 0x45, 0x8b, 0x4b, 0xc8, 0xc9, 0xcb,
	0x90, 0x26, 0x63, 0x3a, 0xb4, 0xc9, 0x84, 0xc9, 0xeb, 0xd2, 0xfe, 0x0d, 0x82, 0xec, 0xcd, 0x7f,
	0xc0, 0xbb, 0xff, 0x46, 0x8f, 0x7b, 0x14, 0x0f, 0x51, 0xda, 0x9b, 0xc7, 0xfd, 0x0b, 0x64, 0x92,
	0x74, 0x9b, 0xda, 0xa8, 0xc8, 0x9e, 0x66, 0xe6, 0xbd, 0xef, 0x4c, 0x3f, 0xef, 0xf5, 0x7d, 0x83,
	0x1f, 0x47, 0x74, 0x0a, 0x82, 0x47, 0x96, 0x4f, 0x67, 0xd6, 0x84, 0x85, 0x0c, 0x08, 0x17, 0x3e,
	0x15, 0x04, 0x84, 0x1b, 0x79, 0x23, 0x6a, 0xc6, 0x82, 0x03, 0x57, 0x1a, 0x85, 0xcc, 0xf4, 0xe9,
	0xec, 0xa4, 0x19, 0xf0, 0x80, 0x67, 0x71, 0x4b, 0xee, 0x72, 0xc9, 0x89, 0x1e, 0x70, 0x1e, 0x4c,
	0xa8, 0x95, 0x9d, 0x86, 0xd3, 0x77, 0x16, 0xb0, 0x90, 0x26, 0xe0, 0x86, 0x71, 0x21, 0x38, 0x2e,
	0xff, 0x54, 0xec, 0x32, 0x41, 0x98, 0xbf, 0xbe, 0x5b, 0x4e, 0x81, 0x70, 0x7d, 0x4a, 0xb6, 0x04,
	0xc6, 0x67, 0x84, 0x9b, 0xaf, 0x25, 0xdd, 0x1b, 0x09, 0xe7, 0xe4, 0x6c, 0xaf, 0xe8, 0x5c, 0x79,
	0x81, 0xef, 0x6d, 0xe9, 0x55, 0xd4, 0x42, 0xed, 0xc6, 0xa9, 0x6a, 0x96, 0x80, 0x4d, 0x47, 0x2a,
	0xce, 0x5d, 0x26, 0xfa, 0x3d, 0xbb, 0x01, 0x37, 0x07, 0x5f, 0x79, 0x86, 0x8f, 0x81, 0x79, 0x63,
	0xc2, 0x22, 0x9f, 0xce, 0x08, 0xb8, 0x63, 0x59, 0x38, 0x27, 0xa1, 0xdc, 0xa8, 0xff, 0xb5, 0x50,
	0xbb, 0x6e, 0x3f, 0x90, 0x82, 0xbe, 0xcc, 0x3b, 0x32, 0xea, 0xf0, 0x81, 0x5c, 0x14, 0x1d, 0x37,
	0x8a, 0x0e, 0x91, 0x31, 0x9d, 0xab, 0xf5, 0x16, 0x6a, 0x1f, 0xd8, 0x18, 0x6e, 0xc0, 0x8c, 0x0f,
	0xfb, 0xf8, 0x68, 0x87, 0x58, 0x39, 0xc3, 0x75, 0x29, 0xcf, 0x21, 0x1f, 0x6d, 0x41, 0x56, 0x95,
	0x67, 0x4b, 0xb5, 0xf2, 0x11, 0xe1, 0xa6, 0xa0, 0x09, 0x15, 0x17, 0x34, 0xc9, 0xd9, 0x88, 0x4f,
	0x23, 0x1e, 0x66, 0x84, 0x07, 0x5d, 0xba, 0x48, 0xf5, 0xda, 0xd7, 0x54, 0x7f, 0x12, 0x30, 0x18,
	0x4d, 0x87, 0xa6, 0xc7, 0x43, 0xcb, 0xe3, 0x49, 0xc8, 0x93, 0x62, 0xe9, 0x24, 0xfe, 0xd8, 0x82,
	0x79, 0x4c, 0x13, 0xb3, 0x1f, 0xc1, 0x8f, 0x54, 0xaf, 0x7c, 0xed, 0x3a, 0xd5, 0x1f, 0xce, 0xdd,
	0x70, 0xf2, 0xdc, 0xa8, 0xca, 0x1a, 0xb6, 0xb2, 0x0e, 0x67, 0x0d, 0xe8, 0xc9, 0xe0, 0x36, 0x19,
	0x94, 0xc8, 0xea, 0xb7, 0x26, 0x83, 0x3f, 0x92, 0x41, 0x25, 0x99, 0xb3, 0x21, 0x7b, 0x8f, 0xf0,
	0x11, 0x70, 0x70, 0x27, 0x5b, 0x0d, 0xfb, 0x3f, 0xc3, 0x22, 0xff, 0x8c, 0xb5, 0xfb, 0xd4, 0x75,
	0xaa, 0xab, 0x39, 0xd3, 0x4e, 0xca, 0xb0, 0x0f, 0xb3, 0xd8, 0xa0, 0x8a, 0xa6, 0xdc, 0xa4, 0x3b,
	0xb7, 0xa3, 0x81, 0xdf, 0xd3, 0xc0, 0x2e, 0x4d, 0xa9, 0x37, 0x03, 0x7c, 0x48, 0x67, 0x31, 0x13,
	0x2e, 0x30, 0x1e, 0x11, 0x69, 0x53, 0x75, 0x2f, 0x1b, 0xc8, 0x13, 0x33, 0xf7, 0xb0, 0xb9, 0xf6,
	0xb0, 0xe9, 0xac, 0x3d, 0xdc, 0xbd, 0xbb, 0x48, 0x75, 0x74, 0xf9, 0x4d, 0x47, 0xf6, 0xfd, 0xcd,
	0x65, 0x99, 0x56, 0x3e, 0x21, 0xdc, 0x8c, 0x05, 0xf3, 0xe8, 0xaf, 0x06, 0xda, 0xcf, 0xea, 0x4b,
	0x8a, 0xfa, 0x9e, 0x96, 0xea, 0x2b, 0xe6, 0xbe, 0xc3, 0x45, 0xb0, 0xde, 0x5b, 0x17, 0xa7, 0xd6,
	0x14, 0xd8, 0x24, 0xb1, 0x42, 0x17, 0x46, 0xe6, 0xb9, 0xa0, 0x5e, 0x8f, 0x7a, 0x72, 0x24, 0xaa,
	0xde, 0xde, 0x8c, 0x44, 0x55, 0xd6, 0xb0, 0x8f, 0xb2, 0x70, 0xd9, 0xb1, 0xdd, 0x97, 0x8b, 0xa5,
	0x86, 0xae, 0x96, 0x1a, 0xfa, 0xbe, 0xd4, 0xd0, 0xe5, 0x4a, 0xab, 0x5d, 0xad, 0xb4, 0xda, 0x97,
	0x95, 0x56, 0x7b, 0xdb, 0xf9, 0x3b, 0xd9, 0x2c, 0xff, 0x32, 0xc9, 0x3f, 0x61, 0xb8, 0x97, 0x75,
	0xe7, 0xec, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x8d, 0xc1, 0xa8, 0x3b, 0x05, 0x00, 0x00,
}

func (m *LimitOrderTrancheKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TrancheKey) > 0 {
		i -= len(m.TrancheKey)
		copy(dAtA[i:], m.TrancheKey)
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(len(m.TrancheKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TickIndexTakerToMaker != 0 {
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(m.TickIndexTakerToMaker))
		i--
		dAtA[i] = 0x10
	}
	if m.TradePairId != nil {
		{
			size, err := m.TradePairId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LimitOrderTranche) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTranche) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTranche) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PriceTakerToMaker.Size()
		i -= size
		if _, err := m.PriceTakerToMaker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.ExpirationTime != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.ExpirationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.TotalTakerDenom.Size()
		i -= size
		if _, err := m.TotalTakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TotalMakerDenom.Size()
		i -= size
		if _, err := m.TotalMakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ReservesTakerDenom.Size()
		i -= size
		if _, err := m.ReservesTakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ReservesMakerDenom.Size()
		i -= size
		if _, err := m.ReservesMakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTranche(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTranche(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradePairId != nil {
		l = m.TradePairId.Size()
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	if m.TickIndexTakerToMaker != 0 {
		n += 1 + sovLimitOrderTranche(uint64(m.TickIndexTakerToMaker))
	}
	l = len(m.TrancheKey)
	if l > 0 {
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	return n
}

func (m *LimitOrderTranche) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	l = m.ReservesMakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.ReservesTakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.TotalMakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.TotalTakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	if m.ExpirationTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime)
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	l = m.PriceTakerToMaker.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	return n
}

func sovLimitOrderTranche(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTranche(x uint64) (n int) {
	return sovLimitOrderTranche(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTranche
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
			return fmt.Errorf("proto: LimitOrderTrancheKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradePairId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TradePairId == nil {
				m.TradePairId = &TradePairID{}
			}
			if err := m.TradePairId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndexTakerToMaker", wireType)
			}
			m.TickIndexTakerToMaker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndexTakerToMaker |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrancheKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrancheKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTranche(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTranche
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
func (m *LimitOrderTranche) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTranche
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
			return fmt.Errorf("proto: LimitOrderTranche: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTranche: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &LimitOrderTrancheKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesMakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesMakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesTakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesTakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalMakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceTakerToMaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceTakerToMaker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTranche(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTranche
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
func skipLimitOrderTranche(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTranche
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
					return 0, ErrIntOverflowLimitOrderTranche
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
					return 0, ErrIntOverflowLimitOrderTranche
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
				return 0, ErrInvalidLengthLimitOrderTranche
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTranche
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTranche
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTranche        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTranche          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTranche = fmt.Errorf("proto: unexpected end of group")
)
