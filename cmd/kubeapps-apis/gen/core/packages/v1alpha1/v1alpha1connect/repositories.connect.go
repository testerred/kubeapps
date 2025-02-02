// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: kubeappsapis/core/packages/v1alpha1/repositories.proto

package v1alpha1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "github.com/vmware-tanzu/kubeapps/cmd/kubeapps-apis/gen/core/packages/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// RepositoriesServiceName is the fully-qualified name of the RepositoriesService service.
	RepositoriesServiceName = "kubeappsapis.core.packages.v1alpha1.RepositoriesService"
)

// RepositoriesServiceClient is a client for the
// kubeappsapis.core.packages.v1alpha1.RepositoriesService service.
type RepositoriesServiceClient interface {
	AddPackageRepository(context.Context, *connect_go.Request[v1alpha1.AddPackageRepositoryRequest]) (*connect_go.Response[v1alpha1.AddPackageRepositoryResponse], error)
	GetPackageRepositoryDetail(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositoryDetailRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryDetailResponse], error)
	GetPackageRepositorySummaries(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositorySummariesRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositorySummariesResponse], error)
	UpdatePackageRepository(context.Context, *connect_go.Request[v1alpha1.UpdatePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.UpdatePackageRepositoryResponse], error)
	DeletePackageRepository(context.Context, *connect_go.Request[v1alpha1.DeletePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.DeletePackageRepositoryResponse], error)
	GetPackageRepositoryPermissions(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositoryPermissionsRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryPermissionsResponse], error)
}

// NewRepositoriesServiceClient constructs a client for the
// kubeappsapis.core.packages.v1alpha1.RepositoriesService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRepositoriesServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RepositoriesServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &repositoriesServiceClient{
		addPackageRepository: connect_go.NewClient[v1alpha1.AddPackageRepositoryRequest, v1alpha1.AddPackageRepositoryResponse](
			httpClient,
			baseURL+"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/AddPackageRepository",
			opts...,
		),
		getPackageRepositoryDetail: connect_go.NewClient[v1alpha1.GetPackageRepositoryDetailRequest, v1alpha1.GetPackageRepositoryDetailResponse](
			httpClient,
			baseURL+"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositoryDetail",
			opts...,
		),
		getPackageRepositorySummaries: connect_go.NewClient[v1alpha1.GetPackageRepositorySummariesRequest, v1alpha1.GetPackageRepositorySummariesResponse](
			httpClient,
			baseURL+"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositorySummaries",
			opts...,
		),
		updatePackageRepository: connect_go.NewClient[v1alpha1.UpdatePackageRepositoryRequest, v1alpha1.UpdatePackageRepositoryResponse](
			httpClient,
			baseURL+"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/UpdatePackageRepository",
			opts...,
		),
		deletePackageRepository: connect_go.NewClient[v1alpha1.DeletePackageRepositoryRequest, v1alpha1.DeletePackageRepositoryResponse](
			httpClient,
			baseURL+"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/DeletePackageRepository",
			opts...,
		),
		getPackageRepositoryPermissions: connect_go.NewClient[v1alpha1.GetPackageRepositoryPermissionsRequest, v1alpha1.GetPackageRepositoryPermissionsResponse](
			httpClient,
			baseURL+"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositoryPermissions",
			opts...,
		),
	}
}

// repositoriesServiceClient implements RepositoriesServiceClient.
type repositoriesServiceClient struct {
	addPackageRepository            *connect_go.Client[v1alpha1.AddPackageRepositoryRequest, v1alpha1.AddPackageRepositoryResponse]
	getPackageRepositoryDetail      *connect_go.Client[v1alpha1.GetPackageRepositoryDetailRequest, v1alpha1.GetPackageRepositoryDetailResponse]
	getPackageRepositorySummaries   *connect_go.Client[v1alpha1.GetPackageRepositorySummariesRequest, v1alpha1.GetPackageRepositorySummariesResponse]
	updatePackageRepository         *connect_go.Client[v1alpha1.UpdatePackageRepositoryRequest, v1alpha1.UpdatePackageRepositoryResponse]
	deletePackageRepository         *connect_go.Client[v1alpha1.DeletePackageRepositoryRequest, v1alpha1.DeletePackageRepositoryResponse]
	getPackageRepositoryPermissions *connect_go.Client[v1alpha1.GetPackageRepositoryPermissionsRequest, v1alpha1.GetPackageRepositoryPermissionsResponse]
}

// AddPackageRepository calls
// kubeappsapis.core.packages.v1alpha1.RepositoriesService.AddPackageRepository.
func (c *repositoriesServiceClient) AddPackageRepository(ctx context.Context, req *connect_go.Request[v1alpha1.AddPackageRepositoryRequest]) (*connect_go.Response[v1alpha1.AddPackageRepositoryResponse], error) {
	return c.addPackageRepository.CallUnary(ctx, req)
}

// GetPackageRepositoryDetail calls
// kubeappsapis.core.packages.v1alpha1.RepositoriesService.GetPackageRepositoryDetail.
func (c *repositoriesServiceClient) GetPackageRepositoryDetail(ctx context.Context, req *connect_go.Request[v1alpha1.GetPackageRepositoryDetailRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryDetailResponse], error) {
	return c.getPackageRepositoryDetail.CallUnary(ctx, req)
}

// GetPackageRepositorySummaries calls
// kubeappsapis.core.packages.v1alpha1.RepositoriesService.GetPackageRepositorySummaries.
func (c *repositoriesServiceClient) GetPackageRepositorySummaries(ctx context.Context, req *connect_go.Request[v1alpha1.GetPackageRepositorySummariesRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositorySummariesResponse], error) {
	return c.getPackageRepositorySummaries.CallUnary(ctx, req)
}

