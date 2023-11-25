package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"net"
)

type GRPCConfig interface {
	GRPCAdress() string
}

type grpcConfig struct {
	Host string
	Port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	cfx := grpcConfig{}

	err := envconfig.Process("grpc", &cfx)
	if err != nil {
		return &grpcConfig{}, errors.New("check suffix 'GRPC' in .env file")
	}

	if len(cfx.Port) == 0 || len(cfx.Host) == 0 {
		return &grpcConfig{}, errors.New("check GrpcHost and GrpcPort in .env file")
	}

	return &grpcConfig{
		Host: cfx.Host,
		Port: cfx.Port,
	}, err
}

func (cfx grpcConfig) GRPCAdress() string {
	return net.JoinHostPort(cfx.Host, cfx.Port)
}
