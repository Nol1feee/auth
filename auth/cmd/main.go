package main

import (
	"context"
	"fmt"
	"github.com/Nol1feee/CLI-chat/auth/internal/repository"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

const (
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedAuthV1Server
}

func (s *server) UpdateInfoUser(context.Context, *desc.UpdateRequest) (*emptypb.Empty, error) {
	logrus.Info("wow, method UPDATE was implemented.. I'm return you id 10")
	return &empty.Empty{}, nil
}
func (s *server) DeleteUser(context.Context, *desc.DeleteRequest) (*emptypb.Empty, error) {
	logrus.Info("wow, method DELETE was implemented.. I'm return you id 10")
	return &empty.Empty{}, nil
}
func (s *server) CreateUser(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error) {
	logrus.Info("wow, method CREATE was implemented.. I'm return you id 10")
	return &desc.CreateResponse{Id: 10}, nil
}

func (s *server) GetUser(context.Context, *desc.GetRequest) (*desc.GetResponse, error) {
	logrus.Info("wow, method GET was implemented.. I'm return you random data!")
	return &desc.GetResponse{
		UserInfo: &desc.UserInfo{
			Name:  gofakeit.BeerName(),
			Email: gofakeit.Email(),
			Role:  desc.Role_user,
		},
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

//type dbConfig struct {
//	Host     string
//	Port     string
//	Name     string
//	User     string
//	Password string
//	SSLMode  string
//}

func main() {

	ctx := context.Background()

	var s repository.Config
	if err := envconfig.Process("db", &s); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s)

	con, err := repository.NewPostgresDB(s, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	user := repository.User{Pool: con}
	id, err := user.CreateUser(ctx, &desc.CreateRequest{UserInfo: &desc.UserInfo{Name: "vlad",
		Email: "vladrazd02@yandex.ru", Role: 0}, Password: "ad10", PasswordConfirm: "ad10"})
	if err != nil {
		log.Fatal(err)
	}

	err = user.DeleteUser(ctx, &desc.DeleteRequest{Id: id.Id})
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info("DELETE user!")

	//resp, err := user.GetUser(ctx, &desc.GetRequest{Id: 14})
	//fmt.Printf("%s", resp.UserInfo.Role)

	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//
	//s := grpc.NewServer()
	////turn on serviceDesc (у сервера можно запросить описание его методов)
	//reflection.Register(s)
	//desc.RegisterAuthV1Server(s, &server{})
	//
	//logrus.Info("Server listening at %v", lis.Addr())
	//
	//if err = s.Serve(lis); err != nil {
	//	logrus.Fatal("Failed to serve %v", err)
	//}
}
