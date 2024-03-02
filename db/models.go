// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bucket struct {
	ID          int64
	Name        string
	Public      bool
	Description pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type File struct {
	ID        int64
	BucketID  int64
	Name      string
	Size      int64
	MimeType  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
