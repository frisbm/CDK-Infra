package config

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/spf13/viper"
)

type Config struct {
	viper *viper.Viper
}

func NewConfig(filepath string) *Config {
	v := viper.New()
	v.SetConfigFile(filepath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	return &Config{
		viper: v,
	}
}

func (c *Config) Get(key string) *string {
	return jsii.String(c.viper.GetString(key))
}
