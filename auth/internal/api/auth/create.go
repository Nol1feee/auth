package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/service/auth/convertor"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.authService.CreateUser(ctx, convertor.ToServiceFromDesc(req))
	if err != nil {
		return &desc.CreateResponse{}, err
	}

	return &desc.CreateResponse{Id: id}, err
}
