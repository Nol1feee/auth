package repository

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
)

type AuthRepository interface {
	Create(ctx context.Context, req *model.UserInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, req *model.UserUpdate) error
}
