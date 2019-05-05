// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pubsubhub/rpc/v1/status.proto

package rpcv1

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Status repesents operation status response.
type Status struct {
	// The status code, which should be an enum value of
	// [pubsubhub.system.v1.Code][pubsubhub.system.v1.Code].
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [pubsubhub.system.v1.Status.details][pubsubhub.system.v1.Status.details] field, or localized
	// by the client.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details              []*types.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_807c87c420d5ef57, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}

func (m *Status) XXX_Size() int {
	return m.Size()
}

func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*types.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "pubsubhub.rpc.v1.Status")
}

func init() { proto.RegisterFile("pubsubhub/rpc/v1/status.proto", fileDescriptor_807c87c420d5ef57) }

var fileDescriptor_807c87c420d5ef57 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x99, 0xc6, 0xb6, 0x38, 0xdd, 0x94, 0xe0, 0x22, 0x16, 0x1a, 0x82, 0xab, 0x6c, 0x7c,
	0x21, 0x7a, 0x02, 0x7b, 0x82, 0x30, 0x42, 0x11, 0xe9, 0x26, 0x33, 0x9d, 0x4e, 0x03, 0x6d, 0x66,
	0x98, 0x3f, 0x81, 0x7a, 0x1c, 0x57, 0xe2, 0x49, 0xc4, 0x95, 0x47, 0xd0, 0x9c, 0xc2, 0xa5, 0x74,
	0x42, 0xaa, 0xb8, 0x7b, 0xbf, 0x79, 0x3f, 0xde, 0xf7, 0x0d, 0x9e, 0x2b, 0x47, 0x8d, 0xa3, 0x5b,
	0x47, 0x33, 0xad, 0x58, 0xd6, 0xe4, 0x99, 0xb1, 0xa5, 0x75, 0x06, 0x94, 0x96, 0x56, 0x86, 0xd3,
	0xd3, 0x1a, 0xb4, 0x62, 0xd0, 0xe4, 0xb3, 0x6b, 0x51, 0xd9, 0x23, 0x32, 0xb9, 0xcf, 0x84, 0x14,
	0x32, 0xf3, 0x22, 0x75, 0x1b, 0x4f, 0x1e, 0xfc, 0xd4, 0x1d, 0x98, 0x5d, 0x0a, 0x29, 0xc5, 0x8e,
	0xff, 0x5a, 0x65, 0x7d, 0xe8, 0x56, 0x57, 0x1b, 0x3c, 0xba, 0xf7, 0x59, 0x61, 0x88, 0xcf, 0x98,
	0x5c, 0xf3, 0x08, 0x25, 0x28, 0x1d, 0x12, 0x3f, 0x87, 0x11, 0x1e, 0xef, 0xb9, 0x31, 0xa5, 0xe0,
	0xd1, 0x20, 0x41, 0xe9, 0x39, 0xe9, 0x31, 0x04, 0x3c, 0x5e, 0x73, 0x5b, 0x56, 0x3b, 0x13, 0x05,
	0x49, 0x90, 0x4e, 0x6e, 0x2e, 0xa0, 0x0b, 0x81, 0x3e, 0x04, 0xee, 0xea, 0x03, 0xe9, 0xa5, 0x85,
	0xfb, 0xfe, 0x8a, 0xd1, 0x4b, 0x1b, 0xa3, 0xb7, 0x36, 0x46, 0x1f, 0x6d, 0x8c, 0x3e, 0xdb, 0x18,
	0xe1, 0xb9, 0xd4, 0x02, 0x9e, 0x78, 0x5d, 0xd9, 0x6d, 0xa9, 0xe1, 0xff, 0x37, 0x17, 0x93, 0xae,
	0x5a, 0x71, 0xbc, 0x58, 0xa0, 0xc7, 0xa1, 0x56, 0xac, 0xc9, 0x9f, 0x07, 0x41, 0x41, 0x1e, 0x5e,
	0x07, 0xd3, 0xe2, 0xa4, 0x13, 0xc5, 0x60, 0x99, 0xbf, 0xff, 0x79, 0x5a, 0x11, 0xc5, 0x56, 0xcb,
	0x9c, 0x8e, 0x7c, 0x9b, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x80, 0xd8, 0xfd, 0x62,
	0x01, 0x00, 0x00,
}

func (this *Status) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if len(this.Details) != len(that1.Details) {
		return false
	}
	for i := range this.Details {
		if !this.Details[i].Equal(that1.Details[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStatus(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.Details) > 0 {
		for _, msg := range m.Details {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintStatus(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func NewPopulatedStatus(r randyStatus, easy bool) *Status {
	this := &Status{}
	this.Code = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Code *= -1
	}
	this.Message = string(randStringStatus(r))
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Details = make([]*types.Any, v1)
		for i := 0; i < v1; i++ {
			this.Details[i] = types.NewPopulatedAny(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedStatus(r, 4)
	}
	return this
}

type randyStatus interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneStatus(r randyStatus) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}

func randStringStatus(r randyStatus) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneStatus(r)
	}
	return string(tmps)
}

func randUnrecognizedStatus(r randyStatus, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldStatus(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}

func randFieldStatus(dAtA []byte, r randyStatus, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateStatus(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateStatus(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateStatus(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateStatus(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateStatus(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateStatus(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}

func encodeVarintPopulateStatus(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovStatus(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	if len(m.Details) > 0 {
		for _, e := range m.Details {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
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
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
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
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = append(m.Details, &types.Any{})
			if err := m.Details[len(m.Details)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatus
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

func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
				return 0, ErrInvalidLengthStatus
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatus
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
				next, err := skipStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthStatus
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
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus   = fmt.Errorf("proto: integer overflow")
)