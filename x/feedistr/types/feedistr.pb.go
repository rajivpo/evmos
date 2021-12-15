// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/feedistr/v1/feedistr.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Distribution defines the fee distribution percentages for block proposer and contract rewards
type Distribution struct {
	// block proposer reward percentage
	ProposerReward github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=proposer_reward,json=proposerReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"proposer_reward"`
	// contract reward percentage
	ContractRewards github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=contract_rewards,json=contractRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"contract_rewards"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_e12a3957ba6fd59a, []int{0}
}
func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return m.Size()
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

// ContractWithdrawAddress
type ContractWithdrawAddress struct {
	// contract hex address registered on the fee distribution module
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// hex address from the registered owner/withdraw address
	WithdrawAddress string `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (m *ContractWithdrawAddress) Reset()         { *m = ContractWithdrawAddress{} }
func (m *ContractWithdrawAddress) String() string { return proto.CompactTextString(m) }
func (*ContractWithdrawAddress) ProtoMessage()    {}
func (*ContractWithdrawAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_e12a3957ba6fd59a, []int{1}
}
func (m *ContractWithdrawAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractWithdrawAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractWithdrawAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractWithdrawAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractWithdrawAddress.Merge(m, src)
}
func (m *ContractWithdrawAddress) XXX_Size() int {
	return m.Size()
}
func (m *ContractWithdrawAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractWithdrawAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ContractWithdrawAddress proto.InternalMessageInfo

func (m *ContractWithdrawAddress) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractWithdrawAddress) GetWithdrawAddress() string {
	if m != nil {
		return m.WithdrawAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Distribution)(nil), "evmos.feedistr.v1.Distribution")
	proto.RegisterType((*ContractWithdrawAddress)(nil), "evmos.feedistr.v1.ContractWithdrawAddress")
}

func init() { proto.RegisterFile("evmos/feedistr/v1/feedistr.proto", fileDescriptor_e12a3957ba6fd59a) }

var fileDescriptor_e12a3957ba6fd59a = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2d, 0xcb, 0xcd,
	0x2f, 0xd6, 0x4f, 0x4b, 0x4d, 0x4d, 0xc9, 0x2c, 0x2e, 0x29, 0xd2, 0x2f, 0x33, 0x84, 0xb3, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xc1, 0x2a, 0xf4, 0xe0, 0xa2, 0x65, 0x86, 0x52, 0x22,
	0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0xe9, 0x10, 0x23, 0x17, 0x8f,
	0x0b, 0x48, 0x49, 0x66, 0x52, 0x69, 0x49, 0x66, 0x7e, 0x9e, 0x50, 0x38, 0x17, 0x7f, 0x41, 0x51,
	0x7e, 0x41, 0x7e, 0x71, 0x6a, 0x51, 0x7c, 0x51, 0x6a, 0x79, 0x62, 0x51, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xa7, 0x93, 0xde, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xab, 0xa5, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0x83, 0x1c, 0x02, 0xa1, 0x74,
	0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x5c, 0x52, 0x93, 0x83, 0xf8, 0x60,
	0xc6, 0x04, 0x81, 0x4d, 0x11, 0x8a, 0xe4, 0x12, 0x48, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e,
	0x81, 0x1a, 0x5c, 0x2c, 0xc1, 0x44, 0x96, 0xc9, 0xfc, 0x30, 0x73, 0x20, 0x26, 0x17, 0x2b, 0xe5,
	0x73, 0x89, 0x3b, 0x43, 0x85, 0xc2, 0x33, 0x4b, 0x32, 0x52, 0x8a, 0x12, 0xcb, 0x1d, 0x53, 0x52,
	0x8a, 0x52, 0x8b, 0x8b, 0x85, 0x34, 0x91, 0x6c, 0x4d, 0x84, 0x88, 0x41, 0xfc, 0x83, 0x30, 0x05,
	0x49, 0x69, 0x39, 0x54, 0x37, 0x5c, 0x29, 0x13, 0x44, 0x69, 0x39, 0xaa, 0xa9, 0x4e, 0xce, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x89, 0xe4, 0x87, 0x92, 0x8c, 0xc4,
	0xa2, 0xe2, 0xcc, 0x62, 0x7d, 0x48, 0x6c, 0x55, 0x20, 0xe2, 0x0b, 0xec, 0x95, 0x24, 0x36, 0x70,
	0x0c, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x6a, 0xbe, 0x6d, 0xce, 0x01, 0x00, 0x00,
}

func (m *Distribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ContractRewards.Size()
		i -= size
		if _, err := m.ContractRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFeedistr(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ProposerReward.Size()
		i -= size
		if _, err := m.ProposerReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFeedistr(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractWithdrawAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractWithdrawAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractWithdrawAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintFeedistr(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintFeedistr(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeedistr(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeedistr(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Distribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ProposerReward.Size()
	n += 1 + l + sovFeedistr(uint64(l))
	l = m.ContractRewards.Size()
	n += 1 + l + sovFeedistr(uint64(l))
	return n
}

func (m *ContractWithdrawAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovFeedistr(uint64(l))
	}
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovFeedistr(uint64(l))
	}
	return n
}

func sovFeedistr(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeedistr(x uint64) (n int) {
	return sovFeedistr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Distribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeedistr
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
			return fmt.Errorf("proto: Distribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeedistr
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
				return ErrInvalidLengthFeedistr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeedistr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProposerReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeedistr
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
				return ErrInvalidLengthFeedistr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeedistr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeedistr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeedistr
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
func (m *ContractWithdrawAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeedistr
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
			return fmt.Errorf("proto: ContractWithdrawAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractWithdrawAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeedistr
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
				return ErrInvalidLengthFeedistr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeedistr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeedistr
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
				return ErrInvalidLengthFeedistr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeedistr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeedistr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeedistr
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
func skipFeedistr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeedistr
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
					return 0, ErrIntOverflowFeedistr
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
					return 0, ErrIntOverflowFeedistr
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
				return 0, ErrInvalidLengthFeedistr
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeedistr
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeedistr
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeedistr        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeedistr          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeedistr = fmt.Errorf("proto: unexpected end of group")
)