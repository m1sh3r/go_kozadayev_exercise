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

	userv1 "go-kozadayev-exercise/task2/api/proto"
	"go-kozadayev-exercise/task2/internal/repository"
	"go-kozadayev-exercise/task2/internal/service"
)

func main() {
	ctx := context.Background()
	conn, _ := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")

	repo := repository.New(conn)
	svc := service.NewUserService(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	userv1.RegisterUserServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) == 3 && parts[0] == "users" && parts[2] == "documents" && r.Method == http.MethodGet {
			if conn == nil {
				http.Error(w, "db unavailable", http.StatusServiceUnavailable)
				return
			}
			userID, _ := strconv.ParseInt(parts[1], 10, 64)
			resp, err := svc.GetUserDocuments(r.Context(), &userv1.GetUserDocumentsRequest{UserId: userID})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
