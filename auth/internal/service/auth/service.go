package service

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/repository"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) repository.AuthRepository {
	return authRepository
}
