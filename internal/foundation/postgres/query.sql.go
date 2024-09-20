// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addTaskToDiary = `-- name: AddTaskToDiary :one
INSERT INTO diary_tasks (task_id, diary_id, status) 
VALUES ($1, $2, $3) 
RETURNING diary_id, task_id, status
`

type AddTaskToDiaryParams struct {
	TaskID  pgtype.UUID
	DiaryID pgtype.UUID
	Status  string
}

func (q *Queries) AddTaskToDiary(ctx context.Context, arg AddTaskToDiaryParams) (DiaryTask, error) {
	row := q.db.QueryRow(ctx, addTaskToDiary, arg.TaskID, arg.DiaryID, arg.Status)
	var i DiaryTask
	err := row.Scan(&i.DiaryID, &i.TaskID, &i.Status)
	return i, err
}

const createDiary = `-- name: CreateDiary :one
INSERT INTO diary (id, day, created_at, updated_at) 
VALUES ($1, $2, $3, $4) 
RETURNING id, day, created_at, updated_at
`

type CreateDiaryParams struct {
	ID        pgtype.UUID
	Day       pgtype.Date
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) CreateDiary(ctx context.Context, arg CreateDiaryParams) (Diary, error) {
	row := q.db.QueryRow(ctx, createDiary,
		arg.ID,
		arg.Day,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Diary
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (id, title, description, tags, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING id, title, description, tags, created_at, updated_at
`

type CreateTaskParams struct {
	ID          pgtype.UUID
	Title       string
	Description string
	Tags        string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, createTask,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Tags,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDiary = `-- name: GetDiary :one
SELECT id, day, created_at, updated_at 
FROM diary 
WHERE id = $1
`

func (q *Queries) GetDiary(ctx context.Context, id pgtype.UUID) (Diary, error) {
	row := q.db.QueryRow(ctx, getDiary, id)
	var i Diary
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDiaryByDay = `-- name: GetDiaryByDay :one
SELECT id, day, created_at, updated_at 
FROM diary 
WHERE day = $1
`

func (q *Queries) GetDiaryByDay(ctx context.Context, day pgtype.Date) (Diary, error) {
	row := q.db.QueryRow(ctx, getDiaryByDay, day)
	var i Diary
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTask = `-- name: GetTask :one
SELECT id, title, description, tags, created_at, updated_at 
FROM tasks 
WHERE id = $1
`

func (q *Queries) GetTask(ctx context.Context, id pgtype.UUID) (Task, error) {
	row := q.db.QueryRow(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTasksByDiary = `-- name: GetTasksByDiary :many
SELECT tasks.id, tasks.title, tasks.description, tasks.tags, tasks.created_at, tasks.updated_at 
FROM tasks 
JOIN diary_tasks ON tasks.id = diary_tasks.task_id 
WHERE diary_tasks.diary_id = $1
`

func (q *Queries) GetTasksByDiary(ctx context.Context, diaryID pgtype.UUID) ([]Task, error) {
	rows, err := q.db.Query(ctx, getTasksByDiary, diaryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Tags,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
