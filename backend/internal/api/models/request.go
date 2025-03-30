package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateListRequest struct {
	Name string `json:"name"`
}

type UpdateListRequest struct {
	Name string `json:"name"`
}

type CreateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Status      bool       `json:"status"`
}

type MoveTodoRequest struct {
	TargetListID uuid.UUID `json:"target_list_id"`
}
