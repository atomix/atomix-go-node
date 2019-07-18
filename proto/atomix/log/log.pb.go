// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/log/log.proto

package log

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProduceRequest struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfa61300f829c393, []int{0}
}

func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRequest.Unmarshal(m, b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return xxx_messageInfo_ProduceRequest.Size(m)
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ProduceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfa61300f829c393, []int{1}
}

func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResponse.Unmarshal(m, b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
}
func (m *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(m, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return xxx_messageInfo_ProduceResponse.Size(m)
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

type ConsumeRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfa61300f829c393, []int{2}
}

func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRequest.Unmarshal(m, b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeRequest.Size(m)
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

func (m *ConsumeRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type LogRecord struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Timestamp            uint64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfa61300f829c393, []int{3}
}

func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRecord.Unmarshal(m, b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return xxx_messageInfo_LogRecord.Size(m)
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *LogRecord) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *LogRecord) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ProduceRequest)(nil), "atomix.log.ProduceRequest")
	proto.RegisterType((*ProduceResponse)(nil), "atomix.log.ProduceResponse")
	proto.RegisterType((*ConsumeRequest)(nil), "atomix.log.ConsumeRequest")
	proto.RegisterType((*LogRecord)(nil), "atomix.log.LogRecord")
}

func init() { proto.RegisterFile("atomix/log/log.proto", fileDescriptor_dfa61300f829c393) }

var fileDescriptor_dfa61300f829c393 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0xcf, 0xc9, 0x4f, 0x07, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x2e,
	0x88, 0xa8, 0x5e, 0x4e, 0x7e, 0xba, 0x92, 0x1a, 0x17, 0x5f, 0x40, 0x51, 0x7e, 0x4a, 0x69, 0x72,
	0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x84, 0xa3, 0x24, 0xc8, 0xc5, 0x0f, 0x57, 0x57,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xa4, 0xc1, 0xc5, 0xe7, 0x9c, 0x9f, 0x57, 0x5c, 0x9a, 0x0b,
	0xd7, 0x2a, 0xc6, 0xc5, 0x96, 0x9f, 0x96, 0x56, 0x9c, 0x5a, 0x02, 0xd6, 0xcb, 0x12, 0x04, 0xe5,
	0x29, 0x85, 0x73, 0x71, 0xfa, 0xe4, 0xa7, 0x07, 0xa5, 0x26, 0xe7, 0x17, 0xa5, 0xe0, 0x52, 0x24,
	0x24, 0xc3, 0xc5, 0x59, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x20, 0xc1, 0x04, 0x96,
	0x42, 0x08, 0x20, 0x5c, 0xc5, 0x8c, 0xe4, 0x2a, 0xa3, 0x69, 0x8c, 0x5c, 0x5c, 0x3e, 0xf9, 0xe9,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x6e, 0x5c, 0xec, 0x50, 0x47, 0x0a, 0x49, 0xe9,
	0x21, 0x3c, 0xa9, 0x87, 0xea, 0x43, 0x29, 0x69, 0xac, 0x72, 0x50, 0x5f, 0x31, 0x68, 0x30, 0x0a,
	0x39, 0x70, 0xb1, 0x43, 0x7d, 0x86, 0x6a, 0x0e, 0xaa, 0x77, 0xa5, 0x44, 0x91, 0xe5, 0xe0, 0x1e,
	0x54, 0x62, 0x30, 0x60, 0x4c, 0x62, 0x03, 0x87, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x36,
	0x05, 0xfb, 0x4c, 0x81, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	Produce(ctx context.Context, opts ...grpc.CallOption) (LogService_ProduceClient, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (LogService_ConsumeClient, error)
}

type logServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogServiceClient(cc *grpc.ClientConn) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Produce(ctx context.Context, opts ...grpc.CallOption) (LogService_ProduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[0], "/atomix.log.LogService/Produce", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceProduceClient{stream}
	return x, nil
}

type LogService_ProduceClient interface {
	Send(*ProduceRequest) error
	CloseAndRecv() (*ProduceResponse, error)
	grpc.ClientStream
}

type logServiceProduceClient struct {
	grpc.ClientStream
}

func (x *logServiceProduceClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logServiceProduceClient) CloseAndRecv() (*ProduceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (LogService_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[1], "/atomix.log.LogService/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceConsumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_ConsumeClient interface {
	Recv() (*LogRecord, error)
	grpc.ClientStream
}

type logServiceConsumeClient struct {
	grpc.ClientStream
}

func (x *logServiceConsumeClient) Recv() (*LogRecord, error) {
	m := new(LogRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	Produce(LogService_ProduceServer) error
	Consume(*ConsumeRequest, LogService_ConsumeServer) error
}

// UnimplementedLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (*UnimplementedLogServiceServer) Produce(srv LogService_ProduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (*UnimplementedLogServiceServer) Consume(req *ConsumeRequest, srv LogService_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_Produce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServiceServer).Produce(&logServiceProduceServer{stream})
}

type LogService_ProduceServer interface {
	SendAndClose(*ProduceResponse) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type logServiceProduceServer struct {
	grpc.ServerStream
}

func (x *logServiceProduceServer) SendAndClose(m *ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logServiceProduceServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LogService_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).Consume(m, &logServiceConsumeServer{stream})
}

type LogService_ConsumeServer interface {
	Send(*LogRecord) error
	grpc.ServerStream
}

type logServiceConsumeServer struct {
	grpc.ServerStream
}

func (x *logServiceConsumeServer) Send(m *LogRecord) error {
	return x.ServerStream.SendMsg(m)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.log.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Produce",
			Handler:       _LogService_Produce_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Consume",
			Handler:       _LogService_Consume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "atomix/log/log.proto",
}
