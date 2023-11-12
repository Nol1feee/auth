package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/convertor"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, i.authService.Update(ctx, convertor.ToUpdateFromDesc(req))
}
