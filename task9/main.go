package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	canbusv1 "go-kozadayev-exercise/task9/api/proto"
	"go-kozadayev-exercise/task9/internal/service"
)

func main() {
	svc := service.NewCanBusService()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	canbusv1.RegisterCanBusServiceServer(s, svc)

	go s.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/alerts", func(w http.ResponseWriter, r *http.Request) {
		start, _ := strconv.ParseInt(r.URL.Query().Get("start_time"), 10, 64)
		end, _ := strconv.ParseInt(r.URL.Query().Get("end_time"), 10, 64)
		resp, _ := svc.GetAlerts(r.Context(), &canbusv1.GetAlertsRequest{StartTime: start, EndTime: end})
		json.NewEncoder(w).Encode(resp)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
