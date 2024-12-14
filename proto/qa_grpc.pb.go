// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: proto/qa.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	QAService_AskQuestion_FullMethodName = "/qaservice.QAService/AskQuestion"
)

// QAServiceClient is the client API for QAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QAServiceClient interface {
	AskQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error)
}

type qAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQAServiceClient(cc grpc.ClientConnInterface) QAServiceClient {
	return &qAServiceClient{cc}
}

func (c *qAServiceClient) AskQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, QAService_AskQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QAServiceServer is the server API for QAService service.
// All implementations must embed UnimplementedQAServiceServer
// for forward compatibility.
type QAServiceServer interface {
	AskQuestion(context.Context, *QuestionRequest) (*QuestionResponse, error)
	mustEmbedUnimplementedQAServiceServer()
}

// UnimplementedQAServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQAServiceServer struct{}

func (UnimplementedQAServiceServer) AskQuestion(context.Context, *QuestionRequest) (*QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskQuestion not implemented")
}
func (UnimplementedQAServiceServer) mustEmbedUnimplementedQAServiceServer() {}
func (UnimplementedQAServiceServer) testEmbeddedByValue()                   {}

// UnsafeQAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QAServiceServer will
// result in compilation errors.
type UnsafeQAServiceServer interface {
	mustEmbedUnimplementedQAServiceServer()
}

func RegisterQAServiceServer(s grpc.ServiceRegistrar, srv QAServiceServer) {
	// If the following call pancis, it indicates UnimplementedQAServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QAService_ServiceDesc, srv)
}

func _QAService_AskQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QAServiceServer).AskQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QAService_AskQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QAServiceServer).AskQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QAService_ServiceDesc is the grpc.ServiceDesc for QAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qaservice.QAService",
	HandlerType: (*QAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskQuestion",
			Handler:    _QAService_AskQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/qa.proto",
}