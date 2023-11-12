package model

import "time"

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type UserInfo struct {
	Name  string
	Email string
	Role  string
}

type User struct {
	UserInfo  *UserInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdate struct {
	UserInfo *UserInfo
	Id       int64
}
