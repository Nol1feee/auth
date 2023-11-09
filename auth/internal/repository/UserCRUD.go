package repository

import (
	"context"
	"fmt"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
)

type User struct {
	*pgxpool.Pool
}

func (u *User) CreateUser(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	dbReq := fmt.Sprintf("insert into %s (name, email, role) values ($1, $2, $3)", usersTable)

	_, err := u.Pool.Exec(ctx, dbReq, req.UserInfo.Name, req.UserInfo.Email, req.UserInfo.Role)
	if err != nil {
		log.Fatal(err, "Exec")
	}

	id, err := u.GetId(ctx, req.UserInfo.Email)

	return &desc.CreateResponse{Id: int64(id)}, err
}

func (u *User) GetId(ctx context.Context, email string) (int64, error) {
	var id int64

	dbReq := fmt.Sprintf("select id from %s where email=$1", usersTable)
	err := u.Pool.QueryRow(ctx, dbReq, email).Scan(&id)

	return id, err
}

func (u *User) GetUser(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	var name, email string
	var role int
	dbReq := fmt.Sprintf("select name, email, role from %s where id=14", usersTable)
	err := u.Pool.QueryRow(ctx, dbReq).Scan(&name, &email, &role)
	if err != nil {
		log.Fatal(err, "query")
	}
	logrus.Info(name, "|", email, "|", role)

	return &desc.GetResponse{
		UserInfo: &desc.UserInfo{
			Name:  name,
			Email: email,
			Role:  desc.Role(role),
		},
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil
}

func (u *User) DeleteUser(ctx context.Context, req *desc.DeleteRequest) error {
	dbReq := fmt.Sprintf("delete from %s where id=$1", usersTable)
	_, err := u.Pool.Exec(ctx, dbReq, req.Id)
	return err
}

//UpdateInfoUser(ctx context.Context, req *desc.UpdateRequest) error
