//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"nacostest/app/user/conf"
	"nacostest/app/user/server"
	"nacostest/app/user/service"
	"nacostest/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
