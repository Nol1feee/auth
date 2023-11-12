package main

import (
	"context"
	"fmt"
	"github.com/Nol1feee/CLI-chat/auth/internal/api/auth"
	authRepo "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth"
	authServ "github.com/Nol1feee/CLI-chat/auth/internal/service/auth"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	grpcPort = 50051
)

func main() {
	ctx := context.Background()

	var config authRepo.Config
	if err := envconfig.Process("db", &config); err != nil {
		logrus.Fatal("main - process", err)
	}

	con, err := authRepo.NewPostgresDB(config, ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	defer con.Close()
	logrus.Info("DB is up")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logrus.Fatal(err)
	}

	authRep := authRepo.NewRepo(con)
	serv := authServ.NewService(authRep)

	s := grpc.NewServer()
	//turn on serviceDesc (у сервера можно запросить описание его методов)
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, auth.NewImplementation(serv))

	logrus.Info("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		logrus.Fatal("Failed to serve %v", err)
	}
}
