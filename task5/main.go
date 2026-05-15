package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	smarthomev1 "go-kozadayev-exercise/task5/api/proto"
	"go-kozadayev-exercise/task5/internal/middleware"
	"go-kozadayev-exercise/task5/internal/service"
)

func main() {
	svc := service.NewSmartHomeService()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
		grpc.StreamInterceptor(middleware.StreamAuthInterceptor),
	)
	smarthomev1.RegisterSmartHomeServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/readings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req smarthomev1.AddReadingRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := svc.AddReading(r.Context(), &req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", middleware.HTTPAuthMiddleware(mux)))
}
