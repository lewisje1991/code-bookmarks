// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package sqlite

import (
	"context"

	"github.com/google/uuid"
)

const createBookmark = `-- name: CreateBookmark :one
INSERT INTO bookmarks (id, url, description, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?) RETURNING id, url, description, tags, created_at, updated_at
`

type CreateBookmarkParams struct {
	ID          uuid.UUID
	Url         string
	Description string
	Tags        string
	CreatedAt   string
	UpdatedAt   string
}

func (q *Queries) CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error) {
	row := q.queryRow(ctx, q.createBookmarkStmt, createBookmark,
		arg.ID,
		arg.Url,
		arg.Description,
		arg.Tags,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Description,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBookmark = `-- name: GetBookmark :one
SELECT id, url, description, tags, created_at, updated_at FROM bookmarks WHERE id = ?
`

func (q *Queries) GetBookmark(ctx context.Context, id uuid.UUID) (Bookmark, error) {
	row := q.queryRow(ctx, q.getBookmarkStmt, getBookmark, id)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Description,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
