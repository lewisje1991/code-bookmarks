// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Diary struct {
	ID        pgtype.UUID
	Day       pgtype.Date
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type DiaryTask struct {
	DiaryID pgtype.UUID
	TaskID  pgtype.UUID
	Status  string
}

type Task struct {
	ID          pgtype.UUID
	Title       string
	Description string
	Tags        string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Worklog struct {
	ID        pgtype.UUID
	TaskID    pgtype.UUID
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
