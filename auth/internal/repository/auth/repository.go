package auth

import (
	"context"
	"fmt"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strings"
	"time"
)

var _ repository.AuthRepository = (*User)(nil)

const (
	UsersTable = "users"

	pgId      = "id"
	pgName    = "name"
	pgEmail   = "email"
	pgRole    = "role"
	roleUser  = "user"
	roleAdmin = "admin"
)

// TODO перенести куда-то

type User struct {
	*pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) repository.AuthRepository {
	return &User{db}
}

func (u *User) Create(ctx context.Context, req *model.UserInfo) (int64, error) {
	dbReq := fmt.Sprintf("insert into %s (%s, %s, %s) values ($1, $2, $3)", UsersTable, pgName, pgEmail, pgRole)

	_, err := u.Pool.Exec(ctx, dbReq, req.Name, req.Email, req.Role)
	if err != nil {
		log.Fatal(err, "Exec")
	}

	id, err := u.getId(ctx, req.Email)

	return id, err
}

func (u *User) getId(ctx context.Context, email string) (int64, error) {
	var id int64

	dbReq := fmt.Sprintf("select %s from %s where %s=$1", pgId, UsersTable, pgEmail)
	err := u.Pool.QueryRow(ctx, dbReq, email).Scan(&id)

	return id, err
}

func (u *User) Delete(ctx context.Context, id int64) error {
	dbReq := fmt.Sprintf("delete from %s where id=$1", UsersTable)

	_, err := u.Pool.Exec(ctx, dbReq, id)
	return err
}

func (u *User) Get(ctx context.Context, id int64) (*model.User, error) {
	var name, email, role string

	dbReq := fmt.Sprintf("select %s, %s, %s from %s where %s=$1", pgName, pgEmail, pgRole, UsersTable, pgId)
	err := u.Pool.QueryRow(ctx, dbReq, id).Scan(&name, &email, &role)

	return &model.User{UserInfo: &model.UserInfo{
		Name:  name,
		Email: email,
		Role:  role,
	}, CreatedAt: time.Now(), UpdatedAt: time.Now()}, err
}

func (u *User) Update(ctx context.Context, req *model.UserUpdate) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if req.UserInfo.Name != "" {
		values = append(values, fmt.Sprintf("name=$%d", argId))
		args = append(args, req.UserInfo.Name)
		argId++
	}

	if req.UserInfo.Email != "" {
		values = append(values, fmt.Sprintf("email=$%d", argId))
		args = append(args, req.UserInfo.Email)
		argId++
	}

	testQuery := strings.Join(values, ", ")

	dbReq := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", UsersTable, testQuery, len(values)+1)

	args = append(args, req.Id)

	_, err := u.Pool.Exec(ctx, dbReq, args...)

	return err
}
