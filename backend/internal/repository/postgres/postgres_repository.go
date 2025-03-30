package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/awnzl/to-do-app/internal/models"
	"github.com/awnzl/to-do-app/internal/repository"
)

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) repository.Repository {
	return &todoRepo{db: db}
}

func (r *todoRepo) CreateList(ctx context.Context, name string) (*models.TodoList, error) {
	list := &models.TodoList{
		ID:   uuid.New(),
		Name: name,
	}
	query := `
		INSERT INTO todo_lists (id, name)
		VALUES ($1, $2)
		RETURNING created_at`

	if err := r.db.GetContext(ctx, &list.CreatedAt, query, list.ID, list.Name); err != nil {
		return nil, fmt.Errorf("failed to create list: %w", err)
	}

	return list, nil
}

func (r *todoRepo) GetList(ctx context.Context, id uuid.UUID) (*models.TodoList, error) {
	list := &models.TodoList{}
	query := `
		SELECT id, name, created_at
		FROM todo_lists
		WHERE id = $1`

	if err := r.db.GetContext(ctx, list, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrListNotFound
		}
		return nil, fmt.Errorf("failed to get list: %w", err)
	}

	return list, nil
}

func (r *todoRepo) UpdateList(ctx context.Context, list *models.TodoList) error {
	query := `
		UPDATE todo_lists
		SET name = $1
		WHERE id = $2`

	if _, err := r.db.ExecContext(ctx, query, list.Name, list.ID); err != nil {
		return fmt.Errorf("failed to update list: %w", err)
	}

	return nil
}

func (r *todoRepo) DeleteList(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM todo_lists
		WHERE id = $1`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete list: %w", err)
	}

	return nil
}

func (r *todoRepo) ListLists(ctx context.Context) ([]*models.TodoList, error) {
	lists := make([]*models.TodoList, 0)
	query := `
		SELECT id, name, created_at
		FROM todo_lists`

	if err := r.db.SelectContext(ctx, &lists, query); err != nil {
		return nil, fmt.Errorf("failed to list lists: %w", err)
	}

	return lists, nil
}

func (r *todoRepo) CreateTodo(ctx context.Context, listID uuid.UUID, title, description string, dueDate *time.Time) (*models.Todo, error) {
	query := `
		INSERT INTO todos (id, list_id, title, description, due_date, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at`

	todo := &models.Todo{
		ID:          uuid.New(),
		ListID:      listID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      false,
	}

	if err := r.db.QueryRowContext(
		ctx,
		query,
		todo.ID,
		todo.ListID,
		todo.Title,
		todo.Description,
		todo.DueDate,
		todo.Status,
	).Scan(&todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return todo, nil
}

func (r *todoRepo) GetTodo(ctx context.Context, id uuid.UUID) (*models.Todo, error) {
	todo := &models.Todo{}
	query := `
		SELECT id, list_id, title, description, due_date, status, created_at, updated_at
		FROM todos
		WHERE id = $1`

	if err := r.db.GetContext(ctx, todo, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrTodoNotFound
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return todo, nil
}

func (r *todoRepo) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	query := `
		UPDATE todos
		SET title = $1, description = $2, due_date = $3, status = $4
		WHERE id = $5`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		todo.Title,
		todo.Description,
		todo.DueDate,
		todo.Status,
		todo.ID,
	); err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	return nil
}

func (r *todoRepo) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM todos
		WHERE id = $1`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	return nil
}

func (r *todoRepo) ListTodos(ctx context.Context, listID uuid.UUID) ([]*models.Todo, error) {
	var todos []*models.Todo
	query := `
		SELECT id, list_id, title, description, due_date, status, created_at, updated_at
		FROM todos
		WHERE list_id = $1`

	if err := r.db.SelectContext(ctx, &todos, query, listID); err != nil {
		return nil, fmt.Errorf("failed to list todos: %w", err)
	}

	return todos, nil
}

func (r *todoRepo) ListOverdueTodos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	query := `
		SELECT id, list_id, title, description, due_date, status, created_at, updated_at
		FROM todos
		WHERE due_date IS NOT NULL AND due_date < NOW() AND status = false`

	if err := r.db.SelectContext(ctx, &todos, query); err != nil {
		return nil, fmt.Errorf("failed to list overdue todos: %w", err)
	}

	return r.ListTodos(ctx, uuid.Nil)
}
