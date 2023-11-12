package auth

import (
	"context"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &empty.Empty{}, i.authService.Delete(ctx, req.GetId())
}
