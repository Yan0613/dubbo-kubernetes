// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/search.proto

package registryv1alpha1

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

const (
	SearchService_SearchUser_FullMethodName                = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchUser"
	SearchService_SearchRepository_FullMethodName          = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchRepository"
	SearchService_SearchLastCommitByContent_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchLastCommitByContent"
	SearchService_SearchCurationPlugin_FullMethodName      = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchCurationPlugin"
	SearchService_SearchTag_FullMethodName                 = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchTag"
	SearchService_SearchDraft_FullMethodName               = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchDraft"
)

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	// SearchUser searches users by username
	SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
	// SearchRepository searches repositories by name or description
	SearchRepository(ctx context.Context, in *SearchRepositoryRequest, opts ...grpc.CallOption) (*SearchRepositoryResponse, error)
	// SearchCommitByContent searches last commit in same repo by idl content
	// that means, for a repo, search results only record last matched commit
	SearchLastCommitByContent(ctx context.Context, in *SearchLastCommitByContentRequest, opts ...grpc.CallOption) (*SearchLastCommitByContentResponse, error)
	// SearchCurationPlugin search plugins by name or description
	SearchCurationPlugin(ctx context.Context, in *SearchCuratedPluginRequest, opts ...grpc.CallOption) (*SearchCuratedPluginResponse, error)
	// SearchTag searches for tags in a repository
	SearchTag(ctx context.Context, in *SearchTagRequest, opts ...grpc.CallOption) (*SearchTagResponse, error)
	// SearchDraft searches for drafts in a repository
	SearchDraft(ctx context.Context, in *SearchDraftRequest, opts ...grpc.CallOption) (*SearchDraftResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchRepository(ctx context.Context, in *SearchRepositoryRequest, opts ...grpc.CallOption) (*SearchRepositoryResponse, error) {
	out := new(SearchRepositoryResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchRepository_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchLastCommitByContent(ctx context.Context, in *SearchLastCommitByContentRequest, opts ...grpc.CallOption) (*SearchLastCommitByContentResponse, error) {
	out := new(SearchLastCommitByContentResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchLastCommitByContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchCurationPlugin(ctx context.Context, in *SearchCuratedPluginRequest, opts ...grpc.CallOption) (*SearchCuratedPluginResponse, error) {
	out := new(SearchCuratedPluginResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchCurationPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchTag(ctx context.Context, in *SearchTagRequest, opts ...grpc.CallOption) (*SearchTagResponse, error) {
	out := new(SearchTagResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchDraft(ctx context.Context, in *SearchDraftRequest, opts ...grpc.CallOption) (*SearchDraftResponse, error) {
	out := new(SearchDraftResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchDraft_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	// SearchUser searches users by username
	SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
	// SearchRepository searches repositories by name or description
	SearchRepository(context.Context, *SearchRepositoryRequest) (*SearchRepositoryResponse, error)
	// SearchCommitByContent searches last commit in same repo by idl content
	// that means, for a repo, search results only record last matched commit
	SearchLastCommitByContent(context.Context, *SearchLastCommitByContentRequest) (*SearchLastCommitByContentResponse, error)
	// SearchCurationPlugin search plugins by name or description
	SearchCurationPlugin(context.Context, *SearchCuratedPluginRequest) (*SearchCuratedPluginResponse, error)
	// SearchTag searches for tags in a repository
	SearchTag(context.Context, *SearchTagRequest) (*SearchTagResponse, error)
	// SearchDraft searches for drafts in a repository
	SearchDraft(context.Context, *SearchDraftRequest) (*SearchDraftResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedSearchServiceServer) SearchRepository(context.Context, *SearchRepositoryRequest) (*SearchRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRepository not implemented")
}
func (UnimplementedSearchServiceServer) SearchLastCommitByContent(context.Context, *SearchLastCommitByContentRequest) (*SearchLastCommitByContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLastCommitByContent not implemented")
}
func (UnimplementedSearchServiceServer) SearchCurationPlugin(context.Context, *SearchCuratedPluginRequest) (*SearchCuratedPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCurationPlugin not implemented")
}
func (UnimplementedSearchServiceServer) SearchTag(context.Context, *SearchTagRequest) (*SearchTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTag not implemented")
}
func (UnimplementedSearchServiceServer) SearchDraft(context.Context, *SearchDraftRequest) (*SearchDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDraft not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchUser(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchRepository(ctx, req.(*SearchRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchLastCommitByContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLastCommitByContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchLastCommitByContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchLastCommitByContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchLastCommitByContent(ctx, req.(*SearchLastCommitByContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchCurationPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCuratedPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchCurationPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchCurationPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchCurationPlugin(ctx, req.(*SearchCuratedPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchTag(ctx, req.(*SearchTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchDraft(ctx, req.(*SearchDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUser",
			Handler:    _SearchService_SearchUser_Handler,
		},
		{
			MethodName: "SearchRepository",
			Handler:    _SearchService_SearchRepository_Handler,
		},
		{
			MethodName: "SearchLastCommitByContent",
			Handler:    _SearchService_SearchLastCommitByContent_Handler,
		},
		{
			MethodName: "SearchCurationPlugin",
			Handler:    _SearchService_SearchCurationPlugin_Handler,
		},
		{
			MethodName: "SearchTag",
			Handler:    _SearchService_SearchTag_Handler,
		},
		{
			MethodName: "SearchDraft",
			Handler:    _SearchService_SearchDraft_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/search.proto",
}
