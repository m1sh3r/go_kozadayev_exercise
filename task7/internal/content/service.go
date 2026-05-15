package content

import (
	"context"
	"fmt"
	"sync"

	searchenginev1 "go-kozadayev-exercise/task7/api/proto"

	"github.com/sony/gobreaker/v2"
)

type ContentService struct {
	searchenginev1.UnimplementedContentServiceServer
	searchClient searchenginev1.SearchServiceClient
	cb           *gobreaker.CircuitBreaker[any]
	mu           sync.RWMutex
	articles     map[string]*searchenginev1.Article
}

func NewContentService(client searchenginev1.SearchServiceClient) *ContentService {
	cb := gobreaker.NewCircuitBreaker[any](gobreaker.Settings{
		Name: "search-service",
	})
	return &ContentService{
		searchClient: client,
		cb:           cb,
		articles:     make(map[string]*searchenginev1.Article),
	}
}

func (s *ContentService) CreateArticle(ctx context.Context, req *searchenginev1.CreateArticleRequest) (*searchenginev1.Article, error) {
	id := fmt.Sprintf("%d", len(s.articles)+1)
	article := &searchenginev1.Article{
		Id:       id,
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: req.AuthorId,
	}

	s.mu.Lock()
	s.articles[id] = article
	s.mu.Unlock()

	_, err := s.cb.Execute(func() (any, error) {
		return s.searchClient.IndexArticle(ctx, &searchenginev1.IndexArticleRequest{
			Article: article,
		})
	})
	if err != nil {
		fmt.Printf("Circuit breaker error: %v\n", err)
	}

	return article, nil
}

func (s *ContentService) GetArticle(ctx context.Context, req *searchenginev1.GetArticleRequest) (*searchenginev1.Article, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.articles[req.Id]
	if !ok {
		return nil, fmt.Errorf("article not found")
	}
	return a, nil
}
