// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jomtx/jomtx/system_info.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type SystemInfo struct {
	FifoHeadIndex uint64 `protobuf:"varint,1,opt,name=fifoHeadIndex,proto3" json:"fifoHeadIndex,omitempty"`
	FifoTailIndex uint64 `protobuf:"varint,2,opt,name=fifoTailIndex,proto3" json:"fifoTailIndex,omitempty"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebcc7a87bf911bee, []int{0}
}
func (m *SystemInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SystemInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemInfo.Merge(m, src)
}
func (m *SystemInfo) XXX_Size() int {
	return m.Size()
}
func (m *SystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemInfo proto.InternalMessageInfo

func (m *SystemInfo) GetFifoHeadIndex() uint64 {
	if m != nil {
		return m.FifoHeadIndex
	}
	return 0
}

func (m *SystemInfo) GetFifoTailIndex() uint64 {
	if m != nil {
		return m.FifoTailIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*SystemInfo)(nil), "jomtx.jomtx.SystemInfo")
}

func init() { proto.RegisterFile("jomtx/jomtx/system_info.proto", fileDescriptor_ebcc7a87bf911bee) }

var fileDescriptor_ebcc7a87bf911bee = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xca, 0xcf, 0x2d,
	0xa9, 0xd0, 0x87, 0x90, 0xc5, 0x95, 0xc5, 0x25, 0xa9, 0xb9, 0xf1, 0x99, 0x79, 0x69, 0xf9, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0x60, 0x09, 0x3d, 0x30, 0xa9, 0x14, 0xc1, 0xc5, 0x15,
	0x0c, 0x56, 0xe1, 0x99, 0x97, 0x96, 0x2f, 0xa4, 0xc2, 0xc5, 0x9b, 0x96, 0x99, 0x96, 0xef, 0x91,
	0x9a, 0x98, 0xe2, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x2a,
	0x08, 0x53, 0x15, 0x92, 0x98, 0x99, 0x03, 0x51, 0xc5, 0x84, 0x50, 0x05, 0x17, 0x74, 0x72, 0x38,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xb5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf2, 0xdc, 0x62, 0x23, 0x53, 0x63, 0x73, 0xa8, 0x63, 0x61,
	0x8e, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x04, 0x43, 0x17, 0x3d, 0xd0, 0x00, 0x00, 0x00,
}

func (m *SystemInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SystemInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SystemInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FifoTailIndex != 0 {
		i = encodeVarintSystemInfo(dAtA, i, uint64(m.FifoTailIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.FifoHeadIndex != 0 {
		i = encodeVarintSystemInfo(dAtA, i, uint64(m.FifoHeadIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSystemInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovSystemInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SystemInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FifoHeadIndex != 0 {
		n += 1 + sovSystemInfo(uint64(m.FifoHeadIndex))
	}
	if m.FifoTailIndex != 0 {
		n += 1 + sovSystemInfo(uint64(m.FifoTailIndex))
	}
	return n
}

func sovSystemInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSystemInfo(x uint64) (n int) {
	return sovSystemInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SystemInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystemInfo
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
			return fmt.Errorf("proto: SystemInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SystemInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoHeadIndex", wireType)
			}
			m.FifoHeadIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FifoHeadIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoTailIndex", wireType)
			}
			m.FifoTailIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FifoTailIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSystemInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSystemInfo
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
func skipSystemInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSystemInfo
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
					return 0, ErrIntOverflowSystemInfo
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
					return 0, ErrIntOverflowSystemInfo
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
				return 0, ErrInvalidLengthSystemInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSystemInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSystemInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSystemInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSystemInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSystemInfo = fmt.Errorf("proto: unexpected end of group")
)