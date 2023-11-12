package auth

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
)

func (s *Serv) Get(ctx context.Context, id int64) (*model.User, error) {
	resp, err := s.AuthRepository.Get(ctx, id)
	return &model.User{UserInfo: &model.UserInfo{
		Name:  resp.UserInfo.Name,
		Email: resp.UserInfo.Email,
		Role:  resp.UserInfo.Role,
	}, CreatedAt: resp.CreatedAt, UpdatedAt: resp.UpdatedAt}, err
}
