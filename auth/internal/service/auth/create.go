package service

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/convertor"
	model "github.com/Nol1feee/CLI-chat/auth/internal/service/auth/model"
)

func (s *serv) CreateUser(ctx context.Context, req *model.UserServ) (int64, error) {
	//convert
	return s.authRepository.CreateUser(ctx, convertor.ToRepoFromService(req))
}
