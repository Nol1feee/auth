package auth

import (
	"context"
	//"github.com/Nol1feee/CLI-chat/auth/internal/api/auth"
	//authService "github.com/Nol1feee/CLI-chat/auth/internal/service/auth"
)

func (s *Serv) Delete(ctx context.Context, id int64) error {
	return s.AuthRepository.Delete(ctx, id)
}
