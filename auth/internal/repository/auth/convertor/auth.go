package convertor

import (
	model "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
	modelServ "github.com/Nol1feee/CLI-chat/auth/internal/service/auth/model"
)

func ToRepoFromService(serv *modelServ.UserServ) *model.User {
	return &model.User{UserInfo: &model.UserInfo{
		Name:  serv.UserInfo.Name,
		Email: serv.UserInfo.Email,
		Role:  serv.UserInfo.Role,
	}, Password: serv.Password, PasswordConfirm: serv.PasswordConfirm}
}
