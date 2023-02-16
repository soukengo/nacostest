// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"nacostest/app/user/conf"
	"nacostest/app/user/server"
	"nacostest/app/user/service"
	"nacostest/pkg/component/discovery"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(env *conf.Env, config *discovery.Config, confServer *conf.Server, logger log.Logger) (*kratos.App, error) {
	userService := service.NewUserService()
	authService := service.NewAuthService()
	grpcServer := server.NewGRPCServer(confServer, userService, authService, logger)
	registrar, err := discovery.NewRegistrar(config)
	if err != nil {
		return nil, err
	}
	app := newApp(env, grpcServer, registrar)
	return app, nil
}