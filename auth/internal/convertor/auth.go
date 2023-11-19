package convertor

import (
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
	modelRepo "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
)

func ToUserInfoFromDesc(req *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role.String(),
	}
}

func ToUserInfoFromService(req *model.UserInfo) *desc.UserInfo {
	var role int
	if req.Role == model.RoleAdmin {
		role = 1
	}
	return &desc.UserInfo{
		Name:  req.Name,
		Email: req.Email,
		Role:  desc.Role(role),
	}
}

func ToDescFromUser(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		UserInfo:  ToUserInfoFromService(user.UserInfo),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}

func ToInfoRepoFromService(req *model.UserInfo) *modelRepo.UserInfo {
	return &modelRepo.UserInfo{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
}

func ToUpdateRepoFromService(req *model.UserUpdate) *modelRepo.UserUpdate {
	return &modelRepo.UserUpdate{
		UserInfo: ToInfoRepoFromService(req.UserInfo),
		Id:       req.Id,
	}
}

func ToUpdateFromDesc(req *desc.UpdateRequest) *model.UserUpdate {
	return &model.UserUpdate{UserInfo: &model.UserInfo{
		Name:  req.Name.GetValue(),
		Email: req.Email.GetValue(),
		Role:  req.Role.String(),
	}, Id: req.Id.GetValue()}
}
