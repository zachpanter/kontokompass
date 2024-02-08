// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package storage

import (
	"database/sql"
	"time"
)

type Classification struct {
	ID   int32
	Name string
}

type Transaction struct {
	ID                 int32
	Postdate           time.Time
	Description        string
	Debit              sql.NullFloat64
	Credit             sql.NullFloat64
	Balance            float32
	ClassificationText string
	ClassificationID   sql.NullInt32
}
