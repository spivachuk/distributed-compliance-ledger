// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/rejected_validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/gogo/protobuf/proto"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

type RejectedDisableValidator struct {
	Address   string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Creator   string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Approvals []*Grant `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
	Rejects   []*Grant `protobuf:"bytes,4,rep,name=rejects,proto3" json:"rejects,omitempty"`
}

func (m *RejectedDisableValidator) Reset()         { *m = RejectedDisableValidator{} }
func (m *RejectedDisableValidator) String() string { return proto.CompactTextString(m) }
func (*RejectedDisableValidator) ProtoMessage()    {}
func (*RejectedDisableValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d82fd5546ec0be, []int{0}
}
func (m *RejectedDisableValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RejectedDisableValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RejectedDisableValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RejectedDisableValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RejectedDisableValidator.Merge(m, src)
}
func (m *RejectedDisableValidator) XXX_Size() int {
	return m.Size()
}
func (m *RejectedDisableValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_RejectedDisableValidator.DiscardUnknown(m)
}

var xxx_messageInfo_RejectedDisableValidator proto.InternalMessageInfo

func (m *RejectedDisableValidator) GetAddress() sdk.ValAddress {
	valAddr, _ := sdk.ValAddressFromBech32(m.Address)
	return valAddr
}

func (m *RejectedDisableValidator) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RejectedDisableValidator) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *RejectedDisableValidator) GetRejects() []*Grant {
	if m != nil {
		return m.Rejects
	}
	return nil
}

func init() {
	proto.RegisterType((*RejectedDisableValidator)(nil), "zigbeealliance.distributedcomplianceledger.validator.RejectedDisableValidator")
}

func init() {
	proto.RegisterFile("validator/rejected_validator.proto", fileDescriptor_e8d82fd5546ec0be)
}

var fileDescriptor_e8d82fd5546ec0be = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xd7, 0xed, 0xe5, 0x1d, 0xab, 0xb7, 0xa2, 0x10, 0x77, 0x08, 0x63, 0xa7, 0x5d, 0x9a,
	0xc0, 0xf4, 0xe6, 0xc9, 0x21, 0x08, 0x1e, 0x27, 0x0a, 0x7a, 0x19, 0x69, 0xf2, 0x50, 0x23, 0x59,
	0x53, 0x92, 0x6c, 0xa8, 0x9f, 0xc2, 0x0f, 0xb3, 0x0f, 0xe1, 0x71, 0x78, 0xf2, 0x28, 0xed, 0x17,
	0x91, 0x2e, 0xed, 0xea, 0x4d, 0xd0, 0x63, 0x9e, 0x3c, 0xff, 0xdf, 0x2f, 0x79, 0x9e, 0x70, 0xbc,
	0x66, 0x4a, 0x0a, 0xe6, 0xb4, 0xa1, 0x06, 0x1e, 0x81, 0x3b, 0x10, 0x8b, 0x7d, 0x89, 0xe4, 0x46,
	0x3b, 0x1d, 0x9d, 0xbe, 0xc8, 0x34, 0x01, 0x60, 0x4a, 0x49, 0x96, 0x71, 0x20, 0x42, 0x5a, 0x67,
	0x64, 0xb2, 0x72, 0x20, 0xb8, 0x5e, 0xe6, 0xbe, 0xaa, 0x40, 0xa4, 0x60, 0xc8, 0x3e, 0x3b, 0x3c,
	0xe6, 0xda, 0x2e, 0xb5, 0x5d, 0xec, 0x18, 0xd4, 0x1f, 0x3c, 0x70, 0x78, 0xd4, 0x4a, 0x53, 0xc3,
	0x32, 0xe7, 0xcb, 0xe3, 0x4d, 0x37, 0x44, 0xf3, 0xfa, 0x11, 0x17, 0xd2, 0xb2, 0x44, 0xc1, 0x6d,
	0xd3, 0x18, 0x4d, 0xc3, 0x3e, 0x13, 0xc2, 0x80, 0xb5, 0x28, 0x18, 0x05, 0x93, 0xc1, 0x0c, 0xbd,
	0x6f, 0xe2, 0xc3, 0x1a, 0x7b, 0xee, 0x6f, 0xae, 0x9d, 0x91, 0x59, 0x3a, 0x6f, 0x1a, 0xab, 0x0c,
	0x37, 0x50, 0xc5, 0x51, 0xf7, 0xa7, 0x4c, 0xdd, 0x18, 0xdd, 0x85, 0x03, 0x96, 0xe7, 0x46, 0xaf,
	0x99, 0xb2, 0xa8, 0x37, 0xea, 0x4d, 0x0e, 0xa6, 0x67, 0xe4, 0x37, 0x03, 0x20, 0x97, 0xd5, 0xd7,
	0xe6, 0x2d, 0x2d, 0xba, 0x09, 0xfb, 0x7e, 0xc6, 0x16, 0xfd, 0xfb, 0x3b, 0xb8, 0x61, 0xcd, 0xc4,
	0x5b, 0x81, 0x83, 0x6d, 0x81, 0x83, 0xcf, 0x02, 0x07, 0xaf, 0x25, 0xee, 0x6c, 0x4b, 0xdc, 0xf9,
	0x28, 0x71, 0xe7, 0xfe, 0x2a, 0x95, 0xee, 0x61, 0x95, 0x10, 0xae, 0x97, 0xd4, 0x9b, 0xe2, 0x46,
	0x45, 0xbf, 0xa9, 0xe2, 0xd6, 0x15, 0x7b, 0x19, 0x7d, 0xa2, 0xed, 0x8a, 0xdc, 0x73, 0x0e, 0x36,
	0xf9, 0xbf, 0xdb, 0xd1, 0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xce, 0x16, 0x1a, 0x31,
	0x02, 0x00, 0x00,
}

func (m *RejectedDisableValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RejectedDisableValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RejectedDisableValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rejects) > 0 {
		for iNdEx := len(m.Rejects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rejects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRejectedValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRejectedValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRejectedValidator(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRejectedValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRejectedValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovRejectedValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RejectedDisableValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRejectedValidator(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRejectedValidator(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovRejectedValidator(uint64(l))
		}
	}
	if len(m.Rejects) > 0 {
		for _, e := range m.Rejects {
			l = e.Size()
			n += 1 + l + sovRejectedValidator(uint64(l))
		}
	}
	return n
}

func sovRejectedValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRejectedValidator(x uint64) (n int) {
	return sovRejectedValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RejectedDisableValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRejectedValidator
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
			return fmt.Errorf("proto: RejectedDisableValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RejectedDisableValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedValidator
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
				return ErrInvalidLengthRejectedValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedValidator
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
				return ErrInvalidLengthRejectedValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedValidator
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
				return ErrInvalidLengthRejectedValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &Grant{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedValidator
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
				return ErrInvalidLengthRejectedValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rejects = append(m.Rejects, &Grant{})
			if err := m.Rejects[len(m.Rejects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRejectedValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRejectedValidator
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
func skipRejectedValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRejectedValidator
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
					return 0, ErrIntOverflowRejectedValidator
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
					return 0, ErrIntOverflowRejectedValidator
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
				return 0, ErrInvalidLengthRejectedValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRejectedValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRejectedValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRejectedValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRejectedValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRejectedValidator = fmt.Errorf("proto: unexpected end of group")
)
