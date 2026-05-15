package repository

import (
	"context"
)

const addDocument = `-- name: AddDocument :one
INSERT INTO documents (title, user_id) VALUES ($1, $2) RETURNING id, title, user_id
`

type AddDocumentParams struct {
	Title  string
	UserID int64
}

func (q *Queries) AddDocument(ctx context.Context, arg AddDocumentParams) (Document, error) {
	row := q.db.QueryRow(ctx, addDocument, arg.Title, arg.UserID)
	var i Document
	err := row.Scan(&i.ID, &i.Title, &i.UserID)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRow(ctx, createUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserDocuments = `-- name: GetUserDocuments :many
SELECT id, title, user_id FROM documents WHERE user_id = $1
`

func (q *Queries) GetUserDocuments(ctx context.Context, userID int64) ([]Document, error) {
	rows, err := q.db.Query(ctx, getUserDocuments, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Document
	for rows.Next() {
		var i Document
		if err := rows.Scan(&i.ID, &i.Title, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
