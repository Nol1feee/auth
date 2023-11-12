package service

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, req *model.UserInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, req *model.UserUpdate) error
}
