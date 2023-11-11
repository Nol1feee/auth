package auth

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/service"
)

type Implementation struct {
	//desc.UnimplementedAuthV1Server
	authService service.AuthService
}

//func NewImplementation(unimplementedAuthV1Server desc.UnimplementedAuthV1Server, authService service.AuthService) *Implementation {
//	return &Implementation{UnimplementedAuthV1Server: unimplementedAuthV1Server, authService: authService}
//}
