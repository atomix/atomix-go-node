// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/storage/protocol/rsm/list/state.proto

package list

import (
	fmt "fmt"
	meta "github.com/atomix/atomix-api/go/atomix/primitive/meta"
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

type ListState struct {
	meta.ObjectMeta `protobuf:"bytes,1,opt,name=meta,proto3,embedded=meta" json:"meta"`
	Value           int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ListState) Reset()         { *m = ListState{} }
func (m *ListState) String() string { return proto.CompactTextString(m) }
func (*ListState) ProtoMessage()    {}
func (*ListState) Descriptor() ([]byte, []int) {
	return fileDescriptor_acad1b7c78cffd2d, []int{0}
}
func (m *ListState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListState.Merge(m, src)
}
func (m *ListState) XXX_Size() int {
	return m.Size()
}
func (m *ListState) XXX_DiscardUnknown() {
	xxx_messageInfo_ListState.DiscardUnknown(m)
}

var xxx_messageInfo_ListState proto.InternalMessageInfo

func (m *ListState) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*ListState)(nil), "atomix.storage.protocol.rsm.list.ListState")
}

func init() {
	proto.RegisterFile("atomix/storage/protocol/rsm/list/state.proto", fileDescriptor_acad1b7c78cffd2d)
}

var fileDescriptor_acad1b7c78cffd2d = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x1b, 0xff, 0xd0, 0xb8, 0x2b, 0xb3, 0x28, 0xb3, 0xc8, 0xd4, 0x59, 0xcd, 0x42, 0xee,
	0x05, 0x7d, 0x00, 0x61, 0xd6, 0x8a, 0x30, 0x3e, 0x41, 0x3a, 0x84, 0x12, 0x69, 0xb8, 0x43, 0x72,
	0x1d, 0x7c, 0x0c, 0x1f, 0x6b, 0x96, 0x5d, 0xba, 0x2a, 0xd2, 0xbe, 0x88, 0x24, 0xa9, 0xee, 0x4e,
	0x4e, 0xbe, 0x73, 0xce, 0x95, 0xf7, 0x9a, 0xc9, 0xd9, 0x4f, 0x0c, 0x4c, 0x5e, 0xb7, 0x06, 0x0f,
	0x9e, 0x98, 0xf6, 0xd4, 0xa1, 0x0f, 0x0e, 0x3b, 0x1b, 0x18, 0x03, 0x6b, 0x36, 0x90, 0xfc, 0xb2,
	0xce, 0x34, 0xcc, 0x34, 0xfc, 0xd1, 0xe0, 0x83, 0x83, 0x48, 0x2f, 0xd7, 0x73, 0xdf, 0xc1, 0x5b,
	0x67, 0xd9, 0x1e, 0x0d, 0x3a, 0xc3, 0x1a, 0xa9, 0x79, 0x37, 0x7b, 0xce, 0xfc, 0x72, 0xd1, 0x52,
	0x4b, 0x49, 0x62, 0x54, 0xd9, 0x5d, 0x37, 0xf2, 0xe6, 0xd9, 0x06, 0x7e, 0x8b, 0x73, 0xe5, 0x93,
	0xbc, 0x88, 0xb9, 0x4a, 0xd4, 0x62, 0x73, 0xfb, 0x70, 0x07, 0xf3, 0xee, 0x7f, 0x2b, 0xc4, 0x5f,
	0x78, 0x4d, 0xad, 0x2f, 0x86, 0xf5, 0xf6, 0xfa, 0x34, 0xac, 0x8a, 0x7e, 0x58, 0x89, 0x5d, 0x0a,
	0x96, 0x0b, 0x79, 0x79, 0xd4, 0xdd, 0x87, 0xa9, 0xce, 0x6a, 0xb1, 0x39, 0xdf, 0xe5, 0xc7, 0xb6,
	0x3a, 0x8d, 0x4a, 0xf4, 0xa3, 0x12, 0x3f, 0xa3, 0x12, 0x5f, 0x93, 0x2a, 0xfa, 0x49, 0x15, 0xdf,
	0x93, 0x2a, 0x9a, 0xab, 0x74, 0xc4, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x02, 0xcc,
	0xe6, 0x10, 0x01, 0x00, 0x00,
}

func (m *ListState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovState(uint64(l))
	if m.Value != 0 {
		n += 1 + sovState(uint64(m.Value))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ListState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)