// UpdatePackageRepository calls
// kubeappsapis.core.packages.v1alpha1.RepositoriesService.UpdatePackageRepository.
func (c *repositoriesServiceClient) UpdatePackageRepository(ctx context.Context, req *connect_go.Request[v1alpha1.UpdatePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.UpdatePackageRepositoryResponse], error) {
	return c.updatePackageRepository.CallUnary(ctx, req)
}

// DeletePackageRepository calls
// kubeappsapis.core.packages.v1alpha1.RepositoriesService.DeletePackageRepository.
func (c *repositoriesServiceClient) DeletePackageRepository(ctx context.Context, req *connect_go.Request[v1alpha1.DeletePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.DeletePackageRepositoryResponse], error) {
	return c.deletePackageRepository.CallUnary(ctx, req)
}

// GetPackageRepositoryPermissions calls
// kubeappsapis.core.packages.v1alpha1.RepositoriesService.GetPackageRepositoryPermissions.
func (c *repositoriesServiceClient) GetPackageRepositoryPermissions(ctx context.Context, req *connect_go.Request[v1alpha1.GetPackageRepositoryPermissionsRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryPermissionsResponse], error) {
	return c.getPackageRepositoryPermissions.CallUnary(ctx, req)
}

// RepositoriesServiceHandler is an implementation of the
// kubeappsapis.core.packages.v1alpha1.RepositoriesService service.
type RepositoriesServiceHandler interface {
	AddPackageRepository(context.Context, *connect_go.Request[v1alpha1.AddPackageRepositoryRequest]) (*connect_go.Response[v1alpha1.AddPackageRepositoryResponse], error)
	GetPackageRepositoryDetail(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositoryDetailRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryDetailResponse], error)
	GetPackageRepositorySummaries(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositorySummariesRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositorySummariesResponse], error)
	UpdatePackageRepository(context.Context, *connect_go.Request[v1alpha1.UpdatePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.UpdatePackageRepositoryResponse], error)
	DeletePackageRepository(context.Context, *connect_go.Request[v1alpha1.DeletePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.DeletePackageRepositoryResponse], error)
	GetPackageRepositoryPermissions(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositoryPermissionsRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryPermissionsResponse], error)
}

// NewRepositoriesServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRepositoriesServiceHandler(svc RepositoriesServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/kubeappsapis.core.packages.v1alpha1.RepositoriesService/AddPackageRepository", connect_go.NewUnaryHandler(
		"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/AddPackageRepository",
		svc.AddPackageRepository,
		opts...,
	))
	mux.Handle("/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositoryDetail", connect_go.NewUnaryHandler(
		"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositoryDetail",
		svc.GetPackageRepositoryDetail,
		opts...,
	))
	mux.Handle("/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositorySummaries", connect_go.NewUnaryHandler(
		"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositorySummaries",
		svc.GetPackageRepositorySummaries,
		opts...,
	))
	mux.Handle("/kubeappsapis.core.packages.v1alpha1.RepositoriesService/UpdatePackageRepository", connect_go.NewUnaryHandler(
		"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/UpdatePackageRepository",
		svc.UpdatePackageRepository,
		opts...,
	))
	mux.Handle("/kubeappsapis.core.packages.v1alpha1.RepositoriesService/DeletePackageRepository", connect_go.NewUnaryHandler(
		"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/DeletePackageRepository",
		svc.DeletePackageRepository,
		opts...,
	))
	mux.Handle("/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositoryPermissions", connect_go.NewUnaryHandler(
		"/kubeappsapis.core.packages.v1alpha1.RepositoriesService/GetPackageRepositoryPermissions",
		svc.GetPackageRepositoryPermissions,
		opts...,
	))
	return "/kubeappsapis.core.packages.v1alpha1.RepositoriesService/", mux
}

// UnimplementedRepositoriesServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRepositoriesServiceHandler struct{}

func (UnimplementedRepositoriesServiceHandler) AddPackageRepository(context.Context, *connect_go.Request[v1alpha1.AddPackageRepositoryRequest]) (*connect_go.Response[v1alpha1.AddPackageRepositoryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("kubeappsapis.core.packages.v1alpha1.RepositoriesService.AddPackageRepository is not implemented"))
}

func (UnimplementedRepositoriesServiceHandler) GetPackageRepositoryDetail(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositoryDetailRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryDetailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("kubeappsapis.core.packages.v1alpha1.RepositoriesService.GetPackageRepositoryDetail is not implemented"))
}

func (UnimplementedRepositoriesServiceHandler) GetPackageRepositorySummaries(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositorySummariesRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositorySummariesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("kubeappsapis.core.packages.v1alpha1.RepositoriesService.GetPackageRepositorySummaries is not implemented"))
}

func (UnimplementedRepositoriesServiceHandler) UpdatePackageRepository(context.Context, *connect_go.Request[v1alpha1.UpdatePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.UpdatePackageRepositoryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("kubeappsapis.core.packages.v1alpha1.RepositoriesService.UpdatePackageRepository is not implemented"))
}

func (UnimplementedRepositoriesServiceHandler) DeletePackageRepository(context.Context, *connect_go.Request[v1alpha1.DeletePackageRepositoryRequest]) (*connect_go.Response[v1alpha1.DeletePackageRepositoryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("kubeappsapis.core.packages.v1alpha1.RepositoriesService.DeletePackageRepository is not implemented"))
}

func (UnimplementedRepositoriesServiceHandler) GetPackageRepositoryPermissions(context.Context, *connect_go.Request[v1alpha1.GetPackageRepositoryPermissionsRequest]) (*connect_go.Response[v1alpha1.GetPackageRepositoryPermissionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("kubeappsapis.core.packages.v1alpha1.RepositoriesService.GetPackageRepositoryPermissions is not implemented"))
}
