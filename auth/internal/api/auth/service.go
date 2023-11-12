package auth

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/service"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
