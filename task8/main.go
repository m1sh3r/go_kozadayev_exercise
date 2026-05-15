package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"

	rentalv1 "go-kozadayev-exercise/task8/api/proto"
	"go-kozadayev-exercise/task8/internal/repository"
	"go-kozadayev-exercise/task8/internal/service"
)

func main() {
	ctx := context.Background()
	conn, _ := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")

	repo := repository.New(conn)
	svc := service.NewRentalService(conn, repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	rentalv1.RegisterRentalServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/bookings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req rentalv1.CreateBookingRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := svc.CreateBooking(r.Context(), &req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
