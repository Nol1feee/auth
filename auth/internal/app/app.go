package app

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/config"
	"github.com/Nol1feee/CLI-chat/auth/internal/interceptor"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/Nol1feee/CLI-chat/auth/pkg/closer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
	"syscall"
)

const (
	//todo пофиксить, тк подобная инфа не должна быть в константe
	envPath = "../.env"
)

type App struct {
	serviceProvider *serviceProvider
	grpc            *grpc.Server
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return &App{}, err
	}

	log.Println("init all deps")

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initCloser,
		a.initGrpc,
		a.initHTTPServer,
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
	a.grpc = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.InterceptorValidate),
	)

	reflection.Register(a.grpc)

	desc.RegisterAuthV1Server(a.grpc, a.serviceProvider.NewImplementation(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterAuthV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GetGRPCConfig().GRPCAdress(), opts)
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.GetHTTPConfig().HTTPAddress(),
		Handler: mux,
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		a.serviceProvider.closer.New(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		a.serviceProvider.closer.Wait()
		log.Println("graceful shutdown, bye!")
		a.grpc.GracefulStop()
		a.httpServer.Shutdown(ctx)
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runGrpcServer()
		if err != nil {
			log.Fatalf("Failed to run GRPC server | %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("Failed to run HTTP server | %s", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) runGrpcServer() error {
	log.Printf("gRPC server is running on %s\n", a.serviceProvider.GetGRPCConfig().GRPCAdress())

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

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.GetHTTPConfig().HTTPAddress())

	err := a.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
