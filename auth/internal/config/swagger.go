package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"net"
)

type swaggerConfig struct {
	Port string
	Host string
}

type SwaggerCfg interface {
	Address() string
}

func NewSwaggerConfig() (*swaggerConfig, error) {
	cfg := swaggerConfig{}

	err := envconfig.Process("swagger", &cfg)
	if err != nil {
		return nil, err
	}

	if len(cfg.Port) == 0 || len(cfg.Host) == 0 {
		return nil, errors.New("swagger host/port can't be nothing, check swagger prefix in .env file")
	}

	return &cfg, nil
}

func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
