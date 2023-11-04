package main

import (
	"context"
	"fmt"
	desc "github.com/Nol1feee/CLI-chat/chat-server/pkg/chatServer_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type server struct {
	desc.UnimplementedChatServerV1Server
}

func (s *server) Create(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error) {
	logrus.Info("wow, method CREATE in chatServer was implemented.. I'm return you id 9")
	return &desc.CreateResponse{Id: 9}, nil
}
func (s *server) Delete(context.Context, *desc.DeleteRequest) (*emptypb.Empty, error) {
	logrus.Info("wow, method DELETE in chatServer was implemented.")
	return &empty.Empty{}, nil
}
func (s *server) SendMessage(context.Context, *desc.SendMessageRequest) (*emptypb.Empty, error) {
	logrus.Info("wow, method SENDMESSAGE in chatServer was implemented.")
	return &empty.Empty{}, nil
}

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()
	//turn on serviceDesc (у сервера можно запросить описание его методов)
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &server{})

	logrus.Info("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		logrus.Fatal("Failed to serve %v", err)
	}
}
