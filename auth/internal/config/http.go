package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"net"
)

type HTTPConfig interface {
	HTTPAddress() string
}

type httpConfig struct {
	Host string
	Port string
}

func NewHTTPConfig() (*httpConfig, error) {
	cfg := httpConfig{}

	err := envconfig.Process("http", &cfg)
	if err != nil {
		return nil, errors.New("check suffix 'HTTP' in .env file")
	}

	if len(cfg.Host) == 0 || len(cfg.Port) == 0 {
		return nil, errors.New("check HttpHhost and HttpPort in .env file")
	}

	return &cfg, nil
}

func (cfg httpConfig) HTTPAddress() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
