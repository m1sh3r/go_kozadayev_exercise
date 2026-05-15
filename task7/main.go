package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	searchenginev1 "go-kozadayev-exercise/task7/api/proto"
	"go-kozadayev-exercise/task7/internal/content"
	"go-kozadayev-exercise/task7/internal/search"
)

func main() {
	searchLis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	searchServer := grpc.NewServer()
	searchenginev1.RegisterSearchServiceServer(searchServer, &search.SearchService{})
	go searchServer.Serve(searchLis)

	searchConn, _ := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	searchClient := searchenginev1.NewSearchServiceClient(searchConn)

	contentSvc := content.NewContentService(searchClient)
	contentLis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	contentServer := grpc.NewServer()
	searchenginev1.RegisterContentServiceServer(contentServer, contentSvc)
	go contentServer.Serve(contentLis)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req searchenginev1.CreateArticleRequest
			json.NewDecoder(r.Body).Decode(&req)
			resp, _ := contentSvc.CreateArticle(r.Context(), &req)
			json.NewEncoder(w).Encode(resp)
		}
	})
	mux.HandleFunc("/v1/articles/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := strings.TrimPrefix(r.URL.Path, "/v1/articles/")
			resp, _ := contentSvc.GetArticle(r.Context(), &searchenginev1.GetArticleRequest{Id: id})
			json.NewEncoder(w).Encode(resp)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
