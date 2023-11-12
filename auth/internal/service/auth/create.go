package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/convertor"
	"github.com/sirupsen/logrus"
)

func (s *Serv) Create(ctx context.Context, req *model.UserInfo) (int64, error) {
	logrus.Info("service-create")
	return s.AuthRepository.Create(ctx, convertor.ToRepoFromService(req))
}
