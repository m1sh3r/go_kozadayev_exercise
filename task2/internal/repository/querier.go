package repository

import (
	"context"
)

type Querier interface {
	AddDocument(ctx context.Context, arg AddDocumentParams) (Document, error)
	CreateUser(ctx context.Context, name string) (User, error)
	GetUserDocuments(ctx context.Context, userID int64) ([]Document, error)
}

var _ Querier = (*Queries)(nil)
