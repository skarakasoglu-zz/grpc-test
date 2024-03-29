// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculatorpb/calculator.proto

package calculatorpb

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

type SumRequest struct {
	Num1                 int32    `protobuf:"zigzag32,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2                 int32    `protobuf:"zigzag32,2,opt,name=num2,proto3" json:"num2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetNum1() int32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *SumRequest) GetNum2() int32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

type SumResponse struct {
	Result               int32    `protobuf:"zigzag32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatePrimeNumberDecompositionRequest struct {
	Number               int32    `protobuf:"zigzag32,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatePrimeNumberDecompositionRequest) Reset() {
	*m = CalculatePrimeNumberDecompositionRequest{}
}
func (m *CalculatePrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatePrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*CalculatePrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{2}
}

func (m *CalculatePrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatePrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *CalculatePrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatePrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *CalculatePrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatePrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *CalculatePrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatePrimeNumberDecompositionRequest.Size(m)
}
func (m *CalculatePrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatePrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatePrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *CalculatePrimeNumberDecompositionRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CalculatePrimeNumberDecompositionResponse struct {
	PrimeNumber          int32    `protobuf:"zigzag32,1,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatePrimeNumberDecompositionResponse) Reset() {
	*m = CalculatePrimeNumberDecompositionResponse{}
}
func (m *CalculatePrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatePrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*CalculatePrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{3}
}

func (m *CalculatePrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatePrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *CalculatePrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatePrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *CalculatePrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatePrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *CalculatePrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatePrimeNumberDecompositionResponse.Size(m)
}
func (m *CalculatePrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatePrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatePrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *CalculatePrimeNumberDecompositionResponse) GetPrimeNumber() int32 {
	if m != nil {
		return m.PrimeNumber
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"zigzag32,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{4}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Average              float32  `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{5}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetAverage() float32 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int32    `protobuf:"zigzag32,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximum              int32    `protobuf:"zigzag32,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e717c78a24322a, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*CalculatePrimeNumberDecompositionRequest)(nil), "calculator.CalculatePrimeNumberDecompositionRequest")
	proto.RegisterType((*CalculatePrimeNumberDecompositionResponse)(nil), "calculator.CalculatePrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
}

func init() { proto.RegisterFile("calculatorpb/calculator.proto", fileDescriptor_87e717c78a24322a) }

var fileDescriptor_87e717c78a24322a = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x34, 0xad, 0x54, 0x78, 0x2d, 0x85, 0xae, 0x18, 0x4b, 0xc0, 0x8f, 0x2e, 0x08, 0x15, 0xa4,
	0xad, 0xb1, 0x82, 0x57, 0x5b, 0xf1, 0x66, 0x91, 0xd6, 0x93, 0x1e, 0x24, 0x8d, 0x0f, 0x09, 0x74,
	0xb3, 0xeb, 0x26, 0x5b, 0xfc, 0x21, 0x82, 0x7f, 0x57, 0x4c, 0x36, 0xdd, 0xc4, 0x5a, 0x5b, 0x6f,
	0x99, 0xc9, 0x9b, 0x99, 0xc7, 0x1b, 0x16, 0x0e, 0x7c, 0x6f, 0xe6, 0xab, 0x99, 0x17, 0x73, 0x29,
	0xa6, 0x5d, 0x03, 0x3a, 0x42, 0xf2, 0x98, 0x13, 0x30, 0x0c, 0xed, 0x03, 0x4c, 0x14, 0x1b, 0xe3,
	0x9b, 0xc2, 0x28, 0x26, 0x04, 0xb6, 0x43, 0xc5, 0xce, 0x9b, 0xd6, 0xb1, 0xd5, 0x6e, 0x8c, 0x93,
	0x6f, 0xcd, 0xb9, 0xcd, 0xd2, 0x82, 0x73, 0xe9, 0x09, 0x54, 0x13, 0x55, 0x24, 0x78, 0x18, 0x21,
	0xb1, 0xa1, 0x22, 0x31, 0x52, 0xb3, 0x58, 0x0b, 0x35, 0xa2, 0x03, 0x68, 0x0f, 0x75, 0x14, 0xde,
	0xcb, 0x80, 0xe1, 0x48, 0xb1, 0x29, 0xca, 0x1b, 0xf4, 0x39, 0x13, 0x3c, 0x0a, 0xe2, 0x80, 0x87,
	0x59, 0xb4, 0x0d, 0x95, 0x30, 0xf9, 0x9b, 0x79, 0xa4, 0x88, 0x8e, 0xe0, 0x74, 0x03, 0x0f, 0xbd,
	0x48, 0x0b, 0x6a, 0xe2, 0x7b, 0xe6, 0xb9, 0x60, 0x55, 0x15, 0x46, 0x47, 0xbb, 0xb0, 0x37, 0xe4,
	0x4c, 0xa8, 0x18, 0xaf, 0xe7, 0x28, 0xbd, 0x57, 0x5c, 0xb7, 0x80, 0x0b, 0xf6, 0x4f, 0x81, 0x4e,
	0x6b, 0xc2, 0x8e, 0x97, 0x52, 0x89, 0xa4, 0x34, 0xce, 0x20, 0x3d, 0x03, 0x72, 0x1b, 0x84, 0x2f,
	0x77, 0xde, 0x7b, 0xc0, 0xcc, 0x75, 0x57, 0x25, 0x74, 0x61, 0xb7, 0x30, 0x6d, 0xec, 0x59, 0x4a,
	0xe9, 0xf9, 0x0c, 0xba, 0x9f, 0x65, 0x68, 0x0c, 0x17, 0x1d, 0x4e, 0x50, 0xce, 0x03, 0x1f, 0xc9,
	0x15, 0x94, 0x27, 0x8a, 0x11, 0xbb, 0x93, 0x2b, 0xdc, 0x74, 0xeb, 0xec, 0x2f, 0xf1, 0x69, 0x0e,
	0xdd, 0x22, 0x1f, 0x16, 0xb4, 0xd6, 0x1e, 0x99, 0xf4, 0xf3, 0x06, 0x9b, 0xf6, 0xea, 0x5c, 0xfe,
	0x53, 0x95, 0x2d, 0xd5, 0xb3, 0xc8, 0x13, 0xd4, 0x8b, 0x97, 0x27, 0xad, 0x82, 0xd9, 0x6f, 0x35,
	0x3a, 0xf4, 0xaf, 0x91, 0xcc, 0xbc, 0x6d, 0x91, 0x07, 0xa8, 0xe6, 0x8e, 0x4e, 0x0e, 0xf3, 0xb2,
	0xe5, 0xee, 0x9c, 0xa3, 0x95, 0xff, 0x8d, 0x67, 0xcf, 0x1a, 0xd4, 0x1f, 0x6b, 0xf9, 0xb7, 0x37,
	0xad, 0x24, 0x2f, 0xee, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xff, 0x9a, 0x89, 0xe1, 0x92, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	CalculatePrimeNumberDecomposition(ctx context.Context, in *CalculatePrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_CalculatePrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculatePrimeNumberDecomposition(ctx context.Context, in *CalculatePrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_CalculatePrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/CalculatePrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculatePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_CalculatePrimeNumberDecompositionClient interface {
	Recv() (*CalculatePrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculatePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculatePrimeNumberDecompositionClient) Recv() (*CalculatePrimeNumberDecompositionResponse, error) {
	m := new(CalculatePrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	CalculatePrimeNumberDecomposition(*CalculatePrimeNumberDecompositionRequest, CalculatorService_CalculatePrimeNumberDecompositionServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) CalculatePrimeNumberDecomposition(req *CalculatePrimeNumberDecompositionRequest, srv CalculatorService_CalculatePrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculatePrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(srv CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculatePrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalculatePrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).CalculatePrimeNumberDecomposition(m, &calculatorServiceCalculatePrimeNumberDecompositionServer{stream})
}

type CalculatorService_CalculatePrimeNumberDecompositionServer interface {
	Send(*CalculatePrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServiceCalculatePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculatePrimeNumberDecompositionServer) Send(m *CalculatePrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalculatePrimeNumberDecomposition",
			Handler:       _CalculatorService_CalculatePrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculatorpb/calculator.proto",
}
