package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type pgConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

type PGConfig interface {
	DSN() string
}

func NewPGConfig() (*pgConfig, error) {
	pgCfx := pgConfig{}
	err := envconfig.Process("db", &pgCfx)
	return &pgCfx, err
}

func (cfg pgConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.Port,
		cfg.Name, cfg.User, cfg.Password, cfg.SSLMode)
}
