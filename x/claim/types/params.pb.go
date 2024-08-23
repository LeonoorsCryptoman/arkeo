// DONTCOVER
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arkeo/claim/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params defines the parameters for the module.
type Params struct {
	AirdropStartTime   time.Time     `protobuf:"bytes,1,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	DurationUntilDecay time.Duration `protobuf:"bytes,2,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay,omitempty" yaml:"duration_until_decay"`
	DurationOfDecay    time.Duration `protobuf:"bytes,3,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay,omitempty" yaml:"duration_of_decay"`
	// denom of claimable asset
	ClaimDenom string `protobuf:"bytes,4,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
	// uarkeo to distribute to arkeo account for gas to make claiming easier
	InitialGasAmount *types.Coin `protobuf:"bytes,5,opt,name=initial_gas_amount,json=initialGasAmount,proto3" json:"initial_gas_amount,omitempty" yaml:"initial_gas_amount"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bdbd8eba92221b6, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *Params) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func (m *Params) GetInitialGasAmount() *types.Coin {
	if m != nil {
		return m.InitialGasAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "arkeo.claim.Params")
}

func init() { proto.RegisterFile("arkeo/claim/params.proto", fileDescriptor_2bdbd8eba92221b6) }

var fileDescriptor_2bdbd8eba92221b6 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x50, 0x2a, 0xe1, 0x0c, 0x14, 0xab, 0x83, 0x9b, 0x0a, 0xbb, 0xb2, 0x84, 0x54,
	0x09, 0x74, 0xa7, 0xc2, 0xc6, 0x46, 0x1a, 0x84, 0x98, 0x40, 0x01, 0x16, 0x16, 0xeb, 0xd9, 0xb9,
	0x98, 0x53, 0x7d, 0x7e, 0xd6, 0xdd, 0x19, 0xc8, 0x57, 0x60, 0xea, 0x84, 0xf8, 0x08, 0x7c, 0x94,
	0x8e, 0x1d, 0x99, 0x02, 0x4a, 0x36, 0xc6, 0x7e, 0x02, 0x74, 0xe7, 0x33, 0x88, 0x26, 0x12, 0x5b,
	0xee, 0xff, 0x7f, 0xef, 0xfd, 0x7f, 0xf1, 0x7b, 0x41, 0x04, 0xea, 0x8c, 0x23, 0x2b, 0x2a, 0x10,
	0x92, 0x35, 0xa0, 0x40, 0x6a, 0xda, 0x28, 0x34, 0x18, 0x0e, 0x9d, 0x43, 0x9d, 0x33, 0xda, 0x2f,
	0xb1, 0x44, 0xa7, 0x33, 0xfb, 0xab, 0x2b, 0x19, 0xc5, 0x25, 0x62, 0x59, 0x71, 0xe6, 0x5e, 0x79,
	0x3b, 0x67, 0xb3, 0x56, 0x81, 0x11, 0x58, 0x7b, 0x3f, 0xb9, 0xee, 0x1b, 0x21, 0xb9, 0x36, 0x20,
	0x9b, 0x7e, 0x40, 0x81, 0x5a, 0xa2, 0x66, 0x39, 0x68, 0xce, 0x3e, 0x9c, 0xe4, 0xdc, 0xc0, 0x09,
	0x2b, 0x50, 0xf8, 0x01, 0xe9, 0xb7, 0x9d, 0x60, 0xf7, 0x95, 0x83, 0x0a, 0x31, 0x08, 0x41, 0xa8,
	0x99, 0xc2, 0x26, 0xd3, 0x06, 0x94, 0xc9, 0xec, 0xac, 0x88, 0x1c, 0x91, 0xe3, 0xe1, 0xa3, 0x11,
	0xed, 0x82, 0x68, 0x1f, 0x44, 0xdf, 0xf4, 0x41, 0xe3, 0xfb, 0x17, 0xcb, 0x64, 0x70, 0xb5, 0x4c,
	0x0e, 0x16, 0x20, 0xab, 0x27, 0xe9, 0xe6, 0x8c, 0xf4, 0xfc, 0x47, 0x42, 0xa6, 0x7b, 0xde, 0x78,
	0x6d, 0x75, 0xdb, 0x1d, 0x7e, 0x21, 0xc1, 0x7e, 0xff, 0x7f, 0xb2, 0xb6, 0x36, 0xa2, 0xca, 0x66,
	0xbc, 0x80, 0x45, 0x74, 0xc3, 0x65, 0x1e, 0x6c, 0x64, 0x4e, 0x7c, 0xf1, 0xf8, 0x85, 0x8d, 0xfc,
	0xb5, 0x4c, 0xe2, 0x6d, 0xed, 0x0f, 0x51, 0x0a, 0xc3, 0x65, 0x63, 0x16, 0x57, 0xcb, 0xe4, 0xb0,
	0x83, 0xda, 0x56, 0x97, 0x7e, 0xb5, 0x58, 0x61, 0x6f, 0xbd, 0xb5, 0xce, 0xc4, 0x1a, 0xe1, 0x67,
	0x12, 0xdc, 0xfd, 0xd3, 0x81, 0x73, 0x4f, 0x75, 0xf3, 0x7f, 0x54, 0xa7, 0x9e, 0xea, 0x70, 0xa3,
	0xf7, 0x1f, 0xa4, 0xe8, 0x1a, 0x52, 0x5f, 0xd4, 0xf1, 0xdc, 0xe9, 0xf5, 0x97, 0xf3, 0x0e, 0x26,
	0x09, 0x86, 0xee, 0x42, 0xb2, 0x19, 0xaf, 0x51, 0x46, 0x3b, 0x47, 0xe4, 0xf8, 0xf6, 0x34, 0x70,
	0xd2, 0xc4, 0x2a, 0xe1, 0x3c, 0x08, 0x45, 0x2d, 0x8c, 0x80, 0x2a, 0x2b, 0x41, 0x67, 0x20, 0xb1,
	0xad, 0x4d, 0x74, 0xcb, 0xd3, 0x76, 0xfb, 0xa7, 0x76, 0xff, 0xd4, 0xef, 0x9f, 0x9e, 0xa2, 0xa8,
	0xc7, 0xf7, 0xfe, 0xae, 0x6c, 0xb3, 0x3d, 0x9d, 0xee, 0x79, 0xf1, 0x39, 0xe8, 0xa7, 0x4e, 0x1a,
	0x3f, 0xbb, 0x58, 0xc5, 0xe4, 0x72, 0x15, 0x93, 0x9f, 0xab, 0x98, 0x9c, 0xaf, 0xe3, 0xc1, 0xe5,
	0x3a, 0x1e, 0x7c, 0x5f, 0xc7, 0x83, 0x77, 0x0f, 0x4a, 0x61, 0xde, 0xb7, 0x39, 0x2d, 0x50, 0x32,
	0x77, 0xd3, 0x35, 0x37, 0x1f, 0x51, 0x9d, 0x75, 0x0f, 0xf6, 0xc9, 0x1f, 0xbf, 0x59, 0x34, 0x5c,
	0xe7, 0xbb, 0xee, 0xc3, 0x3d, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xda, 0x58, 0x9f, 0x2f, 0x18,
	0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitialGasAmount != nil {
		{
			size, err := m.InitialGasAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x22
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.AirdropStartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintParams(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.InitialGasAmount != nil {
		l = m.InitialGasAmount.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialGasAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InitialGasAmount == nil {
				m.InitialGasAmount = &types.Coin{}
			}
			if err := m.InitialGasAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
