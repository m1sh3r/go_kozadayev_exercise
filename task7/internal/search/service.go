package search

import (
	"context"
	"log"

	searchenginev1 "go-kozadayev-exercise/task7/api/proto"
)

type SearchService struct {
	searchenginev1.UnimplementedSearchServiceServer
}

func (s *SearchService) IndexArticle(ctx context.Context, req *searchenginev1.IndexArticleRequest) (*searchenginev1.IndexArticleResponse, error) {
	log.Printf("Indexing article: %s", req.Article.Title)
	return &searchenginev1.IndexArticleResponse{Success: true}, nil
}
