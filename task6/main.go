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

	dashboardv1 "go-kozadayev-exercise/task6/api/proto"
	"go-kozadayev-exercise/task6/internal/repository"
	"go-kozadayev-exercise/task6/internal/service"
)

func main() {
	ctx := context.Background()
	conn, _ := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")

	repo := repository.New(conn)
	svc := service.NewDashboardService(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	dashboardv1.RegisterDashboardServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/owners/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) == 3 && parts[0] == "v1" && parts[1] == "owners" && parts[2] == "dashboard" && r.Method == http.MethodGet {
			id, _ := strconv.ParseInt(parts[1], 10, 64)
			resp, err := svc.GetDashboard(r.Context(), &dashboardv1.GetDashboardRequest{OwnerId: id})
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
