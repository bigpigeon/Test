// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package info

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

type UploadStream struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadStream) Reset()         { *m = UploadStream{} }
func (m *UploadStream) String() string { return proto.CompactTextString(m) }
func (*UploadStream) ProtoMessage()    {}
func (*UploadStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *UploadStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStream.Unmarshal(m, b)
}
func (m *UploadStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStream.Marshal(b, m, deterministic)
}
func (m *UploadStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStream.Merge(m, src)
}
func (m *UploadStream) XXX_Size() int {
	return xxx_messageInfo_UploadStream.Size(m)
}
func (m *UploadStream) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStream.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStream proto.InternalMessageInfo

func (m *UploadStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DownloadRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DownloadStream struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadStream) Reset()         { *m = DownloadStream{} }
func (m *DownloadStream) String() string { return proto.CompactTextString(m) }
func (*DownloadStream) ProtoMessage()    {}
func (*DownloadStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}

func (m *DownloadStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadStream.Unmarshal(m, b)
}
func (m *DownloadStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadStream.Marshal(b, m, deterministic)
}
func (m *DownloadStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadStream.Merge(m, src)
}
func (m *DownloadStream) XXX_Size() int {
	return xxx_messageInfo_DownloadStream.Size(m)
}
func (m *DownloadStream) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadStream.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadStream proto.InternalMessageInfo

func (m *DownloadStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadStream)(nil), "info.UploadStream")
	proto.RegisterType((*UploadResponse)(nil), "info.UploadResponse")
	proto.RegisterType((*DownloadRequest)(nil), "info.DownloadRequest")
	proto.RegisterType((*DownloadStream)(nil), "info.DownloadStream")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0xcc, 0x4b, 0xcb, 0x57,
	0x52, 0xe2, 0xe2, 0x09, 0x2d, 0xc8, 0xc9, 0x4f, 0x4c, 0x09, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x15,
	0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3,
	0x95, 0x94, 0xb8, 0xf8, 0x20, 0x6a, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04,
	0xb8, 0x98, 0xb3, 0x53, 0x2b, 0xc1, 0x8a, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x65, 0x2e, 0x7e, 0x97,
	0xfc, 0xf2, 0x3c, 0x88, 0xaa, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x2c, 0x8a, 0x54, 0xb8, 0xf8, 0x60,
	0x8a, 0x70, 0x5b, 0x67, 0xd4, 0xc4, 0xc8, 0xc5, 0xed, 0x96, 0x99, 0x93, 0x1a, 0x5c, 0x92, 0x5f,
	0x94, 0x98, 0x9e, 0x2a, 0x64, 0xc6, 0xc5, 0x06, 0xb1, 0x5e, 0x48, 0x48, 0x0f, 0xe4, 0x66, 0x3d,
	0x64, 0x07, 0x4b, 0x89, 0x20, 0x8b, 0xc1, 0x1c, 0xa8, 0xc4, 0xa0, 0xc1, 0x28, 0x64, 0xcd, 0xc5,
	0x01, 0xb3, 0x4d, 0x48, 0x14, 0xa2, 0x0a, 0xcd, 0x89, 0x30, 0xcd, 0xa8, 0x8e, 0x52, 0x62, 0x30,
	0x60, 0x4c, 0x62, 0x03, 0x07, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x78, 0xe1, 0xa0,
	0x35, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileStorageClient is the client API for FileStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileStorageClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileStorage_UploadClient, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileStorage_DownloadClient, error)
}

type fileStorageClient struct {
	cc *grpc.ClientConn
}

func NewFileStorageClient(cc *grpc.ClientConn) FileStorageClient {
	return &fileStorageClient{cc}
}

func (c *fileStorageClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileStorage_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileStorage_serviceDesc.Streams[0], "/info.FileStorage/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStorageUploadClient{stream}
	return x, nil
}

type FileStorage_UploadClient interface {
	Send(*UploadStream) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type fileStorageUploadClient struct {
	grpc.ClientStream
}

func (x *fileStorageUploadClient) Send(m *UploadStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileStorageUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileStorageClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileStorage_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileStorage_serviceDesc.Streams[1], "/info.FileStorage/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStorageDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileStorage_DownloadClient interface {
	Recv() (*DownloadStream, error)
	grpc.ClientStream
}

type fileStorageDownloadClient struct {
	grpc.ClientStream
}

func (x *fileStorageDownloadClient) Recv() (*DownloadStream, error) {
	m := new(DownloadStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStorageServer is the server API for FileStorage service.
type FileStorageServer interface {
	Upload(FileStorage_UploadServer) error
	Download(*DownloadRequest, FileStorage_DownloadServer) error
}

// UnimplementedFileStorageServer can be embedded to have forward compatible implementations.
type UnimplementedFileStorageServer struct {
}

func (*UnimplementedFileStorageServer) Upload(srv FileStorage_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedFileStorageServer) Download(req *DownloadRequest, srv FileStorage_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}

func RegisterFileStorageServer(s *grpc.Server, srv FileStorageServer) {
	s.RegisterService(&_FileStorage_serviceDesc, srv)
}

func _FileStorage_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileStorageServer).Upload(&fileStorageUploadServer{stream})
}

type FileStorage_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadStream, error)
	grpc.ServerStream
}

type fileStorageUploadServer struct {
	grpc.ServerStream
}

func (x *fileStorageUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileStorageUploadServer) Recv() (*UploadStream, error) {
	m := new(UploadStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileStorage_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileStorageServer).Download(m, &fileStorageDownloadServer{stream})
}

type FileStorage_DownloadServer interface {
	Send(*DownloadStream) error
	grpc.ServerStream
}

type fileStorageDownloadServer struct {
	grpc.ServerStream
}

func (x *fileStorageDownloadServer) Send(m *DownloadStream) error {
	return x.ServerStream.SendMsg(m)
}

var _FileStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.FileStorage",
	HandlerType: (*FileStorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileStorage_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _FileStorage_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
