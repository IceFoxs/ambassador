// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/original_src/v3alpha/original_src.proto

package envoy_config_filter_http_original_src_v3alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The Original Src filter binds upstream connections to the original source address determined
// for the request. This address could come from something like the Proxy Protocol filter, or it
// could come from trusted http headers.
type OriginalSrc struct {
	// Sets the SO_MARK option on the upstream connection's socket to the provided value. Used to
	// ensure that non-local addresses may be routed back through envoy when binding to the original
	// source address. The option will not be applied if the mark is 0.
	// [#proto-status: experimental]
	Mark                 uint32   `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c93f7450536988, []int{0}
}
func (m *OriginalSrc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginalSrc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginalSrc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginalSrc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalSrc.Merge(m, src)
}
func (m *OriginalSrc) XXX_Size() int {
	return m.Size()
}
func (m *OriginalSrc) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalSrc.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalSrc proto.InternalMessageInfo

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.config.filter.http.original_src.v3alpha.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/original_src/v3alpha/original_src.proto", fileDescriptor_05c93f7450536988)
}

var fileDescriptor_05c93f7450536988 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0x4b, 0xcc, 0x89, 0x2f, 0x2e,
	0x4a, 0xd6, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0x44, 0x11, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xd2, 0x05, 0x9b, 0xa0, 0x07, 0x31, 0x41, 0x0f, 0x62, 0x82, 0x1e, 0xc8, 0x04, 0x3d,
	0x14, 0xc5, 0x50, 0x13, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61,
	0x0c, 0x88, 0x39, 0x4a, 0x8a, 0x5c, 0xdc, 0xfe, 0x50, 0x0d, 0xc1, 0x45, 0xc9, 0x42, 0x42, 0x5c,
	0x2c, 0xb9, 0x89, 0x45, 0xd9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x60, 0xb6, 0x53, 0xec,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xc8, 0x65, 0x9d, 0x99,
	0xaf, 0x07, 0xb6, 0xbb, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x8f, 0x24, 0x67, 0x38, 0x09, 0x20, 0xd9,
	0x15, 0x00, 0xb2, 0x3f, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x10, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xaa, 0x5b, 0x95, 0x14, 0x01, 0x00, 0x00,
}

func (m *OriginalSrc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginalSrc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginalSrc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Mark != 0 {
		i = encodeVarintOriginalSrc(dAtA, i, uint64(m.Mark))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOriginalSrc(dAtA []byte, offset int, v uint64) int {
	offset -= sovOriginalSrc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OriginalSrc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mark != 0 {
		n += 1 + sovOriginalSrc(uint64(m.Mark))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOriginalSrc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOriginalSrc(x uint64) (n int) {
	return sovOriginalSrc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OriginalSrc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOriginalSrc
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
			return fmt.Errorf("proto: OriginalSrc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginalSrc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mark", wireType)
			}
			m.Mark = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginalSrc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mark |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOriginalSrc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOriginalSrc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOriginalSrc
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
func skipOriginalSrc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOriginalSrc
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
					return 0, ErrIntOverflowOriginalSrc
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
					return 0, ErrIntOverflowOriginalSrc
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
				return 0, ErrInvalidLengthOriginalSrc
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOriginalSrc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOriginalSrc
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
				next, err := skipOriginalSrc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOriginalSrc
				}
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
	ErrInvalidLengthOriginalSrc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOriginalSrc   = fmt.Errorf("proto: integer overflow")
)