// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/notifications.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NetworkPolicyNotification struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Yaml                 string   `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkPolicyNotification) Reset()         { *m = NetworkPolicyNotification{} }
func (m *NetworkPolicyNotification) String() string { return proto.CompactTextString(m) }
func (*NetworkPolicyNotification) ProtoMessage()    {}
func (*NetworkPolicyNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e614fc19f85dc80, []int{0}
}
func (m *NetworkPolicyNotification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkPolicyNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkPolicyNotification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkPolicyNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPolicyNotification.Merge(m, src)
}
func (m *NetworkPolicyNotification) XXX_Size() int {
	return m.Size()
}
func (m *NetworkPolicyNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPolicyNotification.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPolicyNotification proto.InternalMessageInfo

func (m *NetworkPolicyNotification) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *NetworkPolicyNotification) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

func (m *NetworkPolicyNotification) MessageClone() proto.Message {
	return m.Clone()
}
func (m *NetworkPolicyNotification) Clone() *NetworkPolicyNotification {
	if m == nil {
		return nil
	}
	cloned := new(NetworkPolicyNotification)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*NetworkPolicyNotification)(nil), "v1.NetworkPolicyNotification")
}

func init() { proto.RegisterFile("api/v1/notifications.proto", fileDescriptor_0e614fc19f85dc80) }

var fileDescriptor_0e614fc19f85dc80 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54, 0xf2, 0xe4, 0x92, 0xf4, 0x4b,
	0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x0e, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0xf4, 0x43, 0x52, 0x27, 0x24,
	0xc1, 0xc5, 0x9e, 0x9c, 0x53, 0x5a, 0x5c, 0x92, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0xe3, 0x0a, 0x09, 0x71, 0xb1, 0x54, 0x26, 0xe6, 0xe6, 0x48, 0x30, 0x81, 0x85, 0xc1, 0x6c,
	0x27, 0xf3, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6,
	0x63, 0x39, 0x06, 0x2e, 0x89, 0xcc, 0x7c, 0xbd, 0xe2, 0x92, 0xc4, 0xe4, 0xec, 0xa2, 0xfc, 0x0a,
	0x88, 0x9d, 0x7a, 0x89, 0x05, 0x99, 0x7a, 0x65, 0x86, 0x51, 0xdc, 0x7a, 0xfa, 0x10, 0x87, 0x59,
	0x97, 0x19, 0x26, 0xb1, 0x81, 0xa5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x74, 0x82, 0xbf,
	0x9f, 0xac, 0x00, 0x00, 0x00,
}

func (m *NetworkPolicyNotification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicyNotification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkPolicyNotification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Yaml) > 0 {
		i -= len(m.Yaml)
		copy(dAtA[i:], m.Yaml)
		i = encodeVarintNotifications(dAtA, i, uint64(len(m.Yaml)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = encodeVarintNotifications(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNotifications(dAtA []byte, offset int, v uint64) int {
	offset -= sovNotifications(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetworkPolicyNotification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovNotifications(uint64(l))
	}
	l = len(m.Yaml)
	if l > 0 {
		n += 1 + l + sovNotifications(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNotifications(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNotifications(x uint64) (n int) {
	return sovNotifications(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkPolicyNotification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotifications
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
			return fmt.Errorf("proto: NetworkPolicyNotification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicyNotification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotifications
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
				return ErrInvalidLengthNotifications
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotifications
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Yaml", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotifications
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
				return ErrInvalidLengthNotifications
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotifications
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Yaml = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotifications(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNotifications
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNotifications(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotifications
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
					return 0, ErrIntOverflowNotifications
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
					return 0, ErrIntOverflowNotifications
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
				return 0, ErrInvalidLengthNotifications
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNotifications
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNotifications
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNotifications        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotifications          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNotifications = fmt.Errorf("proto: unexpected end of group")
)
