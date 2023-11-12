package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/convertor"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	modelUser, err := i.authService.Get(ctx, req.GetId())
	return convertor.ToDescFromUser(modelUser), err
}
