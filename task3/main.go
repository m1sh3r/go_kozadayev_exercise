package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"

	assemblyv1 "go-kozadayev-exercise/task3/api/proto"
	"go-kozadayev-exercise/task3/internal/repository"
	"go-kozadayev-exercise/task3/internal/service"
)

func main() {
	ctx := context.Background()
	conn, _ := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")

	repo := repository.New(conn)
	svc := service.NewAssemblyService(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	assemblyv1.RegisterAssemblyServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/assemble", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req assemblyv1.AssembleCarRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := svc.AssembleCar(r.Context(), &req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	mux.HandleFunc("/v1/cars/", func(w http.ResponseWriter, r *http.Request) {
		// URL: /v1/cars/{id}/spec -> parts = ["v1", "cars", "{id}", "spec"]
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) == 4 && parts[0] == "v1" && parts[1] == "cars" && parts[3] == "spec" && r.Method == http.MethodGet {
			id, _ := strconv.ParseInt(parts[2], 10, 64)
			resp, err := svc.GetCarSpec(r.Context(), &assemblyv1.GetCarSpecRequest{CarId: id})
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
