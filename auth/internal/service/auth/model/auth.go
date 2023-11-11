package model

const (
	RoleUserServ  = "user"
	RoleAdminServ = "admin"
)

type UserInfoServ struct {
	Name  string
	Email string
	Role  string
}

type UserServ struct {
	UserInfo        *UserInfoServ
	Password        string
	PasswordConfirm string
}
