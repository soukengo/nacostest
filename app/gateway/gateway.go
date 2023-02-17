package gateway

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"nacostest"
	"nacostest/app/gateway/conf"
)

// New new a new gateway Application
func New(version string) (app *kratos.App, err error) {
	cfg, err := conf.Load()
	if err != nil {
		return
	}
	version = nacostest.SetVersion(version)
	cfg.Env.Version = version
	cfg.Discovery.AppId = cfg.Env.AppId
	app, err = wireApp(cfg.Env, cfg.Discovery, cfg.Server, log.DefaultLogger)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func newApp(env *conf.Env, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(env.AppId),
		kratos.Version(env.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
		),
	)
}
