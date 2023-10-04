// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package grpc

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ScheduleBroadcastTxRequest struct {
	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *ScheduleBroadcastTxRequest) Reset()         { *m = ScheduleBroadcastTxRequest{} }
func (m *ScheduleBroadcastTxRequest) String() string { return proto.CompactTextString(m) }
func (*ScheduleBroadcastTxRequest) ProtoMessage()    {}
func (*ScheduleBroadcastTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *ScheduleBroadcastTxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleBroadcastTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduleBroadcastTxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduleBroadcastTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleBroadcastTxRequest.Merge(m, src)
}
func (m *ScheduleBroadcastTxRequest) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleBroadcastTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleBroadcastTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleBroadcastTxRequest proto.InternalMessageInfo

func (m *ScheduleBroadcastTxRequest) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type ScheduleBroadcastTxResponse struct {
}

func (m *ScheduleBroadcastTxResponse) Reset()         { *m = ScheduleBroadcastTxResponse{} }
func (m *ScheduleBroadcastTxResponse) String() string { return proto.CompactTextString(m) }
func (*ScheduleBroadcastTxResponse) ProtoMessage()    {}
func (*ScheduleBroadcastTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *ScheduleBroadcastTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleBroadcastTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduleBroadcastTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduleBroadcastTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleBroadcastTxResponse.Merge(m, src)
}
func (m *ScheduleBroadcastTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleBroadcastTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleBroadcastTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleBroadcastTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScheduleBroadcastTxRequest)(nil), "ScheduleBroadcastTxRequest")
	proto.RegisterType((*ScheduleBroadcastTxResponse)(nil), "ScheduleBroadcastTxResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe1, 0x92, 0x0a, 0x4e, 0xce,
	0x48, 0x4d, 0x29, 0xcd, 0x49, 0x75, 0x2a, 0xca, 0x4f, 0x4c, 0x49, 0x4e, 0x2c, 0x2e, 0x09, 0xa9,
	0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0x2a, 0xa9, 0x90, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x09, 0x62, 0x2a, 0xa9, 0x50, 0x92, 0xe5, 0x92, 0xc6, 0xaa, 0xba, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0xd5, 0x28, 0x9e, 0x8b, 0x1b, 0x2e, 0x9c, 0x5a, 0x24, 0x14, 0xc0, 0x25, 0x8c,
	0x45, 0xb5, 0x90, 0xb4, 0x1e, 0x6e, 0x1b, 0xa5, 0x64, 0xf4, 0xf0, 0x58, 0xe0, 0x14, 0x74, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70,
	0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x16, 0xe9, 0x99, 0x25, 0x39, 0x89, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x45, 0x89, 0x45, 0x99, 0xb9, 0xf9, 0xfa, 0x49, 0x08, 0xa7, 0xe8,
	0x16, 0x97, 0x25, 0xeb, 0x67, 0xe6, 0x95, 0xa4, 0x16, 0xe5, 0x25, 0xe6, 0xe8, 0x43, 0x43, 0xa0,
	0x58, 0x3f, 0xbd, 0xa8, 0x20, 0x39, 0x89, 0x0d, 0x1c, 0x10, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0xdc, 0xe1, 0xf2, 0x19, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BroadcasterClient is the client API for Broadcaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BroadcasterClient interface {
	ScheduleBroadcastTx(ctx context.Context, in *ScheduleBroadcastTxRequest, opts ...grpc.CallOption) (*ScheduleBroadcastTxResponse, error)
}

type broadcasterClient struct {
	cc grpc1.ClientConn
}

func NewBroadcasterClient(cc grpc1.ClientConn) BroadcasterClient {
	return &broadcasterClient{cc}
}

func (c *broadcasterClient) ScheduleBroadcastTx(ctx context.Context, in *ScheduleBroadcastTxRequest, opts ...grpc.CallOption) (*ScheduleBroadcastTxResponse, error) {
	out := new(ScheduleBroadcastTxResponse)
	err := c.cc.Invoke(ctx, "/Broadcaster/ScheduleBroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcasterServer is the server API for Broadcaster service.
type BroadcasterServer interface {
	ScheduleBroadcastTx(context.Context, *ScheduleBroadcastTxRequest) (*ScheduleBroadcastTxResponse, error)
}

// UnimplementedBroadcasterServer can be embedded to have forward compatible implementations.
type UnimplementedBroadcasterServer struct {
}

func (*UnimplementedBroadcasterServer) ScheduleBroadcastTx(ctx context.Context, req *ScheduleBroadcastTxRequest) (*ScheduleBroadcastTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleBroadcastTx not implemented")
}

func RegisterBroadcasterServer(s grpc1.Server, srv BroadcasterServer) {
	s.RegisterService(&_Broadcaster_serviceDesc, srv)
}

func _Broadcaster_ScheduleBroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleBroadcastTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcasterServer).ScheduleBroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Broadcaster/ScheduleBroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcasterServer).ScheduleBroadcastTx(ctx, req.(*ScheduleBroadcastTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broadcaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Broadcaster",
	HandlerType: (*BroadcasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleBroadcastTx",
			Handler:    _Broadcaster_ScheduleBroadcastTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func (m *ScheduleBroadcastTxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleBroadcastTxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleBroadcastTxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintService(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScheduleBroadcastTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleBroadcastTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleBroadcastTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduleBroadcastTxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *ScheduleBroadcastTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduleBroadcastTxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ScheduleBroadcastTxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleBroadcastTxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *ScheduleBroadcastTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ScheduleBroadcastTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleBroadcastTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
