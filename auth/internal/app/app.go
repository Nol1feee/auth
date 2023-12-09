package app

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/config"
	"github.com/Nol1feee/CLI-chat/auth/internal/interceptor"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/Nol1feee/CLI-chat/auth/pkg/closer"
	_ "github.com/Nol1feee/CLI-chat/auth/statik"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"io"
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
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swaggerServer   *http.Server
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
		a.initSwaggerServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initSwaggerServer(ctx context.Context) error {
	staticFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(staticFs)))
	mux.HandleFunc("/api.swagger.json", serverSwaggerFile("/api.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.serviceProvider.GetSwaggerConfig().Address(),
		Handler: mux,
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
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.InterceptorValidate),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.NewImplementation(ctx))

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

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.GetHTTPConfig().HTTPAddress(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		a.serviceProvider.closer.New(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		a.serviceProvider.closer.Wait()
		log.Println("graceful shutdown, bye!")
		a.grpcServer.GracefulStop()
		a.httpServer.Shutdown(ctx)
		a.swaggerServer.Shutdown(ctx)
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)

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

	go func() {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("Failed to run SWAGGER server | %s", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) runSwaggerServer() error {
	log.Printf("swagger server is running on %s\n", a.serviceProvider.GetSwaggerConfig().Address())

	err := a.swaggerServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (a *App) runGrpcServer() error {
	log.Printf("gRPC server is running on %s\n", a.serviceProvider.GetGRPCConfig().GRPCAdress())

	lis, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().GRPCAdress())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(lis)
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

func serverSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving swagger file: %s", path)

		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Open swagger file: %s", path)

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		log.Printf("Read swagger file: %s", path)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Write swagger file: %s", path)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Served swagger file: %s", path)
	}
}
