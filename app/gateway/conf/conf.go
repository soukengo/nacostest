package conf

import (
	"flag"
	"nacostest/api"
	"nacostest/pkg/component/config"
	"nacostest/pkg/component/discovery"
	"time"
)

var (
	configPath = ""
)

func init() {
	flag.StringVar(&configPath, "conf", "configs/gateway.yaml", "config path, eg: -conf gateway.yaml")
}

func Load() (conf *Config, err error) {
	conf = &Config{
		Env: &Env{AppId: api.AppUser},
	}
	source := config.NewFileSource(configPath)
	defer source.Close()
	loader := config.NewLoader(source)
	err = loader.Load(conf)
	return
}

type Config struct {
	Env       *Env
	Server    *Server
	Discovery *discovery.Config
}

type Env struct {
	AppId   string
	Env     string
	Version string
}

type Server struct {
	Http *Http
}

type Http struct {
	Network string
	Addr    string
	Timeout time.Duration
}
