package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/awnzl/to-do-app/internal/models"
	"github.com/awnzl/to-do-app/internal/repository"
)

type todoService struct {
	repo repository.Repository
	txm  repository.TransactionManager
}

func NewTodoService(repo repository.Repository, txm repository.TransactionManager) *todoService {
	return &todoService{
		repo: repo,
		txm:  txm,
	}
}

func (s *todoService) CreateList(ctx context.Context, name string) (*models.TodoList, error) {
	var list *models.TodoList
	err := s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		list, err = s.repo.CreateList(ctx, name)
		return err
	})
	return list, err
}

func (s *todoService) GetList(ctx context.Context, id uuid.UUID) (*models.TodoList, error) {
	// read operations don't need transactions
	return s.repo.GetList(ctx, id)
}

func (s *todoService) UpdateList(ctx context.Context, list *models.TodoList) error {
	return s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return s.repo.UpdateList(ctx, list)
	})
}

func (s *todoService) DeleteList(ctx context.Context, id uuid.UUID) error {
	return s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return s.repo.DeleteList(ctx, id)
	})
}

func (s *todoService) ListLists(ctx context.Context) ([]*models.TodoList, error) {
	// read operations don't need transactions
	return s.repo.ListLists(ctx)
}

func (s *todoService) CreateTodo(
	ctx context.Context, listID uuid.UUID, title, description string, dueDate *time.Time,
) (*models.Todo, error) {
	var newTodo *models.Todo
	err := s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		newTodo, err = s.repo.CreateTodo(ctx, listID, title, description, dueDate)
		return err
	})
	return newTodo, err
}

func (s *todoService) GetTodo(ctx context.Context, id uuid.UUID) (*models.Todo, error) {
	// read operations don't need transactions
	return s.repo.GetTodo(ctx, id)
}

func (s *todoService) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	return s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return s.repo.UpdateTodo(ctx, todo)
	})
}

func (s *todoService) MoveTodoToList(ctx context.Context, todoID, newListID uuid.UUID) error {
	todo, err := s.GetTodo(ctx, todoID)
	if err != nil {
		return fmt.Errorf("getting todo '%s': %w", todoID.String(), err)
	}
	todo.ListID = newListID
	return s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return s.repo.UpdateTodo(ctx, todo)
	})
}

func (s *todoService) CompleteTodo(ctx context.Context, todoID uuid.UUID) error {
	todo, err := s.GetTodo(ctx, todoID)
	if err != nil {
		return fmt.Errorf("getting todo '%s': %w", todoID.String(), err)
	}
	todo.Status = true
	return s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return s.repo.UpdateTodo(ctx, todo)
	})
}

func (s *todoService) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	return s.txm.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return s.repo.DeleteTodo(ctx, id)
	})
}

func (s *todoService) ListTodos(ctx context.Context, listID uuid.UUID) ([]*models.Todo, error) {
	// read operations don't need transactions
	return s.repo.ListTodos(ctx, listID)
}

func (s *todoService) ListOverdueTodos(ctx context.Context) ([]*models.Todo, error) {
	return s.repo.ListOverdueTodos(ctx)
}
