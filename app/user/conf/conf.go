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
	flag.StringVar(&configPath, "conf", "configs/user.yaml", "config path, eg: -conf user.yaml")
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
	Grpc *Grpc
}

type Grpc struct {
	Network string
	Addr    string
	Timeout time.Duration
}
