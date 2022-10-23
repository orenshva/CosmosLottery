// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoslottery/bet_chart.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type BetChart struct {
	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	Bet         uint64 `protobuf:"varint,2,opt,name=bet,proto3" json:"bet,omitempty"`
}

func (m *BetChart) Reset()         { *m = BetChart{} }
func (m *BetChart) String() string { return proto.CompactTextString(m) }
func (*BetChart) ProtoMessage()    {}
func (*BetChart) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d5754f0cb63b79a, []int{0}
}
func (m *BetChart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BetChart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetChart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BetChart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetChart.Merge(m, src)
}
func (m *BetChart) XXX_Size() int {
	return m.Size()
}
func (m *BetChart) XXX_DiscardUnknown() {
	xxx_messageInfo_BetChart.DiscardUnknown(m)
}

var xxx_messageInfo_BetChart proto.InternalMessageInfo

func (m *BetChart) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *BetChart) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func init() {
	proto.RegisterType((*BetChart)(nil), "orenshva.cosmoslottery.cosmoslottery.BetChart")
}

func init() { proto.RegisterFile("cosmoslottery/bet_chart.proto", fileDescriptor_5d5754f0cb63b79a) }

var fileDescriptor_5d5754f0cb63b79a = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xce, 0xc9, 0x2f, 0x29, 0x49, 0x2d, 0xaa, 0xd4, 0x4f, 0x4a, 0x2d, 0x89, 0x4f, 0xce,
	0x48, 0x2c, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xc9, 0x2f, 0x4a, 0xcd, 0x2b,
	0xce, 0x28, 0x4b, 0xd4, 0x43, 0x51, 0x87, 0xca, 0x53, 0xb2, 0xe3, 0xe2, 0x70, 0x4a, 0x2d, 0x71,
	0x06, 0xe9, 0x13, 0x52, 0xe0, 0xe2, 0x4e, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xf1, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x16, 0x12, 0x12, 0xe0, 0x62, 0x4e, 0x4a,
	0x2d, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x02, 0x31, 0x9d, 0x02, 0x4f, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x3c, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0xe6, 0x14, 0x7d, 0x67, 0xb0, 0xe5, 0x3e, 0x50, 0x27, 0x57, 0xe8, 0xa3, 0x7a,
	0xa1, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0c, 0xb4, 0x9e, 0xcd, 0xe0, 0x00, 0x00, 0x00,
}

func (m *BetChart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetChart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetChart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bet != 0 {
		i = encodeVarintBetChart(dAtA, i, uint64(m.Bet))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AccountName) > 0 {
		i -= len(m.AccountName)
		copy(dAtA[i:], m.AccountName)
		i = encodeVarintBetChart(dAtA, i, uint64(len(m.AccountName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBetChart(dAtA []byte, offset int, v uint64) int {
	offset -= sovBetChart(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BetChart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountName)
	if l > 0 {
		n += 1 + l + sovBetChart(uint64(l))
	}
	if m.Bet != 0 {
		n += 1 + sovBetChart(uint64(m.Bet))
	}
	return n
}

func sovBetChart(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBetChart(x uint64) (n int) {
	return sovBetChart(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BetChart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBetChart
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
			return fmt.Errorf("proto: BetChart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetChart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetChart
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
				return ErrInvalidLengthBetChart
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetChart
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			m.Bet = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetChart
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bet |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBetChart(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBetChart
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
func skipBetChart(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBetChart
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
					return 0, ErrIntOverflowBetChart
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
					return 0, ErrIntOverflowBetChart
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
				return 0, ErrInvalidLengthBetChart
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBetChart
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBetChart
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBetChart        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBetChart          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBetChart = fmt.Errorf("proto: unexpected end of group")
)
