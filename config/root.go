package config

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type RootConfig struct {
	Port int `yaml:"port"`
}

// GetConfig 获取配置
func GetConfig() *RootConfig {
	d := &RootConfig{}
	data := g.Cfg(configFile).MustData(context.Background())
	if err := gconv.Struct(data, &d); err != nil {
		return d
	}
	return d
}
