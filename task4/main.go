package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	inventoryv1 "go-kozadayev-exercise/task4/api/proto"
	"go-kozadayev-exercise/task4/internal/service"
)

func main() {
	svc := service.NewInventoryService()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	inventoryv1.RegisterInventoryServiceServer(s, svc)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req inventoryv1.CreateOrderRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := svc.CreateOrder(r.Context(), &req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
