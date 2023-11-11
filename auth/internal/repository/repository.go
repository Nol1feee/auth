package repository

import (
	"context"
	model "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
)

type AuthRepository interface {
	//CreateUser(ctx context.Context, req *desc.CreateRequest, pool pgxpool.Pool) (*desc.CreateResponse, error)

	CreateUser(ctx context.Context, req *model.User) (int64, error)
	//GetUser(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
	//DeleteUser(ctx context.Context, req *desc.DeleteRequest) error
	//UpdateInfoUser(ctx context.Context, req *desc.UpdateRequest) error
}
