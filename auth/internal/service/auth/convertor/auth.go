package convertor

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/service/auth/model"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
)

func ToServiceFromDesc(req *desc.CreateRequest) *model.UserServ {
	var userRole string = model.RoleUserServ

	if req.UserInfo.GetRole() == 1 {
		userRole = model.RoleUserServ
	}

	return &model.UserServ{UserInfo: &model.UserInfoServ{
		Name:  req.UserInfo.Name,
		Email: req.UserInfo.Email,
		Role:  userRole,
	}, Password: req.Password, PasswordConfirm: req.PasswordConfirm}
}
