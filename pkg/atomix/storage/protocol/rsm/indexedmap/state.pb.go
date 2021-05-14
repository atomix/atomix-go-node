// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/storage/protocol/rsm/indexedmap/state.proto

package indexedmap

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

type IndexedMapState struct {
	meta.ObjectMeta `protobuf:"bytes,1,opt,name=meta,proto3,embedded=meta" json:"meta"`
	Value           int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *IndexedMapState) Reset()         { *m = IndexedMapState{} }
func (m *IndexedMapState) String() string { return proto.CompactTextString(m) }
func (*IndexedMapState) ProtoMessage()    {}
func (*IndexedMapState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac3bdf108813c09, []int{0}
}
func (m *IndexedMapState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedMapState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedMapState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedMapState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedMapState.Merge(m, src)
}
func (m *IndexedMapState) XXX_Size() int {
	return m.Size()
}
func (m *IndexedMapState) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedMapState.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedMapState proto.InternalMessageInfo

func (m *IndexedMapState) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*IndexedMapState)(nil), "atomix.storage.protocol.rsm.indexedmap.IndexedMapState")
}

func init() {
	proto.RegisterFile("atomix/storage/protocol/rsm/indexedmap/state.proto", fileDescriptor_7ac3bdf108813c09)
}

var fileDescriptor_7ac3bdf108813c09 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9,
	0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x2a, 0xce, 0xd5, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x48, 0x4d, 0xc9,
	0x4d, 0x2c, 0xd0, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x03, 0xcb, 0x0a, 0xa9, 0x41, 0xf4, 0xe8,
	0x41, 0xf5, 0xe8, 0xc1, 0xf4, 0xe8, 0x15, 0x15, 0xe7, 0xea, 0x21, 0xf4, 0x48, 0x29, 0x41, 0xcd,
	0x2e, 0x28, 0xca, 0xcc, 0xcd, 0x2c, 0xc9, 0x2c, 0x4b, 0xd5, 0xcf, 0x4d, 0x2d, 0x49, 0xd4, 0xcf,
	0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x81, 0xe8, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5,
	0x41, 0x2c, 0x88, 0xa8, 0x52, 0x06, 0x17, 0xbf, 0x27, 0xc4, 0x1c, 0xdf, 0xc4, 0x82, 0x60, 0x90,
	0xd5, 0x42, 0xf6, 0x5c, 0x2c, 0x20, 0xdd, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x8a, 0x7a,
	0x50, 0x37, 0xc0, 0xcd, 0xd6, 0x03, 0xc9, 0xea, 0xf9, 0x83, 0xcd, 0xf6, 0x4d, 0x2d, 0x49, 0x74,
	0xe2, 0x38, 0x71, 0x4f, 0x9e, 0xe1, 0xc2, 0x3d, 0x79, 0xc6, 0x20, 0xb0, 0x46, 0x21, 0x11, 0x2e,
	0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0xc7, 0x49,
	0xe2, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x4e, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb3, 0xcb, 0xce, 0x22, 0x01, 0x00, 0x00,
}

func (m *IndexedMapState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedMapState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedMapState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *IndexedMapState) Size() (n int) {
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
func (m *IndexedMapState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IndexedMapState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedMapState: illegal tag %d (wire type %d)", fieldNum, wire)
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
