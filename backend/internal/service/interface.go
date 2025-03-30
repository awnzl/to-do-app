package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/awnzl/to-do-app/internal/models"
)

type TodoService interface {
	// List operations
	CreateList(ctx context.Context, name string) (*models.TodoList, error)
	GetList(ctx context.Context, id uuid.UUID) (*models.TodoList, error)
	UpdateList(ctx context.Context, list *models.TodoList) error
	DeleteList(ctx context.Context, id uuid.UUID) error
	ListLists(ctx context.Context) ([]*models.TodoList, error)

	// Todo operations
	CreateTodo(ctx context.Context, listID uuid.UUID, title, description string, dueDate *time.Time) (*models.Todo, error)
	GetTodo(ctx context.Context, id uuid.UUID) (*models.Todo, error)
	UpdateTodo(ctx context.Context, todo *models.Todo) error
	MoveTodoToList(ctx context.Context, todoID, newListID uuid.UUID) error
	CompleteTodo(ctx context.Context, todoID uuid.UUID) error
	DeleteTodo(ctx context.Context, id uuid.UUID) error
	ListTodos(ctx context.Context, listID uuid.UUID) ([]*models.Todo, error)
	ListOverdueTodos(ctx context.Context) ([]*models.Todo, error)
}
