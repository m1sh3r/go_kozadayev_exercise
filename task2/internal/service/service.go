package service

import (
	"context"

	userv1 "go-kozadayev-exercise/task2/api/proto"
	"go-kozadayev-exercise/task2/internal/repository"
)

type UserService struct {
	userv1.UnimplementedUserServiceServer
	repo repository.Querier
}

func NewUserService(repo repository.Querier) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	u, err := s.repo.CreateUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &userv1.CreateUserResponse{
		User: &userv1.User{
			Id:   u.ID,
			Name: u.Name,
		},
	}, nil
}

func (s *UserService) AddDocumentToUser(ctx context.Context, req *userv1.AddDocumentToUserRequest) (*userv1.AddDocumentToUserResponse, error) {
	d, err := s.repo.AddDocument(ctx, repository.AddDocumentParams{
		Title:  req.Title,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &userv1.AddDocumentToUserResponse{
		Document: &userv1.Document{
			Id:     d.ID,
			Title:  d.Title,
			UserId: d.UserID,
		},
	}, nil
}

func (s *UserService) GetUserDocuments(ctx context.Context, req *userv1.GetUserDocumentsRequest) (*userv1.GetUserDocumentsResponse, error) {
	docs, err := s.repo.GetUserDocuments(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res := make([]*userv1.Document, len(docs))
	for i, d := range docs {
		res[i] = &userv1.Document{
			Id:     d.ID,
			Title:  d.Title,
			UserId: d.UserID,
		}
	}
	return &userv1.GetUserDocumentsResponse{Documents: res}, nil
}
