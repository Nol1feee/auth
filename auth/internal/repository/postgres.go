package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

const dbDSN = "host=127.0.0.1 port=5432 dbname=postgres user=postgres password=superpassword sslmode=disable"

func NewPostgresDB(cfg Config, ctx context.Context) (*pgxpool.Pool, error) {
	dbDSN := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.Port,
		cfg.Name, cfg.User, cfg.Password, cfg.SSLMode)

	con, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		return &pgxpool.Pool{}, errors.New(fmt.Sprintf("repository - postgres - connect to DB | %s", err))
	}

	logrus.Info("DB is up")

	return con, nil
}
