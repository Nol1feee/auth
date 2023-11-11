package auth

import (
	"context"
	"errors"
	"fmt"
	model "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

// TODO перенести куда-то
type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

// TODO перенести куда-то
func NewPostgresDB(cfg Config, ctx context.Context) (*pgxpool.Pool, error) {
	dbDSN := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.Port,
		cfg.Name, cfg.User, cfg.Password, cfg.SSLMode)

	con, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		return &pgxpool.Pool{}, errors.New(fmt.Sprintf("repository - postgres - connect to DB | %s", err))
	}

	return con, nil
}

type User struct {
	*pgxpool.Pool
}

func (u *User) CreateUser(ctx context.Context, req *model.UserInfo) (int64, error) {
	dbReq := fmt.Sprintf("insert into %s (name, email, role) values ($1, $2, $3)", model.UsersTable)

	_, err := u.Pool.Exec(ctx, dbReq, req.Name, req.Email, req.Role)
	if err != nil {
		log.Fatal(err, "Exec")
	}

	id, err := u.getId(ctx, req.Email)

	return id, err
}

func (u *User) getId(ctx context.Context, email string) (int64, error) {
	var id int64

	dbReq := fmt.Sprintf("select id from %s where email=$1", model.UsersTable)
	err := u.Pool.QueryRow(ctx, dbReq, email).Scan(&id)

	return id, err
}

//
//func (u *User) GetUser(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
//	var name, email string
//	var role int
//	dbReq := fmt.Sprintf("select name, email, role from %s where id=14", repository.UsersTable)
//	err := u.Pool.QueryRow(ctx, dbReq).Scan(&name, &email, &role)
//	if err != nil {
//		log.Fatal(err, "query")
//	}
//	logrus.Info(name, "|", email, "|", role)
//
//	return &desc.GetResponse{
//		UserInfo: &desc.UserInfo{
//			Name:  name,
//			Email: email,
//			Role:  desc.Role(role),
//		},
//		CreatedAt: nil,
//		UpdatedAt: nil,
//	}, nil
//}
//
//func (u *User) DeleteUser(ctx context.Context, req *desc.DeleteRequest) error {
//	dbReq := fmt.Sprintf("delete from %s where id=$1", repository.UsersTable)
//	_, err := u.Pool.Exec(ctx, dbReq, req.Id)
//	return err
//}
//
//func (u *User) UpdateInfoUser(ctx context.Context, req *desc.UpdateRequest) error {
//	values := make([]string, 0)
//	args := make([]interface{}, 0)
//	argId := 1
//
//	if req.Name != nil {
//		values = append(values, fmt.Sprintf("name=$%d", argId))
//		args = append(args, req.Name.GetValue())
//		argId++
//	}
//
//	if req.Email != nil {
//		values = append(values, fmt.Sprintf("email=$%d", argId))
//		args = append(args, req.Email.GetValue())
//		argId++
//	}
//
//	testQuery := strings.Join(values, ", ")
//
//	dbReq := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", repository.UsersTable, testQuery, len(values)+1)
//
//	args = append(args, req.Id.GetValue())
//
//	_, err := u.Pool.Exec(ctx, dbReq, args...)
//
//	return err
//}
