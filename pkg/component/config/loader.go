package config

import (
	"github.com/spf13/viper"
)

type Loader interface {
	Load(container interface{}) error
}

type viperLoader struct {
	vp     *viper.Viper
	source Source
}

func NewLoader(source Source) Loader {
	return &viperLoader{
		vp:     viper.New(),
		source: source,
	}
}

func (l *viperLoader) Load(v any) (err error) {
	if err = l.source.Error(); err != nil {
		return
	}
	l.vp.SetConfigType(l.source.ConfigType())
	err = l.vp.ReadConfig(l.source.From())
	if err != nil {
		return
	}
	return l.vp.Unmarshal(v)
}
