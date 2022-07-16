package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	//DBURL    string `envconfig:"DBURL" default:""`
	HostName string `envconfig:"HOST_NAME" defaul:""`
	Port     string `envconfig:"PORT" defaul:""`
	Protocol string `envconfig:"PROTOCOL" defaul:""`
}

const prefix string = ""

func NewConfig() *Config {
	conf := new(Config)
	envconfig.MustProcess(prefix, conf)
	return conf
}
