// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Note      string    `json:"note"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
}
