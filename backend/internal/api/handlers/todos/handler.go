package todos

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/awnzl/to-do-app/internal/api/models"
	"github.com/awnzl/to-do-app/internal/service"
)

type Handler struct {
	svc service.TodoService
}

func NewHandler(svc service.TodoService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterCreateRoute(r chi.Router) {
	r.Post("/", h.Create)
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/{todoID}", func(r chi.Router) {
		r.Get("/", h.GetByID)
		r.Put("/", h.Update)
		r.Delete("/", h.Delete)
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	listID, err := uuid.Parse(chi.URLParam(r, "listID"))
	if err != nil {
		http.Error(w, "invalid list ID", http.StatusBadRequest)
		return
	}

	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.svc.CreateTodo(r.Context(), listID, req.Title, req.Description, req.DueDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	todoID, err := uuid.Parse(chi.URLParam(r, "todoID"))
	if err != nil {
		http.Error(w, "invalid todo ID", http.StatusBadRequest)
		return
	}

	todo, err := h.svc.GetTodo(r.Context(), todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	todoID, err := uuid.Parse(chi.URLParam(r, "todoID"))
	if err != nil {
		http.Error(w, "invalid todo ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.svc.GetTodo(r.Context(), todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.DueDate = req.DueDate
	todo.Status = req.Status

	if err := h.svc.UpdateTodo(r.Context(), todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	todoID, err := uuid.Parse(chi.URLParam(r, "todoID"))
	if err != nil {
		http.Error(w, "invalid todo ID", http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteTodo(r.Context(), todoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
