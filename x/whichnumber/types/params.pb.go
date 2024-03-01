// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: whichnumber/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	CommitTimeout     uint64     `protobuf:"varint,1,opt,name=commit_timeout,json=commitTimeout,proto3" json:"commit_timeout,omitempty"`
	RevealTimeout     uint64     `protobuf:"varint,2,opt,name=reveal_timeout,json=revealTimeout,proto3" json:"reveal_timeout,omitempty"`
	MaxPlayersPerGame uint64     `protobuf:"varint,3,opt,name=max_players_per_game,json=maxPlayersPerGame,proto3" json:"max_players_per_game,omitempty"`
	MinDistanceToWin  uint64     `protobuf:"varint,4,opt,name=min_distance_to_win,json=minDistanceToWin,proto3" json:"min_distance_to_win,omitempty"`
	MinReward         types.Coin `protobuf:"bytes,5,opt,name=min_reward,json=minReward,proto3" json:"min_reward"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5a5d056c93d32ee, []int{0}
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

func (m *Params) GetCommitTimeout() uint64 {
	if m != nil {
		return m.CommitTimeout
	}
	return 0
}

func (m *Params) GetRevealTimeout() uint64 {
	if m != nil {
		return m.RevealTimeout
	}
	return 0
}

func (m *Params) GetMaxPlayersPerGame() uint64 {
	if m != nil {
		return m.MaxPlayersPerGame
	}
	return 0
}

func (m *Params) GetMinDistanceToWin() uint64 {
	if m != nil {
		return m.MinDistanceToWin
	}
	return 0
}

func (m *Params) GetMinReward() types.Coin {
	if m != nil {
		return m.MinReward
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "zale144.whichnumber.whichnumber.Params")
}

func init() { proto.RegisterFile("whichnumber/params.proto", fileDescriptor_b5a5d056c93d32ee) }

var fileDescriptor_b5a5d056c93d32ee = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0x77, 0xbf, 0x6f, 0x0d, 0xb8, 0xa2, 0xe8, 0x9a, 0x62, 0x4d, 0xb1, 0x09, 0x82, 0x90,
	0xc6, 0x1d, 0x12, 0x53, 0x59, 0x58, 0x44, 0xc1, 0xc6, 0x22, 0x84, 0x80, 0x60, 0x33, 0xcc, 0x6e,
	0x2e, 0x9b, 0x81, 0xdc, 0x99, 0x65, 0x66, 0xf2, 0xcf, 0xa7, 0xb0, 0xb4, 0xf4, 0x71, 0x52, 0xa6,
	0xb4, 0x12, 0x49, 0x5e, 0xc3, 0x42, 0x76, 0x27, 0x4a, 0xec, 0xee, 0x9c, 0xf3, 0xbb, 0x0c, 0xe7,
	0x5c, 0x3f, 0x9c, 0x8d, 0x78, 0x3a, 0x12, 0x13, 0x4c, 0x40, 0x91, 0x9c, 0x29, 0x86, 0x3a, 0xce,
	0x95, 0x34, 0x32, 0xa8, 0x3f, 0xb3, 0x31, 0xb4, 0x3a, 0x9d, 0x78, 0x87, 0xd8, 0x9d, 0x6b, 0xd5,
	0x4c, 0x66, 0xb2, 0x64, 0x49, 0x31, 0xd9, 0xb5, 0x5a, 0x94, 0x4a, 0x8d, 0x52, 0x93, 0x84, 0x69,
	0x20, 0xd3, 0x56, 0x02, 0x86, 0xb5, 0x48, 0x2a, 0xb9, 0xb0, 0xfe, 0xf9, 0x97, 0xeb, 0x57, 0x7a,
	0xe5, 0x3f, 0xc1, 0x85, 0x7f, 0x94, 0x4a, 0x44, 0x6e, 0xa8, 0xe1, 0x08, 0x72, 0x62, 0x42, 0xb7,
	0xe1, 0x36, 0xbd, 0xfe, 0xa1, 0x55, 0x07, 0x56, 0x2c, 0x30, 0x05, 0x53, 0x60, 0xe3, 0x5f, 0xec,
	0x9f, 0xc5, 0xac, 0xfa, 0x83, 0x11, 0xbf, 0x8a, 0x6c, 0x4e, 0xf3, 0x31, 0x5b, 0x80, 0xd2, 0x34,
	0x07, 0x45, 0x33, 0x86, 0x10, 0xfe, 0x2f, 0xe1, 0x13, 0x64, 0xf3, 0x9e, 0xb5, 0x7a, 0xa0, 0xee,
	0x19, 0x42, 0x70, 0xe9, 0x9f, 0x22, 0x17, 0x74, 0xc8, 0xb5, 0x61, 0x22, 0x05, 0x6a, 0x24, 0x9d,
	0x71, 0x11, 0x7a, 0x25, 0x7f, 0x8c, 0x5c, 0xdc, 0x6d, 0x9d, 0x81, 0x7c, 0xe4, 0x22, 0xb8, 0xf1,
	0xfd, 0x02, 0x57, 0x30, 0x63, 0x6a, 0x18, 0xee, 0x35, 0xdc, 0xe6, 0x41, 0xfb, 0x2c, 0xb6, 0x69,
	0xe3, 0x22, 0x6d, 0xbc, 0x4d, 0x1b, 0xdf, 0x4a, 0x2e, 0xba, 0xde, 0xf2, 0xa3, 0xee, 0xf4, 0xf7,
	0x91, 0x8b, 0x7e, 0xb9, 0x71, 0xed, 0xbd, 0xbe, 0xd5, 0x9d, 0xee, 0xc3, 0x72, 0x1d, 0xb9, 0xab,
	0x75, 0xe4, 0x7e, 0xae, 0x23, 0xf7, 0x65, 0x13, 0x39, 0xab, 0x4d, 0xe4, 0xbc, 0x6f, 0x22, 0xe7,
	0xa9, 0x9d, 0x71, 0x33, 0x9a, 0x24, 0x71, 0x2a, 0x91, 0x6c, 0xab, 0x27, 0xbb, 0xc7, 0x99, 0xff,
	0x79, 0x99, 0x45, 0x0e, 0x3a, 0xa9, 0x94, 0x9d, 0x5e, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0xe3, 0x4f, 0x56, 0xc6, 0x01, 0x00, 0x00,
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
	{
		size, err := m.MinReward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.MinDistanceToWin != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinDistanceToWin))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxPlayersPerGame != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPlayersPerGame))
		i--
		dAtA[i] = 0x18
	}
	if m.RevealTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RevealTimeout))
		i--
		dAtA[i] = 0x10
	}
	if m.CommitTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommitTimeout))
		i--
		dAtA[i] = 0x8
	}
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
	if m.CommitTimeout != 0 {
		n += 1 + sovParams(uint64(m.CommitTimeout))
	}
	if m.RevealTimeout != 0 {
		n += 1 + sovParams(uint64(m.RevealTimeout))
	}
	if m.MaxPlayersPerGame != 0 {
		n += 1 + sovParams(uint64(m.MaxPlayersPerGame))
	}
	if m.MinDistanceToWin != 0 {
		n += 1 + sovParams(uint64(m.MinDistanceToWin))
	}
	l = m.MinReward.Size()
	n += 1 + l + sovParams(uint64(l))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitTimeout", wireType)
			}
			m.CommitTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealTimeout", wireType)
			}
			m.RevealTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevealTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPlayersPerGame", wireType)
			}
			m.MaxPlayersPerGame = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPlayersPerGame |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDistanceToWin", wireType)
			}
			m.MinDistanceToWin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDistanceToWin |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinReward", wireType)
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
			if err := m.MinReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
