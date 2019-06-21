// Code generated by protoc-gen-gogo. DO NOT EDIT.
// deprecated/deprecated.proto is a deprecated file.

// package deprecated contains only deprecated messages and services.

package deprecated

import (
	context "context"
	fmt "fmt"
	proto "github.com/frankee/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// DeprecatedEnum contains deprecated values.
type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	// DEPRECATED is the iota value of this enum.
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) String() string {
	return proto.EnumName(DeprecatedEnum_name, int32(x))
}

func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{0}
}

// DeprecatedRequest is a request to DeprecatedCall.
//
// Deprecated: Do not use.
type DeprecatedRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeprecatedRequest) Reset()         { *m = DeprecatedRequest{} }
func (m *DeprecatedRequest) String() string { return proto.CompactTextString(m) }
func (*DeprecatedRequest) ProtoMessage()    {}
func (*DeprecatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{0}
}
func (m *DeprecatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedRequest.Unmarshal(m, b)
}
func (m *DeprecatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedRequest.Marshal(b, m, deterministic)
}
func (m *DeprecatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedRequest.Merge(m, src)
}
func (m *DeprecatedRequest) XXX_Size() int {
	return xxx_messageInfo_DeprecatedRequest.Size(m)
}
func (m *DeprecatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedRequest proto.InternalMessageInfo

// Deprecated: Do not use.
type DeprecatedResponse struct {
	// DeprecatedField contains a DeprecatedEnum.
	DeprecatedField DeprecatedEnum `protobuf:"varint,1,opt,name=deprecated_field,json=deprecatedField,proto3,enum=deprecated.DeprecatedEnum" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	// DeprecatedOneof contains a deprecated field.
	//
	// Types that are valid to be assigned to DeprecatedOneof:
	//	*DeprecatedResponse_DeprecatedOneofField
	DeprecatedOneof      isDeprecatedResponse_DeprecatedOneof `protobuf_oneof:"deprecated_oneof"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DeprecatedResponse) Reset()         { *m = DeprecatedResponse{} }
func (m *DeprecatedResponse) String() string { return proto.CompactTextString(m) }
func (*DeprecatedResponse) ProtoMessage()    {}
func (*DeprecatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{1}
}
func (m *DeprecatedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedResponse.Unmarshal(m, b)
}
func (m *DeprecatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedResponse.Marshal(b, m, deterministic)
}
func (m *DeprecatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedResponse.Merge(m, src)
}
func (m *DeprecatedResponse) XXX_Size() int {
	return xxx_messageInfo_DeprecatedResponse.Size(m)
}
func (m *DeprecatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedResponse proto.InternalMessageInfo

type isDeprecatedResponse_DeprecatedOneof interface {
	isDeprecatedResponse_DeprecatedOneof()
}

type DeprecatedResponse_DeprecatedOneofField struct {
	DeprecatedOneofField string `protobuf:"bytes,2,opt,name=deprecated_oneof_field,json=deprecatedOneofField,proto3,oneof"`
}

func (*DeprecatedResponse_DeprecatedOneofField) isDeprecatedResponse_DeprecatedOneof() {}

func (m *DeprecatedResponse) GetDeprecatedOneof() isDeprecatedResponse_DeprecatedOneof {
	if m != nil {
		return m.DeprecatedOneof
	}
	return nil
}

// Deprecated: Do not use.
func (m *DeprecatedResponse) GetDeprecatedField() DeprecatedEnum {
	if m != nil {
		return m.DeprecatedField
	}
	return DeprecatedEnum_DEPRECATED
}

// Deprecated: Do not use.
func (m *DeprecatedResponse) GetDeprecatedOneofField() string {
	if x, ok := m.GetDeprecatedOneof().(*DeprecatedResponse_DeprecatedOneofField); ok {
		return x.DeprecatedOneofField
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeprecatedResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeprecatedResponse_OneofMarshaler, _DeprecatedResponse_OneofUnmarshaler, _DeprecatedResponse_OneofSizer, []interface{}{
		(*DeprecatedResponse_DeprecatedOneofField)(nil),
	}
}

func _DeprecatedResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeprecatedResponse)
	// deprecated_oneof
	switch x := m.DeprecatedOneof.(type) {
	case *DeprecatedResponse_DeprecatedOneofField:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.DeprecatedOneofField)
	case nil:
	default:
		return fmt.Errorf("DeprecatedResponse.DeprecatedOneof has unexpected type %T", x)
	}
	return nil
}

func _DeprecatedResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeprecatedResponse)
	switch tag {
	case 2: // deprecated_oneof.deprecated_oneof_field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DeprecatedOneof = &DeprecatedResponse_DeprecatedOneofField{x}
		return true, err
	default:
		return false, nil
	}
}

func _DeprecatedResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeprecatedResponse)
	// deprecated_oneof
	switch x := m.DeprecatedOneof.(type) {
	case *DeprecatedResponse_DeprecatedOneofField:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.DeprecatedOneofField)))
		n += len(x.DeprecatedOneofField)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("deprecated.DeprecatedEnum", DeprecatedEnum_name, DeprecatedEnum_value)
	proto.RegisterType((*DeprecatedRequest)(nil), "deprecated.DeprecatedRequest")
	proto.RegisterType((*DeprecatedResponse)(nil), "deprecated.DeprecatedResponse")
}

func init() { proto.RegisterFile("deprecated/deprecated.proto", fileDescriptor_f64ba265cd7eae3f) }

