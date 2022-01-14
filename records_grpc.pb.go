// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package record

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordServiceClient interface {
	//Simple RPC
	GetRecord(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error)
	//Server side streaming RPC
	ListRecords(ctx context.Context, in *User, opts ...grpc.CallOption) (RecordService_ListRecordsClient, error)
	//Client side streaming RPC
	SetRecords(ctx context.Context, opts ...grpc.CallOption) (RecordService_SetRecordsClient, error)
	//Bi directional side streaming RPC
	RecordPong(ctx context.Context, opts ...grpc.CallOption) (RecordService_RecordPongClient, error)
}

type recordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceClient(cc grpc.ClientConnInterface) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) GetRecord(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/record.RecordService/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) ListRecords(ctx context.Context, in *User, opts ...grpc.CallOption) (RecordService_ListRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RecordService_ServiceDesc.Streams[0], "/record.RecordService/ListRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordServiceListRecordsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordService_ListRecordsClient interface {
	Recv() (*RecordResponse, error)
	grpc.ClientStream
}

type recordServiceListRecordsClient struct {
	grpc.ClientStream
}

func (x *recordServiceListRecordsClient) Recv() (*RecordResponse, error) {
	m := new(RecordResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordServiceClient) SetRecords(ctx context.Context, opts ...grpc.CallOption) (RecordService_SetRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RecordService_ServiceDesc.Streams[1], "/record.RecordService/SetRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordServiceSetRecordsClient{stream}
	return x, nil
}

type RecordService_SetRecordsClient interface {
	Send(*RecordRequest) error
	CloseAndRecv() (*Error, error)
	grpc.ClientStream
}

type recordServiceSetRecordsClient struct {
	grpc.ClientStream
}

func (x *recordServiceSetRecordsClient) Send(m *RecordRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recordServiceSetRecordsClient) CloseAndRecv() (*Error, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Error)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordServiceClient) RecordPong(ctx context.Context, opts ...grpc.CallOption) (RecordService_RecordPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &RecordService_ServiceDesc.Streams[2], "/record.RecordService/RecordPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordServiceRecordPongClient{stream}
	return x, nil
}

type RecordService_RecordPongClient interface {
	Send(*RecordRequest) error
	Recv() (*RecordResponse, error)
	grpc.ClientStream
}

type recordServiceRecordPongClient struct {
	grpc.ClientStream
}

func (x *recordServiceRecordPongClient) Send(m *RecordRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recordServiceRecordPongClient) Recv() (*RecordResponse, error) {
	m := new(RecordResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordServiceServer is the server API for RecordService service.
// All implementations must embed UnimplementedRecordServiceServer
// for forward compatibility
type RecordServiceServer interface {
	//Simple RPC
	GetRecord(context.Context, *RecordRequest) (*RecordResponse, error)
	//Server side streaming RPC
	ListRecords(*User, RecordService_ListRecordsServer) error
	//Client side streaming RPC
	SetRecords(RecordService_SetRecordsServer) error
	//Bi directional side streaming RPC
	RecordPong(RecordService_RecordPongServer) error
	mustEmbedUnimplementedRecordServiceServer()
}

// UnimplementedRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (UnimplementedRecordServiceServer) GetRecord(context.Context, *RecordRequest) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedRecordServiceServer) ListRecords(*User, RecordService_ListRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListRecords not implemented")
}
func (UnimplementedRecordServiceServer) SetRecords(RecordService_SetRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method SetRecords not implemented")
}
func (UnimplementedRecordServiceServer) RecordPong(RecordService_RecordPongServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordPong not implemented")
}
func (UnimplementedRecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {}

// UnsafeRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServiceServer will
// result in compilation errors.
type UnsafeRecordServiceServer interface {
	mustEmbedUnimplementedRecordServiceServer()
}

func RegisterRecordServiceServer(s grpc.ServiceRegistrar, srv RecordServiceServer) {
	s.RegisterService(&RecordService_ServiceDesc, srv)
}

func _RecordService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record.RecordService/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).GetRecord(ctx, req.(*RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_ListRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordServiceServer).ListRecords(m, &recordServiceListRecordsServer{stream})
}

type RecordService_ListRecordsServer interface {
	Send(*RecordResponse) error
	grpc.ServerStream
}

type recordServiceListRecordsServer struct {
	grpc.ServerStream
}

func (x *recordServiceListRecordsServer) Send(m *RecordResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RecordService_SetRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecordServiceServer).SetRecords(&recordServiceSetRecordsServer{stream})
}

type RecordService_SetRecordsServer interface {
	SendAndClose(*Error) error
	Recv() (*RecordRequest, error)
	grpc.ServerStream
}

type recordServiceSetRecordsServer struct {
	grpc.ServerStream
}

func (x *recordServiceSetRecordsServer) SendAndClose(m *Error) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recordServiceSetRecordsServer) Recv() (*RecordRequest, error) {
	m := new(RecordRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RecordService_RecordPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecordServiceServer).RecordPong(&recordServiceRecordPongServer{stream})
}

type RecordService_RecordPongServer interface {
	Send(*RecordResponse) error
	Recv() (*RecordRequest, error)
	grpc.ServerStream
}

type recordServiceRecordPongServer struct {
	grpc.ServerStream
}

func (x *recordServiceRecordPongServer) Send(m *RecordResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recordServiceRecordPongServer) Recv() (*RecordRequest, error) {
	m := new(RecordRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordService_ServiceDesc is the grpc.ServiceDesc for RecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "record.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecord",
			Handler:    _RecordService_GetRecord_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListRecords",
			Handler:       _RecordService_ListRecords_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SetRecords",
			Handler:       _RecordService_SetRecords_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RecordPong",
			Handler:       _RecordService_RecordPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "records.proto",
}
