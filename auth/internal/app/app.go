package app

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/closer"
	"github.com/Nol1feee/CLI-chat/auth/internal/config"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"syscall"
)

const (
	//todo пофиксить, тк подобная инфа не должна быть в константe
	envPath = "../.env"
)

type App struct {
	serviceProvider *serviceProvider
	grpc            *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return &App{}, err
	}

	logrus.Info("init all deps")

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initCloser,
		a.initGrpc,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initCloser(_ context.Context) error {
	a.serviceProvider.closer = closer.NewCloser()
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	return config.Load(envPath)
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
func (a *App) initGrpc(ctx context.Context) error {
	a.grpc = grpc.NewServer()

	reflection.Register(a.grpc)

	desc.RegisterAuthV1Server(a.grpc, a.serviceProvider.NewImplementation(ctx))

	return nil
}

func (a *App) Run() error {
	go func() {
		a.serviceProvider.closer.New(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		a.serviceProvider.closer.Wait()
		logrus.Info("успешно реализовал shutdown grace")
		a.grpc.GracefulStop()
	}()

	return a.runGrpcServer()
}

func (a *App) runGrpcServer() error {
	logrus.Info("listening %s adress", a.serviceProvider.GetGRPCConfig().GRPCAdress())

	lis, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().GRPCAdress())
	if err != nil {
		return err
	}

	err = a.grpc.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

//todo add logs
