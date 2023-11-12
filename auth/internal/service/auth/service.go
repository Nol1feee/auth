package auth

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/repository"
	"github.com/Nol1feee/CLI-chat/auth/internal/service"
)

var _ service.AuthService = (*Serv)(nil)

type Serv struct {
	AuthRepository repository.AuthRepository
}

func NewService(repo repository.AuthRepository) service.AuthService {
	return &Serv{AuthRepository: repo}
}
