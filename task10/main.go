package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	maintenancev1 "go-kozadayev-exercise/task10/api/proto"
	"go-kozadayev-exercise/task10/internal/middleware"
	"go-kozadayev-exercise/task10/internal/service"
)

func main() {
	svc := service.NewMaintenanceService()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChangelogInterceptor),
	)
	maintenancev1.RegisterMaintenanceServiceServer(s, svc)

	go s.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/workorders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req maintenancev1.CreateWorkOrderRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
				return
			}
			resp, err := svc.CreateWorkOrder(r.Context(), &req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	})
	mux.HandleFunc("/v1/workorders/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := strings.TrimPrefix(r.URL.Path, "/v1/workorders/")
			resp, _ := svc.GetWorkOrder(r.Context(), &maintenancev1.GetWorkOrderRequest{Id: id})
			json.NewEncoder(w).Encode(resp)
		}
	})
	mux.HandleFunc("/v1/export", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := svc.ExportToExcel(r.Context(), &maintenancev1.ExportRequest{})
		json.NewEncoder(w).Encode(resp)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
