package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/convertor"
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
)

func (s *Serv) Update(ctx context.Context, req *model.UserUpdate) error {
	return s.AuthRepository.Update(ctx, convertor.ToUpdateRepoFromService(req))
}
