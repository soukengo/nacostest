package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"nacostest/api"
	v1 "nacostest/api/user/v1"
	"nacostest/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewUserClient, NewAuthClient)

func NewUserClient(r registry.Discovery) (v1.UserClient, error) {
	conn, err := discovery.NewGrpcClient(api.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewUserClient(conn), nil
}
func NewAuthClient(r registry.Discovery) (v1.AuthClient, error) {
	conn, err := discovery.NewGrpcClient(api.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthClient(conn), nil
}
