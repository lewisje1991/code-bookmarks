-- name: CreateBookmark :one
INSERT INTO bookmarks (id, url, description, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetBookmark :one
SELECT * FROM bookmarks WHERE id = $1;

-- name: CreateNote :one
INSERT INTO notes (id, title, content, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetNote :one
SELECT * FROM notes WHERE id = $1;

-- name: CreateDiary :one
INSERT INTO diary (id, day, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *;
