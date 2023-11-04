package main

import (
	"context"
	"fmt"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

const (
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedAuthV1Server
}

func (s *server) Update(context.Context, *desc.UpdateRequest) (*emptypb.Empty, error) {
	logrus.Info("wow, method UPDATE was implemented.. I'm return you id 10")
	return &empty.Empty{}, nil
}
func (s *server) Delete(context.Context, *desc.DeleteRequest) (*emptypb.Empty, error) {
	logrus.Info("wow, method DELETE was implemented.. I'm return you id 10")
	return &empty.Empty{}, nil
}
func (s *server) Create(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error) {
	logrus.Info("wow, method CREATE was implemented.. I'm return you id 10")
	return &desc.CreateResponse{Id: 10}, nil
}

func (s *server) Get(context.Context, *desc.GetRequest) (*desc.GetResponse, error) {
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

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()
	//turn on serviceDesc (у сервера можно запросить описание его методов)
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	logrus.Info("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		logrus.Fatal("Failed to serve %v", err)
	}
}
