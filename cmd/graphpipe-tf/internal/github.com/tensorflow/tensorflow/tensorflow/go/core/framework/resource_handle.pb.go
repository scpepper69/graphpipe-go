// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandleProto) Reset()         { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()    {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptorResourceHandle, []int{0}
}

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}
func (m *ResourceHandleProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceHandleProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	if len(m.Container) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Container)))
		i += copy(dAtA[i:], m.Container)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.HashCode != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(m.HashCode))
	}
	if len(m.MaybeTypeName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.MaybeTypeName)))
		i += copy(dAtA[i:], m.MaybeTypeName)
	}
	return i, nil
}

func encodeVarintResourceHandle(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceHandleProto) Size() (n int) {
	var l int
	_ = l
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Container)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	if m.HashCode != 0 {
		n += 1 + sovResourceHandle(uint64(m.HashCode))
	}
	l = len(m.MaybeTypeName)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	return n
}

func sovResourceHandle(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourceHandle(x uint64) (n int) {
	return sovResourceHandle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceHandleProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceHandle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceHandleProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceHandleProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Container = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashCode", wireType)
			}
			m.HashCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HashCode |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaybeTypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaybeTypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceHandle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceHandle
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
func skipResourceHandle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResourceHandle
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResourceHandle
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourceHandle
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipResourceHandle(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthResourceHandle = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceHandle   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptorResourceHandle)
}

var fileDescriptorResourceHandle = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0x19, 0x8d, 0x87, 0x59, 0x50, 0x61, 0x05, 0x09, 0x28, 0xe1, 0xb0, 0x90, 0xab, 0x92,
	0xc2, 0xda, 0xe6, 0x6c, 0xac, 0x44, 0x82, 0x95, 0x4d, 0xd8, 0x6c, 0xe6, 0x92, 0xe0, 0xed, 0x4e,
	0xd8, 0xec, 0x79, 0xa4, 0xf1, 0x69, 0x7c, 0x18, 0x4b, 0x4b, 0x4b, 0xc9, 0x53, 0x58, 0x8a, 0xc3,
	0x61, 0xe4, 0xb8, 0x6e, 0xe6, 0xfb, 0x53, 0x7c, 0x3f, 0x91, 0x7a, 0xb4, 0x1d, 0xb9, 0xc5, 0x92,
	0xd6, 0xa9, 0x26, 0x87, 0xe9, 0xc2, 0x29, 0x83, 0x6b, 0x72, 0xcf, 0xa9, 0xc3, 0x8e, 0x56, 0x4e,
	0x63, 0x5e, 0x2b, 0x5b, 0x2e, 0x31, 0x69, 0x1d, 0x79, 0x92, 0x62, 0x2c, 0x5c, 0xbe, 0x81, 0x38,
	0xcd, 0x36, 0xa9, 0x3b, 0x0e, 0x3d, 0x70, 0xe6, 0x4c, 0x4c, 0x4a, 0x7c, 0x69, 0x34, 0x46, 0x30,
	0x85, 0x59, 0x98, 0x6d, 0x3e, 0x79, 0x21, 0x42, 0x4d, 0xd6, 0xab, 0xc6, 0xa2, 0x8b, 0xf6, 0xd8,
	0x1a, 0x05, 0x29, 0x45, 0x60, 0x95, 0xc1, 0x68, 0x9f, 0x0d, 0xbe, 0xe5, 0xb9, 0x08, 0x6b, 0xd5,
	0xd5, 0xb9, 0xa6, 0x12, 0xa3, 0x60, 0x0a, 0xb3, 0x20, 0x3b, 0xfc, 0x15, 0x6e, 0xa9, 0x44, 0x79,
	0x25, 0x4e, 0x8c, 0xea, 0x0b, 0xcc, 0x7d, 0xdf, 0x62, 0xce, 0xdd, 0x03, 0xee, 0x1e, 0xb1, 0xfc,
	0xd8, 0xb7, 0x78, 0xaf, 0x0c, 0xce, 0x5f, 0xdf, 0x87, 0x18, 0x3e, 0x86, 0x18, 0x3e, 0x87, 0x18,
	0xbe, 0x86, 0x18, 0x44, 0x44, 0xae, 0x4a, 0xc6, 0x21, 0xc9, 0xdf, 0xe8, 0xf9, 0xf1, 0xd6, 0x1e,
	0x78, 0xba, 0xa9, 0x1a, 0x5f, 0xaf, 0x8a, 0x44, 0x93, 0xf9, 0x0f, 0x6b, 0xf7, 0x59, 0xd1, 0x16,
	0xc5, 0x6f, 0x80, 0x62, 0xc2, 0xe4, 0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xab, 0xae, 0xe0,
	0xe3, 0x6c, 0x01, 0x00, 0x00,
}