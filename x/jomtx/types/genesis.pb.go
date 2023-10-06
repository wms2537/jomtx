// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jomtx/jomtx/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the jomtx module's genesis state.
type GenesisState struct {
	Params     Params      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TxnList    []Txn       `protobuf:"bytes,2,rep,name=txnList,proto3" json:"txnList"`
	TxnCount   uint64      `protobuf:"varint,3,opt,name=txnCount,proto3" json:"txnCount,omitempty"`
	SystemInfo *SystemInfo `protobuf:"bytes,4,opt,name=systemInfo,proto3" json:"systemInfo,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d402d7f1cd269b11, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTxnList() []Txn {
	if m != nil {
		return m.TxnList
	}
	return nil
}

func (m *GenesisState) GetTxnCount() uint64 {
	if m != nil {
		return m.TxnCount
	}
	return 0
}

func (m *GenesisState) GetSystemInfo() *SystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "jomtx.jomtx.GenesisState")
}

func init() { proto.RegisterFile("jomtx/jomtx/genesis.proto", fileDescriptor_d402d7f1cd269b11) }

var fileDescriptor_d402d7f1cd269b11 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xca, 0xcf, 0x2d,
	0xa9, 0xd0, 0x87, 0x90, 0xe9, 0xa9, 0x79, 0xa9, 0xc5, 0x99, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xdc, 0x60, 0x41, 0x3d, 0x30, 0x29, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x16, 0xd7,
	0x07, 0xb1, 0x20, 0x4a, 0xa4, 0x24, 0x90, 0x75, 0x17, 0x24, 0x16, 0x25, 0xe6, 0x42, 0x35, 0x4b,
	0x89, 0x22, 0xcb, 0x94, 0x54, 0xe4, 0x41, 0x85, 0x65, 0x91, 0x85, 0x8b, 0x2b, 0x8b, 0x4b, 0x52,
	0x73, 0xe3, 0x33, 0xf3, 0xd2, 0xa0, 0xe6, 0x29, 0x9d, 0x60, 0xe4, 0xe2, 0x71, 0x87, 0x38, 0x22,
	0xb8, 0x24, 0xb1, 0x24, 0x55, 0xc8, 0x90, 0x8b, 0x0d, 0x62, 0xac, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0xb0, 0x1e, 0x92, 0xa3, 0xf4, 0x02, 0xc0, 0x52, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33,
	0x04, 0x41, 0x15, 0x0a, 0x19, 0x70, 0xb1, 0x97, 0x54, 0xe4, 0xf9, 0x64, 0x16, 0x97, 0x48, 0x30,
	0x29, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xa0, 0xe8, 0x09, 0xa9, 0xc8, 0x83, 0x6a, 0x80, 0x29, 0x13,
	0x92, 0xe2, 0xe2, 0x28, 0xa9, 0xc8, 0x73, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x56, 0x60, 0xd4,
	0x60, 0x09, 0x82, 0xf3, 0x85, 0xcc, 0xb9, 0xb8, 0x20, 0xce, 0xf4, 0xcc, 0x4b, 0xcb, 0x97, 0x60,
	0x01, 0x3b, 0x42, 0x1c, 0xc5, 0xc0, 0x60, 0xb8, 0x74, 0x10, 0x92, 0x52, 0x27, 0x87, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xcf, 0x2d, 0x36, 0x32, 0x35, 0x36, 0x87, 0x06, 0x08, 0x3c, 0xbc,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x61, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0xf6, 0x6f, 0xa7, 0xa3, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SystemInfo != nil {
		{
			size, err := m.SystemInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TxnCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TxnCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TxnList) > 0 {
		for iNdEx := len(m.TxnList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxnList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TxnList) > 0 {
		for _, e := range m.TxnList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TxnCount != 0 {
		n += 1 + sovGenesis(uint64(m.TxnCount))
	}
	if m.SystemInfo != nil {
		l = m.SystemInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxnList = append(m.TxnList, Txn{})
			if err := m.TxnList[len(m.TxnList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnCount", wireType)
			}
			m.TxnCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxnCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemInfo == nil {
				m.SystemInfo = &SystemInfo{}
			}
			if err := m.SystemInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
