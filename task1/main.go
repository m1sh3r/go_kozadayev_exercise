package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	carv1 "go-kozadayev-exercise/task1/api/proto"
	"go-kozadayev-exercise/task1/internal/repository"
	"go-kozadayev-exercise/task1/internal/service"
)

func main() {
	repo := repository.NewMemoryRepository()
	svc := service.NewCarService(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	carv1.RegisterCarServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	gwmux := runtime.NewServeMux()
	if err := carv1.RegisterCarServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", gwmux))
}
