package main

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/api/auth"
	"github.com/Nol1feee/CLI-chat/auth/internal/config"
	authRepo "github.com/Nol1feee/CLI-chat/auth/internal/repository/auth"
	authServ "github.com/Nol1feee/CLI-chat/auth/internal/service/auth"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	ctx := context.Background()

	err := config.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	pgCfg, err := config.NewPGConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfxGRPC, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}

	con, err := pgxpool.Connect(ctx, pgCfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()
	logrus.Info("DB is up")

	lis, err := net.Listen("tcp", cfxGRPC.GRPCAdress())
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
