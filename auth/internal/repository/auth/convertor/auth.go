package convertor

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
	modelRepo "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
)

func ToRepoFromService(req *model.UserInfo) *modelRepo.UserInfo {
	return &modelRepo.UserInfo{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
}
