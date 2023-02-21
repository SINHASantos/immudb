// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package schemav2

import (
	context "context"
	schema "github.com/codenotary/immudb/pkg/api/schema"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImmuServiceV2Client is the client API for ImmuServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImmuServiceV2Client interface {
	LoginV2(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponseV2, error)
	DocumentInsert(ctx context.Context, in *DocumentInsertRequest, opts ...grpc.CallOption) (*schema.VerifiableTx, error)
	DocumentSearch(ctx context.Context, in *DocumentSearchRequest, opts ...grpc.CallOption) (*DocumentSearchResponse, error)
	CollectionCreate(ctx context.Context, in *CollectionCreateRequest, opts ...grpc.CallOption) (*CollectionInformation, error)
	CollectionGet(ctx context.Context, in *CollectionGetRequest, opts ...grpc.CallOption) (*CollectionInformation, error)
	CollectionList(ctx context.Context, in *CollectionListRequest, opts ...grpc.CallOption) (*CollectionListResponse, error)
	CollectionDelete(ctx context.Context, in *CollectionDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type immuServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewImmuServiceV2Client(cc grpc.ClientConnInterface) ImmuServiceV2Client {
	return &immuServiceV2Client{cc}
}

func (c *immuServiceV2Client) LoginV2(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponseV2, error) {
	out := new(LoginResponseV2)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/LoginV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceV2Client) DocumentInsert(ctx context.Context, in *DocumentInsertRequest, opts ...grpc.CallOption) (*schema.VerifiableTx, error) {
	out := new(schema.VerifiableTx)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/DocumentInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceV2Client) DocumentSearch(ctx context.Context, in *DocumentSearchRequest, opts ...grpc.CallOption) (*DocumentSearchResponse, error) {
	out := new(DocumentSearchResponse)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/DocumentSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceV2Client) CollectionCreate(ctx context.Context, in *CollectionCreateRequest, opts ...grpc.CallOption) (*CollectionInformation, error) {
	out := new(CollectionInformation)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/CollectionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceV2Client) CollectionGet(ctx context.Context, in *CollectionGetRequest, opts ...grpc.CallOption) (*CollectionInformation, error) {
	out := new(CollectionInformation)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/CollectionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceV2Client) CollectionList(ctx context.Context, in *CollectionListRequest, opts ...grpc.CallOption) (*CollectionListResponse, error) {
	out := new(CollectionListResponse)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/CollectionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceV2Client) CollectionDelete(ctx context.Context, in *CollectionDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuServiceV2/CollectionDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmuServiceV2Server is the server API for ImmuServiceV2 service.
// All implementations should embed UnimplementedImmuServiceV2Server
// for forward compatibility
type ImmuServiceV2Server interface {
	LoginV2(context.Context, *LoginRequest) (*LoginResponseV2, error)
	DocumentInsert(context.Context, *DocumentInsertRequest) (*schema.VerifiableTx, error)
	DocumentSearch(context.Context, *DocumentSearchRequest) (*DocumentSearchResponse, error)
	CollectionCreate(context.Context, *CollectionCreateRequest) (*CollectionInformation, error)
	CollectionGet(context.Context, *CollectionGetRequest) (*CollectionInformation, error)
	CollectionList(context.Context, *CollectionListRequest) (*CollectionListResponse, error)
	CollectionDelete(context.Context, *CollectionDeleteRequest) (*empty.Empty, error)
}

// UnimplementedImmuServiceV2Server should be embedded to have forward compatible implementations.
type UnimplementedImmuServiceV2Server struct {
}

func (UnimplementedImmuServiceV2Server) LoginV2(context.Context, *LoginRequest) (*LoginResponseV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginV2 not implemented")
}
func (UnimplementedImmuServiceV2Server) DocumentInsert(context.Context, *DocumentInsertRequest) (*schema.VerifiableTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentInsert not implemented")
}
func (UnimplementedImmuServiceV2Server) DocumentSearch(context.Context, *DocumentSearchRequest) (*DocumentSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentSearch not implemented")
}
func (UnimplementedImmuServiceV2Server) CollectionCreate(context.Context, *CollectionCreateRequest) (*CollectionInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionCreate not implemented")
}
func (UnimplementedImmuServiceV2Server) CollectionGet(context.Context, *CollectionGetRequest) (*CollectionInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionGet not implemented")
}
func (UnimplementedImmuServiceV2Server) CollectionList(context.Context, *CollectionListRequest) (*CollectionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionList not implemented")
}
func (UnimplementedImmuServiceV2Server) CollectionDelete(context.Context, *CollectionDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionDelete not implemented")
}

// UnsafeImmuServiceV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImmuServiceV2Server will
// result in compilation errors.
type UnsafeImmuServiceV2Server interface {
	mustEmbedUnimplementedImmuServiceV2Server()
}

func RegisterImmuServiceV2Server(s grpc.ServiceRegistrar, srv ImmuServiceV2Server) {
	s.RegisterService(&ImmuServiceV2_ServiceDesc, srv)
}

func _ImmuServiceV2_LoginV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).LoginV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/LoginV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).LoginV2(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuServiceV2_DocumentInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentInsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).DocumentInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/DocumentInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).DocumentInsert(ctx, req.(*DocumentInsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuServiceV2_DocumentSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).DocumentSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/DocumentSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).DocumentSearch(ctx, req.(*DocumentSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuServiceV2_CollectionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).CollectionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/CollectionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).CollectionCreate(ctx, req.(*CollectionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuServiceV2_CollectionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).CollectionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/CollectionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).CollectionGet(ctx, req.(*CollectionGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuServiceV2_CollectionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).CollectionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/CollectionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).CollectionList(ctx, req.(*CollectionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuServiceV2_CollectionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceV2Server).CollectionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuServiceV2/CollectionDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceV2Server).CollectionDelete(ctx, req.(*CollectionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImmuServiceV2_ServiceDesc is the grpc.ServiceDesc for ImmuServiceV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImmuServiceV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "immudb.schemav2.ImmuServiceV2",
	HandlerType: (*ImmuServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginV2",
			Handler:    _ImmuServiceV2_LoginV2_Handler,
		},
		{
			MethodName: "DocumentInsert",
			Handler:    _ImmuServiceV2_DocumentInsert_Handler,
		},
		{
			MethodName: "DocumentSearch",
			Handler:    _ImmuServiceV2_DocumentSearch_Handler,
		},
		{
			MethodName: "CollectionCreate",
			Handler:    _ImmuServiceV2_CollectionCreate_Handler,
		},
		{
			MethodName: "CollectionGet",
			Handler:    _ImmuServiceV2_CollectionGet_Handler,
		},
		{
			MethodName: "CollectionList",
			Handler:    _ImmuServiceV2_CollectionList_Handler,
		},
		{
			MethodName: "CollectionDelete",
			Handler:    _ImmuServiceV2_CollectionDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schemav2.proto",
}