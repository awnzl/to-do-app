package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ListID      uuid.UUID  `db:"list_id" json:"list_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description,omitempty"`
	DueDate     *time.Time `db:"due_date" json:"due_date,omitempty"`
	Status      bool       `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}
