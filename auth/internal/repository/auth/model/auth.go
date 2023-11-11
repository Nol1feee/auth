package model

const (
	UsersTable = "users"

	roleUser  = "user"
	roleAdmin = "admin"
)

type UserInfo struct {
	Name  string
	Email string
	Role  string
}

type User struct {
	UserInfo        *UserInfo
	Password        string
	PasswordConfirm string
}