var fileDescriptor_f64ba265cd7eae3f = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x7b, 0xe6, 0x83, 0x0f, 0x9d, 0x45, 0xad, 0x83, 0x68, 0x88, 0x28, 0x25, 0xab, 0x20,
	0x34, 0x81, 0xba, 0xeb, 0x46, 0x4c, 0x13, 0xd1, 0x95, 0x12, 0xbb, 0x72, 0x23, 0xf9, 0x39, 0x89,
	0xc1, 0x74, 0x26, 0x26, 0x13, 0xaf, 0xc1, 0xfb, 0x71, 0xe3, 0xe5, 0xc9, 0xa4, 0xc5, 0x99, 0x82,
	0x6e, 0xc2, 0xc9, 0xbc, 0xef, 0x73, 0x7e, 0xe9, 0x69, 0x8e, 0x4d, 0x8b, 0x59, 0x22, 0x31, 0xf7,
	0x75, 0xe8, 0x35, 0xad, 0x90, 0x82, 0x51, 0xfd, 0xe2, 0x9c, 0xd0, 0xc3, 0xf0, 0xe7, 0x2f, 0xc6,
	0xb7, 0x1e, 0x3b, 0xb9, 0x20, 0x16, 0x38, 0x9f, 0x40, 0x99, 0xa9, 0x74, 0x8d, 0xe0, 0x1d, 0xb2,
	0x3b, 0x3a, 0xd1, 0xf4, 0x73, 0x51, 0x61, 0x9d, 0x5b, 0x30, 0x05, 0x77, 0x3c, 0xb7, 0x3d, 0xa3,
	0x90, 0x26, 0x23, 0xde, 0xaf, 0x03, 0x62, 0x41, 0x7c, 0xa0, 0xe5, 0x1b, 0x85, 0xb1, 0x05, 0x3d,
	0x36, 0x52, 0x09, 0x8e, 0xa2, 0xd8, 0x26, 0x24, 0x53, 0x70, 0xf7, 0x15, 0x74, 0x3b, 0x8a, 0x8f,
	0xb4, 0xe7, 0x5e, 0x59, 0x06, 0x56, 0x75, 0x18, 0xb0, 0x9d, 0x56, 0x06, 0xfe, 0xc2, 0xa5, 0xe3,
	0xdd, 0xd2, 0x8c, 0x51, 0x1a, 0x46, 0x0f, 0x71, 0xb4, 0xbc, 0x5e, 0x45, 0xe1, 0x64, 0x64, 0x93,
	0x3d, 0xb0, 0x89, 0x05, 0x73, 0x6e, 0x0e, 0xfe, 0x88, 0xed, 0x7b, 0x95, 0x21, 0x5b, 0x99, 0xf8,
	0x32, 0xa9, 0x6b, 0x76, 0xf6, 0xfb, 0x54, 0xdb, 0x4d, 0xd9, 0xe7, 0x7f, 0xc9, 0x9b, 0x75, 0x39,
	0xff, 0x3e, 0x08, 0xd8, 0xea, 0x13, 0x84, 0x4f, 0x57, 0x65, 0x25, 0x5f, 0xfa, 0xd4, 0xcb, 0xc4,
	0xda, 0x2f, 0xda, 0x84, 0xbf, 0x22, 0xfa, 0xc3, 0x41, 0xd2, 0xbe, 0xd8, 0x04, 0xd9, 0xac, 0x44,
	0x3e, 0x2b, 0x45, 0x29, 0x7c, 0x89, 0x9d, 0xcc, 0x13, 0x99, 0x18, 0xd7, 0xfb, 0x02, 0x48, 0xff,
	0x0f, 0xbe, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x46, 0x95, 0x5b, 0xe0, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeprecatedServiceClient is the client API for DeprecatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type DeprecatedServiceClient interface {
	// DeprecatedCall takes a DeprecatedRequest and returns a DeprecatedResponse.
	DeprecatedCall(ctx context.Context, in *DeprecatedRequest, opts ...grpc.CallOption) (*DeprecatedResponse, error)
}

type deprecatedServiceClient struct {
	cc *grpc.ClientConn
}

// Deprecated: Do not use.
func NewDeprecatedServiceClient(cc *grpc.ClientConn) DeprecatedServiceClient {
	return &deprecatedServiceClient{cc}
}

// Deprecated: Do not use.
func (c *deprecatedServiceClient) DeprecatedCall(ctx context.Context, in *DeprecatedRequest, opts ...grpc.CallOption) (*DeprecatedResponse, error) {
	out := new(DeprecatedResponse)
	err := c.cc.Invoke(ctx, "/deprecated.DeprecatedService/DeprecatedCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeprecatedServiceServer is the server API for DeprecatedService service.
//
// Deprecated: Do not use.
type DeprecatedServiceServer interface {
	// DeprecatedCall takes a DeprecatedRequest and returns a DeprecatedResponse.
	DeprecatedCall(context.Context, *DeprecatedRequest) (*DeprecatedResponse, error)
}

// Deprecated: Do not use.
// UnimplementedDeprecatedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeprecatedServiceServer struct {
}

func (*UnimplementedDeprecatedServiceServer) DeprecatedCall(ctx context.Context, req *DeprecatedRequest) (*DeprecatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprecatedCall not implemented")
}

// Deprecated: Do not use.
func RegisterDeprecatedServiceServer(s *grpc.Server, srv DeprecatedServiceServer) {
	s.RegisterService(&_DeprecatedService_serviceDesc, srv)
}

func _DeprecatedService_DeprecatedCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeprecatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeprecatedServiceServer).DeprecatedCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deprecated.DeprecatedService/DeprecatedCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeprecatedServiceServer).DeprecatedCall(ctx, req.(*DeprecatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeprecatedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deprecated.DeprecatedService",
	HandlerType: (*DeprecatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeprecatedCall",
			Handler:    _DeprecatedService_DeprecatedCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deprecated/deprecated.proto",
}
