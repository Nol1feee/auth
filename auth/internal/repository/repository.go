package repository

import (
	"context"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	UserCRUD interface {
		CreateUser(ctx context.Context, req *desc.CreateRequest, pool pgxpool.Pool) (*desc.CreateResponse, error)
		GetUser(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
		DeleteUser(ctx context.Context, req *desc.DeleteRequest) error
		UpdateInfoUser(ctx context.Context, req *desc.UpdateRequest) error
		GetId(ctx context.Context, email string) (int64, error)
	}
)
