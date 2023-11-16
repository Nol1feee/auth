package app

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/closer"

	//"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/api/auth"
	"github.com/Nol1feee/CLI-chat/auth/internal/config"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository"
	authRepo "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth"
	"github.com/Nol1feee/CLI-chat/auth/internal/service"
	authService "github.com/Nol1feee/CLI-chat/auth/internal/service/auth"
	"github.com/jackc/pgx/v4/pgxpool"
	//"google.golang.org/genproto/googleapis/appengine/v1"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pool *pgxpool.Pool

	authService    service.AuthService
	authRepository repository.AuthRepository

	authImplementation *auth.Implementation

	closer *closer.Closer
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {
	pgCfg, err := config.NewPGConfig()
	if err != nil {
		//fatal, т.к. смысл как-то обрабатывать ошибку/не фаталить, если мы не смогли запустить приложение
		log.Fatal(err)
	}
	return pgCfg
}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	cfgGRPC, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}

	return cfgGRPC
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pool == nil {
		con, err := pgxpool.Connect(ctx, s.GetPGConfig().DSN())
		if err != nil {
			//fatal, т.к. сервис еще не запущен
			log.Fatalf("connect to DB error | %s ", err)
		}

		err = con.Ping(ctx)
		if err != nil {
			log.Fatalf("ping DB error | %s ", err)
		}

		s.closer.Add(func() error {
			con.Close()
			return nil
		})

		s.pool = con
	}

	return s.pool
}

func (s *serviceProvider) AuthRepo(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepo.NewRepo(s.PgPool(ctx))
	}
	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepo(ctx))
	}
	return s.authService
}

func (s *serviceProvider) NewImplementation(ctx context.Context) *auth.Implementation {
	if s.authImplementation == nil {
		s.authImplementation = auth.NewImplementation(s.AuthService(ctx))
	}
	return s.authImplementation
}
