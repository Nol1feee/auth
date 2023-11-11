package service

import (
	"context"
	model "github.com/Nol1feee/CLI-chat/auth/internal/service/auth/model"
)

type AuthService interface {
	CreateUser(ctx context.Context, req *model.UserServ) (int64, error)
}
