// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueClient interface {
	// enqueue value
	Enqueue(ctx context.Context, in *EnqueueRequest, opts ...grpc.CallOption) (*EnqueueResponse, error)
	// dequeue value
	Dequeue(ctx context.Context, in *DequeueRequest, opts ...grpc.CallOption) (*DequeueResponse, error)
}

type queueClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueClient(cc grpc.ClientConnInterface) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) Enqueue(ctx context.Context, in *EnqueueRequest, opts ...grpc.CallOption) (*EnqueueResponse, error) {
	out := new(EnqueueResponse)
	err := c.cc.Invoke(ctx, "/Queue/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Dequeue(ctx context.Context, in *DequeueRequest, opts ...grpc.CallOption) (*DequeueResponse, error) {
	out := new(DequeueResponse)
	err := c.cc.Invoke(ctx, "/Queue/Dequeue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServer is the server API for Queue service.
// All implementations must embed UnimplementedQueueServer
// for forward compatibility
type QueueServer interface {
	// enqueue value
	Enqueue(context.Context, *EnqueueRequest) (*EnqueueResponse, error)
	// dequeue value
	Dequeue(context.Context, *DequeueRequest) (*DequeueResponse, error)
	mustEmbedUnimplementedQueueServer()
}

// UnimplementedQueueServer must be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (UnimplementedQueueServer) Enqueue(context.Context, *EnqueueRequest) (*EnqueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (UnimplementedQueueServer) Dequeue(context.Context, *DequeueRequest) (*DequeueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dequeue not implemented")
}
func (UnimplementedQueueServer) mustEmbedUnimplementedQueueServer() {}

// UnsafeQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServer will
// result in compilation errors.
type UnsafeQueueServer interface {
	mustEmbedUnimplementedQueueServer()
}

func RegisterQueueServer(s grpc.ServiceRegistrar, srv QueueServer) {
	s.RegisterService(&Queue_ServiceDesc, srv)
}

func _Queue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Queue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Enqueue(ctx, req.(*EnqueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Dequeue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DequeueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Dequeue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Queue/Dequeue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Dequeue(ctx, req.(*DequeueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Queue_ServiceDesc is the grpc.ServiceDesc for Queue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Queue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _Queue_Enqueue_Handler,
		},
		{
			MethodName: "Dequeue",
			Handler:    _Queue_Dequeue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/queue.proto",
}
