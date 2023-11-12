package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/convertor"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
)

// todo добавить обработку кейса, когда email уже существует
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.authService.Create(ctx, convertor.ToUserInfoFromDesc(req.GetUserInfo()))
	return &desc.CreateResponse{Id: id}, err
}
