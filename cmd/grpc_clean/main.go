package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AlbertPuwadol/grpc-clean/config"
	handler "github.com/AlbertPuwadol/grpc-clean/pkg/handler/grpc"
	"github.com/AlbertPuwadol/grpc-clean/pkg/repository"
	"github.com/AlbertPuwadol/grpc-clean/pkg/service"
	"github.com/AlbertPuwadol/grpc-clean/pkg/usecase"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	jwtManager := service.NewJWTManager(cfg.Secret)
	interceptor := service.NewAuthInterceptor(jwtManager, cfg.GuessMe)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))

	repository := repository.NewRepository()
	usecase := usecase.NewUsecase(repository)

	handler.NewHandler(usecase, grpcServer)

	log.Printf("Listening on port: %d\n", cfg.Port)
	log.Fatal(grpcServer.Serve(lis))
}
