package config

import "langgo/config/plugins"

// Configuration 配置文件中所有字段对应的结构体
type Configuration struct {
	App      *App                `mapstructure:"app" json:"app" yaml:"app"`
	Log      *Log                `mapstructure:"log" json:"log" yaml:"log"`
	Database []*plugins.Database `mapstructure:"database" json:"database" yaml:"database"`
	Redis    *plugins.Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Minio    *plugins.Minio      `mapstructure:"minio" json:"minio" yaml:"minio"`
	ES       *plugins.ES         `mapstructure:"es" json:"es" yaml:"es"`
}

func (conf *Configuration) IsEsEnable() bool {
	if conf.ES != nil && conf.ES.Enable {
		return true
	}
	return false
}

func (conf *Configuration) IsRedisEnable() bool {
	if conf.Redis != nil && conf.Redis.Enable {
		return true
	}
	return false
}

func (conf *Configuration) IsMinioEnable() bool {
	if conf.Minio != nil && conf.Minio.Enable {
		return true
	}
	return false
}
